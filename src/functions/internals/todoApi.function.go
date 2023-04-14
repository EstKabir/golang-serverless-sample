package internals

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"golang-serverless-sample/src/config"
	"golang-serverless-sample/src/database/dynamodb"
	"golang-serverless-sample/src/domain/todoDomain"
	"golang-serverless-sample/src/services"
	"net/http"
	"regexp"
)

type TodoApiFunction struct {
	config  config.Config
	service todoDomain.Service
	path    string
}

func (function *TodoApiFunction) handleGetTodo(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	id := req.PathParameters["id"]
	getTodo, err := function.service.FindById(id)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

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

func (function *TodoApiFunction) handleDeleteTodo(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	id := req.PathParameters["id"]
	err := function.service.Delete(id)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Deleted",
	}, nil
}

func (function *TodoApiFunction) handleGetTodos(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	todos, err := function.service.FindAll()
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

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
	var todoModel todoDomain.Model
	err := json.Unmarshal([]byte(req.Body), &todoModel)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	todoCreated, err := function.service.Create(&todoModel)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}
	js, err := json.Marshal(todoCreated)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       http.StatusText(http.StatusInternalServerError),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(js),
	}, nil
}

func (function *TodoApiFunction) Router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	hasID, _ := regexp.MatchString(function.path+"/.+", req.Path)
	if req.HTTPMethod == http.MethodGet && hasID {
		return function.handleGetTodo(req)
	} else if req.HTTPMethod == http.MethodGet && req.Path == function.path {
		return function.handleGetTodos(req)
	} else if req.HTTPMethod == http.MethodPost && req.Path == function.path {
		return function.handleCreateTodo(req)
	} else if req.HTTPMethod == http.MethodDelete && hasID {
		return function.handleDeleteTodo(req)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       http.StatusText(http.StatusMethodNotAllowed),
	}, nil
}

func NewTodoApiFunction(config config.Config) (*TodoApiFunction, error) {
	todoRepository, err := dynamodb.NewTodoRepository(config)
	if err != nil {
		return nil, err
	}
	todoService := services.NewTodoService(todoRepository)
	return &TodoApiFunction{
		config:  config,
		service: todoService,
		path:    "/todos",
	}, nil
}
