/*
    * IMS
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// 批量添加镜像成员body
type BatchAddMembersRequestBody struct {
	// 镜像ID列表
	Images []string `json:"images"`
	// 项目ID列表
	Projects []string `json:"projects"`
}
