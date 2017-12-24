output "counter_name" {
  value = "${aws_dynamodb_table.counter.id}"
}
