package provider

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccVictoropsTeam(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVictoropsTeamConfig,
			},
		}})
}

var testAccVictoropsTeamConfig = `

resource "victorops_team" "team" {
  name = "test_team"
}

`
