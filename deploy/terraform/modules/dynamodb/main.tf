resource "aws_dynamodb_table" "table" {
  name             = "${var.project_name}-${var.table_name}-${var.environment}"
  hash_key         = "id"
  billing_mode     = "PAY_PER_REQUEST"
  stream_enabled   = true
  stream_view_type = "NEW_AND_OLD_IMAGES"

  attribute {
    name = "id"
    type = "S"
  }

  replica {
    region_name = "us-east-2"
  }
}

resource "aws_iam_policy" "table_read_policy" {
  name = "${var.project_name}-${var.table_name}-dynamodb-read-${var.environment}"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        Action : [
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:GetRecords",
          "dynamodb:ListTables",
          "dynamodb:Query",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:GetRecords",
          "dynamodb:ListTables",
          "dynamodb:Query",
          "dynamodb:BatchGetItem",
          "dynamodb:Scan",
        ],
        Effect : "Allow",
        Resource : aws_dynamodb_table.table.arn
      }
    ]
  })
}

resource "aws_iam_policy" "table_write_policy" {
  name = "${var.project_name}-${var.table_name}-dynamodb-write-${var.environment}"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        Action : [
          "dynamodb:DeleteItem",
          "dynamodb:PutItem",
          "dynamodb:UpdateItem",
          "dynamodb:BatchExecuteStatement",
          "dynamodb:BatchWriteItem"
        ],
        Effect : "Allow",
        Resource : aws_dynamodb_table.table.arn
      }
    ]
  })
}

output "aws_iam_policy_table_read_dynamo_arn" {
  value = aws_iam_policy.table_read_policy.arn
}

output "aws_iam_policy_table_write_dynamo_arn" {
  value = aws_iam_policy.table_write_policy.arn
}
