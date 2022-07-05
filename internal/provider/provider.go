package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"foo": {
                Type:        schema.TypeString,
                Required:    true,
            },
		},

		ResourcesMap: map[string]*schema.Resource{
			"dummy_resource": dummyResource(),
		},
	}
}