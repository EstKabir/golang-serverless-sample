package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"golang-serverless-sample/src/config"
	"golang-serverless-sample/src/domain/todo"
)

type TodoDynamoDbRepository struct {
	Database  *dynamodb.DynamoDB
	TableName string
}

// GetTodo retrieves one Todo from the DB based on its ID
func (repository *TodoDynamoDbRepository) GetTodo(uuid string) (todo.Model, error) {

	// Prepares the input to retrieve the item with the given ID
	input := &dynamodb.GetItemInput{
		TableName: aws.String(repository.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(uuid),
			},
		},
	}

	// Retrieves the item
	result, err := repository.Database.GetItem(input)
	if err != nil {
		return todo.Model{}, err
	}
	if result.Item == nil {
		return todo.Model{}, nil
	}

	// Unmarshals the object retrieved into a domain struct
	var todoModel todo.Model
	err = dynamodbattribute.UnmarshalMap(result.Item, &todoModel)
	if err != nil {
		return todo.Model{}, err
	}

	return todoModel, nil
}

// GetTodos retrieves all the Todos from the DB
func (repository *TodoDynamoDbRepository) GetTodos() ([]todo.Model, error) {

	// Prepares the input to scan the whole table
	input := &dynamodb.ScanInput{
		TableName: aws.String(repository.TableName),
	}
	result, err := repository.Database.Scan(input)
	if err != nil {
		return []todo.Model{}, err
	}
	if len(result.Items) == 0 {
		return []todo.Model{}, nil
	}

	// Unmarshals the array retrieved into a domain struct's slice
	var todos []todo.Model
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &todos)
	if err != nil {
		return []todo.Model{}, err
	}

	return todos, nil
}

// CreateTodo inserts a new Todo item to the table.
func (repository *TodoDynamoDbRepository) CreateTodo(todo todo.Model) error {

	// Generates a new random ID
	id := uuid.New().String()

	// Creates the item that's going to be inserted
	input := &dynamodb.PutItemInput{
		TableName: aws.String(repository.TableName),
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

	_, err := repository.Database.PutItem(input)
	return err
}

func NewTodoRepository(config config.Config) (*TodoDynamoDbRepository, error) {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000")},
	)
	if err != nil {
		return &TodoDynamoDbRepository{}, err
	}
	// Create DynamoDB client
	return &TodoDynamoDbRepository{
		Database:  dynamodb.New(sess),
		TableName: "go-serverless-api",
	}, nil
}
