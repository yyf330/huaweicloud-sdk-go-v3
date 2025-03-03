package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateRoutesResponse struct {

	// 路由列表
	Routes         *[]Route `json:"routes,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o UpdateRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoutesResponse struct{}"
	}

	return strings.Join([]string{"UpdateRoutesResponse", string(data)}, " ")
}
