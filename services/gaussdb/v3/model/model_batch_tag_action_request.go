package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchTagActionRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *BatchOperateInstanceTagRequestBody `json:"body,omitempty"`
}

func (o BatchTagActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagActionRequest struct{}"
	}

	return strings.Join([]string{"BatchTagActionRequest", string(data)}, " ")
}
