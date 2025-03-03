package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSetAttributesReq struct {

	// 临时文件ID
	FileTempId *int64 `json:"file_temp_id,omitempty"`

	// 自定义属性集合
	Attributes *[]AttributeReq `json:"attributes,omitempty"`
}

func (o BatchSetAttributesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetAttributesReq struct{}"
	}

	return strings.Join([]string{"BatchSetAttributesReq", string(data)}, " ")
}
