resource "aws_api_gateway_resource" "hello" {
  rest_api_id = "${aws_api_gateway_rest_api.main.id}"
  parent_id   = "${aws_api_gateway_rest_api.main.root_resource_id}"
  path_part   = "hello"
}

resource "aws_api_gateway_method" "get_hello" {
  rest_api_id   = "${aws_api_gateway_rest_api.main.id}"
  resource_id   = "${aws_api_gateway_resource.hello.id}"
  http_method   = "GET"
  authorization = "NONE"
}

module "get_hello_integration" {
  source = "lambda_integration"

  rest_api_id   = "${aws_api_gateway_rest_api.main.id}"
  resource_id   = "${aws_api_gateway_resource.hello.id}"
  http_method   = "${aws_api_gateway_method.get_hello.http_method}"
  function_arn  = "${var.hello_arn}"
  resource_path = "${aws_api_gateway_resource.hello.path}"
}
