variable "name" {}
variable "ext" {}

resource "local_file" "foo" {
  content  = "foo!"
  filename = "/tmp/${var.name}.${var.ext}"
}

output "filename" {
  value = "${local_file.foo.filename}"
}
