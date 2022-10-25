package tests

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

// ValidateBucketLocation get bucket location
func ValidateBucketLocation(t *testing.T, svc *s3.S3, bucketName string, region string, verboseOutput bool) {
	t.Helper()

	getBucketLocationInput := &s3.GetBucketLocationInput{
		Bucket: aws.String(bucketName),
	}

	getBucketLocationResult, err1 := svc.GetBucketLocation(getBucketLocationInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
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
		fmt.Println(*getBucketLocationResult.LocationConstraint)
	}
}

// ValidateBucketPolicy get bucket policy
func ValidateBucketPolicy(t *testing.T, svc *s3.S3, bucketName string, policyJSON string, verboseOutput bool) {
	t.Helper()

	getBucketPolicyInput := &s3.GetBucketPolicyInput{
		Bucket: aws.String(bucketName),
	}

	getBucketPolicyResult, err1 := svc.GetBucketPolicy(getBucketPolicyInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
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
		fmt.Println(getBucketPolicyResult.String())
	}

	assert.JSONEq(t, policyJSON, *getBucketPolicyResult.Policy)
}

// ValidateBucketACL get bucket acl
func ValidateBucketACL(t *testing.T, svc *s3.S3, bucketName string, verboseOutput bool) {
	t.Helper()

	getBucketACLInput := &s3.GetBucketAclInput{
		Bucket: aws.String(bucketName),
	}

	getBucketACLResult, err1 := svc.GetBucketAcl(getBucketACLInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
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
		fmt.Println(getBucketACLResult.String())
	}

	assert.NotEmpty(t, getBucketACLResult.String())
}

// ValidateBucketEncryption get bucket encryption
func ValidateBucketEncryption(t *testing.T, svc *s3.S3, bucketName string, encryptionType string, verboseOutput bool) {
	t.Helper()

	getBucketEncryptionInput := &s3.GetBucketEncryptionInput{
		Bucket: aws.String(bucketName),
	}

	getBucketEncryptionResult, err1 := svc.GetBucketEncryption(getBucketEncryptionInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
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
		fmt.Println(getBucketEncryptionResult.String())
	}

	assert.Equal(t, encryptionType, *getBucketEncryptionResult.ServerSideEncryptionConfiguration.Rules[0].ApplyServerSideEncryptionByDefault.SSEAlgorithm)
}

// ValidateBucketLifecycleConfiguration get bucket LifecycleConfiguration
func ValidateBucketLifecycleConfiguration(t *testing.T, svc *s3.S3, bucketName string, ruleID string, expiration int64, status string, verboseOutput bool) {
	t.Helper()

	getBucketLifecycleConfigurationInput := &s3.GetBucketLifecycleConfigurationInput{
		Bucket: aws.String(bucketName),
	}

	getBucketLifecycleConfigurationResult, err1 := svc.GetBucketLifecycleConfiguration(getBucketLifecycleConfigurationInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
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
		fmt.Println(getBucketLifecycleConfigurationResult.String())
	}

	assert.Equal(t, ruleID, *getBucketLifecycleConfigurationResult.Rules[0].ID)
	assert.Equal(t, status, *getBucketLifecycleConfigurationResult.Rules[0].Status)
	assert.Equal(t, expiration, *getBucketLifecycleConfigurationResult.Rules[0].Expiration.Days)
	assert.Equal(t, expiration, *getBucketLifecycleConfigurationResult.Rules[0].NoncurrentVersionExpiration.NoncurrentDays)
}

// ValidateBucketReplication get bucket Replication
func ValidateBucketReplication(t *testing.T, svc *s3.S3, bucketName string, roleArn string, acl string, status string, destinationBucket string, storageClass string, idDestination string, destinationAccountID string, verboseOutput bool) {
	t.Helper()

	getBucketReplicationInput := &s3.GetBucketReplicationInput{
		Bucket: aws.String(bucketName),
	}

	getBucketReplicationResult, err1 := svc.GetBucketReplication(getBucketReplicationInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
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
		fmt.Println(getBucketReplicationResult.String())
	}

	assert.Equal(t, roleArn, *getBucketReplicationResult.ReplicationConfiguration.Role)
	assert.Equal(t, destinationAccountID, *getBucketReplicationResult.ReplicationConfiguration.Rules[0].Destination.Account)
	assert.Equal(t, acl, *getBucketReplicationResult.ReplicationConfiguration.Rules[0].Destination.AccessControlTranslation.Owner)
	assert.Equal(t, status, *getBucketReplicationResult.ReplicationConfiguration.Rules[0].Status)
	assert.Equal(t, destinationBucket, *getBucketReplicationResult.ReplicationConfiguration.Rules[0].Destination.Bucket)
	assert.Equal(t, storageClass, *getBucketReplicationResult.ReplicationConfiguration.Rules[0].Destination.StorageClass)
	assert.Equal(t, idDestination, *getBucketReplicationResult.ReplicationConfiguration.Rules[0].ID)
}

// ValidateBucketVersioning get bucket Versioning
func ValidateBucketVersioning(t *testing.T, svc *s3.S3, bucketName string, status string, verboseOutput bool) {
	t.Helper()

	getBucketVersioningInput := &s3.GetBucketVersioningInput{
		Bucket: aws.String(bucketName),
	}

	getBucketVersioningResult, err1 := svc.GetBucketVersioning(getBucketVersioningInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
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
		fmt.Println(getBucketVersioningResult.String())
	}

	assert.Equal(t, status, *getBucketVersioningResult.Status)
}

// ValidateBucketTagging get bucket Tagging
func ValidateBucketTagging(t *testing.T, svc *s3.S3, bucketName string, tagValues []string, verboseOutput bool) {
	t.Helper()

	getBucketTaggingInput := &s3.GetBucketTaggingInput{
		Bucket: aws.String(bucketName),
	}

	getBucketTaggingResult, err1 := svc.GetBucketTagging(getBucketTaggingInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
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
		fmt.Println(getBucketTaggingResult.String())
	}

	// validate tags
	for i := 0; i < len(tagValues); i++ {
		assert.Contains(t, getBucketTaggingResult.String(), tagValues[i])

		if verboseOutput {
			fmt.Println(tagValues[i])
		}
	}
}

// ValidatePublicAccessBlock get bucket PublicAccessBlock
func ValidatePublicAccessBlock(t *testing.T, svc *s3.S3, bucketName string, blockPublicAcls bool, blockPublicPolicy bool, ignorePublicAcls bool, restrictPublicBuckets bool, verboseOutput bool) {
	t.Helper()

	getPublicAccessBlockInput := &s3.GetPublicAccessBlockInput{
		Bucket: aws.String(bucketName),
	}

	getPublicAccessBlockResult, err1 := svc.GetPublicAccessBlock(getPublicAccessBlockInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case kms.ErrCodeNotFoundException:
				fmt.Println(kms.ErrCodeNotFoundException, aerr.Error())
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
		fmt.Println(getPublicAccessBlockResult.String())
	}

	assert.Equal(t, blockPublicAcls, *getPublicAccessBlockResult.PublicAccessBlockConfiguration.BlockPublicAcls)
	assert.Equal(t, blockPublicPolicy, *getPublicAccessBlockResult.PublicAccessBlockConfiguration.BlockPublicAcls)
	assert.Equal(t, ignorePublicAcls, *getPublicAccessBlockResult.PublicAccessBlockConfiguration.BlockPublicAcls)
	assert.Equal(t, restrictPublicBuckets, *getPublicAccessBlockResult.PublicAccessBlockConfiguration.BlockPublicAcls)
}
