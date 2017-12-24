resource "aws_dynamodb_table" "counter" {
  name           = "${var.name}_counter"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "UserIP"
  range_key      = "Timestamp"

  attribute {
    name = "UserIP"
    type = "S"
  }

  attribute {
    name = "Timestamp"
    type = "S"
  }
}
