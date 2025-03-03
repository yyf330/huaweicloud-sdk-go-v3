package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLiveDataApiV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询到的API列表
	Apis           *[]LdApiInfo `json:"apis,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListLiveDataApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLiveDataApiV2Response struct{}"
	}

	return strings.Join([]string{"ListLiveDataApiV2Response", string(data)}, " ")
}
