package provider

import (
	"github.com/erikvanbrakel/terraform-provider-victorops/go-victorops-swagger/client/teams"
	"github.com/erikvanbrakel/terraform-provider-victorops/go-victorops-swagger/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVictorOpsTeam() *schema.Resource {
	return &schema.Resource{
		Create: resourceVictorOpsTeamCreate,
		Read: resourceVictorOpsTeamRead,
		Delete: resourceVictorOpsTeamDelete,
		Update: resourceVictorOpsTeamUpdate,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"is_default_team": {
				Type: schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceVictorOpsTeamCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(VictorOpsClient)

	params := teams.NewPostAPIPublicV1TeamParams()

	name := d.Get("name").(string)
	params.SetBody(&models.AddTeamPayload{
		Name: &name,
	})
	params.SetXVOAPIID(c.ApiId)
	params.SetXVOAPIKey(c.ApiKey)

	response, err := c.Client.Teams.PostAPIPublicV1Team(params)

	if err != nil {
		return err
	}

	d.SetId(response.Payload.Slug)

	return resourceVictorOpsTeamUpdate(d, meta)
}

func resourceVictorOpsTeamUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(VictorOpsClient)

	params := teams.NewPutAPIPublicV1TeamTeamParams()

	params.SetTeam(d.Id())

	name := d.Get("name").(string)
	params.SetBody(&models.AddTeamPayload{
		Name: &name,
	})

	params.SetXVOAPIID(c.ApiId)
	params.SetXVOAPIKey(c.ApiKey)

	_, err := c.Client.Teams.PutAPIPublicV1TeamTeam(params)

	if err != nil {
		return err
	}

	return resourceVictorOpsTeamRead(d, meta)
}

func resourceVictorOpsTeamRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(VictorOpsClient)

	params := teams.NewGetAPIPublicV1TeamTeamParams()

	params.SetTeam(d.Id())
	params.SetXVOAPIID(c.ApiId)
	params.SetXVOAPIKey(c.ApiKey)

	response ,err := c.Client.Teams.GetAPIPublicV1TeamTeam(params)

	if err != nil {
		return err
	}

	d.Set("name", response.Payload.Name)
	d.Set("is_default_team", response.Payload.IsDefaultTeam)

	return nil
}

func resourceVictorOpsTeamDelete(d *schema.ResourceData, meta interface{}) error {

	c := meta.(VictorOpsClient)

	params := teams.NewDeleteAPIPublicV1TeamTeamParams()
	params.SetXVOAPIID(c.ApiId)
	params.SetXVOAPIKey(c.ApiKey)

	params.SetTeam(d.Id())

	_, err := c.Client.Teams.DeleteAPIPublicV1TeamTeam(params)
	return err
}