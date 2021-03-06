## v0.4.0

ENHANCEMENTS:

* provider: Upgrade to Go v1.15
* provider: Upgrade dependencies

OTHER:

* provider/docs: Add HTML documentation powered by GitHub Pages and Terraform Registry

## 0.3.1

ENHANCEMENTS:

* provider: Upgraded to Go v1.13 and Terraform v0.12.18

BUG FIXES:

* resource/server: Fixed edge case issue

## 0.3.0

ENHANCEMENTS:

* provider: Moved API functions and structures to `clouddk` package
* resource/server: Added upgrade/downgrade support

BUG FIXES:

* provider: Improved error handling
* resource/server: Fixed issue which caused the provider to crash when no network interfaces are available

## 0.2.1

BUG FIXES:

* provider: Fixed compatibility issues with Terraform 0.12+

## 0.2.0

FEATURES:

* **New Resource:** `clouddk_disk`
* **New Resource:** `clouddk_ip_address`

ENHANCEMENTS:

* provider: Improved error handling
* provider: Improved stability
* resource/server: Added `primary_network_interface_default_firewall_rule` argument
* resource/server: Added `primary_network_interface_label` argument

## 0.1.0

FEATURES:

* **New Data Source:** `clouddk_disk`
* **New Data Source:** `clouddk_disks`
* **New Data Source:** `clouddk_firewall_rule`
* **New Data Source:** `clouddk_firewall_rules`
* **New Data Source:** `clouddk_ip_addresses`
* **New Data Source:** `clouddk_locations`
* **New Data Source:** `clouddk_network_interface`
* **New Data Source:** `clouddk_network_interfaces`
* **New Data Source:** `clouddk_packages`
* **New Data Source:** `clouddk_server`
* **New Data Source:** `clouddk_servers`
* **New Data Source:** `clouddk_templates`
* **New Resource:** `clouddk_firewall_rule`
* **New Resource:** `clouddk_server`
