terraform {
  required_providers {
    onsched = {
      source = "level2.com/level2/onsched"
    }
  }
}

provider "onsched" {}

resource "onsched_webhook" "webhooks" {
  customer_webhook_url = "TEST"
  resource_webhook_url = "TEST"
  reminder_webhook_url = "TEST"
  booking_webhook_url  = "TEST"
}
