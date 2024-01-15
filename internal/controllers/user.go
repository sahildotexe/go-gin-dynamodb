package controllers

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/sahildotexe/go-gin-dynamodb/config"
	"github.com/sahildotexe/go-gin-dynamodb/constants"
)

func DescribeTable(c *gin.Context) {
	// get table information
	table, err := config.DB.DescribeTable(
		context.TODO(),
		&dynamodb.DescribeTableInput{
			TableName: aws.String(constants.UserTable),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"tableName": table.Table.TableName,
	})
}
