locals {
  network_security_group_name = module.resource_names["network_security_group"].standard
  resource_group_name         = module.resource_names["resource_group"].standard
}
