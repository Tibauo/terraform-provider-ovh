package ovh

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers/hashcode"
)

func datasourceCloudProjectWorkflowsBackup() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceCloudProjectWorkflowsBackupRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVH_CLOUD_PROJECT_SERVICE", nil),
				Description: "Service name of the resource representing the id of the cloud project.",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Region of the workflow_backup.",
			},
			// Computed
			"workflows": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cron": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"executions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"executed_ad": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state_info": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backup_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceCloudProjectWorkflowsBackupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	regionName := d.Get("region").(string)

	workflows := make([]CloudProjectWorkflowsBackup, 0)

	endpoint := fmt.Sprintf(
		"/cloud/project/%s/region/%s/workflow/backup",
		url.PathEscape(serviceName),
		url.PathEscape(regionName),
	)
	log.Printf("[DEBUG] Response %v, %v, %s", config, workflows, endpoint)
	if err := config.OVHClient.Get(endpoint, &workflows); err != nil {
		return diag.FromErr(helpers.CheckDeleted(d, err, endpoint))
	}
	mapWorkflows := make([]map[string]interface{}, len(workflows))
	ids := make([]string, len(workflows))

	for i, workflow := range workflows {
		mapWorkflows[i] = workflow.ToMap()
		mapWorkflows[i]["id"] = workflow.Id
		ids = append(ids, workflow.Id)
	}
	d.SetId(hashcode.Strings(ids))
	d.Set("workflows", mapWorkflows)

	return nil
}
