package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSubCustomerBudgetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSubCustomerBudgetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubCustomerBudgetResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubCustomerBudgetResponse", string(data)}, " ")
}
