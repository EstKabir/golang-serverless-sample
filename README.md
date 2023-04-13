# Golang serverless framework

## Description
Golang serverless framework is a framework for building serverless applications using golang.

## Requirements
- [terraform](https://www.terraform.io/downloads.html)
- [aws cli](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)
- [golang](https://golang.org/doc/install)
- [docker](https://docs.docker.com/get-docker/)
- [docker-compose](https://docs.docker.com/compose/install/)

## Environment variables
```shell
FUNCTION_NAME=todo_api;DATABASE_URL=http://localhost:8000
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

## terraform
Run terraform to deploy the lambda function to aws and other resources.

### start project
```bash
terraform init
```

```bash
terraform plan
terraform apply
```

## DynamoDb 
Run docker compose to start dynamodb locally
```bash
docker-compose up -d
```

## Roadmap
- [x] lambda using api gateway
- [x] lambda using dynamodb
- [ ] terraform for dynamodb and lambda
- [ ] lambda using cloudwatch events, logs, alarms, metrics
- [ ] Service layer
- [ ] unit tests
- [ ] Run locally mode
- [ ] lambda using s3
- [ ] lambda using sqs
- [ ] lambda using sns
- [ ] lambda using ses
- [ ] lambda using step functions
- [ ] implement x-ray
- [ ] lambda using ssm
- [ ] lambda using secrets manager
- [ ] lambda using cognito
- [ ] aws chat bot
    