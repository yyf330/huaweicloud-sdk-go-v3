package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteMqsInstanceTopicRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *BatchDeleteMqsInstanceTopicReq `json:"body,omitempty"`
}

func (o BatchDeleteMqsInstanceTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMqsInstanceTopicRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteMqsInstanceTopicRequest", string(data)}, " ")
}
