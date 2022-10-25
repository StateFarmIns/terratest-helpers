package tests

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/wafv2"
	"github.com/stretchr/testify/assert"
)

// ValidateWAFV2WebACL validate base parameters of a WAFv2 Web ACL
func ValidateWAFV2WebACL(t *testing.T, svc *wafv2.WAFV2, webACLID string, webACLName string, webACLScope string, webACLARN string, verboseOutput bool) {
	t.Helper()

	getWebACLResult, err := svc.GetWebACL(
		&wafv2.GetWebACLInput{
			Id:    aws.String(webACLID),
			Name:  aws.String(webACLName),
			Scope: aws.String(webACLScope),
		},
	)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case wafv2.ErrCodeWAFInternalErrorException:
				fmt.Println(wafv2.ErrCodeWAFInternalErrorException, aerr.Error())
			case wafv2.ErrCodeWAFNonexistentItemException:
				fmt.Println(wafv2.ErrCodeWAFNonexistentItemException, aerr.Error())
			case wafv2.ErrCodeWAFInvalidParameterException:
				fmt.Println(wafv2.ErrCodeWAFInvalidParameterException, aerr.Error())
			case wafv2.ErrCodeWAFUnavailableEntityException:
				fmt.Println(wafv2.ErrCodeWAFUnavailableEntityException, aerr.Error())
			case wafv2.ErrCodeWAFInvalidOperationException:
				fmt.Println(wafv2.ErrCodeWAFInvalidOperationException, aerr.Error())
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
		fmt.Println(getWebACLResult.String())
	}

	assert.Equal(t, webACLARN, *getWebACLResult.WebACL.ARN)
	assert.Equal(t, webACLID, *getWebACLResult.WebACL.Id)
	assert.Equal(t, webACLName, *getWebACLResult.WebACL.Name)
}

// ValidateWAFV2WebACLRulesByName validate the expected names of rules are associated to a WAFv2 Web ACL
func ValidateWAFV2WebACLRulesByName(t *testing.T, svc *wafv2.WAFV2, webACLID string, webACLName string, webACLScope string, expectedRuleNameList []string, verboseOutput bool) {
	t.Helper()

	getWebACLResult, err := svc.GetWebACL(
		&wafv2.GetWebACLInput{
			Id:    aws.String(webACLID),
			Name:  aws.String(webACLName),
			Scope: aws.String(webACLScope),
		},
	)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case wafv2.ErrCodeWAFInternalErrorException:
				fmt.Println(wafv2.ErrCodeWAFInternalErrorException, aerr.Error())
			case wafv2.ErrCodeWAFNonexistentItemException:
				fmt.Println(wafv2.ErrCodeWAFNonexistentItemException, aerr.Error())
			case wafv2.ErrCodeWAFInvalidParameterException:
				fmt.Println(wafv2.ErrCodeWAFInvalidParameterException, aerr.Error())
			case wafv2.ErrCodeWAFUnavailableEntityException:
				fmt.Println(wafv2.ErrCodeWAFUnavailableEntityException, aerr.Error())
			case wafv2.ErrCodeWAFInvalidOperationException:
				fmt.Println(wafv2.ErrCodeWAFInvalidOperationException, aerr.Error())
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
		fmt.Println(getWebACLResult.String())
	}

	resultRuleNameList := []string{}
	for _, rule := range getWebACLResult.WebACL.Rules {
		resultRuleNameList = append(resultRuleNameList, *rule.Name)
	}

	assert.ElementsMatch(t, expectedRuleNameList, resultRuleNameList)
}

// ValidateResourceAssociatedToWAFV2WebACL validate a REGIONAL qualified resource ARN is associated to a WAFv2 Web ACL
func ValidateResourceAssociatedToWAFV2WebACL(t *testing.T, svc *wafv2.WAFV2, resourceARN string, webACLARN string, verboseOutput bool) {
	t.Helper()

	getWebACLForResourceResult, err := svc.GetWebACLForResource(
		&wafv2.GetWebACLForResourceInput{
			ResourceArn: aws.String(resourceARN),
		},
	)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case wafv2.ErrCodeWAFInternalErrorException:
				fmt.Println(wafv2.ErrCodeWAFInternalErrorException, aerr.Error())
			case wafv2.ErrCodeWAFNonexistentItemException:
				fmt.Println(wafv2.ErrCodeWAFNonexistentItemException, aerr.Error())
			case wafv2.ErrCodeWAFInvalidParameterException:
				fmt.Println(wafv2.ErrCodeWAFInvalidParameterException, aerr.Error())
			case wafv2.ErrCodeWAFUnavailableEntityException:
				fmt.Println(wafv2.ErrCodeWAFUnavailableEntityException, aerr.Error())
			case wafv2.ErrCodeWAFInvalidOperationException:
				fmt.Println(wafv2.ErrCodeWAFInvalidOperationException, aerr.Error())
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
		fmt.Println(fmt.Println(getWebACLForResourceResult.String()))
	}

	assert.Equal(t, webACLARN, *getWebACLForResourceResult.WebACL.ARN)
}
