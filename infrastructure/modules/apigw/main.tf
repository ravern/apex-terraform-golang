resource "aws_api_gateway_rest_api" "main" {
  name = "${var.name}"
}

resource "aws_api_gateway_deployment" "main" {
  depends_on = [
    "module.hello",
  ]

  rest_api_id = "${aws_api_gateway_rest_api.main.id}"
  stage_name  = "${var.env}"
}
