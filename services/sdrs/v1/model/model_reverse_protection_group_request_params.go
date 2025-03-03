package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 保护组切换请求参数数据结构
type ReverseProtectionGroupRequestParams struct {

	// 切换方向。target：表示从创建保护组时指定的生产站点切换到创建保护组时指定的容灾站点。source：表示从创建保护组时指定的容灾站点切换到创建保护组时指定的生产站点。
	PriorityStation ReverseProtectionGroupRequestParamsPriorityStation `json:"priority_station"`
}

func (o ReverseProtectionGroupRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReverseProtectionGroupRequestParams struct{}"
	}

	return strings.Join([]string{"ReverseProtectionGroupRequestParams", string(data)}, " ")
}

type ReverseProtectionGroupRequestParamsPriorityStation struct {
	value string
}

type ReverseProtectionGroupRequestParamsPriorityStationEnum struct {
	TARGET ReverseProtectionGroupRequestParamsPriorityStation
	SOURCE ReverseProtectionGroupRequestParamsPriorityStation
}

func GetReverseProtectionGroupRequestParamsPriorityStationEnum() ReverseProtectionGroupRequestParamsPriorityStationEnum {
	return ReverseProtectionGroupRequestParamsPriorityStationEnum{
		TARGET: ReverseProtectionGroupRequestParamsPriorityStation{
			value: "target",
		},
		SOURCE: ReverseProtectionGroupRequestParamsPriorityStation{
			value: "source",
		},
	}
}

func (c ReverseProtectionGroupRequestParamsPriorityStation) Value() string {
	return c.value
}

func (c ReverseProtectionGroupRequestParamsPriorityStation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReverseProtectionGroupRequestParamsPriorityStation) UnmarshalJSON(b []byte) error {
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
