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
