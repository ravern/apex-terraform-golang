resource "aws_iam_role" "main" {
  name = "${var.name}"

  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "${var.principal}"
      },
      "Effect": "Allow"
    }
  ]
}
POLICY
}

resource "aws_iam_role_policy_attachment" "policy" {
  count      = "${var.policy_arns_count}"
  role       = "${var.name}"
  policy_arn = "${element(var.policy_arns, count.index)}"
}
