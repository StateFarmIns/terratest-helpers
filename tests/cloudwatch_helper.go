package tests

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/stretchr/testify/assert"
)

// ValidateCloudWatchLogGroupName validate a Cloud Watch Log Group by name
func ValidateCloudWatchLogGroupName(t *testing.T, svc *cloudwatchlogs.CloudWatchLogs, groupName string, verboseOutput bool) {
	t.Helper()

	describeLogGroupsResult, err := svc.DescribeLogGroups(
		&cloudwatchlogs.DescribeLogGroupsInput{
			LogGroupNamePrefix: aws.String(groupName),
		},
	)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudwatchlogs.ErrCodeInvalidParameterException:
				fmt.Println(cloudwatchlogs.ErrCodeInvalidParameterException, aerr.Error())
			case cloudwatchlogs.ErrCodeServiceUnavailableException:
				fmt.Println(cloudwatchlogs.ErrCodeServiceUnavailableException, aerr.Error())
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
		fmt.Println(fmt.Println(describeLogGroupsResult.String()))
	}

	assert.Equal(t, groupName, *describeLogGroupsResult.LogGroups[0].LogGroupName)
}

// ValidateCloudWatchLogGroupsByPrefix validate a list of Cloud Watch Log Groups by summarized prefix
func ValidateCloudWatchLogGroupsByPrefix(t *testing.T, svc *cloudwatchlogs.CloudWatchLogs, groupPrefix string, expectedGroupNameList []string, verboseOutput bool) {
	t.Helper()

	describeLogGroupsResult, err := svc.DescribeLogGroups(
		&cloudwatchlogs.DescribeLogGroupsInput{
			LogGroupNamePrefix: aws.String(groupPrefix),
		},
	)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudwatchlogs.ErrCodeInvalidParameterException:
				fmt.Println(cloudwatchlogs.ErrCodeInvalidParameterException, aerr.Error())
			case cloudwatchlogs.ErrCodeServiceUnavailableException:
				fmt.Println(cloudwatchlogs.ErrCodeServiceUnavailableException, aerr.Error())
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
		fmt.Println(fmt.Println(describeLogGroupsResult.String()))
	}

	resultGroupNameList := []string{}
	for _, logGroup := range describeLogGroupsResult.LogGroups {
		resultGroupNameList = append(resultGroupNameList, *logGroup.LogGroupName)
	}

	assert.ElementsMatch(t, expectedGroupNameList, resultGroupNameList)
}

// ValidateCloudWatchEventRule gets the event rule and validates its details
func ValidateCloudWatchEventRule(t *testing.T, svc *cloudwatchevents.CloudWatchEvents, ruleName string, ruleArn string, ruleEventPatternJSON string, ruleState string, verboseOutput bool) {
	t.Helper()

	describeRuleInput := &cloudwatchevents.DescribeRuleInput{
		Name: aws.String(ruleName),
	}

	describeRuleResult, err := svc.DescribeRule(describeRuleInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudwatchevents.InternalException{}.String():
				fmt.Println(cloudwatchevents.InternalException{}.String(), aerr.Error())
			case cloudwatchevents.ResourceNotFoundException{}.String():
				fmt.Println(cloudwatchevents.ResourceNotFoundException{}.String(), aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(describeRuleResult)
	}

	assert.Equal(t, ruleName, *describeRuleResult.Name)
	assert.Equal(t, ruleArn, *describeRuleResult.Arn)
	assert.JSONEq(t, ruleEventPatternJSON, *describeRuleResult.EventPattern)
	assert.Equal(t, ruleState, *describeRuleResult.State)
}

// ValidateCloudWatchEventRuleTarget get the event rule target and validates its details
func ValidateCloudWatchEventRuleTarget(t *testing.T, svc *cloudwatchevents.CloudWatchEvents, ruleName string, roleArn string, eventBusArn string, verboseOutput bool) {
	t.Helper()

	listTargetsByRuleInput := &cloudwatchevents.ListTargetsByRuleInput{
		Rule: aws.String(ruleName),
	}

	listTargetsByRuleResult, err := svc.ListTargetsByRule(listTargetsByRuleInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case cloudwatchevents.InternalException{}.String():
				fmt.Println(cloudwatchevents.InternalException{}.String(), aerr.Error())
			case cloudwatchevents.ResourceNotFoundException{}.String():
				fmt.Println(cloudwatchevents.ResourceNotFoundException{}.String(), aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(listTargetsByRuleResult)
	}

	assert.Equal(t, roleArn, *listTargetsByRuleResult.Targets[0].RoleArn)
	assert.Equal(t, eventBusArn, *listTargetsByRuleResult.Targets[0].Arn)
}
