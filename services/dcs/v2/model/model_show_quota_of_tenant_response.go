package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowQuotaOfTenantResponse struct {
	Quotas         *QueryTenantQuotaRespQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowQuotaOfTenantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaOfTenantResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaOfTenantResponse", string(data)}, " ")
}
