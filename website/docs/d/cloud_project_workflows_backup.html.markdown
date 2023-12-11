---
subcategory : "VM Instances"
---

# ovh_cloud_project_workflows_backup

List all workflows that schedules backups of public cloud instance.

## Example Usage

```hcl
data "ovh_cloud_project_workflows_backup" "backup" {
    region = "xxxxx"
}

output "returns_workflows" {
    value = data.ovh_cloud_project_workflows_backup.backup.workflows
}
```

## Argument Reference

The following arguments are supported:

* `service_name` - (Optional) The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.

* `region_name` - (Mandatory) The name of the openstack region. 


## Attributes Reference
- `workflows` - The list of workflows backup of a public cloud project.
  - `id` - The ID of a the workflow.
  - `name` - The name of a the workflow.
  - `cron` - The cron periodicity at which the backup workflow is scheduled.
  - `executions` - A list of workflows execution.
    - `id` - ID of the execution.
    - `state` - The state of the execution.
    - `stateInfo` - The state information of the execution.
    - `executedAt` - The last execution of the workflow.
  - `createAd` - When the workflow backup was created.
  - `backupName` - The name of the workflow backup.
  - `instanceId` - The instance ID use by the workflow backup.
