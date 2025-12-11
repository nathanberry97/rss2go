module "lightsail" {
  source = "./modules/lightsail"

  container_image = var.container_image
  app_user        = var.app_user
  app_pass        = var.app_pass
}

module "cdn" {
  source = "./modules/cdn"

  domain_name     = var.domain_name
  route53_zone_id = var.route53_zone_id
  origin_ip       = module.lightsail.instance_ip
}
