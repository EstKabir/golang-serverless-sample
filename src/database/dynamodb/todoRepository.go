package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"golang-serverless-sample/src/config"
	"golang-serverless-sample/src/domain/todoDomain"
)

type TodoDynamoDbRepository struct {
	database  *dynamodb.DynamoDB
	tableName string
}

func (repository *TodoDynamoDbRepository) GetById(id string) (*todoDomain.Model, error) {

	input := &dynamodb.GetItemInput{
		TableName: aws.String(repository.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := repository.database.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	var todoModel todoDomain.Model
	err = dynamodbattribute.UnmarshalMap(result.Item, &todoModel)
	if err != nil {
		return nil, err
	}

	return &todoModel, nil
}

func (repository *TodoDynamoDbRepository) FindAll() (*[]todoDomain.Model, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String(repository.tableName),
	}
	result, err := repository.database.Scan(input)
	if err != nil {
		return nil, err
	}

	var todos []todoDomain.Model
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &todos)
	if err != nil {
		return nil, err
	}

	return &todos, nil
}

func (repository *TodoDynamoDbRepository) Create(todo *todoDomain.Model) (*todoDomain.Model, error) {

	// Generates a new random ID
	id := uuid.New().String()

	input := &dynamodb.PutItemInput{
		TableName: aws.String(repository.tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(fmt.Sprintf("%v", id)),
			},
			"title": {
				S: aws.String(todo.Title),
			},
			"description": {
				S: aws.String(todo.Description),
			},
		},
	}

	_, err := repository.database.PutItem(input)
	if err != nil {
		return nil, err
	}
	todo.Id = id
	return todo, nil
}

func (repository *TodoDynamoDbRepository) Update(id string, todo *todoDomain.Model) (*todoDomain.Model, error) {
	panic("implement me")
}

func (repository *TodoDynamoDbRepository) Delete(id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
		TableName: aws.String(repository.tableName),
	}
	_, err := repository.database.DeleteItem(input)
	if err != nil {
		return err
	}
	return nil
}

func NewTodoRepository(config config.Config) (todoDomain.Repository, error) {
	var dynamoSess *session.Session
	if config.Database.Url == "" {
		dynamoSess = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	} else {
		sess, err := session.NewSession(&aws.Config{
			Endpoint: aws.String(config.Database.Url)},
		)
		if err != nil {
			return nil, err
		}
		dynamoSess = sess
	}

	return &TodoDynamoDbRepository{
		database:  dynamodb.New(dynamoSess),
		tableName: config.ProjectName + "-todo-" + config.Environment,
	}, nil
}
