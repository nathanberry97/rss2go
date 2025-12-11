output "lightsail_ip" {
  value = module.lightsail.instance_ip
}

output "cloudfront_domain" {
  value = module.cdn.cloudfront_domain
}
