terraform {
  required_version = ">= 1.0.0"
  experiments      = [module_variable_optional_attrs]
  required_providers {
    dummy = {
      source  = "hashicorp.com/marshall7m/dummy"
      version = ">= 0.0.1"
    }
  }
}
provider "dummy" {
    foo = "do"
}

resource "dummy_resource" "this" {
  
}