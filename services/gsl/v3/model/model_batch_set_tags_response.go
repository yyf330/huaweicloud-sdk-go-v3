package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchSetTagsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchSetTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchSetTagsResponse", string(data)}, " ")
}
