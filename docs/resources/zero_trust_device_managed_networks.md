---
page_title: "cloudflare_zero_trust_device_managed_networks Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_device_managed_networks (Resource)



## Example Usage

```terraform
resource "cloudflare_zero_trust_device_managed_networks" "example_zero_trust_device_managed_networks" {
  account_id = "699d98642c564d2e855e9661899b7252"
  config = {
    tls_sockaddr = "foo.bar:1234"
    sha256 = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
  }
  name = "managed-network-1"
  type = "tls"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String)
- `config` (Attributes) The configuration object containing information for the WARP client to detect the managed network. (see [below for nested schema](#nestedatt--config))
- `name` (String) The name of the device managed network. This name must be unique.
- `type` (String) The type of device managed network.
Available values: "tls".

### Read-Only

- `id` (String) API UUID.
- `network_id` (String) API UUID.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Required:

- `tls_sockaddr` (String) A network address of the form "host:port" that the WARP client will use to detect the presence of a TLS host.

Optional:

- `sha256` (String) The SHA-256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_zero_trust_device_managed_networks.example '<account_id>/<network_id>'
```
