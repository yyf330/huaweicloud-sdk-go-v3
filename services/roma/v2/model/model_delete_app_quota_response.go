package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAppQuotaResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppQuotaResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppQuotaResponse", string(data)}, " ")
}
