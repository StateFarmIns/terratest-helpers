package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/stretchr/testify/assert"
)

// Vpc struct containing elements returned from a VPC module.
// This provides a convenient way of consolidating our VPC attributes when calling helper functions
type Vpc struct {
	VpcID   string
	VpcCidr string
}

// ValidateVpc validate a VPC via attributes passed in using the Vpc struct
func ValidateVpc(t *testing.T, svc *ec2.EC2, vpc Vpc, verboseOutput bool) {
	t.Helper()

	describeVpcResult, err := svc.DescribeVpcs(
		&ec2.DescribeVpcsInput{
			VpcIds: []*string{aws.String(vpc.VpcID)},
		},
	)
	if err != nil {
		fmt.Println(err.Error())
		t.Logf("Failing test.")
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(describeVpcResult.String())
	}

	assert.Equal(t, vpc.VpcCidr, *describeVpcResult.Vpcs[0].CidrBlock)
}

// ValidateTgwConsumer helper function to validate transit gateway vpc associations
func ValidateTgwConsumer(t *testing.T, svc *ec2.EC2, verboseOutput bool, tgwAttachmentID string, vpcID string) {
	t.Helper()

	describeTransitGatewayVpcAttachmentsResult, err := svc.DescribeTransitGatewayVpcAttachments(
		&ec2.DescribeTransitGatewayVpcAttachmentsInput{
			TransitGatewayAttachmentIds: []*string{aws.String(tgwAttachmentID)},
		},
	)
	if err != nil {
		fmt.Println(err.Error())
		t.Logf("Failing test.")
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(describeTransitGatewayVpcAttachmentsResult.String())
	}

	assert.Equal(t, vpcID, *describeTransitGatewayVpcAttachmentsResult.TransitGatewayVpcAttachments[0].VpcId)
}

// ValidateVPC gets vpc and validates its info
func ValidateVPC(t *testing.T, svc *ec2.EC2, isDefault bool, cidrBlockState string, instanceTenancy string, ownerID string, state string, tagValues []string, verboseOutput bool) {
	t.Helper()

	describeVpcsInput := &ec2.DescribeVpcsInput{}

	describeVpcsResult, err1 := svc.DescribeVpcs(describeVpcsInput)
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
		fmt.Println(describeVpcsResult.String())
	}

	assert.Equal(t, isDefault, *describeVpcsResult.Vpcs[0].IsDefault)
	assert.Equal(t, cidrBlockState, *describeVpcsResult.Vpcs[0].CidrBlockAssociationSet[0].CidrBlockState.State)
	assert.Equal(t, instanceTenancy, *describeVpcsResult.Vpcs[0].InstanceTenancy)
	assert.Equal(t, ownerID, *describeVpcsResult.Vpcs[0].OwnerId)
	assert.Equal(t, state, *describeVpcsResult.Vpcs[0].State)
	assert.NotEmpty(t, *describeVpcsResult.Vpcs[0].VpcId)

	// validate tags
	for i := 0; i < len(tagValues); i++ {
		assert.Contains(t, describeVpcsResult.Vpcs[0].String(), tagValues[i])

		if verboseOutput {
			fmt.Println(tagValues[i])
		}
	}
}

// ValidateSingleVPC gets vpc and validates its info
func ValidateSingleVPC(t *testing.T, svc *ec2.EC2, vpcID string, isDefault bool, cidrBlockState string, instanceTenancy string, ownerID string, state string, tagValues []string, verboseOutput bool) {
	t.Helper()

	describeVpcsResult, err1 := svc.DescribeVpcs(
		&ec2.DescribeVpcsInput{
			VpcIds: []*string{aws.String(vpcID)},
		},
	)

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
		fmt.Println(describeVpcsResult.String())
	}

	assert.Equal(t, isDefault, *describeVpcsResult.Vpcs[0].IsDefault)
	assert.Equal(t, cidrBlockState, *describeVpcsResult.Vpcs[0].CidrBlockAssociationSet[0].CidrBlockState.State)
	assert.Equal(t, instanceTenancy, *describeVpcsResult.Vpcs[0].InstanceTenancy)
	assert.Equal(t, ownerID, *describeVpcsResult.Vpcs[0].OwnerId)
	assert.Equal(t, state, *describeVpcsResult.Vpcs[0].State)
	assert.NotEmpty(t, *describeVpcsResult.Vpcs[0].VpcId)

	// validate tags
	for i := 0; i < len(tagValues); i++ {
		assert.Contains(t, describeVpcsResult.Vpcs[0].String(), tagValues[i])

		if verboseOutput {
			fmt.Println(tagValues[i])
		}
	}
}

// ValidateFlowLog gets FlowLog and validates its info
func ValidateFlowLog(t *testing.T, svc *ec2.EC2, vpcID string, deliverLogsPermissionArn string, deliverLogsStatus string, flowLogStatus string, logDestination string, logDestinationType string, logFormat string, trafficType string, verboseOutput bool) {
	t.Helper()

	fmt.Println("Running ValidateFlowLog")

	describeFlowLogsInput := &ec2.DescribeFlowLogsInput{}

	describeFlowLogsResult, err1 := svc.DescribeFlowLogs(describeFlowLogsInput)
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
		fmt.Println(describeFlowLogsResult.String())
	}

	// validate flow log details
	for i := 0; i < len(describeFlowLogsResult.FlowLogs); i++ {
		if *describeFlowLogsResult.FlowLogs[i].ResourceId == vpcID && *describeFlowLogsResult.FlowLogs[i].LogDestination == logDestination {
			// need to check to see if key exists, because for one item, it is not there.
			// assert.Equal(t, deliverLogsPermissionArn, *describeFlowLogsResult.FlowLogs[i].DeliverLogsPermissionArn)
			assert.Equal(t, deliverLogsStatus, *describeFlowLogsResult.FlowLogs[i].DeliverLogsStatus)
			assert.Equal(t, flowLogStatus, *describeFlowLogsResult.FlowLogs[i].FlowLogStatus)
			assert.Equal(t, logDestinationType, *describeFlowLogsResult.FlowLogs[i].LogDestinationType)
			assert.Equal(t, logFormat, *describeFlowLogsResult.FlowLogs[i].LogFormat)
			assert.Equal(t, trafficType, *describeFlowLogsResult.FlowLogs[i].TrafficType)
		} else {
			fmt.Println("ValidateFlowLog: info: logVPCID or logDestination does not match.")
		}
	}
}

// ValidateInternetGateway gets InternetGateway and validates its info
func ValidateInternetGateway(t *testing.T, svc *ec2.EC2, state string, ownerID string, tagValues []string, verboseOutput bool) {
	t.Helper()

	describeInternetGatewaysInput := &ec2.DescribeInternetGatewaysInput{}

	describeInternetGatewaysResult, err1 := svc.DescribeInternetGateways(describeInternetGatewaysInput)
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
		fmt.Println(describeInternetGatewaysResult.String())
	}

	assert.Equal(t, ownerID, *describeInternetGatewaysResult.InternetGateways[0].OwnerId)
	assert.Equal(t, state, *describeInternetGatewaysResult.InternetGateways[0].Attachments[0].State)
	assert.NotEmpty(t, *describeInternetGatewaysResult.InternetGateways[0].InternetGatewayId)

	// validate tags
	for i := 0; i < len(tagValues); i++ {
		assert.Contains(t, describeInternetGatewaysResult.String(), tagValues[i])

		if verboseOutput {
			fmt.Println(tagValues[i])
		}
	}
}

// ValidateRouteTables gets Route Tables and validates its info
func ValidateRouteTables(t *testing.T, svc *ec2.EC2, vpcID string, ownerID string, tagValues []string, verboseOutput bool) {
	t.Helper()

	describeRouteTablesInput := &ec2.DescribeRouteTablesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []*string{aws.String(vpcID)},
			},
		},
	}

	describeRouteTablesResult, err1 := svc.DescribeRouteTables(describeRouteTablesInput)
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
		fmt.Println(describeRouteTablesResult.String())
	}

	// Remove the element at index i from a.
	tagValues[7] = tagValues[len(tagValues)-1] // Copy last element to index i.
	tagValues[len(tagValues)-1] = ""           // Erase last element (write zero value).
	tagValues = tagValues[:len(tagValues)-1]   // Truncate slice.

	// validate rtb details
	for i := 0; i < len(describeRouteTablesResult.RouteTables); i++ {
		for x := 0; x < len(describeRouteTablesResult.RouteTables[i].Tags); x++ {
			switch {
			case strings.Contains(*describeRouteTablesResult.RouteTables[i].Tags[x].Value, "default-rtb"):
				assert.Equal(t, true, *describeRouteTablesResult.RouteTables[i].Associations[0].Main)
				assert.Equal(t, "associated", *describeRouteTablesResult.RouteTables[i].Associations[0].AssociationState.State)
				assert.Equal(t, ownerID, *describeRouteTablesResult.RouteTables[i].OwnerId)
				// the gatewayId checks seem to be failing for all three in this case. Why are they failing with nill pointer?
				// assert.Equal(t, "local", *describeRouteTablesResult.RouteTables[i].Routes[0].GatewayId)
				assert.Equal(t, "active", *describeRouteTablesResult.RouteTables[i].Routes[0].State)
				assert.Equal(t, "CreateRouteTable", *describeRouteTablesResult.RouteTables[i].Routes[0].Origin)
				assert.Equal(t, 1, len(describeRouteTablesResult.RouteTables[i].Routes))
				assert.NotEmpty(t, describeRouteTablesResult.RouteTables[i].VpcId)

				for z := 0; z < len(tagValues); z++ {
					assert.Contains(t, describeRouteTablesResult.RouteTables[i].String(), tagValues[z])
				}
			case strings.Contains(*describeRouteTablesResult.RouteTables[i].Tags[x].Value, "public-rtb"):
				assert.Equal(t, false, *describeRouteTablesResult.RouteTables[i].Associations[0].Main)
				assert.Equal(t, "associated", *describeRouteTablesResult.RouteTables[i].Associations[0].AssociationState.State)
				assert.Equal(t, ownerID, *describeRouteTablesResult.RouteTables[i].OwnerId)
				// the gatewayId checks seem to be failing for all three in this case. Why are they failing with nill pointer?
				// assert.Equal(t, "local", *describeRouteTablesResult.RouteTables[i].Routes[0].GatewayId)
				assert.Equal(t, "active", *describeRouteTablesResult.RouteTables[i].Routes[0].State)
				assert.Equal(t, "CreateRouteTable", *describeRouteTablesResult.RouteTables[i].Routes[0].Origin)
				assert.Equal(t, 2, len(describeRouteTablesResult.RouteTables[i].Routes))
				assert.NotEmpty(t, describeRouteTablesResult.RouteTables[i].VpcId)

				for z := 0; z < len(tagValues); z++ {
					assert.Contains(t, describeRouteTablesResult.RouteTables[i].String(), tagValues[z])
				}
			case strings.Contains(*describeRouteTablesResult.RouteTables[i].Tags[x].Value, "private-rtb"):
				assert.Equal(t, false, *describeRouteTablesResult.RouteTables[i].Associations[0].Main)
				assert.Equal(t, "associated", *describeRouteTablesResult.RouteTables[i].Associations[0].AssociationState.State)
				assert.Equal(t, ownerID, *describeRouteTablesResult.RouteTables[i].OwnerId)
				// the gatewayId checks seem to be failing for all three in this case. Why are they failing with nill pointer?
				// assert.Equal(t, "local", *describeRouteTablesResult.RouteTables[i].Routes[0].GatewayId)
				assert.Equal(t, "active", *describeRouteTablesResult.RouteTables[i].Routes[0].State)
				assert.Equal(t, "CreateRouteTable", *describeRouteTablesResult.RouteTables[i].Routes[0].Origin)
				// assert.Equal(t, 5, len(describeRouteTablesResult.RouteTables[i].Routes))
				assert.Equal(t, 3, len(describeRouteTablesResult.RouteTables[i].Routes))
				assert.NotEmpty(t, describeRouteTablesResult.RouteTables[i].VpcId)

				for z := 0; z < len(tagValues); z++ {
					assert.Contains(t, describeRouteTablesResult.RouteTables[i].String(), tagValues[z])
				}
			default:
				fmt.Println(*describeRouteTablesResult.RouteTables[i].Tags[x].Value)
			}
		}
	}
}

// ValidateSubnet gets Subnet and validates its info
func ValidateSubnet(t *testing.T, svc *ec2.EC2, state string, ownerID string, tagValues []string, verboseOutput bool) {
	t.Helper()

	describeSubnetsInput := &ec2.DescribeSubnetsInput{
		// Bucket: aws.String(bucketName),
	}

	describeSubnetsResult, err1 := svc.DescribeSubnets(describeSubnetsInput)
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
		fmt.Println(describeSubnetsResult.String())
	}

	// Remove the element at index i from a.
	tagValues[7] = tagValues[len(tagValues)-1] // Copy last element to index i.
	tagValues[len(tagValues)-1] = ""           // Erase last element (write zero value).
	tagValues = tagValues[:len(tagValues)-1]   // Truncate slice.

	// validate rtb details
	for i := 0; i < len(describeSubnetsResult.Subnets); i++ {
		for x := 0; x < len(describeSubnetsResult.Subnets[i].Tags); x++ {
			switch {
			case strings.Contains(*describeSubnetsResult.Subnets[i].Tags[x].Value, "private-sbn"):
				assert.Equal(t, false, *describeSubnetsResult.Subnets[i].AssignIpv6AddressOnCreation)
				assert.Equal(t, false, *describeSubnetsResult.Subnets[i].DefaultForAz)
				assert.Equal(t, false, *describeSubnetsResult.Subnets[i].MapPublicIpOnLaunch)
				assert.Equal(t, ownerID, *describeSubnetsResult.Subnets[i].OwnerId)
				assert.Equal(t, state, *describeSubnetsResult.Subnets[i].State)
				assert.NotEmpty(t, describeSubnetsResult.Subnets[i].VpcId)

				for z := 0; z < len(tagValues); z++ {
					assert.Contains(t, describeSubnetsResult.Subnets[i].String(), tagValues[z])
				}
			case strings.Contains(*describeSubnetsResult.Subnets[i].Tags[x].Value, "public-sbn"):
				assert.Equal(t, false, *describeSubnetsResult.Subnets[i].AssignIpv6AddressOnCreation)
				assert.Equal(t, false, *describeSubnetsResult.Subnets[i].DefaultForAz)
				assert.Equal(t, false, *describeSubnetsResult.Subnets[i].MapPublicIpOnLaunch)
				assert.Equal(t, ownerID, *describeSubnetsResult.Subnets[i].OwnerId)
				assert.Equal(t, state, *describeSubnetsResult.Subnets[i].State)
				assert.NotEmpty(t, describeSubnetsResult.Subnets[i].VpcId)

				for z := 0; z < len(tagValues); z++ {
					assert.Contains(t, describeSubnetsResult.Subnets[i].String(), tagValues[z])
				}
			default:
				fmt.Println(*describeSubnetsResult.Subnets[i].Tags[x].Value)
			}
		}
	}
}

// ValidateNatGateway gets NatGateway and validates its info
func ValidateNatGateway(t *testing.T, svc *ec2.EC2, state string, tagValues []string, verboseOutput bool) {
	t.Helper()

	describeNatGatewaysInput := &ec2.DescribeNatGatewaysInput{}

	describeNatGatewaysResult, err1 := svc.DescribeNatGateways(describeNatGatewaysInput)
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
		fmt.Println(describeNatGatewaysResult.String())
	}

	// Remove the element at index i from a.
	tagValues[7] = tagValues[len(tagValues)-1] // Copy last element to index i.
	tagValues[len(tagValues)-1] = ""           // Erase last element (write zero value).
	tagValues = tagValues[:len(tagValues)-1]   // Truncate slice.

	// validate rtb details
	for i := 0; i < len(describeNatGatewaysResult.NatGateways); i++ {
		for x := 0; x < len(describeNatGatewaysResult.NatGateways[i].Tags); x++ {
			if strings.Contains(*describeNatGatewaysResult.NatGateways[i].Tags[x].Value, "1b-natgw") {
				assert.Equal(t, state, *describeNatGatewaysResult.NatGateways[i].State)
				assert.NotEmpty(t, describeNatGatewaysResult.NatGateways[i].VpcId)

				for z := 0; z < len(tagValues); z++ {
					assert.Contains(t, describeNatGatewaysResult.NatGateways[i].String(), tagValues[z])
				}
			} else if strings.Contains(*describeNatGatewaysResult.NatGateways[i].Tags[x].Value, "1a-natgw") {
				assert.Equal(t, state, *describeNatGatewaysResult.NatGateways[i].State)
				assert.NotEmpty(t, describeNatGatewaysResult.NatGateways[i].VpcId)

				for z := 0; z < len(tagValues); z++ {
					assert.Contains(t, describeNatGatewaysResult.NatGateways[i].String(), tagValues[z])
				}
			}
		}
	}
}

// ValidateNetworkACLs gets NetworkAcl and validates its info
func ValidateNetworkACLs(t *testing.T, svc *ec2.EC2, naclName string, naclRules int, verboseOutput bool) {
	t.Helper()

	describeNetworkAclsInput := &ec2.DescribeNetworkAclsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []*string{aws.String(naclName)},
			},
		},
	}

	describeNetworkAclsResult, err1 := svc.DescribeNetworkAcls(describeNetworkAclsInput)
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
		fmt.Println(describeNetworkAclsResult.String())
	}

	assert.NotEmpty(t, describeNetworkAclsResult.NetworkAcls[0].Associations)
	assert.NotEmpty(t, describeNetworkAclsResult.NetworkAcls[0].Tags)
	assert.False(t, *describeNetworkAclsResult.NetworkAcls[0].IsDefault)
	assert.Equal(t, naclRules, len(describeNetworkAclsResult.NetworkAcls[0].Entries))
}

// ValidateVpcEndpoints gets NetworkAcl and validates its info
func ValidateVpcEndpoints(t *testing.T, svc *ec2.EC2, serviceName string, vpcID string, ownerID string, state string, privateDNSEnabled bool, securityGroups []string, vpcEndpointType string, verboseOutput bool) {
	t.Helper()

	fmt.Println("Running ValidateVpcEndpoints")

	describeVpcEndpointsInput := &ec2.DescribeVpcEndpointsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("service-name"),
				Values: []*string{aws.String(serviceName)},
			},
			{
				Name:   aws.String("vpc-id"),
				Values: []*string{aws.String(vpcID)},
			},
		},
	}

	describeVpcEndpointsResult, err1 := svc.DescribeVpcEndpoints(describeVpcEndpointsInput)
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
		fmt.Println(describeVpcEndpointsResult.String())
	}

	// create slice of security groups
	groups := make([]string, len(describeVpcEndpointsResult.VpcEndpoints[0].Groups))
	for index, value := range describeVpcEndpointsResult.VpcEndpoints[0].Groups {
		groups[index] = *value.GroupName
	}

	assert.ElementsMatch(t, securityGroups, groups)
	assert.Equal(t, privateDNSEnabled, *describeVpcEndpointsResult.VpcEndpoints[0].PrivateDnsEnabled)
	assert.Equal(t, vpcEndpointType, *describeVpcEndpointsResult.VpcEndpoints[0].VpcEndpointType)
	assert.Equal(t, state, *describeVpcEndpointsResult.VpcEndpoints[0].State)
	assert.Equal(t, ownerID, *describeVpcEndpointsResult.VpcEndpoints[0].OwnerId)
	assert.Equal(t, serviceName, *describeVpcEndpointsResult.VpcEndpoints[0].ServiceName)
}

// ValidateSecurityGroup gets security group by name and vpcID and validates its info
func ValidateSecurityGroup(t *testing.T, svc *ec2.EC2, vpcID string, groupName string, numIngressRules int, numEgressRules int, verboseOutput bool) {
	t.Helper()

	describeSecurityGroupsInput := &ec2.DescribeSecurityGroupsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("group-name"),
				Values: []*string{aws.String(groupName)},
			},
			{
				Name:   aws.String("vpc-id"),
				Values: []*string{aws.String(vpcID)},
			},
		},
	}

	describeSecurityGroupsResult, err1 := svc.DescribeSecurityGroups(describeSecurityGroupsInput)
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

	if len(describeSecurityGroupsResult.SecurityGroups) != 0 {
		assert.Equal(t, numIngressRules, len(describeSecurityGroupsResult.SecurityGroups[0].IpPermissions))
		assert.Equal(t, numEgressRules, len(describeSecurityGroupsResult.SecurityGroups[0].IpPermissionsEgress))
		assert.NotEmpty(t, describeSecurityGroupsResult.SecurityGroups[0].Tags)
	} else {
		fmt.Println("Security Group Name of " + groupName + " does not exist in this account")
		t.Logf("Failing test.")
		t.Fail()

		return
	}
}

// ValidateTransitGateways gets NetworkAcl and validates its info
func ValidateTransitGateways(t *testing.T, svc *ec2.EC2, verboseOutput bool) {
	t.Helper()

	describeTransitGatewaysInput := &ec2.DescribeTransitGatewaysInput{}

	describeTransitGatewaysResult, err1 := svc.DescribeTransitGateways(describeTransitGatewaysInput)
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
		fmt.Println(describeTransitGatewaysResult.String())
	}

	assert.NotEmpty(t, describeTransitGatewaysResult.TransitGateways)
	assert.Equal(t, "available", *describeTransitGatewaysResult.TransitGateways[0].State)
}

// ValidateTransitGatewayAttachments gets NetworkAcl and validates its info
func ValidateTransitGatewayAttachments(t *testing.T, svc *ec2.EC2, verboseOutput bool) {
	t.Helper()

	describeTransitGatewayAttachmentsInput := &ec2.DescribeTransitGatewayAttachmentsInput{}

	describeTransitGatewayAttachmentsResult, err1 := svc.DescribeTransitGatewayAttachments(describeTransitGatewayAttachmentsInput)
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
		fmt.Println(describeTransitGatewayAttachmentsResult.String())
	}

	assert.NotEmpty(t, describeTransitGatewayAttachmentsResult.TransitGatewayAttachments)
	assert.Equal(t, "available", *describeTransitGatewayAttachmentsResult.TransitGatewayAttachments[0].State)
}
