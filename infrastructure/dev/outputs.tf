output "lambda_hello_role_arn" {
  value = "${module.iam.lambda_hello_arn}"
}

output "apigw_hello_endpoint" {
  value = "${module.apigw.hello_endpoint}"
}
