package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type DeleteNamespaceAuthRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType DeleteNamespaceAuthRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace string `json:"namespace"`

	Body *[]string `json:"body,omitempty"`
}

func (o DeleteNamespaceAuthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNamespaceAuthRequest struct{}"
	}

	return strings.Join([]string{"DeleteNamespaceAuthRequest", string(data)}, " ")
}

type DeleteNamespaceAuthRequestContentType struct {
	value string
}

type DeleteNamespaceAuthRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 DeleteNamespaceAuthRequestContentType
	APPLICATION_JSON             DeleteNamespaceAuthRequestContentType
}

func GetDeleteNamespaceAuthRequestContentTypeEnum() DeleteNamespaceAuthRequestContentTypeEnum {
	return DeleteNamespaceAuthRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: DeleteNamespaceAuthRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: DeleteNamespaceAuthRequestContentType{
			value: "application/json",
		},
	}
}

func (c DeleteNamespaceAuthRequestContentType) Value() string {
	return c.value
}

func (c DeleteNamespaceAuthRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteNamespaceAuthRequestContentType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
