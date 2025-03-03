package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateDatakeyWithoutPlaintextRequest struct {

	// API版本号
	VersionId string `json:"version_id"`

	Body *CreateDatakeyRequestBody `json:"body,omitempty"`
}

func (o CreateDatakeyWithoutPlaintextRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatakeyWithoutPlaintextRequest struct{}"
	}

	return strings.Join([]string{"CreateDatakeyWithoutPlaintextRequest", string(data)}, " ")
}
