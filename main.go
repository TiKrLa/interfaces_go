package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "eu-west-1"
		return nil
	})
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	out, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("my-table"),
		Item: map[string]types.AttributeValue{
			"id":    &types.AttributeValueMemberS{Value: "12346"},
			"name":  &types.AttributeValueMemberS{Value: "Tia La"},
			"email": &types.AttributeValueMemberS{Value: "tia.lapinjoki14@gmail.com"},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(out.Attributes)

}
