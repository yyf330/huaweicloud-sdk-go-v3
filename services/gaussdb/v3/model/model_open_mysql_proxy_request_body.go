package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OpenMysqlProxyRequestBody struct {

	// 代理规格码。
	FlavorRef string `json:"flavor_ref"`

	// 代理实例节点数，取值整数2-32。
	NodeNum int32 `json:"node_num"`

	// 代理实例名称。 用于表示实例的名称，同一租户下，同类型的实例名可重名。取值范围：4~64个字符之间，必须以字母开头，区分大小写，可以包含字母、数字、中划线或者下划线，不能包含其他的特殊字符。
	ProxyName *string `json:"proxy_name,omitempty"`

	// 代理实例类型。默认类型为readwrite。
	ProxyMode *OpenMysqlProxyRequestBodyProxyMode `json:"proxy_mode,omitempty"`

	// 数据库节点的读权重设置。
	NodesReadWeight *[]NodesWeight `json:"nodes_read_weight,omitempty"`
}

func (o OpenMysqlProxyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenMysqlProxyRequestBody struct{}"
	}

	return strings.Join([]string{"OpenMysqlProxyRequestBody", string(data)}, " ")
}

type OpenMysqlProxyRequestBodyProxyMode struct {
	value string
}

type OpenMysqlProxyRequestBodyProxyModeEnum struct {
	READWRITE OpenMysqlProxyRequestBodyProxyMode
	READONLY  OpenMysqlProxyRequestBodyProxyMode
}

func GetOpenMysqlProxyRequestBodyProxyModeEnum() OpenMysqlProxyRequestBodyProxyModeEnum {
	return OpenMysqlProxyRequestBodyProxyModeEnum{
		READWRITE: OpenMysqlProxyRequestBodyProxyMode{
			value: "readwrite",
		},
		READONLY: OpenMysqlProxyRequestBodyProxyMode{
			value: "readonly",
		},
	}
}

func (c OpenMysqlProxyRequestBodyProxyMode) Value() string {
	return c.value
}

func (c OpenMysqlProxyRequestBodyProxyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenMysqlProxyRequestBodyProxyMode) UnmarshalJSON(b []byte) error {
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
