package shell

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() func() *schema.Provider {
	return func() *schema.Provider {

		return &schema.Provider{
			Schema:         map[string]*schema.Schema{},
			DataSourcesMap: map[string]*schema.Resource{},
			ResourcesMap:   map[string]*schema.Resource{},
		}
	}
}