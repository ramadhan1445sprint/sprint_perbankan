terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.40.0"
    }
  }

  backend "s3" {
    key            = "project_sprint/paimon_bank/terraform.tfstate"
    bucket         = "remote-state-tf-1523"
    region         = "ap-southeast-1"
    dynamodb_table = "tf-locks-table"
    encrypt        = true
  }
}

provider "aws" {
  region     = "ap-southeast-1"
  access_key = var.ecs_ak_id
  secret_key = var.ecs_ak_secret

  default_tags {
    tags = {
      Project = "sprint_perbankan"
    }
  }
}
