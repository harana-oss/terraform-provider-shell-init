package shell

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {

	log.Println("[INFO] Hello World Init")
}

// Provider returns a terraform.ResourceProvider.
func Provider() func() *schema.Provider {
	return func() *schema.Provider {

		log.Println("[INFO] Hello World Provider")

		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"environment": {
					Type:     schema.TypeMap,
					Optional: true,
					Elem:     schema.TypeString,
				},
				"sensitive_environment": {
					Type:      schema.TypeMap,
					Optional:  true,
					Sensitive: true,
					Elem:      schema.TypeString,
				},
				"shell_script": {
					Type:     schema.TypeString,
					Required: true,
				},
			},

			DataSourcesMap: map[string]*schema.Resource{},
			ResourcesMap:   map[string]*schema.Resource{},
		}

		p.ConfigureContextFunc = configure(p)

		return p
	}
}

type client struct {
}

func configure(p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// var diags diag.Diagnostics
		// var environment map[string]interface{}
		// if v, ok := d.GetOk("environment"); ok {
		// 	environment = v.(map[string]interface{})
		// }
		// var sensitiveEnvironment map[string]interface{}
		// if v, ok := d.GetOk("sensitive_environment"); ok {
		// 	sensitiveEnvironment = v.(map[string]interface{})
		// }
		// shellScript := d.Get("shell_script").(string)

		log.Println("[INFO] Hello World")

		// config := Config{
		// 	Environment:          environment,
		// 	SensitiveEnvironment: sensitiveEnvironment,
		// 	ShellScript:          shellScript,
		// }

		return &client{}, diag.Diagnostics{
			{
				Severity: diag.Warning,
				Summary:  "Warning Diagnostic",
				Detail:   "This is a warning.",
			},
		}
	}
}
