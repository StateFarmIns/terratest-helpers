package tests

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/stretchr/testify/assert"
)

// ValidateCreateAccountSCP validate create account scp module
func ValidateCreateAccountSCP(t *testing.T, svc *organizations.Organizations, policyName string, policyID string, verboseOutput bool) {
	t.Helper()

	describePolicyInput := &organizations.DescribePolicyInput{
		PolicyId: aws.String(policyID),
	}

	describePolicyResult, err := svc.DescribePolicy(describePolicyInput)
	if err != nil {
		fmt.Println(err.Error())
		t.Logf("Failing test.")
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(describePolicyResult.String())
	}

	assert.Equal(t, policyName, *describePolicyResult.Policy.PolicySummary.Name)
}
