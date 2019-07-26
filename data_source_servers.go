package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
)

const DataSourceServersFilterKey = "filter"
const DataSourceServersFilterHostnameKey = "hostname"
const DataSourceServersHostnamesKey = "hostnames"
const DataSourceServersIdsKey = "ids"
const DataSourceServersLabelsKey = "labels"
const DataSourceServersLocationsKey = "locations"
const DataSourceServersPackagesKey = "packages"
const DataSourceServersTemplatesKey = "templates"

// dataSourceServers() retrieves a list of servers.
func dataSourceServers() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			DataSourceServersFilterKey: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						DataSourceServersFilterHostnameKey: &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The hostname filter",
							ForceNew:    true,
						},
					},
				},
				MaxItems: 1,
			},
			DataSourceServersHostnamesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersIdsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersLabelsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersLocationsKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersPackagesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			DataSourceServersTemplatesKey: &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},

		Read: dataSourceServersRead,
	}
}

// dataSourceServersRead() reads information about servers.
func dataSourceServersRead(d *schema.ResourceData, m interface{}) error {
	filter := d.Get(DataSourceServersFilterKey).([]interface{})
	filterHostname := ""

	if len(filter) > 0 {
		filterData := filter[0].(map[string]interface{})
		filterHostname = filterData[DataSourceServersFilterHostnameKey].(string)
	}

	// Prepare the relative path based on the filters.
	path := "cloudservers?per-page=1000"

	if len(filterHostname) > 0 {
		path = fmt.Sprintf("%s&hostname=%s", path, url.QueryEscape(filterHostname))
	}

	// Retrieve the list of templates by invoking the API action.
	clientSettings := m.(ClientSettings)
	req, reqErr := getClientRequestObject(&clientSettings, "GET", path, new(bytes.Buffer))

	if reqErr != nil {
		return reqErr
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return resErr
	}

	list := make(ServerListBody, 0)
	json.NewDecoder(res.Body).Decode(&list)

	hostnames := make([]interface{}, len(list))
	ids := make([]interface{}, len(list))
	labels := make([]interface{}, len(list))
	locations := make([]interface{}, len(list))
	packages := make([]interface{}, len(list))
	templates := make([]interface{}, len(list))

	for i, v := range list {
		hostnames[i] = v.Hostname
		ids[i] = v.Identifier
		labels[i] = v.Label
		locations[i] = v.Location.Identifier
		packages[i] = v.Package.Identifier
		templates[i] = v.Template.Identifier
	}

	d.SetId("servers")

	d.Set(DataSourceServersHostnamesKey, hostnames)
	d.Set(DataSourceServersIdsKey, ids)
	d.Set(DataSourceServersLabelsKey, labels)
	d.Set(DataSourceServersLocationsKey, locations)
	d.Set(DataSourceServersPackagesKey, packages)
	d.Set(DataSourceServersTemplatesKey, templates)

	return nil
}