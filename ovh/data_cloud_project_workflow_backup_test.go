package ovh

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testAccDataCloudProjectWorkflowBackupRead_basic = `
resource "ovh_cloud_project_workflow_backup" "my_backup" {
	service_name = "%s"
	region_name         = "%s"
	cron                = "50 4 * * *"
	instance_id         = "%s"
	max_execution_count = "0"
	name                = "Backup workflow for instance"
	rotation            = "7"
  }

data "ovh_cloud_project_workflows_backup" "backup" {
    region = ovh_cloud_project_workflow_backup.my_backup.region_name
}

output "returns_workflows" {
    value = contains([for workflow in data.ovh_cloud_project_workflows_backup.backup.workflows: workflow.name], "Backup workflow for instance")
}
`

func TestAccDataCloudProjectWorkflowBackupRead_basic(t *testing.T) {
	serviceName := os.Getenv("OVH_CLOUD_PROJECT_SERVICE_TEST")
	regionName := os.Getenv("OVH_CLOUD_PROJECT_WORKFLOW_BACKUP_REGION_TEST")
	instanceId := os.Getenv("OVH_CLOUD_PROJECT_WORKFLOW_BACKUP_INSTANCE_ID_TEST")

	config := fmt.Sprintf(testAccDataCloudProjectWorkflowBackupRead_basic, serviceName, regionName, instanceId)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCloud(t); testAccCheckCloudProjectExists(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("returns_workflows", "true"),
				),
			},
		},
	})
}
