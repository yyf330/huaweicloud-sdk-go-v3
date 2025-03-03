package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCommandsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 服务ID
	ServiceId string `json:"service_id"`

	// 每页显示条目数量，最大数量999，超过999后只返回999
	Limit *int32 `json:"limit,omitempty"`

	// 命令ID
	CommandId *int32 `json:"command_id,omitempty"`

	// 命令名称
	CommandName *string `json:"command_name,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListCommandsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommandsRequest struct{}"
	}

	return strings.Join([]string{"ListCommandsRequest", string(data)}, " ")
}
