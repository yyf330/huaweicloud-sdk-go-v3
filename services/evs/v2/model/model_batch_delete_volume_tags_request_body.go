/*
    * EVS
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// This is a auto create Body Object
type BatchDeleteVolumeTagsRequestBody struct {
	// 操作标识，当前支持的取值如下：  删除标签：delete
	Action string `json:"action"`
	// 标签列表。
	Tags []DeleteTagsOption `json:"tags"`
}
