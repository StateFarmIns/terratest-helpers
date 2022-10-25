package terratests

import (
	"testing"

	// "github.com/StateFarmIns/terratest-helpers/tests"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

const terraformDir string = "test/"
const region string = "us-east-1"

func TestS3TerratestExample(t *testing.T) {

	terraformOptions := &terraform.Options{
		TerraformDir: terraformDir,
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	bucketName := terraform.Output(t, terraformOptions, "s3_bucket_name")

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	s3svc := s3.New(sess)

	tests.ValidateBucketExample(t, s3svc, bucketName, true)
}
