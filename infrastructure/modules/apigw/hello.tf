module "hello" {
  source = "lambda_resource"

  rest_api_id = "${aws_api_gateway_rest_api.main.id}"
  parent_id   = "${aws_api_gateway_rest_api.main.root_resource_id}"
  path_part   = "hello"

  methods_count = 1

  methods = [
    {
      http_method  = "GET"
      function_arn = "${var.hello_arn}"
    },
  ]
}
