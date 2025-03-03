package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type UnbindEipReq struct {

	// 弹性公网IP的ID。
	EipId string `json:"eipId"`
}

func (o UnbindEipReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindEipReq struct{}"
	}

	return strings.Join([]string{"UnbindEipReq", string(data)}, " ")
}
