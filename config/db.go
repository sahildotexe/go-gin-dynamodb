package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DB *dynamodb.Client

func NewAWSConfig() aws.Config {
	// get config from environment variables
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")
	// setup aws credential provider
	credProvider := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
		awsAccessKey,
		awsSecretAccessKey,
		"",
	))
	conf, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credProvider),
	)
	if err != nil {
		panic(err)
	}
	return conf
}

func NewDynamoDBClient(sdkConfig aws.Config) *dynamodb.Client {
	// initialize new dynamodb client from aws config and return it
	return dynamodb.NewFromConfig(sdkConfig)
}

func ConnectDB() {
	config := NewAWSConfig()
	client := NewDynamoDBClient(config)
	log.Printf("Connected to DynamoDB")
	DB = client
}
