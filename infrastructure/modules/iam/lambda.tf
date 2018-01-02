module "lambda_hello" {
  source = "role"

  name      = "${var.name}_lambda_hello"
  principal = "lambda.amazonaws.com"

  policies = [
    "${aws_iam_policy.create_put_cw_logs.arn}",
    "${aws_iam_policy.scan_put_item_dynamodb.arn}",
  ]
}
