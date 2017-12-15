provider "aws" {
  region = "${var.aws_region}"
}

module "iam" {
  source = "../modules/iam"

  name = "${var.name}"
}
