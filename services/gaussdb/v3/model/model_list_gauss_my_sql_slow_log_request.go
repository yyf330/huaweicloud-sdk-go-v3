package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListGaussMySqlSlowLogRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	StartDate string `json:"start_date"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	EndDate string `json:"end_date"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100
	Limit *int32 `json:"limit,omitempty"`

	// 语句类型，取空值，表示查询所有语句类型，也可指定如下日志类型：INSERT、UPDATE、SELECT、DELETE和CREATE
	Type *string `json:"type,omitempty"`

	// 节点ID
	NodeId string `json:"node_id"`
}

func (o ListGaussMySqlSlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlSlowLogRequest struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlSlowLogRequest", string(data)}, " ")
}
