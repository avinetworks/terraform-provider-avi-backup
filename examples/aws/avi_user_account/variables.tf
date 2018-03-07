variable "aws_access_key" {}

variable "aws_secret_key" {}

variable "aws_region" {
  type    = "string"
  default = "us-west-2"
}

variable "avi_username" {
  type    = "string"
  default = "admin"
}

variable "avi_old_password" {}

variable "avi_password" {}

variable "project_name" {}
