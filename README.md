# Golang serverless framework

## Description
This is a serverless framework for golang. It uses the [aws sam cli](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html) to deploy the application to aws lambda. It also uses [terraform](https://www.terraform.io/) to deploy the infrastructure to aws. The application is built using [golang](https://golang.org/) and [aws lambda go](https://docs.aws.amazon.com/lambda/latest/dg/go-programming-model.html). The database is [dynamodb](https://aws.amazon.com/dynamodb/).

## Requirements
- [terraform](https://www.terraform.io/downloads.html)
- [aws cli](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)
- [golang](https://golang.org/doc/install)
- [docker](https://docs.docker.com/get-docker/)
- [docker-compose](https://docs.docker.com/compose/install/)
- [sam-cli](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Environment variables
```shell
FUNCTION_NAME=TODO_API;DATABASE_URL=http://localhost:8000
_LAMBDA_SERVER_PORT=9000;AWS_LAMBDA_RUNTIME_API=http://localhost:9001

```
### FUNCTION_NAME
The name of the function to be executed. This is the name of the function in the `src/functions` directory.
- TODO_API = The api for the todo application

### DATABASE_URL
The url of the database to connect to. This is the url of the database in the `src/database` directory.

### _LAMBDA_SERVER_PORT (only local)
The port to run the lambda server on. This is only used when running the lambda server locally.

### AWS_LAMBDA_RUNTIME_API (only local)
The url of the lambda runtime api. This is only used when running the lambda server locally.


## DynamoDb 
Run docker compose to start dynamodb locally
```bash
docker-compose up -d
```
