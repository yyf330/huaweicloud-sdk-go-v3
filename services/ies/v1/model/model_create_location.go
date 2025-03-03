package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 场地信息
type CreateLocation struct {

	// 场地名称，最大支持长度为64个字节。只包含中文字符、英文字母（a-z、A-Z）、数字（0-9）、下划线（_）、中划线（-）
	Name string `json:"name"`

	// 场地所在省/自治区/直辖市
	Province string `json:"province"`

	// 场地所在市/区
	City string `json:"city"`

	// 场地所在区/县
	District string `json:"district"`

	// 场地所在国家
	Country *CreateLocationCountry `json:"country,omitempty"`

	Condition *Condition `json:"condition"`

	// 场地描述，最大支持长度为255个字节，不允许包含<>
	Description *string `json:"description,omitempty"`
}

func (o CreateLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLocation struct{}"
	}

	return strings.Join([]string{"CreateLocation", string(data)}, " ")
}

type CreateLocationCountry struct {
	value string
}

type CreateLocationCountryEnum struct {
	CN CreateLocationCountry
}

func GetCreateLocationCountryEnum() CreateLocationCountryEnum {
	return CreateLocationCountryEnum{
		CN: CreateLocationCountry{
			value: "CN",
		},
	}
}

func (c CreateLocationCountry) Value() string {
	return c.value
}

func (c CreateLocationCountry) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLocationCountry) UnmarshalJSON(b []byte) error {
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
