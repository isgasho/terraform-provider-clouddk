package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

// TestDataSourceServersInstantiation() tests whether the dataSourceServers instance can be instantiated.
func TestDataSourceServersInstantiation(t *testing.T) {
	s := dataSourceServers()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceServers")
	}
}

// TestDataSourceServersSchema() tests the dataSourceServers schema.
func TestDataSourceServersSchema(t *testing.T) {
	s := dataSourceServers()

	if s.Schema[DataSourceServersHostnamesKey] == nil {
		t.Fatalf("Error in dataSourceServers.Schema: Missing attribute \"%s\"", DataSourceServersHostnamesKey)
	}

	if s.Schema[DataSourceServersHostnamesKey].Computed != true {
		t.Fatalf("Error in dataSourceServers.Schema: Attribute \"%s\" is not computed", DataSourceServersHostnamesKey)
	}

	if s.Schema[DataSourceServersIdsKey] == nil {
		t.Fatalf("Error in dataSourceServers.Schema: Missing attribute \"%s\"", DataSourceServersIdsKey)
	}

	if s.Schema[DataSourceServersIdsKey].Computed != true {
		t.Fatalf("Error in dataSourceServers.Schema: Attribute \"%s\" is not computed", DataSourceServersIdsKey)
	}

	if s.Schema[DataSourceServersLabelsKey] == nil {
		t.Fatalf("Error in dataSourceServers.Schema: Missing attribute \"%s\"", DataSourceServersLabelsKey)
	}

	if s.Schema[DataSourceServersLabelsKey].Computed != true {
		t.Fatalf("Error in dataSourceServers.Schema: Attribute \"%s\" is not computed", DataSourceServersLabelsKey)
	}

	if s.Schema[DataSourceServersLocationsKey] == nil {
		t.Fatalf("Error in dataSourceServers.Schema: Missing attribute \"%s\"", DataSourceServersLocationsKey)
	}

	if s.Schema[DataSourceServersLocationsKey].Computed != true {
		t.Fatalf("Error in dataSourceServers.Schema: Attribute \"%s\" is not computed", DataSourceServersLocationsKey)
	}

	if s.Schema[DataSourceServersPackagesKey] == nil {
		t.Fatalf("Error in dataSourceServers.Schema: Missing attribute \"%s\"", DataSourceServersPackagesKey)
	}

	if s.Schema[DataSourceServersPackagesKey].Computed != true {
		t.Fatalf("Error in dataSourceServers.Schema: Attribute \"%s\" is not computed", DataSourceServersPackagesKey)
	}

	if s.Schema[DataSourceServersTemplatesKey] == nil {
		t.Fatalf("Error in dataSourceServers.Schema: Missing attribute \"%s\"", DataSourceServersTemplatesKey)
	}

	if s.Schema[DataSourceServersTemplatesKey].Computed != true {
		t.Fatalf("Error in dataSourceServers.Schema: Attribute \"%s\" is not computed", DataSourceServersTemplatesKey)
	}
}

// TestDataSourceServersSchemaFilter() tests the dataSourceServers.Filter schema.
func TestDataSourceServersSchemaFilter(t *testing.T) {
	s := dataSourceServers()

	if s.Schema[DataSourceServersFilterKey] == nil {
		t.Fatalf("Error in dataSourceServers.Schema: Missing block \"%s\"", DataSourceServersFilterKey)
	}

	if s.Schema[DataSourceServersFilterKey].Optional != true {
		t.Fatalf("Error in dataSourceServers.Schema: Block \"%s\" is not optional", DataSourceServersFilterKey)
	}

	if s.Schema[DataSourceServersFilterKey].Type != schema.TypeList {
		t.Fatalf("Error in dataSourceServers.Schema: Block \"%s\" is not a list", DataSourceServersFilterKey)
	}

	if s.Schema[DataSourceServersFilterKey].MaxItems != 1 {
		t.Fatalf("Error in dataSourceServers.Schema: Block \"%s\" is not limited to a single definition", DataSourceServersFilterKey)
	}

	if s.Schema[DataSourceServersFilterKey].Elem == nil {
		t.Fatalf("Error in dataSourceServers.Schema: Missing element for block \"%s\"", DataSourceServersFilterKey)
	}

	blockElement, blockElementCasted := s.Schema[DataSourceServersFilterKey].Elem.(*schema.Resource)

	if !blockElementCasted {
		t.Fatalf("Error in dataSourceServers.Schema: Element for block \"%s\" is not a pointer to schema.Resource", DataSourceServersFilterKey)
	}

	if blockElement.Schema[DataSourceServersFilterHostnameKey] == nil {
		t.Fatalf("Error in dataSourceServers.Schema.subscriber: Missing argument \"%s\"", DataSourceServersFilterHostnameKey)
	}

	if blockElement.Schema[DataSourceServersFilterHostnameKey].Optional != true {
		t.Fatalf("Error in dataSourceServers.Schema.subscriber: Argument \"%s\" is not optional", DataSourceServersFilterHostnameKey)
	}
}