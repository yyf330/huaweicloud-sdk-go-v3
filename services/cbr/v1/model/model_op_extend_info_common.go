/*
    * Cbr
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

type OpExtendInfoCommon struct {
	// 进度，取值为0-100
	Progress int32 `json:"progress"`
	// 请求id
	RequestId string `json:"request_id"`
	// 备份任务id
	TaskId string `json:"task_id,omitempty"`
}
