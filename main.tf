variable "name" {}

locals {
  name = "${element(split(".", var.name), 0)}"
  ext  = "${element(split(".", var.name), 1)}"
}

module "mymod" {
  source = "./mymod"

  name = "${local.name}"
  ext  = "${local.ext}"
}

output "file" {
  value = "${module.mymod.filename}"
}
