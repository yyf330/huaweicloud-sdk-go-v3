package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunNerDomainResponse struct {

	// 命名实体识别结果。调用失败时无此字段。
	NamedEntities *[]DomainNamedEntity `json:"named_entities,omitempty"`

	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty"`

	// 调用失败时的错误信息。调用成功时无此字段。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunNerDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunNerDomainResponse struct{}"
	}

	return strings.Join([]string{"RunNerDomainResponse", string(data)}, " ")
}
