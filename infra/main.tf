terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.40.0"
    }
  }

  backend "s3" {
    key            = var.s3_key
    bucket         = var.s3_bucket
    region         = "ap-southeast-1"
    dynamodb_table = "tf-locks-table"
    encrypt        = true
  }
}

provider "aws" {
  region = "ap-southeast-1"
  access_key = var.ecs_ak_id
  secret_key = var.ecs_ak_secret

  default_tags {
    tags = {
      Project = "sprint_perbankan"
    }
  }
}