# Terraform Provider Graylog

[![Build Status](https://travis-ci.org/rebuy-de/terraform-provider-graylog.svg?branch=master)](https://travis-ci.org/rebuy-de/terraform-provider-graylog)
[![license](https://img.shields.io/github/license/rebuy-de/terraform-provider-graylog.svg)]()
[![GitHub release](https://img.shields.io/github/release/rebuy-de/terraform-provider-graylog.svg)]()

A [Terraform](https://www.terraform.io/) provider for configuring
[Graylog](https://www.graylog.org/).

> **Development Status** *terraform-provider-graylog* is in an early
> development stage. By far not all possible resources are covered yet and the
> existing ones might be incomplete. You are encouraged to pinpoint missing
> features by creating Issues or Pull Requests.

## Use Case

The *terraform-provider-graylog* can be used to provision a new (or even an
existing) Graylog installation. Putting the configuration into code makes it
easily reproducable, less error prone and changes reviewable. Furthermore
Terraform became a very convinient way for doing Infrastructure-as-Code
everywhere.

## Installation

1. Download the [latest
   binary](https://github.com/rebuy-de/terraform-provider-graylog/releases) (or
   compile it from source).
2. Copy the binary either to `$HOME/.terraform.d/plugins/` (for global use) or
   to `./terraform.d/plugins` relative to your Terraform templates.
3. Make the binary executable: eg `chmod +x
   $HOME/.terraform.d/plugins/terraform-provider-graylog-v0.1.0`.

After that you can write [Terraform
templates](https://www.terraform.io/intro/index.html) as usual.


## Example

At the beginning you need to setup the Graylog server URL and its credentials:

```hcl
provider "graylog" {
    server_url = "http://localhost:9000"
    username   = "admin"
    password   = "admin"
}
```

Alternatively you can define these parameters via environment variables
`GRAYLOG_SERVER_URL`, `GRAYLOG_USERNAME` and `GRAYLOG_PASSWORD`.

The next step is to create actual Graylog resource. For example a GELF UDP
input like in this example:

```hcl
resource "graylog_input" "gelf_udp" {
  title  = "GELF UDP"
  global = true

  gelf_udp {
    bind_address = "0.0.0.0"
    port         = 22201
  }
}
```

## Reference

The following arguments are supported in the provider block:

* *server_url* - This is the URL to the Graylog API, eg
  `http://localhost:9000`. This value can be set by the environment variable
  `GRAYLOG_SERVER_URL`.
* *username* - This is the username for the API user. This value cn be set by
  the environment variable `GRAYLOG_USERNAME`.
* *password* - This is the password for the API user. This value cn be set by
  the environment variable `GRAYLOG_PASSWORD`.

### graylog_input

**Arguments**

The following arguments are supported:

* *title* - The title of the input.
* *global* - (Optional) If *true*, the input will be installed on all nodes.
* *node* - (Optional) The ID of the node where to install the input. Required,
  if *global* is set to `true`.
* *gelf_udp* - (Optional) The configuration as a GELF UDP input (documented
  below). Only on input type is allowed.
* *gelf_tcp* - (Optional) The configuration as a GELF UDP input (documented
  below). Only on input type is allowed.
* *beats* - (Optional) The configuration as a GELF UDP input (documented
  below). Only on input type is allowed.

The `gelf_udp` block supports the following:

* *port* - The port where the input should listen to.
* *bind_address* - (Optional) The IP address where to bind the input to.
  Defaults to `0.0.0.0`.

The `gelf_tcp` block supports the following:

* *port* - The port where the input should listen to.
* *bind_address* - (Optional) The IP address where to bind the input to.
  Defaults to `0.0.0.0`.

The `beats` block supports the following:

* *port* - The port where the input should listen to.
* *bind_address* - (Optional) The IP address where to bind the input to.
  Defaults to `0.0.0.0`.

**Attributes**

The following attributes are exported:

* **name** - The name of the input, eg `GELF UDP`.
* **type** - The type of the input, eg `org.graylog2.inputs.gelf.udp.GELFUDPInput`.

## Developing

### Local Development

For local development you might want to test your binaries against a real
Graylog instance. To do this, you can use the Dockerfile that is used by the
e2e test.

You need to build the image and then run it:

```
docker build -t graylog-all-in-one e2e/docker
docker run --rm -it -p 127.0.0.1:9000:9000 graylog-all-in-one
```

The URL of the Graylog instance is `http://localhost:9000` and username and
password are both `admin`.

Afterwards you can use this Graylog instance with the example templates in
`e2e/example`. Make sure that the provider is built and Terraform finds it:

```
./buildutil
cd e2e/example
terraform init -plugin-dir ../../dist
```
