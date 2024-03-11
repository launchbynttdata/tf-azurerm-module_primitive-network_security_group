// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

resource "azurerm_network_security_group" "network_security_group" {
  name                = var.name
  location            = var.location
  resource_group_name = var.resource_group_name
  tags                = local.tags

  dynamic "security_rule" {
    for_each = var.security_rules != null ? var.security_rules : []
    content {
      name                                       = security_rule.value.name
      priority                                   = security_rule.value.priority
      direction                                  = security_rule.value.direction
      access                                     = security_rule.value.access
      protocol                                   = security_rule.value.protocol
      source_port_range                          = security_rule.value.source_port_range
      destination_port_range                     = security_rule.value.destination_port_range
      source_address_prefix                      = security_rule.value.source_address_prefix
      destination_address_prefix                 = security_rule.value.destination_address_prefix
      description                                = security_rule.value.description
      source_port_ranges                         = security_rule.value.source_port_ranges
      destination_port_ranges                    = security_rule.value.destination_port_ranges
      source_address_prefixes                    = security_rule.value.source_address_prefixes
      destination_address_prefixes               = security_rule.value.destination_address_prefixes
      source_application_security_group_ids      = security_rule.value.source_application_security_group_ids
      destination_application_security_group_ids = security_rule.value.destination_application_security_group_ids
    }

  }
}
