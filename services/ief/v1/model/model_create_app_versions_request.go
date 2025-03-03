package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAppVersionsRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 应用模板ID
	AppId string `json:"app_id"`

	Body *Version `json:"body,omitempty"`
}

func (o CreateAppVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppVersionsRequest struct{}"
	}

	return strings.Join([]string{"CreateAppVersionsRequest", string(data)}, " ")
}
