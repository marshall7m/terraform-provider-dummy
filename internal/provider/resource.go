package provider

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dummyResource() *schema.Resource {
	return &schema.Resource{
		Description: `The resource uses the standard resource lifecycle but does nothing.`,

		Create: resourceCreate,
		Read:   resourceRead,
		Delete: resourceDelete,
	}
}

func resourceCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(fmt.Sprintf("%d", rand.Int()))
	return nil
}

func resourceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}