package tests

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/stretchr/testify/assert"
)

// ValidateKmsKey get the KMS key
func ValidateKmsKey(t *testing.T, svc *kms.KMS, keyAlias string, accountID string, verboseOutput bool) {
	t.Helper()

	keyInput := &kms.DescribeKeyInput{
		KeyId: aws.String(keyAlias),
	}

	keyResult, err1 := svc.DescribeKey(keyInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
			case kms.ErrCodeInvalidArnException:
				fmt.Println(kms.ErrCodeInvalidArnException, aerr.Error())
			case kms.ErrCodeDependencyTimeoutException:
				fmt.Println(kms.ErrCodeDependencyTimeoutException, aerr.Error())
			case kms.ErrCodeInternalException:
				fmt.Println(kms.ErrCodeInternalException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err1.Error())
		}

		t.Logf("Failing test.")
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(keyResult.String())
	}

	assert.Equal(t, accountID, *keyResult.KeyMetadata.AWSAccountId)
	assert.Equal(t, "SYMMETRIC_DEFAULT", *keyResult.KeyMetadata.KeySpec)
	assert.Equal(t, true, *keyResult.KeyMetadata.Enabled)
	assert.Equal(t, "SYMMETRIC_DEFAULT", *keyResult.KeyMetadata.EncryptionAlgorithms[0])
	assert.Equal(t, "CUSTOMER", *keyResult.KeyMetadata.KeyManager)
	assert.Equal(t, "Enabled", *keyResult.KeyMetadata.KeyState)
	assert.Equal(t, "ENCRYPT_DECRYPT", *keyResult.KeyMetadata.KeyUsage)
	assert.Equal(t, "AWS_KMS", *keyResult.KeyMetadata.Origin)
}

// ValidateKmsKeyPolicy get the KMS key policy
func ValidateKmsKeyPolicy(t *testing.T, svc *kms.KMS, keyArn string, verboseOutput bool) {
	t.Helper()

	keyPolicyInput := &kms.GetKeyPolicyInput{
		KeyId:      aws.String(keyArn),
		PolicyName: aws.String("default"),
	}

	keyPolicyResult, err1 := svc.GetKeyPolicy(keyPolicyInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
			case kms.ErrCodeInvalidArnException:
				fmt.Println(kms.ErrCodeInvalidArnException, aerr.Error())
			case kms.ErrCodeDependencyTimeoutException:
				fmt.Println(kms.ErrCodeDependencyTimeoutException, aerr.Error())
			case kms.ErrCodeInternalException:
				fmt.Println(kms.ErrCodeInternalException, aerr.Error())
			case kms.ErrCodeInvalidStateException:
				fmt.Println(kms.ErrCodeInvalidStateException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err1.Error())
		}

		t.Logf("Failing test.")
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(keyPolicyResult.String())
	}

	assert.NotEmpty(t, *keyPolicyResult.Policy)
}

// ValidateKmsKeyTags gets tags and validates them
func ValidateKmsKeyTags(t *testing.T, svc *kms.KMS, keyArn string, tags []string, verboseOutput bool) {
	t.Helper()

	keyTagsInput := &kms.ListResourceTagsInput{
		KeyId: aws.String(keyArn),
	}

	keyTagsResult, err1 := svc.ListResourceTags(keyTagsInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
			case kms.ErrCodeInvalidArnException:
				fmt.Println(kms.ErrCodeInvalidArnException, aerr.Error())
			case kms.ErrCodeInvalidMarkerException:
				fmt.Println(kms.ErrCodeInvalidMarkerException, aerr.Error())
			case kms.ErrCodeInternalException:
				fmt.Println(kms.ErrCodeInternalException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err1.Error())
		}

		t.Logf("Failing test.")
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(keyTagsResult.String())
	}

	// validate tags
	for i := 0; i < len(tags); i++ {
		assert.Contains(t, keyTagsResult.String(), tags[i])

		if verboseOutput {
			fmt.Println(tags[i])
		}
	}
}

// ValidateKmsKeyRotationStatus get the KMS key rotation status
func ValidateKmsKeyRotationStatus(t *testing.T, svc *kms.KMS, keyArn string, keyRotationStatus bool, verboseOutput bool) {
	t.Helper()

	getKeyRotationStatusInput := &kms.GetKeyRotationStatusInput{
		KeyId: aws.String(keyArn),
	}

	keyRotationStatusResult, err1 := svc.GetKeyRotationStatus(getKeyRotationStatusInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
			case kms.ErrCodeInvalidArnException:
				fmt.Println(kms.ErrCodeInvalidArnException, aerr.Error())
			case kms.ErrCodeDependencyTimeoutException:
				fmt.Println(kms.ErrCodeDependencyTimeoutException, aerr.Error())
			case kms.ErrCodeInternalException:
				fmt.Println(kms.ErrCodeInternalException, aerr.Error())
			case kms.ErrCodeInvalidStateException:
				fmt.Println(kms.ErrCodeInvalidStateException, aerr.Error())
			case kms.ErrCodeUnsupportedOperationException:
				fmt.Println(kms.ErrCodeUnsupportedOperationException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err1.Error())
		}

		t.Logf("Failing test.")
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(keyRotationStatusResult.String())
	}

	assert.Equal(t, keyRotationStatus, *keyRotationStatusResult.KeyRotationEnabled)
}

// ValidateKmsGrant get the KMS key rotation status
func ValidateKmsGrant(t *testing.T, svc *kms.KMS, kmsKeyID string, terraformGrantID string, grantName string, granteePrincipal string, issuingAccount string, keyIDArn string, operations []string, verboseOutput bool) {
	t.Helper()

	input := &kms.ListGrantsInput{
		KeyId: aws.String(kmsKeyID),
	}

	result, err := svc.ListGrants(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
			case kms.ErrCodeDependencyTimeoutException:
				fmt.Println(kms.ErrCodeDependencyTimeoutException, aerr.Error())
			case kms.ErrCodeInvalidMarkerException:
				fmt.Println(kms.ErrCodeInvalidMarkerException, aerr.Error())
			case kms.ErrCodeInvalidArnException:
				fmt.Println(kms.ErrCodeInvalidArnException, aerr.Error())
			case kms.ErrCodeInternalException:
				fmt.Println(kms.ErrCodeInternalException, aerr.Error())
			case kms.ErrCodeInvalidStateException:
				fmt.Println(kms.ErrCodeInvalidStateException, aerr.Error())
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
		fmt.Println(result.String())
	}

	for i := 0; i < len(result.Grants); i++ {
		grantID := *result.Grants[i].GrantId
		if grantID == terraformGrantID {
			if verboseOutput {
				fmt.Println(result.String())
			}

			fmt.Println(result.Grants[i].String())
			assert.Equal(t, grantName, *result.Grants[i].Name)
			assert.Equal(t, granteePrincipal, *result.Grants[i].GranteePrincipal)
			assert.Equal(t, issuingAccount, *result.Grants[i].IssuingAccount)
			assert.Equal(t, keyIDArn, *result.Grants[i].KeyId)
			// validate operations
			for j := 0; j < len(operations); j++ {
				assert.Contains(t, result.Grants[i].String(), operations[j])

				if verboseOutput {
					fmt.Println(operations[j])
				}
			}
		}
	}
}
