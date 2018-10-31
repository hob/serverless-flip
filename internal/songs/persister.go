package songs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"rak.dettelback/flip/pkg/models"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

type Persister interface {
	CreateSongInDB(song models.Song) error
}

type DynamoPersister struct {
}

func (p *DynamoPersister) CreateSongInDB(song models.Song) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return err
	}
	client := dynamodb.New(sess)
	marshalled,err:=dynamodbattribute.MarshalMap(song)
	songsTableName:=os.Getenv("SONGS_TABLE")
	_, err = client.PutItem(&dynamodb.PutItemInput{
		TableName: &songsTableName,
		Item:marshalled,
	})
	if err != nil {
		return err
	}
	return nil
}
