/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviPool() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPoolRead,
		Schema: map[string]*schema.Schema{
			"analytics_policy": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolAnalyticsPolicySchema(),
			},
			"analytics_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apic_epg_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_persistence_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autoscale_launch_config_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autoscale_networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"autoscale_policy_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capacity_estimation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"capacity_estimation_ttfb_thresh": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"cloud_config_cksum": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conn_pool_properties": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConnPoolPropertiesSchema(),
			},
			"connection_ramp_duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_server_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"delete_server_on_dns_refresh": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"east_west": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"external_autoscale_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fail_action": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFailActionSchema(),
			},
			"fewest_tasks_feedback_delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"graceful_disable_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"health_monitor_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"host_check_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"inline_health_monitor": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ipaddrgroup_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lb_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LB_ALGORITHM_LEAST_CONNECTIONS",
			},
			"lb_algorithm_consistent_hash_hdr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lb_algorithm_core_nonaffinity": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"lb_algorithm_hash": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS",
			},
			"lookup_server_by_name": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"max_concurrent_connections_per_server": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_conn_rate_per_server": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"min_health_monitors_up": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_servers_up": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceNetworkFilterSchema(),
			},
			"nsx_securitygroup": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pki_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"placement_networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePlacementNetworkSchema(),
			},
			"request_queue_depth": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"request_queue_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rewrite_host_header_to_server_name": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rewrite_host_header_to_sni": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_reselect": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPServerReselectSchema(),
			},
			"server_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"servers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerSchema(),
			},
			"service_metadata": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sni_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ssl_key_and_certificate_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_service_port": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
