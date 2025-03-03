package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowWorkItemWrokflowConfigResponse struct {

	// 流转数据
	Workflows      *[]WorkItemStatusFlowVo `json:"workflows,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowWorkItemWrokflowConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkItemWrokflowConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkItemWrokflowConfigResponse", string(data)}, " ")
}
