package provider

import (
	"github.com/erikvanbrakel/terraform-provider-victorops/go-victorops-swagger/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  os.Getenv("VICTOROPS_API_ID"),
			},
			"api_key": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  os.Getenv("VICTOROPS_API_KEY"),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"victorops_team": resourceVictorOpsTeam(),
		},
		
		DataSourcesMap: map[string]*schema.Resource {
			"victorops_team": dataSourceVictorOpsUser(),
		},

		ConfigureFunc: providerConfigure,
	}
}

type VictorOpsClient struct {
	Client *client.VictorOps
	ApiId string
	ApiKey string
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	return VictorOpsClient{
		Client: client.Default,
		ApiId: d.Get("api_id").(string),
		ApiKey: d.Get("api_key").(string),
	}, nil
}
