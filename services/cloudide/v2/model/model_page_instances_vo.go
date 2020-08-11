/*
    * cloudide
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

type PageInstancesVo struct {
	// 是否为空
	Empty bool `json:"empty,omitempty"`
	// 列表详情
	Items []InstancesVo `json:"items,omitempty"`
	// 偏移量，表示从此偏移量开始查询
	ItemsBefore int64 `json:"items_before,omitempty"`
	// 每页显示的条目数量
	Size int64 `json:"size,omitempty"`
	// 总数
	TotalItemsCount int64 `json:"total_items_count,omitempty"`
}
