/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"testing"
)

// TestResourceFirewallRuleInstantiation tests whether the resourceFirewallRule instance can be instantiated.
func TestResourceFirewallRuleInstantiation(t *testing.T) {
	s := resourceFirewallRule()

	if s == nil {
		t.Fatalf("Cannot instantiate resourceFirewallRule")
	}
}

// TestResourceFirewallRuleSchema tests the resourceFirewallRule schema.
func TestResourceFirewallRuleSchema(t *testing.T) {
	s := resourceFirewallRule()

	requiredKeys := []string{
		dataSourceFirewallRuleAddressKey,
		dataSourceFirewallRuleCommandKey,
		dataSourceFirewallRuleNetworkInterfaceIDKey,
		dataSourceFirewallRulePortKey,
		dataSourceFirewallRuleProtocolKey,
		dataSourceFirewallRuleServerIDKey,
	}

	for _, v := range requiredKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in resourceFirewallRule.Schema: Missing argument \"%s\"", v)
		}

		if s.Schema[v].Required != true {
			t.Fatalf("Error in resourceFirewallRule.Schema: Argument \"%s\" is not required", v)
		}
	}
}
