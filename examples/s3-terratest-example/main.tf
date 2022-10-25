resource "aws_s3_bucket" "main" {
  bucket = "terratest-example-bucket"
  tags = {
    Name = "terratest-example-bucket"
  }
}

