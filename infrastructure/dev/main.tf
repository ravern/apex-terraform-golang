provider "aws" {
  region = "${var.aws_region}"
}

module "iam" {
  source = "../modules/iam"

  name = "${var.name}"
}

module "apigw" {
  source = "../modules/apigw"

  name = "${var.name}"

  hello_arn = "${var.apex_function_hello}"
}
