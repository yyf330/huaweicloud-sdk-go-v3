/*
    * EVS
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

type Resource struct {
	// 资源ID。
	ResourceId string `json:"resource_id"`
	// 资源名称。
	ResourceName string `json:"resource_name,omitempty"`
	ResourceDetail *VolumeDetail `json:"resource_detail"`
	// 标签列表。
	Tags []map[string]string `json:"tags"`
}
