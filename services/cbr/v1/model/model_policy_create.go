/*
    * Cbr
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

type PolicyCreate struct {
	// 是否启用策略
	Enabled bool `json:"enabled,omitempty"`
	// 策略名称，长度限制：1- 64，只能由中文、字母、数字、“_”、“-”组成。
	Name string `json:"name"`
	OperationDefinition *PolicyoOdCreate `json:"operation_definition"`
	// 策略类型，如备份，复制 Enum:[ backup，replication]
	OperationType string `json:"operation_type"`
	Trigger *PolicyTriggerReq `json:"trigger"`
}
