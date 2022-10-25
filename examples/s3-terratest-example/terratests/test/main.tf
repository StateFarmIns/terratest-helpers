provider "aws" {
  region = var.region
}

module "test_suite_terratest_sample" {
  source = "./../../"
  region = var.region
}