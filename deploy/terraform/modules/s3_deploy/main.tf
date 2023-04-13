# Create a bucket
resource "aws_s3_bucket" "s3_deploy_bucket" {
  bucket = "deploy-lambda-${var.project_name}-${var.environment}"
  tags = {
    Name        = "Deploy bucket"
    Environment = var.environment
  }
}

resource "aws_s3_bucket_acl" "s3_deploy_bucket_acl" {
  bucket = aws_s3_bucket.s3_deploy_bucket.id
  acl    = "private"
}

data "archive_file" "zip_app" {
  type        = "zip"
  source_file = "${var.deploy_source_folder}/${var.deploy_source_file}"
  output_path = "${var.deploy_source_folder}/${var.deploy_source_file}.zip"
}

# Upload an object
resource "aws_s3_object" "object" {
  source = "${var.deploy_source_folder}/${var.deploy_source_file}.zip"
  bucket = aws_s3_bucket.s3_deploy_bucket.id
  key = "${var.deploy_source_file}.zip"
  depends_on = [aws_s3_bucket.s3_deploy_bucket]
}

output "aws_s3_bucket_name" {
  value = aws_s3_bucket.s3_deploy_bucket.bucket
}