package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJobRequest struct {

	// 任务ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	JobId string `json:"job_id"`
}

func (o ShowJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobRequest struct{}"
	}

	return strings.Join([]string{"ShowJobRequest", string(data)}, " ")
}
