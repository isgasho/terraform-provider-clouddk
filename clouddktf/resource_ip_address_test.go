/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package clouddktf

import (
	"testing"
)

// TestResourceIPAddressInstantiation tests whether the resourceIPAddress instance can be instantiated.
func TestResourceIPAddressInstantiation(t *testing.T) {
	s := resourceIPAddress()

	if s == nil {
		t.Fatalf("Cannot instantiate resourceIPAddress")
	}
}

// TestResourceIPAddressSchema tests the resourceIPAddress schema.
func TestResourceIPAddressSchema(t *testing.T) {
	s := resourceIPAddress()

	requiredKeys := []string{
		resourceIPAddressServerIDKey,
	}

	for _, v := range requiredKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in resourceIPAddress.Schema: Missing argument \"%s\"", v)
		}

		if s.Schema[v].Required != true {
			t.Fatalf("Error in resourceIPAddress.Schema: Argument \"%s\" is not required", v)
		}
	}

	attributeKeys := []string{
		resourceIPAddressAddressKey,
		resourceIPAddressGatewayKey,
		resourceIPAddressNetmaskKey,
		resourceIPAddressNetworkKey,
		resourceIPAddressNetworkInterfaceIDKey,
	}

	for _, v := range attributeKeys {
		if s.Schema[v] == nil {
			t.Fatalf("Error in resourceIPAddress.Schema: Missing attribute \"%s\"", v)
		}

		if s.Schema[v].Computed != true {
			t.Fatalf("Error in resourceIPAddress.Schema: Attribute \"%s\" is not computed", v)
		}
	}
}
