/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danitso/terraform-provider-clouddk/clouddk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	dataSourcePackagesIdsKey   = "ids"
	dataSourcePackagesNamesKey = "names"
)

// dataSourcePackages retrieves a list of server packages.
func dataSourcePackages() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			dataSourcePackagesIdsKey: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			dataSourcePackagesNamesKey: {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourcePackagesRead,
	}
}

// dataSourcePackagesRead reads information about server packages.
func dataSourcePackagesRead(d *schema.ResourceData, m interface{}) error {
	clientSettings := m.(clouddk.ClientSettings)
	req, err := clouddk.GetClientRequestObject(&clientSettings, "GET", "cloudservers/get-packages", new(bytes.Buffer))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	} else if res.StatusCode != 200 {
		return fmt.Errorf("Failed to read the information about the packages - Reason: The API responded with HTTP %s", res.Status)
	}

	list := make(clouddk.PackageeListBody, 0)
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

	d.SetId("locations")

	d.Set(dataSourcePackagesIdsKey, ids)
	d.Set(dataSourcePackagesNamesKey, names)

	return nil
}
