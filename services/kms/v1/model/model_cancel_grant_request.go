package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelGrantRequest struct {

	// API版本号
	VersionId string `json:"version_id"`

	Body *RevokeGrantRequestBody `json:"body,omitempty"`
}

func (o CancelGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelGrantRequest struct{}"
	}

	return strings.Join([]string{"CancelGrantRequest", string(data)}, " ")
}
