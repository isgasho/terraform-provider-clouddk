/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	dataSourceTemplatesFilterKey     = "filter"
	dataSourceTemplatesFilterNameKey = "name"
	dataSourceTemplatesIdsKey        = "ids"
	dataSourceTemplatesNamesKey      = "names"
)

// dataSourceTemplates retrieves a list of OS templates.
func dataSourceTemplates() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourceTemplatesFilterKey: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dataSourceTemplatesFilterNameKey: {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The name filter",
							ForceNew:    true,
						},
					},
				},
				MaxItems: 1,
			},
			dataSourceTemplatesIdsKey: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dataSourceTemplatesNamesKey: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourceTemplatesRead,
	}
}

// dataSourceTemplatesRead reads information about OS templates.
func dataSourceTemplatesRead(d *schema.ResourceData, m interface{}) error {
	filter := d.Get(dataSourceTemplatesFilterKey).([]interface{})
	filterName := ""

	if len(filter) > 0 {
		filterData := filter[0].(map[string]interface{})
		filterName = filterData[dataSourceTemplatesFilterNameKey].(string)
	}

	// Prepare the relative path based on the filters.
	path := "templates?per-page=1000"

	if len(filterName) > 0 {
		path = fmt.Sprintf("%s&name=%s", path, url.QueryEscape(filterName))
	}

	// Retrieve the list of templates by invoking the API action.
	clientSettings := m.(clouddk.ClientSettings)
	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", path, new(bytes.Buffer))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the templates - Reason: The API responded with HTTP %s", res.Status)
	}

	list := make(clouddk.TemplateListBody, 0)
	err = json.NewDecoder(res.Body).Decode(&list)

	if err != nil {
		return err
	}

	ids := make([]interface{}, len(list))
	names := make([]interface{}, len(list))

	for i, v := range list {
		ids[i] = v.Identifier
		names[i] = v.Name
	}

	d.SetId("templates")

	d.Set(dataSourceTemplatesIdsKey, ids)
	d.Set(dataSourceTemplatesNamesKey, names)

	return nil
}
