/*
    * Cbr
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// Response Object
type ListOpLogsResponse struct {
	// 任务列表
	OperationLogs []OperationLog `json:"operation_logs,omitempty"`
	// 任务个数
	Count int32 `json:"count,omitempty"`
}
