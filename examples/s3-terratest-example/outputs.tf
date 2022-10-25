output "s3_bucket_name" {
  value = aws_s3_bucket.main.bucket
}

output "s3_bucket_tags" {
  value = aws_s3_bucket.main.tags
}
