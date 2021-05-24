package pterodactyl

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	croc "github.com/krickler/crocgodyl"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PTERODACTYL_TOKEN", nil),
			},
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PTERODACTYL_ENDPOINT", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"pterodactyl_node": dataSourceNode(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	endpoint := d.Get("endpoint").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (token != "") && (endpoint != "") {
		panel, err := croc.NewApp(endpoint, token)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return panel, diags
	}
	panel, err := croc.NewApp(token, endpoint)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return panel, diags
}
