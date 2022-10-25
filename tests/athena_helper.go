package tests

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/stretchr/testify/assert"
)

func ValidateDatabaseExists(t *testing.T, svc *athena.Athena, databaseName string, catalogName string, verboseOutput bool) {
	t.Helper()

	input := &athena.GetDatabaseInput{
		CatalogName:  aws.String(catalogName),
		DatabaseName: aws.String(databaseName),
	}

	fmt.Println("Validating Database: ", databaseName)

	result, err := svc.GetDatabase(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case athena.ErrCodeInternalServerException:
				fmt.Println(athena.ErrCodeInternalServerException, aerr.Error())
			case athena.ErrCodeInvalidRequestException:
				fmt.Println(athena.ErrCodeInvalidRequestException, aerr.Error())
			case athena.ErrCodeMetadataException:
				fmt.Println(athena.ErrCodeMetadataException, aerr.Error())
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
		fmt.Println(fmt.Println(result.String()))
	}

	assert.Equal(t, databaseName, *result.Database.Name)
}

func ValidateTableOrViewExists(t *testing.T, svc *athena.Athena, databaseName string, catalogName string, tableName string, verboseOutput bool) {
	t.Helper()

	input := &athena.GetTableMetadataInput{
		CatalogName:  aws.String(catalogName),
		DatabaseName: aws.String(databaseName),
		TableName:    aws.String(tableName),
	}

	fmt.Println("Validating Table: ", tableName)

	result, err := svc.GetTableMetadata(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case athena.ErrCodeInternalServerException:
				t.Logf("Failing test.")
				fmt.Println(athena.ErrCodeInternalServerException, aerr.Error())
			case athena.ErrCodeInvalidRequestException:
				t.Logf("Failing test.")
				fmt.Println(athena.ErrCodeInvalidRequestException, aerr.Error())
			case athena.ErrCodeMetadataException:
				t.Logf("Failing test.")
				fmt.Println(athena.ErrCodeMetadataException, aerr.Error())
			default:
				t.Logf("Failing test.")
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
		fmt.Println(fmt.Println(result.String()))
	}

	assert.Equal(t, tableName, *result.TableMetadata.Name)
}
