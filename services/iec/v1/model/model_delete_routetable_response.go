package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteRoutetableResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRoutetableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRoutetableResponse struct{}"
	}

	return strings.Join([]string{"DeleteRoutetableResponse", string(data)}, " ")
}
