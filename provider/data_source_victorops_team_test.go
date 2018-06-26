package provider

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccVictoropsTeamDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVictoropsTeamDataSourceConfig,
			},
		}})
}

var testAccVictoropsTeamDataSourceConfig = `

resource "victorops_team" "team" {
  name = "test_team"
}

data "victorops_team" "team_data" {
  team = "${victorops_team.team.id}"
}

`

