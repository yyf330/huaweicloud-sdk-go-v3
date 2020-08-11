/*
    * Cbr
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// Request Object
type ListOpLogsRequest struct {
	EndTime string `json:"end_time,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	Offset int32 `json:"offset,omitempty"`
	OperationType string `json:"operation_type,omitempty"`
	ProviderId string `json:"provider_id,omitempty"`
	ResourceId string `json:"resource_id,omitempty"`
	ResourceName string `json:"resource_name,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	Status string `json:"status,omitempty"`
	VaultId string `json:"vault_id,omitempty"`
	VaultName string `json:"vault_name,omitempty"`
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}
