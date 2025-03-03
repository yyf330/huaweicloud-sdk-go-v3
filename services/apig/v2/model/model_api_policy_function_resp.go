package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiPolicyFunctionResp struct {

	// 函数URN
	FunctionUrn string `json:"function_urn"`

	// 调用类型 - async： 异步 - sync：同步
	InvocationType ApiPolicyFunctionRespInvocationType `json:"invocation_type"`

	// 版本。字符长度不超过64
	Version *string `json:"version,omitempty"`

	// API网关请求后端服务的超时时间。最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000。  单位：毫秒。
	Timeout *int32 `json:"timeout,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`

	// 关联的策略组合模式： - ALL：满足全部条件 - ANY：满足任一条件
	EffectMode ApiPolicyFunctionRespEffectMode `json:"effect_mode"`

	// 策略后端名称。字符串由中文、英文字母、数字、下划线组成，且只能以中文或英文开头。
	Name string `json:"name"`

	// 后端参数列表
	BackendParams *[]BackendParam `json:"backend_params,omitempty"`

	// 策略条件列表
	Conditions []CoditionResp `json:"conditions"`

	// 后端自定义认证对象的ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`
}

func (o ApiPolicyFunctionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiPolicyFunctionResp struct{}"
	}

	return strings.Join([]string{"ApiPolicyFunctionResp", string(data)}, " ")
}

type ApiPolicyFunctionRespInvocationType struct {
	value string
}

type ApiPolicyFunctionRespInvocationTypeEnum struct {
	ASYNC ApiPolicyFunctionRespInvocationType
	SYNC  ApiPolicyFunctionRespInvocationType
}

func GetApiPolicyFunctionRespInvocationTypeEnum() ApiPolicyFunctionRespInvocationTypeEnum {
	return ApiPolicyFunctionRespInvocationTypeEnum{
		ASYNC: ApiPolicyFunctionRespInvocationType{
			value: "async",
		},
		SYNC: ApiPolicyFunctionRespInvocationType{
			value: "sync",
		},
	}
}

func (c ApiPolicyFunctionRespInvocationType) Value() string {
	return c.value
}

func (c ApiPolicyFunctionRespInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionRespInvocationType) UnmarshalJSON(b []byte) error {
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

type ApiPolicyFunctionRespEffectMode struct {
	value string
}

type ApiPolicyFunctionRespEffectModeEnum struct {
	ALL ApiPolicyFunctionRespEffectMode
	ANY ApiPolicyFunctionRespEffectMode
}

func GetApiPolicyFunctionRespEffectModeEnum() ApiPolicyFunctionRespEffectModeEnum {
	return ApiPolicyFunctionRespEffectModeEnum{
		ALL: ApiPolicyFunctionRespEffectMode{
			value: "ALL",
		},
		ANY: ApiPolicyFunctionRespEffectMode{
			value: "ANY",
		},
	}
}

func (c ApiPolicyFunctionRespEffectMode) Value() string {
	return c.value
}

func (c ApiPolicyFunctionRespEffectMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiPolicyFunctionRespEffectMode) UnmarshalJSON(b []byte) error {
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
