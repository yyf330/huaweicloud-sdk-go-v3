package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisassociateSubnetResponse struct {
	Routetable     *Routetable `json:"routetable,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o DisassociateSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateSubnetResponse struct{}"
	}

	return strings.Join([]string{"DisassociateSubnetResponse", string(data)}, " ")
}
