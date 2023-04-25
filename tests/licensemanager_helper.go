package tests

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/licensemanager"
	"github.com/stretchr/testify/assert"
)

func ValidateLicenseManagerGrant(t *testing.T, svc *licensemanager.LicenseManager, grantName string, grantArn string, licenseArn string, grantStatus string, verboseOutput bool) {
	t.Helper()

	receivedGrantInput := &licensemanager.ListReceivedGrantsInput{
		Filters: []*licensemanager.Filter{
			{
				Name:   aws.String("LicenseArn"),
				Values: []*string{aws.String(licenseArn)},
			},
		},
		GrantArns: []*string{
			aws.String(grantArn),
		},
	}

	listReceivedGrantsResult, err := svc.ListReceivedGrants(receivedGrantInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case licensemanager.ErrCodeValidationException:
				fmt.Println(licensemanager.ErrCodeValidationException, aerr.Error())
			case licensemanager.ErrCodeInvalidParameterValueException:
				fmt.Println(licensemanager.ErrCodeInvalidParameterValueException, aerr.Error())
			case licensemanager.ErrCodeResourceLimitExceededException:
				fmt.Println(licensemanager.ErrCodeResourceLimitExceededException, aerr.Error())
			case licensemanager.ErrCodeServerInternalException:
				fmt.Println(licensemanager.ErrCodeServerInternalException, aerr.Error())
			case licensemanager.ErrCodeAuthorizationException:
				fmt.Println(licensemanager.ErrCodeAuthorizationException, aerr.Error())
			case licensemanager.ErrCodeAccessDeniedException:
				fmt.Println(licensemanager.ErrCodeAccessDeniedException, aerr.Error())
			case licensemanager.ErrCodeRateLimitExceededException:
				fmt.Println(licensemanager.ErrCodeRateLimitExceededException, aerr.Error())
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
		fmt.Println(listReceivedGrantsResult.String())
	}

	assert.Equal(t, grantName, *listReceivedGrantsResult.Grants[0].GrantName)
	assert.Equal(t, grantArn, *listReceivedGrantsResult.Grants[0].GrantArn)
	assert.Equal(t, licenseArn, *listReceivedGrantsResult.Grants[0].LicenseArn)
	assert.Equal(t, grantStatus, *listReceivedGrantsResult.Grants[0].GrantStatus)
}
