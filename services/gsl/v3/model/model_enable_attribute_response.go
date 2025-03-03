package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type EnableAttributeResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableAttributeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAttributeResponse struct{}"
	}

	return strings.Join([]string{"EnableAttributeResponse", string(data)}, " ")
}
