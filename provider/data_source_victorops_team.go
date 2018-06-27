package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/erikvanbrakel/terraform-provider-victorops/go-victorops-swagger/client/teams"
)

func dataSourceVictorOpsUser() *schema.Resource {

	return &schema.Resource {
		Read: dataSourceVictorOpsUserRead,

		Schema: map[string]*schema.Schema {
			"team": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_default_team": {
				Type: schema.TypeBool,
				Computed: true,
			},

		},
	}
}

func dataSourceVictorOpsUserRead(d *schema.ResourceData, meta interface {}) error {
	c := meta.(VictorOpsClient)

	params := teams.NewGetAPIPublicV1TeamTeamParams()

	params.SetTeam(d.Get("team").(string))
	params.SetXVOAPIID(c.ApiId)
	params.SetXVOAPIKey(c.ApiKey)

	response, err := c.Client.Teams.GetAPIPublicV1TeamTeam(params)

	if err != nil {
		return err
	}

	d.SetId(response.Payload.Slug)
	d.Set("name", response.Payload.Name)
	d.Set("is_default_team", response.Payload.IsDefaultTeam)
	d.Set("team", response.Payload.Slug)

	return nil
}