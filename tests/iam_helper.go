package tests

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/stretchr/testify/assert"
)

// ValidatePolicy gets Polcy by arn and validates its data
func ValidatePolicy(t *testing.T, svc *iam.IAM, policyArn string, policyName string, verboseOutput bool) {
	t.Helper()

	policyInput := &iam.GetPolicyInput{
		PolicyArn: aws.String(policyArn),
	}

	policyResult, err1 := svc.GetPolicy(policyInput)
	if err1 != nil {
		if aerr, ok := err1.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(policyResult.String())
	}

	assert.Equal(t, policyName, *policyResult.Policy.PolicyName)
}

// ValidateUserDetails get user details
func ValidateUserDetails(t *testing.T, svc *iam.IAM, userName string, userArn string, verboseOutput bool) {
	t.Helper()

	input := &iam.GetUserInput{
		UserName: aws.String(userName),
	}

	userResult, err := svc.GetUser(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(userResult.String())
	}

	assert.Equal(t, userName, *userResult.User.UserName)
	assert.Equal(t, userArn, *userResult.User.Arn)
}

// ValidateUserDetailsWTags get user details
func ValidateUserDetailsWTags(t *testing.T, svc *iam.IAM, userName string, userArn string, tags []string, verboseOutput bool) {
	t.Helper()

	input := &iam.GetUserInput{
		UserName: aws.String(userName),
	}

	userResult, err := svc.GetUser(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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

	// validate tags
	for i := 0; i < len(tags); i++ {
		assert.Contains(t, userResult.String(), tags[i])

		if verboseOutput {
			fmt.Println(tags[i])
		}
	}

	if verboseOutput {
		fmt.Println(userResult.String())
	}

	assert.Equal(t, userName, *userResult.User.UserName)
	assert.Equal(t, userArn, *userResult.User.Arn)
}

// ValidateRoleArn Validate the ARN of an IAM role by querying the Role Name
func ValidateRoleArn(t *testing.T, svc *iam.IAM, roleName string, roleArn string, verboseOutput bool) {
	t.Helper()

	getRoleResult, err := svc.GetRole(
		&iam.GetRoleInput{
			RoleName: aws.String(roleName),
		},
	)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(getRoleResult.String())
	}

	assert.Equal(t, roleArn, *getRoleResult.Role.Arn)
	t.Log("Assertion passed. Role ARNs match for Role:", roleName)
}

// ValidatePolicyIsAttachedToASpecificGroup gets policy and checks it is attached to a specific group
func ValidatePolicyIsAttachedToASpecificGroup(t *testing.T, svc *iam.IAM, policyArn string, groupName string, verboseOutput bool) {
	t.Helper()

	policyGroupInput := &iam.ListEntitiesForPolicyInput{
		PolicyArn: aws.String(policyArn),
	}

	policyGroupResult, err := svc.ListEntitiesForPolicy(policyGroupInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(policyGroupResult.String())
	}

	assert.Equal(t, groupName, *policyGroupResult.PolicyGroups[0].GroupName)
	assert.Equal(t, 1, len(policyGroupResult.PolicyGroups))
}

// ValidateGroup gets Group by name and validates its arn
func ValidateGroup(t *testing.T, svc *iam.IAM, groupName string, groupArn string, verboseOutput bool) {
	t.Helper()

	groupResult, err := svc.GetGroup(
		&iam.GetGroupInput{
			GroupName: aws.String(groupName),
		},
	)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(groupResult.String())
	}

	assert.Equal(t, groupArn, *groupResult.Group.Arn)
	assert.Equal(t, groupName, *groupResult.Group.GroupName)
}

// ValidateGroupIsAttachedToASpecificUser get the group and the user attached
func ValidateGroupIsAttachedToASpecificUser(t *testing.T, svc *iam.IAM, groupName string, groupArn string, userName string, userArn string, verboseOutput bool) {
	t.Helper()

	groupInput := &iam.GetGroupInput{
		GroupName: aws.String(groupName),
	}

	groupResult, err := svc.GetGroup(groupInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(groupResult.String())
	}

	assert.Equal(t, userName, *groupResult.Users[0].UserName)
	assert.Equal(t, userArn, *groupResult.Users[0].Arn)
	assert.Equal(t, 1, len(groupResult.Users))

	assert.Equal(t, groupName, *groupResult.Group.GroupName)
	assert.Equal(t, groupArn, *groupResult.Group.Arn)
}

// ValidatePolicyIsAttachedToARole get polcy by arn and validates that at least one role is attached
func ValidatePolicyIsAttachedToARole(t *testing.T, svc *iam.IAM, policyArn string, verboseOutput bool) {
	t.Helper()

	policyRolesInput := &iam.ListEntitiesForPolicyInput{
		PolicyArn: aws.String(policyArn),
	}

	policyRolesResult, err := svc.ListEntitiesForPolicy(policyRolesInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(policyRolesResult.String())
	}

	assert.NotEmpty(t, policyRolesResult.PolicyRoles)
}

// ValidatePolicyIsAttachedToASpecificRole get polcy by arn and validates that the specified role is attached
func ValidatePolicyIsAttachedToASpecificRole(t *testing.T, svc *iam.IAM, policyArn string, roleName string, verboseOutput bool) {
	t.Helper()

	policyRolesInput := &iam.ListEntitiesForPolicyInput{
		PolicyArn: aws.String(policyArn),
	}

	policyRolesResult, err := svc.ListEntitiesForPolicy(policyRolesInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(policyRolesResult.String())
	}

	// Test that the aqua policy is attached to the aqua role
	assert.NotEmpty(t, policyRolesResult.PolicyRoles)
	assert.Contains(t, policyRolesResult.String(), roleName)
}

// ValidateRoleHasManagedPolicyAttached get role by name and validates that the specified role has managed policy attached
func ValidateRoleHasManagedPolicyAttached(t *testing.T, svc *iam.IAM, policyArn string, roleName string, verboseOutput bool) {
	t.Helper()

	policyRolesInput := &iam.ListAttachedRolePoliciesInput{
		RoleName: aws.String(roleName),
	}

	policyRolesResult, err := svc.ListAttachedRolePolicies(policyRolesInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(policyRolesResult.String())
	}

	assert.Contains(t, policyRolesResult.String(), policyArn)
}

// ValidateAccountPasswordPolicy gets the account password policy and validates it
func ValidateAccountPasswordPolicy(t *testing.T, svc *iam.IAM, verboseOutput bool) {
	t.Helper()

	accountPasswordpolicyInput := &iam.GetAccountPasswordPolicyInput{}

	accountPasswordPolicyResult, err := svc.GetAccountPasswordPolicy(accountPasswordpolicyInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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

	assert.Equal(t, true, *accountPasswordPolicyResult.PasswordPolicy.AllowUsersToChangePassword)
	assert.Equal(t, true, *accountPasswordPolicyResult.PasswordPolicy.HardExpiry)
	assert.Equal(t, int64(90), *accountPasswordPolicyResult.PasswordPolicy.MaxPasswordAge)
	assert.Equal(t, int64(8), *accountPasswordPolicyResult.PasswordPolicy.MinimumPasswordLength)
	assert.Equal(t, int64(3), *accountPasswordPolicyResult.PasswordPolicy.PasswordReusePrevention)
	assert.Equal(t, true, *accountPasswordPolicyResult.PasswordPolicy.RequireLowercaseCharacters)
	assert.Equal(t, true, *accountPasswordPolicyResult.PasswordPolicy.RequireNumbers)
	assert.Equal(t, true, *accountPasswordPolicyResult.PasswordPolicy.RequireSymbols)
	assert.Equal(t, true, *accountPasswordPolicyResult.PasswordPolicy.RequireUppercaseCharacters)
}

// ValidatePolicyDetails gets the polcy by arn and validates that the JSON permissions are correct
func ValidatePolicyDetails(t *testing.T, svc *iam.IAM, policyArn string, policyJSON string, verboseOutput bool) {
	t.Helper()

	versionID := GetNewestPolicyVersion(t, svc, policyArn, verboseOutput)

	policyDetailsInput := &iam.GetPolicyVersionInput{
		PolicyArn: aws.String(policyArn),
		VersionId: aws.String(versionID),
	}

	policyDetailsResult, err := svc.GetPolicyVersion(policyDetailsInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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

	decodedValue, err := url.QueryUnescape(*policyDetailsResult.PolicyVersion.Document)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(decodedValue)
		fmt.Println(policyDetailsResult.String())
	}

	assert.JSONEq(t, decodedValue, policyJSON)
}

// ValidateRoleDetails get the role by name and validates the details on it
func ValidateRoleDetails(t *testing.T, svc *iam.IAM, roleName string, roleArn string, trustRelationshipJSON string, tags []string, verboseOutput bool) {
	t.Helper()

	roleInput := &iam.GetRoleInput{
		RoleName: aws.String(roleName),
	}

	roleResult, err := svc.GetRole(roleInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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

	decodedValue, err := url.QueryUnescape(*roleResult.Role.AssumeRolePolicyDocument)
	if err != nil {
		fmt.Println(err.Error())
		t.Logf("Failing test.")
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(decodedValue)
		fmt.Println(roleResult.String())
	}

	assert.Equal(t, roleName, *roleResult.Role.RoleName)
	assert.Equal(t, roleArn, *roleResult.Role.Arn)

	// validate tags
	for i := 0; i < len(tags); i++ {
		assert.Contains(t, roleResult.String(), tags[i])

		if verboseOutput {
			fmt.Println(tags[i])
		}
	}

	assert.JSONEq(t, decodedValue, trustRelationshipJSON)
}

// ValidateRoleInlinePolicy get the role by name and validates the inline policy on it
func ValidateRoleInlinePolicy(t *testing.T, svc *iam.IAM, roleName string, policyName string, policyJSON string, verboseOutput bool) {
	t.Helper()

	roleInput := &iam.GetRolePolicyInput{
		RoleName:   aws.String(roleName),
		PolicyName: aws.String(policyName),
	}

	roleResult, err := svc.GetRolePolicy(roleInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

		t.Fail()

		return
	}

	decodedValue, err := url.QueryUnescape(*roleResult.PolicyDocument)
	if err != nil {
		fmt.Println(err.Error())
		t.Logf("Failing test.")
		t.Fail()

		return
	}

	if verboseOutput {
		fmt.Println(decodedValue)
		fmt.Println(roleResult.String())
	}

	assert.Equal(t, roleName, *roleResult.RoleName)
	assert.Equal(t, policyName, *roleResult.PolicyName)
	assert.JSONEq(t, decodedValue, policyJSON)
}

func ValidateRolePermissionsBoundary(t *testing.T, svc *iam.IAM, roleName string, permissionsBoundaryArn string, verboseOutput bool) {
	t.Helper()

	getRoleResult, err := svc.GetRole(
		&iam.GetRoleInput{
			RoleName: aws.String(roleName),
		},
	)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(getRoleResult.String())
	}

	assert.Equal(t, permissionsBoundaryArn, *getRoleResult.Role.PermissionsBoundary.PermissionsBoundaryArn)
}

// ValidateInstanceProfileDetails get the role by name and validates the details on it
func ValidateInstanceProfileDetails(t *testing.T, svc *iam.IAM, instanceProfileName string, instanceProfileArn string, roleName string, roleArn string, verboseOutput bool) {
	t.Helper()

	instanceProfileInput := &iam.GetInstanceProfileInput{
		InstanceProfileName: aws.String(instanceProfileName),
	}

	instanceProfileResult, err := svc.GetInstanceProfile(instanceProfileInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(instanceProfileResult.String())
	}

	assert.Equal(t, instanceProfileName, *instanceProfileResult.InstanceProfile.InstanceProfileName)
	assert.Equal(t, instanceProfileArn, *instanceProfileResult.InstanceProfile.Arn)
	assert.Equal(t, roleName, *instanceProfileResult.InstanceProfile.Roles[0].RoleName)
	assert.Equal(t, roleArn, *instanceProfileResult.InstanceProfile.Roles[0].Arn)
}

// ValidateAccountAlias gets the account alias and verifies it is what you set it to be
func ValidateAccountAlias(t *testing.T, svc *iam.IAM, accountProfile string, verboseOutput bool) {
	t.Helper()

	accountAliasInput := &iam.ListAccountAliasesInput{}

	accountAliasResult, err := svc.ListAccountAliases(accountAliasInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(accountAliasResult.String())
	}

	assert.Equal(t, accountProfile, *accountAliasResult.AccountAliases[0])
}

// ValidateSAMLProvider get the saml provider
func ValidateSAMLProvider(t *testing.T, svc *iam.IAM, providerArn string, verboseOutput bool) {
	t.Helper()

	samlProviderInput := &iam.ListSAMLProvidersInput{}

	samlProviderResult, err := svc.ListSAMLProviders(samlProviderInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(samlProviderResult.String())
	}

	assert.Equal(t, 1, len(samlProviderResult.SAMLProviderList))
	assert.Equal(t, providerArn, *samlProviderResult.SAMLProviderList[0].Arn)
}

// ValidateNumberOfAttachedRolePolicies get the role by name and validates that the correct number of policies are attached to the role
func ValidateNumberOfAttachedRolePolicies(t *testing.T, svc *iam.IAM, roleName string, roleArn string, numberOfPolicies int, verboseOutput bool) {
	t.Helper()

	listAttachedRolePoliciesInput := &iam.ListAttachedRolePoliciesInput{
		RoleName: aws.String(roleName),
	}

	listAttachedRolePoliciesResult, err := svc.ListAttachedRolePolicies(listAttachedRolePoliciesInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeInvalidInputException:
				fmt.Println(iam.ErrCodeInvalidInputException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
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
		fmt.Println(listAttachedRolePoliciesResult.String())
	}

	assert.Equal(t, numberOfPolicies, len(listAttachedRolePoliciesResult.AttachedPolicies))
}

// GetNewestPolicyVersion gets the newest policy version
func GetNewestPolicyVersion(t *testing.T, svc *iam.IAM, policyArn string, verboseOutput bool) string {
	t.Helper()

	policyInput := &iam.ListPolicyVersionsInput{
		PolicyArn: aws.String(policyArn),
	}

	versionResults, err := svc.ListPolicyVersions(policyInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeNoSuchEntityException:
				fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			case iam.ErrCodeInvalidInputException:
				fmt.Println(iam.ErrCodeInvalidInputException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

		t.Logf("Failing test.")
		t.Fail()

		return "FAIL"
	}

	versionID := ""

	for k, v := range versionResults.Versions {
		if *v.IsDefaultVersion {
			versionID = *v.VersionId

			if verboseOutput {
				fmt.Println(k)
			}

			break
		}
	}

	if verboseOutput {
		fmt.Println(versionResults.Versions)
		fmt.Println(versionID)
	}

	return versionID
}
