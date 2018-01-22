resource "aws_api_gateway_resource" "main" {
  rest_api_id = "${var.rest_api_id}"
  parent_id   = "${var.parent_id}"
  path_part   = "${var.path_part}"
}

resource "aws_api_gateway_method" "main" {
  count         = "${var.methods_count}"
  rest_api_id   = "${var.rest_api_id}"
  resource_id   = "${aws_api_gateway_resource.main.id}"
  http_method   = "${lookup(var.methods[count.index], "http_method")}"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "main" {
  depends_on = [
    "aws_api_gateway_method.main",
  ]

  count                   = "${var.methods_count}"
  rest_api_id             = "${var.rest_api_id}"
  resource_id             = "${aws_api_gateway_resource.main.id}"
  http_method             = "${lookup(var.methods[count.index], "http_method")}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${lookup(var.methods[count.index], "function_arn")}/invocations"
}

resource "aws_lambda_permission" "main" {
  count         = "${var.methods_count}"
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = "${lookup(var.methods[count.index], "function_arn")}"
  principal     = "apigateway.amazonaws.com"
  source_arn    = "arn:aws:execute-api:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${var.rest_api_id}/*/${lookup(var.methods[count.index], "http_method")}${aws_api_gateway_resource.main.path}"
}
