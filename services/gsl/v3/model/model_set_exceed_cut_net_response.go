package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SetExceedCutNetResponse struct {

	// 业务受理单号
	WorkOrderId    *int64 `json:"work_order_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SetExceedCutNetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetExceedCutNetResponse struct{}"
	}

	return strings.Join([]string{"SetExceedCutNetResponse", string(data)}, " ")
}
