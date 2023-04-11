package functions

import (
	"github.com/aws/aws-lambda-go/lambda"
	"golang-serverless-sample/src/config"
	"golang-serverless-sample/src/functions/internals"
)

const (
	todoApi string = "TODO_API"
)

func StartFunctions(config config.Config) {
	switch config.FunctionName {
	case todoApi:
		todoApiFunction, err := internals.NewTodoApiFunction(config)
		if err != nil {
			println(err)
			return
		}
		lambda.Start(todoApiFunction.Router)
		return
	}
	return
}
