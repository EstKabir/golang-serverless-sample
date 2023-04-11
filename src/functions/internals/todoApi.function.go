package internals

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"golang-serverless-sample/src/config"
	"golang-serverless-sample/src/database/dynamodb"
	"golang-serverless-sample/src/domain/todo"
	"net/http"
	"regexp"
)

type TodoApiFunction struct {
	config   config.Config
	database *dynamodb.TodoDynamoDbRepository
}

func (function *TodoApiFunction) handleGetTodo(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Retrieves the ID from the URL
	id := req.PathParameters["id"]

	// Fetches the requested Todo
	getTodo, err := function.database.GetTodo(id)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	// Marshals the struct so the API Gateway is able to proccess it
	js, err := json.Marshal(getTodo)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func (function *TodoApiFunction) handleGetTodos(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Fetches all the Todos
	todos, err := function.database.GetTodos()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	// Marshals the struct so the API Gateway is able to proccess it
	js, err := json.Marshal(todos)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func (function *TodoApiFunction) handleCreateTodo(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Unmarshals the request's body
	var todoModel todo.Model
	err := json.Unmarshal([]byte(req.Body), &todoModel)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	// Inserts the Todo into the table
	err = function.database.CreateTodo(todoModel)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "Created",
	}, nil
}

func (function *TodoApiFunction) Router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Routes the application to the correct handler based in the request's path
	if req.HTTPMethod == "GET" {
		hasID, _ := regexp.MatchString("/todos/.+", req.Path)
		if hasID {
			return function.handleGetTodo(req)
		}

		if req.Path == "/todos" {
			return function.handleGetTodos(req)
		}
	}
	if req.HTTPMethod == "POST" {
		if req.Path == "/todos" {
			return function.handleCreateTodo(req)
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       http.StatusText(http.StatusMethodNotAllowed),
	}, nil
}

func NewTodoApiFunction(config config.Config) (*TodoApiFunction, error) {
	dynamodbTodo, err := dynamodb.NewTodoRepository(config)
	if err != nil {
		return nil, err
	}
	return &TodoApiFunction{
		config:   config,
		database: dynamodbTodo,
	}, nil
}
