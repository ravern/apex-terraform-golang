module "lambda_hello" {
  source = "role"

  name      = "${var.name}_lambda_hello"
  principal = "lambda.amazonaws.com"
}

resource "aws_iam_role_policy_attachment" "lambda_hello_create_put_cw_logs" {
  role       = "${module.lambda_hello.name}"
  policy_arn = "${aws_iam_policy.create_put_cw_logs.arn}"
}
