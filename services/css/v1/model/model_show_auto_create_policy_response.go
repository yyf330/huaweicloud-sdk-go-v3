package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAutoCreatePolicyResponse struct {

	// 快照保留的天数。
	Keepday *int32 `json:"keepday,omitempty"`

	// 每天快照创建时刻。
	Period *string `json:"period,omitempty"`

	// 快照命名前缀。
	Prefix *string `json:"prefix,omitempty"`

	// 快照存放的OBS桶。
	Bucket *string `json:"bucket,omitempty"`

	// 快照在OBS桶中的存放路径。
	BasePath *string `json:"basePath,omitempty"`

	// 访问OBS桶用到的委托。
	Agency *string `json:"agency,omitempty"`

	// 是否开启自动创建快照策略。
	Enable         *string `json:"enable,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAutoCreatePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoCreatePolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoCreatePolicyResponse", string(data)}, " ")
}
