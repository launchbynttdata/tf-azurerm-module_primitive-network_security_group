location                = "eastus"
logical_product_service = "nsg"
class_env               = "demo"
security_rules = [
  {
    name                       = "test123-ib"
    description                = "test123-ib"
    protocol                   = "Tcp"
    source_port_range          = 80
    destination_port_range     = 80
    source_address_prefix      = "*"
    destination_address_prefix = "*"
    access                     = "Allow"
    priority                   = 100
    direction                  = "Inbound"
  },
  {
    name                       = "test123-ob"
    description                = "test123-ob"
    protocol                   = "Tcp"
    source_port_range          = 80
    destination_port_range     = 80
    source_address_prefix      = "*"
    destination_address_prefix = "*"
    access                     = "Allow"
    priority                   = 100
    direction                  = "Outbound"
  }
]
