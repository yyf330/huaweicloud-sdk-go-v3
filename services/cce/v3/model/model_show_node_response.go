package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowNodeResponse struct {

	// API类型，固定值“Node”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	// API版本，固定值“v3”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *NodeMetadata `json:"metadata,omitempty"`

	Spec *NodeSpec `json:"spec,omitempty"`

	Status         *NodeStatus `json:"status,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeResponse struct{}"
	}

	return strings.Join([]string{"ShowNodeResponse", string(data)}, " ")
}
