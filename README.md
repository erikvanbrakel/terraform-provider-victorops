# terraform-provider-victorops

[![Build Status](https://travis-ci.org/erikvanbrakel/terraform-provider-victorops.svg?branch=master)](https://travis-ci.org/erikvanbrakel/terraform-provider-victorops)

Terraform provider for https://victorops.com


## Building the provider

In this section you will learn how to build and run terraform-provider-victorops
locally. Please follow the steps below:

Requirements
------------

 - [Terraform] 0.9.x
 - [Go] 1.8 (to build the provider plugin)
 - [VictorOps] Must have an account to be able to
   get api id and key.

Setting your environment
---------------------
Create the couple environment variables below:

`$GOROOT = $HOME/go`

`$GOPATH = $GOROOT/bin`

The GOPATH can be set wherever you want but please read this [topic][0]
to understand how they work.

Clone repository to: `$GOPATH/src/github.com/erikvanbrakel/terraform-provider-victorops`

## Installation
Download the binary for your platform and architecture from the [releases page].
Unpack the zip, and place the `terraform-provider-victorops` binary in the same
directory as `terraform` binary or add a `.terraformrc` file with the provider
stanza:

```hcl
providers {
  victorops = "/PATH/TO/MODULE/ARCH/terraform-provider-victorops"
}
```

## Usage
See the [documentation].

[VictorOps]: https://victorops.com/pricing
[Go]: https://golang.org/doc/install
[Terraform]: https://www.terraform.io/downloads.html
[documentation]: docs/README.md
[releases page]: https://github.com/erikvanbrakel/terraform-provider-victorops/releases

[0]: https://stackoverflow.com/questions/7970390/what-should-be-the-values-of-gopath-and-goroot
