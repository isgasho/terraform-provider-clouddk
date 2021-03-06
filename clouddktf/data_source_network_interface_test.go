/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"testing"
)

// TestDataSourceNetworkInterfaceInstantiation tests whether the dataSourceNetworkInterface instance can be instantiated.
func TestDataSourceNetworkInterfaceInstantiation(t *testing.T) {
	s := dataSourceNetworkInterface()

	if s == nil {
		t.Fatalf("Cannot instantiate dataSourceNetworkInterface")
	}
}

// TestDataSourceNetworkInterfaceSchema tests the dataSourceNetworkInterface schema.
func TestDataSourceNetworkInterfaceSchema(t *testing.T) {
	s := dataSourceNetworkInterface()

	idKeys := []string{
		dataSourceNetworkInterfaceIDKey,
		dataSourceNetworkInterfaceServerIDKey,
	}

	for _, v := range idKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceNetworkInterface.Schema: Missing argument \"%s\"", v)
		}

		if s.Schema[v].Required != true {
			t.Fatalf("Error in dataSourceNetworkInterface.Schema: Argument \"%s\" is not required", v)
		}
	}

	attributeKeys := []string{
		dataSourceNetworkInterfaceAddressesKey,
		dataSourceNetworkInterfaceDefaultFirewallRuleKey,
		dataSourceNetworkInterfaceFirewallRulesAddressesKey,
		dataSourceNetworkInterfaceFirewallRulesCommandsKey,
		dataSourceNetworkInterfaceFirewallRulesIdsKey,
		dataSourceNetworkInterfaceFirewallRulesPortsKey,
		dataSourceNetworkInterfaceFirewallRulesProtocolsKey,
		dataSourceNetworkInterfaceGatewaysKey,
		dataSourceNetworkInterfaceLabelKey,
		dataSourceNetworkInterfaceNetmasksKey,
		dataSourceNetworkInterfaceNetworksKey,
		dataSourceNetworkInterfacePrimaryKey,
		dataSourceNetworkInterfaceRateLimitKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in dataSourceNetworkInterface.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in dataSourceNetworkInterface.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
