terraform {
  required_providers {
    pterodactyl = {
      version = "0.2"
      source  = "gl.reindruecken.de/Krickler/terraform-provider-pterodactyl"
    }
  }
}

provider "pterodactyl" {}

data "pterodactyl_node" "node" {
  id = 1
}

output "node" {
  value = data.pterodactyl_node.node
}
