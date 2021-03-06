---
layout: "icingamon"
page_title: "Icinga2: service"
sidebar_current: "docs-icingamon-resource-service"
description: |-
  Configures a service resource. This allows service to be configured, updated and deleted.
---

# icingamon\_service

Configures an Icinga2 service resource. This allows service to be configured, updated,
and deleted.

## Example Usage

```hcl
# Configure a new service to be monitored by an Icinga2 Server
provider "icingamon" {
  api_url = "https://192.168.33.5:5665/v1"
}

resource "icingamon_service" "my-service" {
  name          = "ssh"
  hostname      = "c1-mysql-1"
  check_command = "ssh"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Service object.
* `check_command` - (Required) The name of an existing Icinga2 CheckCommand object that is used to determine if the service is available on the host.
* `hostname` - (Required) The host to check the service's status on

