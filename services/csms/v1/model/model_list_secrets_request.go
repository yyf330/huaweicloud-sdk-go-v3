package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSecretsRequest struct {

	// 每页返回的个数。  默认值：50。
	Limit *string `json:"limit,omitempty"`

	// 分页查询起始的资源id，为空时为查询第一页
	Marker *string `json:"marker,omitempty"`
}

func (o ListSecretsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretsRequest struct{}"
	}

	return strings.Join([]string{"ListSecretsRequest", string(data)}, " ")
}
