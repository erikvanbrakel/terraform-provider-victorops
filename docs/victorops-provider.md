# VictorOps Provider
This provider is used to manage your VictorOps setup.. The provider needs to be
configured with the proper credentials before it can be used.

## Example Usage
```hcl
# Configure the VictorOps Provider
provider "victorops" {
    api_id   = "${var.victorops_api_id}"
    api_key  = "${var.victorops_api_key}"
}
```

## Authentication
The provider offers a flexible means of providing credentials for
authentication. The following methods are supported and explained below:

 - Static credentials
 - Environment variables

### Static credentials
Static credentials can be provided by adding an `api_id` and `api_key` in-line
in the provider block:

Usage:
```hcl
provider "victorops" {
    api_id   = "your-api-id"
    api_key  = "your-api-key"
}
```

### Environment variables
You can provide your credentials via the `VICTOROPS_API_ID` and
`VICTOROPS_API_KEY` environment variables, representing your VictorOps API ID
and VictorOps API Key, respectively.

Usage:
```hcl
provider "victorops" { }
```

```bash
$ export VICTOROPS_API_ID="your-api-id"
$ export VICTOROPS_API_KEY="your-api-id"
$ terraform plan
```

## Argument Reference
 - `api_id` - (Optional) This is the VictorOps API ID. It must be provided, but
   it can also be source from the VICTOROPS_API_ID environment variable.
 - `api_key` - (Optional) This is the VictorOps API Key. It must be provided, but
   it can also be sourced from the VICTOROPS_API_KEY variable.

[Back to Index][0]

[0]: README.md
