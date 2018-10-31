package main

import (
	"context"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"rak.dettelback/flip/internal/songs"
	"rak.dettelback/flip/pkg/models"
)

func main() {
	lambda.Start(Handler)
}

func Handler(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	var song models.Song
	json.Unmarshal([]byte(request.Body), &song)
	persister := songs.DynamoPersister{}
	persister.CreateSongInDB(song)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       request.Body,
	}, nil
}
