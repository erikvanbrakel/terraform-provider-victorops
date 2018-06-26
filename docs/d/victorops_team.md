# victorops_team
Provides a VictorOps Team.

## Example Usage
```hcl
resource "victorops_team" "my_ops_team" {
    name = "Operations"
}

data "victorops_team" "my_ops_team" {
    team = "${victorops_team.my_ops_team.id}"
}
```

## Argument reference
 - `team` - The ID (slug) of the team. This is required.

## Attributes reference
 - `name` - Name of the team.
 - `is_default_team` - Self-explanatory.

[Back to Index](../README.md)
