variable "aws_region" { default = "eu-west-2" }
variable "instance_name" { default = "rss2go-app" }
variable "instance_plan" { default = "nano_2_0" }
variable "disk_size_gb" { default = 1 }
variable "container_image" {}
variable "enable_auth" { default = "true" }
variable "app_user" {}
variable "app_pass" {}
