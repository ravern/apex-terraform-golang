resource "aws_api_gateway_rest_api" "main" {
  name = "${var.name}"
}

resource "aws_api_gateway_deployment" "api" {
  depends_on = [
    "aws_api_gateway_method.get_hello",
  ]

  rest_api_id = "${aws_api_gateway_rest_api.main.id}"
  stage_name  = "${var.env}"
}
