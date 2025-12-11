variable "container_image" {
  type        = string
  default     = "nathanberry97/rss2go:latest"
  description = "Docker image to deploy for the RSS2Go app"
}
variable "app_user" {}
variable "app_pass" {}
variable "domain_name" {}
variable "route53_zone_id" {}
