package tests

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/stretchr/testify/assert"
)

func ValidateGlueCrawlerExists(t *testing.T, svc *glue.Glue, crawlerName string, schedule string, testSchedule bool, verboseOutput bool) {
	t.Helper()

	getCrawlerResult, err := svc.GetCrawler(
		&glue.GetCrawlerInput{
			Name: aws.String(crawlerName),
		},
	)

	fmt.Println("Validating Crawler: ", crawlerName)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glue.ErrCodeEntityNotFoundException:
				fmt.Println(glue.ErrCodeEntityNotFoundException, aerr.Error())
			case glue.ErrCodeOperationTimeoutException:
				fmt.Println(glue.ErrCodeOperationTimeoutException, aerr.Error())
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
		fmt.Println(fmt.Println(getCrawlerResult.String()))
	}

	assert.Equal(t, crawlerName, *getCrawlerResult.Crawler.Name)

	if testSchedule {
		assert.Equal(t, schedule, *getCrawlerResult.Crawler.Schedule.ScheduleExpression)
	}
}

func ValidateGlueJobExists(t *testing.T, svc *glue.Glue, jobName string, verboseOutput bool) {
	t.Helper()

	getJobResult, err := svc.GetJob(
		&glue.GetJobInput{
			JobName: aws.String(jobName),
		},
	)

	fmt.Println("Validating Job: ", jobName)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glue.ErrCodeInvalidInputException:
				fmt.Println(glue.ErrCodeInvalidInputException, aerr.Error())
			case glue.ErrCodeEntityNotFoundException:
				fmt.Println(glue.ErrCodeEntityNotFoundException, aerr.Error())
			case glue.ErrCodeInternalServiceException:
				fmt.Println(glue.ErrCodeInternalServiceException, aerr.Error())
			case glue.ErrCodeOperationTimeoutException:
				fmt.Println(glue.ErrCodeOperationTimeoutException, aerr.Error())
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
		fmt.Println(fmt.Println(getJobResult.String()))
	}

	assert.Equal(t, jobName, *getJobResult.Job.Name)
}

func ValidateGlueConnectionExists(t *testing.T, svc *glue.Glue, connectionName string, hidePassword bool, verboseOutput bool) {
	t.Helper()

	getConnectionResult, err := svc.GetConnection(
		&glue.GetConnectionInput{
			HidePassword: aws.Bool(hidePassword),
			Name:         aws.String(connectionName),
		},
	)

	fmt.Println("Validating Connection: ", connectionName)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glue.ErrCodeInvalidInputException:
				fmt.Println(glue.ErrCodeInvalidInputException, aerr.Error())
			case glue.ErrCodeEntityNotFoundException:
				fmt.Println(glue.ErrCodeEntityNotFoundException, aerr.Error())
			case glue.ErrCodeEncryptionException:
				fmt.Println(glue.ErrCodeEncryptionException, aerr.Error())
			case glue.ErrCodeOperationTimeoutException:
				fmt.Println(glue.ErrCodeOperationTimeoutException, aerr.Error())
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
		fmt.Println(fmt.Println(getConnectionResult.String()))
	}

	assert.Equal(t, connectionName, *getConnectionResult.Connection.Name)
}

func ValidateGlueJobTriggerExists(t *testing.T, svc *glue.Glue, triggerName string, verboseOutput bool) {
	t.Helper()

	getTriggerResult, err := svc.GetTrigger(
		&glue.GetTriggerInput{
			Name: aws.String(triggerName),
		},
	)

	fmt.Println("Validating Trigger: ", triggerName)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case glue.ErrCodeInvalidInputException:
				fmt.Println(glue.ErrCodeInvalidInputException, aerr.Error())
			case glue.ErrCodeEntityNotFoundException:
				fmt.Println(glue.ErrCodeEntityNotFoundException, aerr.Error())
			case glue.ErrCodeInternalServiceException:
				fmt.Println(glue.ErrCodeInternalServiceException, aerr.Error())
			case glue.ErrCodeOperationTimeoutException:
				fmt.Println(glue.ErrCodeOperationTimeoutException, aerr.Error())
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
		fmt.Println(fmt.Println(getTriggerResult.String()))
	}

	assert.Equal(t, triggerName, *getTriggerResult.Trigger.Name)
}
