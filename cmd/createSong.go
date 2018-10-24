package main

import (
	"context"

	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"rak.dettelback/flip/internal/songs"
	"rak.dettelback/flip/pkg/models"
)

func main() {
	lambda.Start(Handler)
}

func Handler(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body, _ := ioutil.ReadAll(request.Body)
	var song models.Song
	json.Unmarshal(body, &song)
	persister := songs.DynamoPersister{}
	persister.CreateSong(song)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       request.Body,
	}, nil
}
