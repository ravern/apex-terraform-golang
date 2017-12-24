provider "aws" {
  region = "${var.aws_region}"
}

module "iam" {
  source = "../modules/iam"

  name = "${var.name}_${var.apex_environment}"
}

module "apigw" {
  source = "../modules/apigw"

  name = "${var.name}_${var.apex_environment}"
  env  = "${var.apex_environment}"

  hello_arn = "${var.apex_function_hello}"
}

module "dynamodb" {
  source = "../modules/dynamodb"

  name = "${var.name}_${var.apex_environment}"
}
