package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateResponsePropertyResponse struct {

	// 属性ID
	PropertyId *int32 `json:"property_id,omitempty"`

	// 属性名称，首位必须为字母，支持大小写字母，数字，中划线及下划线，长度2-50
	PropertyName *string `json:"property_name,omitempty"`

	// 属性描述，长度0-200
	Description *string `json:"description,omitempty"`

	// 属性数据类型，枚举值大小写敏感；number格式为数字，范围±1.0 x 10^-28 to ±7.9228 x 10^28；sting为字符串；integer为整数；datetime为时间，格式为yyyyMMddTHHmmss；json为自定义json格式
	DataType *CreateResponsePropertyResponseDataType `json:"data_type,omitempty"`

	// 是否必填 0-非必填 1-必填
	Required *CreateResponsePropertyResponseRequired `json:"required,omitempty"`

	// 最小值，当data_type为integer或number时有效
	Min *string `json:"min,omitempty"`

	// 最大值，当data_type为integer或number时有效
	Max *string `json:"max,omitempty"`

	// 步长，当data_type为integer或number时有效
	Step *string `json:"step,omitempty"`

	// 字符串最大长度，当data_type为string, datetime, json时有效
	MaxLength *int32 `json:"max_length,omitempty"`

	// 属性单位
	Unit *string `json:"unit,omitempty"`

	// string的枚举值数组，使用逗号分隔
	EnumList       *string `json:"enum_list,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResponsePropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResponsePropertyResponse struct{}"
	}

	return strings.Join([]string{"CreateResponsePropertyResponse", string(data)}, " ")
}

type CreateResponsePropertyResponseDataType struct {
	value string
}

type CreateResponsePropertyResponseDataTypeEnum struct {
	INTEGER  CreateResponsePropertyResponseDataType
	NUMBER   CreateResponsePropertyResponseDataType
	STRING   CreateResponsePropertyResponseDataType
	DATETIME CreateResponsePropertyResponseDataType
	JSON     CreateResponsePropertyResponseDataType
}

func GetCreateResponsePropertyResponseDataTypeEnum() CreateResponsePropertyResponseDataTypeEnum {
	return CreateResponsePropertyResponseDataTypeEnum{
		INTEGER: CreateResponsePropertyResponseDataType{
			value: "integer",
		},
		NUMBER: CreateResponsePropertyResponseDataType{
			value: "number",
		},
		STRING: CreateResponsePropertyResponseDataType{
			value: "string",
		},
		DATETIME: CreateResponsePropertyResponseDataType{
			value: "datetime",
		},
		JSON: CreateResponsePropertyResponseDataType{
			value: "json",
		},
	}
}

func (c CreateResponsePropertyResponseDataType) Value() string {
	return c.value
}

func (c CreateResponsePropertyResponseDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateResponsePropertyResponseDataType) UnmarshalJSON(b []byte) error {
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

type CreateResponsePropertyResponseRequired struct {
	value int32
}

type CreateResponsePropertyResponseRequiredEnum struct {
	E_0 CreateResponsePropertyResponseRequired
	E_1 CreateResponsePropertyResponseRequired
}

func GetCreateResponsePropertyResponseRequiredEnum() CreateResponsePropertyResponseRequiredEnum {
	return CreateResponsePropertyResponseRequiredEnum{
		E_0: CreateResponsePropertyResponseRequired{
			value: 0,
		}, E_1: CreateResponsePropertyResponseRequired{
			value: 1,
		},
	}
}

func (c CreateResponsePropertyResponseRequired) Value() int32 {
	return c.value
}

func (c CreateResponsePropertyResponseRequired) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateResponsePropertyResponseRequired) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
