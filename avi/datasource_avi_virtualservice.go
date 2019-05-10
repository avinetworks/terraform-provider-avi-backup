/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviVirtualService() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviVirtualServiceRead,
		Schema: map[string]*schema.Schema{
			"active_standby_se_tag": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ACTIVE_STANDBY_SE_1"},
			"allow_invalid_client_cert": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"analytics_policy": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAnalyticsPolicySchema(),
			},
			"analytics_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apic_contract_graph": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bulk_sync_kvcache": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"client_auth": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPClientAuthenticationParamsSchema(),
			},
			"close_client_conn_on_config_update": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
			"cloud_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CLOUD_NONE"},
			"connections_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"content_rewrite": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceContentRewriteProfileSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"delay_fairness": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsInfoSchema(),
			},
			"dns_policies": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsPoliciesSchema(),
			},
			"east_west_placement": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_autogw": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enable_rhi": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_rhi_snat": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"error_page_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flow_dist": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOAD_AWARE"},
			"flow_label_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NO_LABEL"},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_name_xlate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_policies": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPPoliciesSchema(),
			},
			"ign_pool_net_reach": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"l4_policies": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceL4PoliciesSchema(),
			},
			"limit_doser": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"max_cps_per_client": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"microservice_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_pools_up": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_security_policy_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsx_securitygroup": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"performance_limits": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePerformanceLimitsSchema(),
			},
			"pool_group_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remove_listening_port_on_vs_down": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"requests_rate_limit": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"saml_sp_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSAMLSPConfigSchema(),
			},
			"scaleout_ecmp": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_group_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_policy_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_network_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_metadata": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_pool_select": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServicePoolSelectorSchema(),
			},
			"services": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServiceSchema(),
			},
			"sideband_profile": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSidebandProfileSchema(),
			},
			"snat_ip": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ssl_key_and_certificate_refs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ssl_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_profile_selectors": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSSLProfileSelectorSchema(),
			},
			"ssl_sess_cache_avg_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"sso_policy_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"static_dns_records": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsRecordSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"topology_policies": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsPoliciesSchema(),
			},
			"traffic_clone_profile_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VS_TYPE_NORMAL"},
			"use_bridge_ip_as_vip": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_vip_as_snat": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vh_domain_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vh_parent_vs_uuid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSchema(),
			},
			"vrf_context_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_datascripts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVSDataScriptsSchema(),
			},
			"vsvip_cloud_config_cksum": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vsvip_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"waf_policy_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
		},
	}
}
