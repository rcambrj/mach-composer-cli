# This file is auto-generated by MACH composer
# site: test-1
terraform {
  backend "s3" {
    bucket         = "state-bucket"
    key            = "mach-composer/test-1"
    region         = "eu-west-1"
    dynamodb_table = "lock-table"
    encrypt        = true
  }
  required_providers {
    aws = {
      version = "~> 3.74.1"
    }
  }
}

# File sources
# Resources
# Configuring AWS
provider "aws" {
  region = "eu-west-1"
}

locals {
  tags = {
    Site        = "test-1"
    Environment = "test"
  }
}

# Component: component-1
module "component-1" {
  source = "{{ .PWD }}/testdata/modules/application"
}

output "component-1" {
  description = "The module outputs for component-1"
  sensitive   = true
  value       = module.component-1
}
