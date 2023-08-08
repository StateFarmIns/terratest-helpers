package tests

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/stretchr/testify/assert"
)

func ValidateLambdaFunctionExists(t *testing.T, svc *lambda.Lambda, functionName string, layerName string, testLayer bool, verboseOutput bool) {
	t.Helper()

	getFunctionInput := &lambda.GetFunctionInput{
		FunctionName: aws.String(functionName),
	}

	fmt.Println("Validating Lambda: ", functionName)

	getFunctionResult, err := svc.GetFunction(getFunctionInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lambda.ErrCodeServiceException:
				fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
			case lambda.ErrCodeResourceNotFoundException:
				fmt.Println(lambda.ErrCodeResourceNotFoundException, aerr.Error())
			case lambda.ErrCodeTooManyRequestsException:
				fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
			case lambda.ErrCodeInvalidParameterValueException:
				fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

		t.Logf("Failing test.")
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(fmt.Println(getFunctionResult.String()))
	}

	assert.Equal(t, functionName, *getFunctionResult.Configuration.FunctionName)

	if testLayer {
		assert.Contains(t, *getFunctionResult.Configuration.Layers[0].Arn, layerName)
	}
}

func ValidateLambdaFunctionConfiguration(t *testing.T, svc *lambda.Lambda, functionName string, architecture string, handlerName string, layerNames []string, memorySize int64, packageType string, role string, runtime string, state string, timeout int64, vpcID string, subnets []string, securityGroups []string, verboseOutput bool) {
	t.Helper()

	getFunctionInput := &lambda.GetFunctionInput{
		FunctionName: aws.String(functionName),
	}

	fmt.Println("Validating Lambda: ", functionName)

	getFunctionResult, err := svc.GetFunction(getFunctionInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lambda.ErrCodeServiceException:
				fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
			case lambda.ErrCodeResourceNotFoundException:
				fmt.Println(lambda.ErrCodeResourceNotFoundException, aerr.Error())
			case lambda.ErrCodeTooManyRequestsException:
				fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
			case lambda.ErrCodeInvalidParameterValueException:
				fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

		t.Logf("Failing test.")
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(fmt.Println(getFunctionResult.String()))
	}

	assert.Equal(t, functionName, *getFunctionResult.Configuration.FunctionName)
	assert.Equal(t, architecture, *getFunctionResult.Configuration.Architectures[0])
	assert.Equal(t, handlerName, *getFunctionResult.Configuration.Handler)
	assert.Equal(t, *aws.Int64(memorySize), *getFunctionResult.Configuration.MemorySize)
	assert.Equal(t, packageType, *getFunctionResult.Configuration.PackageType)
	assert.Equal(t, role, *getFunctionResult.Configuration.Role)
	assert.Equal(t, runtime, *getFunctionResult.Configuration.Runtime)
	assert.Equal(t, state, *getFunctionResult.Configuration.State)
	assert.Equal(t, *aws.Int64(timeout), *getFunctionResult.Configuration.Timeout)
	assert.Equal(t, vpcID, *getFunctionResult.Configuration.VpcConfig.VpcId)
	// validate subnets
	for i := 0; i < len(subnets); i++ {
		assert.Contains(t, getFunctionResult.String(), subnets[i])

		if verboseOutput {
			fmt.Println(subnets[i])
		}
	}
	// validate securityGroups
	for j := 0; j < len(securityGroups); j++ {
		assert.Contains(t, getFunctionResult.String(), securityGroups[j])

		if verboseOutput {
			fmt.Println(securityGroups[j])
		}
	}

	// validate layers attached
	if len(layerNames) != 0 {
		for x := 0; x < len(layerNames); x++ {
			assert.Contains(t, getFunctionResult.String(), layerNames[x])

			if verboseOutput {
				fmt.Println(layerNames[x])
			}
		}
	}
}