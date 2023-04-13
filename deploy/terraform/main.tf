terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

module "s3_deploy_dev" {
  source          = "./modules/s3_deploy"
  environment     = "dev"
  project_name    = var.project_name
  aws_region      = var.aws_region
  deploy_source_file = var.deploy_source_file
  deploy_source_folder = "./../../build"
}

module "dynamo_db_todo_dev" {
  source          = "./modules/dynamodb"
  environment     = "dev"
  project_name    = var.project_name
  aws_region      = var.aws_region
  table_name      = "todo"
}

module "lambda_api_todo_dev" {
  source                                          = "./modules/lambda_api"
  environment                                     = "dev"
  project_name                                    = var.project_name
  project_version                                 = var.project_version
  aws_region                                      = var.aws_region
  function_name                                   = "todo_api"
  deploy_bucket = module.s3_deploy_dev.aws_s3_bucket_name
  deploy_file = "${var.deploy_source_file}.zip"
  aws_iam_policy_todo_table_read_dynamo_arn = module.dynamo_db_todo_dev.aws_iam_policy_table_read_dynamo_arn
  aws_iam_policy_todo_table_write_dynamo_arn = module.dynamo_db_todo_dev.aws_iam_policy_table_write_dynamo_arn
}