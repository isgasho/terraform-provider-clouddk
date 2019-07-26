data "clouddk_servers" "example" {}

output "data_clouddk_servers_example_hostnames" {
  description = "The server hostnames"
  value       = "${data.clouddk_servers.example.hostnames}"
}

output "data_clouddk_servers_example_ids" {
  description = "The server identifiers"
  value       = "${data.clouddk_servers.example.ids}"
}

output "data_clouddk_servers_example_labels" {
  description = "The server labels"
  value       = "${data.clouddk_servers.example.labels}"
}

output "data_clouddk_servers_example_location_ids" {
  description = "The server location identifiers"
  value       = "${data.clouddk_servers.example.location_ids}"
}

output "data_clouddk_servers_example_package_ids" {
  description = "The server package identifiers"
  value       = "${data.clouddk_servers.example.package_ids}"
}

output "data_clouddk_servers_example_template_ids" {
  description = "The server template identifiers"
  value       = "${data.clouddk_servers.example.template_ids}"
}
#==============================================================================
data "clouddk_servers" "example_filter" {
  filter {
    hostname = "terraform"
  }
}

output "data_clouddk_servers_example_filter_hostnames" {
  description = "The server hostnames"
  value       = "${data.clouddk_servers.example_filter.hostnames}"
}

output "data_clouddk_servers_example_filter_ids" {
  description = "The server identifiers"
  value       = "${data.clouddk_servers.example_filter.ids}"
}

output "data_clouddk_servers_example_filter_labels" {
  description = "The server labels"
  value       = "${data.clouddk_servers.example_filter.labels}"
}

output "data_clouddk_servers_example_filter_location_ids" {
  description = "The server location identifiers"
  value       = "${data.clouddk_servers.example_filter.location_ids}"
}

output "data_clouddk_servers_example_filter_package_ids" {
  description = "The server package identifiers"
  value       = "${data.clouddk_servers.example_filter.package_ids}"
}

output "data_clouddk_servers_example_filter_template_ids" {
  description = "The server template identifiers"
  value       = "${data.clouddk_servers.example_filter.template_ids}"
}
