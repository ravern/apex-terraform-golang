output "hello_endpoint" {
  value = "${aws_api_gateway_deployment.main.invoke_url}${module.hello.path}"
}
