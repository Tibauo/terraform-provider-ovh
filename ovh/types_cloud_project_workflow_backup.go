package ovh

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers"
)

type CloudProjectWorkflowBackupCreateOpts struct {
	Cron              *string `json:"cron"`
	InstanceId        *string `json:"instanceId"`
	MaxExecutionCount *int64  `json:"maxExecutionCount,omitempty"`
	Name              *string `json:"name"`
	Rotation          *int64  `json:"rotation"`
}

type CloudProjectWorkflowBackupResponse struct {
	BackupName string `json:"backupName"`
	CreatedAt  string `json:"createdAt"`
	Cron       string `json:"cron"`
	Id         string `json:"id"`
	InstanceId string `json:"instanceId"`
	Name       string `json:"name"`
}

func (opts *CloudProjectWorkflowBackupCreateOpts) FromResource(d *schema.ResourceData) *CloudProjectWorkflowBackupCreateOpts {
	opts.Cron = helpers.GetNilStringPointerFromData(d, "cron")
	opts.InstanceId = helpers.GetNilStringPointerFromData(d, "instance_id")
	opts.MaxExecutionCount = helpers.GetNilInt64PointerFromData(d, "max_execution_count")
	opts.Name = helpers.GetNilStringPointerFromData(d, "name")
	opts.Rotation = helpers.GetNilInt64PointerFromData(d, "rotation")
	return opts
}

func (v CloudProjectWorkflowBackupResponse) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["backup_name"] = v.BackupName
	obj["created_at"] = v.CreatedAt
	obj["cron"] = v.Cron
	obj["id"] = v.Id
	obj["instance_id"] = v.InstanceId
	obj["name"] = v.Name
	return obj
}

type CloudProjectWorkflowsBackup struct {
	Id         string                                   `json:"id"`
	Name       string                                   `json:"name"`
	Cron       string                                   `json:"cron"`
	Executions []*CloudProjectWorkflowsBackupExecutions `json:"executions"`
	CreatedAt  string                                   `json:"createtAd"`
	InstanceId string                                   `json:"instanceId"`
	BackupName string                                   `json:"backupName"`
}

type CloudProjectWorkflowsBackupExecutions struct {
	Id         string `json:"id"`
	ExecutedAd string `json:"executedAd"`
	State      string `json:"state"`
	StateInfo  string `json:"stateInfo"`
}

func (u *CloudProjectWorkflowsBackup) String() string {
	return fmt.Sprintf("Id: %v, Name: %s, Cron: %s, CreateAd: %s, InstanceId: %s, BackupName: %s", u.Id, u.Name, u.Cron, u.CreatedAt, u.InstanceId, u.BackupName)
}
func (e CloudProjectWorkflowsBackupExecutions) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["executed_ad"] = e.ExecutedAd
	obj["id"] = e.Id
	obj["state"] = e.State
	obj["state_info"] = e.StateInfo
	return obj
}
func (u CloudProjectWorkflowsBackup) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["id"] = u.Id
	obj["name"] = u.Name
	obj["cron"] = u.Cron
	obj["executions"] = u.Executions
	obj["created_at"] = u.CreatedAt
	obj["instance_id"] = u.InstanceId
	obj["backup_name"] = u.BackupName
	var executions []map[string]interface{}
	for _, e := range u.Executions {
		executions = append(executions, e.ToMap())
	}
	obj["executions"] = executions
	return obj
}
