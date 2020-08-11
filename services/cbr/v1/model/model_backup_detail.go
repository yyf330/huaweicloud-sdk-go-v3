/*
    * Cbr
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model
import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
)

type BackupDetail struct {
	// 还原点ID
	CheckpointId string `json:"checkpoint_id"`
	// 创建时间，例如:\"2020-02-05T10:38:34.209782\"
	CreatedAt *sdktime.SdkTime `json:"created_at"`
	// 备份描述
	Description string `json:"description"`
	// 过期时间，例如:\"2020-02-05T10:38:34.209782\"
	ExpiredAt *sdktime.SdkTime `json:"expired_at"`
	ExtendInfo *BackupExtendInfo `json:"extend_info"`
	// 备份ID
	Id string `json:"id"`
	// 备份类型
	ImageType string `json:"image_type"`
	// 备份名称
	Name string `json:"name"`
	// 父备份ID
	ParentId string `json:"parent_id"`
	// 项目ID
	ProjectId string `json:"project_id"`
	// 备份时间
	ProtectedAt string `json:"protected_at"`
	// 资源可用区
	ResourceAz string `json:"resource_az"`
	// 资源ID
	ResourceId string `json:"resource_id"`
	// 资源名称
	ResourceName string `json:"resource_name"`
	// 资源大小，单位为GB
	ResourceSize string `json:"resource_size"`
	// 资源类型
	ResourceType string `json:"resource_type"`
	// 备份状态
	Status string `json:"status"`
	// 更新时间，例如:\"2020-02-05T10:38:34.209782\"
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
	// 存储库ID
	VaultId string `json:"vault_id"`
	// 复制记录
	ReplicationRecords []ReplicationRecordGet `json:"replication_records"`
	// 企业项目id,默认为‘0’。
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// 
	Children []BackupResp `json:"children"`
}
