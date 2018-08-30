package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"context"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode:200,
		Body:"Flipped!!!",
	} ,nil
}

func main() {
	lambda.Start(Handler)
}
