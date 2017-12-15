module "hello_lambda" {
  source = "role"

  name      = "${var.name}_hello_lambda"
  principal = "lambda.amazonaws.com"
}
