package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/stretchr/testify/assert"
)

// ValidateRoute53HostedZone Validate the Hosted Zone was created
func ValidateRoute53HostedZone(t *testing.T, svc *route53.Route53, hostedZoneID string, hostedZoneName string, privateZone bool, verboseOutput bool) {
	t.Helper()

	getHostedZoneInput := &route53.GetHostedZoneInput{
		Id: aws.String(hostedZoneID),
	}

	getHostedZoneResult, err := svc.GetHostedZone(getHostedZoneInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
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
		fmt.Println(getHostedZoneResult.String())
	}

	assert.Equal(t, hostedZoneName, *getHostedZoneResult.HostedZone.Name)
	assert.Equal(t, hostedZoneID, strings.Split(*getHostedZoneResult.HostedZone.Id, "/")[2])
	assert.Equal(t, privateZone, *getHostedZoneResult.HostedZone.Config.PrivateZone)
}

// ValidateRoute53ResolverRuleAssociation Validate a rule association exists
func ValidateRoute53ResolverRuleAssociation(t *testing.T, svc *route53resolver.Route53Resolver, vpcID string, ruleAssociationID string, verboseOutput bool) {
	t.Helper()

	getResolverRuleAssociationInput := &route53resolver.GetResolverRuleAssociationInput{
		ResolverRuleAssociationId: aws.String(ruleAssociationID),
	}

	getResolverRuleAssociationResult, err := svc.GetResolverRuleAssociation(getResolverRuleAssociationInput)
	if err != nil {
		fmt.Println(err.Error())
		t.Logf("Failing test.")
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(getResolverRuleAssociationResult.String())
	}

	assert.Equal(t, vpcID, *getResolverRuleAssociationResult.ResolverRuleAssociation.VPCId)
}
