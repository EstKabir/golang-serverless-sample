
resource "aws_lambda_function" "lambda" {
  function_name = "${var.project_name}-${var.function_name}-${var.environment}"
  role          = aws_iam_role.iam_role_lambda.arn
  handler       = var.deploy_file
  runtime       = "go1.x"
  s3_bucket     = var.deploy_bucket
  s3_key        = "${var.deploy_file}.zip"
  memory_size   = 256

  environment {
    variables = {
      FUNCTION_NAME = var.function_name
    }
  }

  ephemeral_storage {
    size = 512 # Min 512 MB and the Max 10240 MB
  }
}

resource "aws_lambda_function_url" "lambda_function_url" {
  function_name      = aws_lambda_function.lambda.function_name
  authorization_type = "NONE"
  cors {
    allow_credentials = false
    allow_origins     = ["*"]
    allow_methods     = ["*"]
    max_age           = 0
  }
}

// ##################### aws_cloudwatch_log_group ##########################

resource "aws_cloudwatch_log_group" "function_log_group" {
  name              = "/aws/lambda/${aws_lambda_function.lambda.function_name}"
  retention_in_days = var.environment != "prod" ? 7 : 30
  lifecycle {
    prevent_destroy = false
  }
}

// ##################### aws_iam_role ##########################

resource "aws_iam_role" "iam_role_lambda" {
  name = "${var.project_name}-${var.function_name}-${var.environment}"

  assume_role_policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        Action : "sts:AssumeRole",
        Effect : "Allow",
        Principal : {
          "Service" : "lambda.amazonaws.com"
        }
      }
    ]
  })
}

// ##################### aws_iam_policy ##########################

resource "aws_iam_policy" "function_logging_policy_lambda_api" {
  name = "logging-${aws_lambda_function.lambda.function_name}"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        Action : [
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Effect : "Allow",
        Resource : "arn:aws:logs:*:*:*"
      }
    ]
  })
}

// ##################### aws_iam_role_policy_attachment ##########################

resource "aws_iam_role_policy_attachment" "function_logging_policy_attachment" {
  role       = aws_iam_role.iam_role_lambda.id
  policy_arn = aws_iam_policy.function_logging_policy_lambda_api.arn
}

resource "aws_iam_role_policy_attachment" "dynamo_db_subscribe_table_read_policy_attachment" {
  role       = aws_iam_role.iam_role_lambda.id
  policy_arn = var.aws_iam_policy_todo_table_read_dynamo_arn
}

resource "aws_iam_role_policy_attachment" "dynamo_db_subscribe_table_write_policy_attachment" {
  role       = aws_iam_role.iam_role_lambda.id
  policy_arn = var.aws_iam_policy_todo_table_write_dynamo_arn
}