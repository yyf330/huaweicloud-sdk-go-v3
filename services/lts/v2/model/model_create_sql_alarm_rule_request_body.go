package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateSqlAlarmRuleRequestBody struct {

	// SQL告警名称
	SqlAlarmRuleName string `json:"sql_alarm_rule_name"`

	// SQL告警信息描述
	SqlAlarmRuleDescription *string `json:"sql_alarm_rule_description,omitempty"`

	// SQL详细信息
	SqlRequests []SqlRequest `json:"sql_requests"`

	// 告警统计周期
	Frequency *Frequency `json:"frequency"`

	// 条件表达式
	ConditionExpression string `json:"condition_expression"`

	// 告警级别
	SqlAlarmLevel CreateSqlAlarmRuleRequestBodySqlAlarmLevel `json:"sql_alarm_level"`

	// 是否发送
	SqlAlarmSend bool `json:"sql_alarm_send"`

	// domainId
	DomainId string `json:"domain_id"`

	// 通知主题
	NotificationSaveRule *NotificationSaveRule `json:"notification_save_rule,omitempty"`
}

func (o CreateSqlAlarmRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlAlarmRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSqlAlarmRuleRequestBody", string(data)}, " ")
}

type CreateSqlAlarmRuleRequestBodySqlAlarmLevel struct {
	value string
}

type CreateSqlAlarmRuleRequestBodySqlAlarmLevelEnum struct {
	INFO     CreateSqlAlarmRuleRequestBodySqlAlarmLevel
	MINOR    CreateSqlAlarmRuleRequestBodySqlAlarmLevel
	MAJOR    CreateSqlAlarmRuleRequestBodySqlAlarmLevel
	CRITICAL CreateSqlAlarmRuleRequestBodySqlAlarmLevel
}

func GetCreateSqlAlarmRuleRequestBodySqlAlarmLevelEnum() CreateSqlAlarmRuleRequestBodySqlAlarmLevelEnum {
	return CreateSqlAlarmRuleRequestBodySqlAlarmLevelEnum{
		INFO: CreateSqlAlarmRuleRequestBodySqlAlarmLevel{
			value: "Info",
		},
		MINOR: CreateSqlAlarmRuleRequestBodySqlAlarmLevel{
			value: "Minor",
		},
		MAJOR: CreateSqlAlarmRuleRequestBodySqlAlarmLevel{
			value: "Major",
		},
		CRITICAL: CreateSqlAlarmRuleRequestBodySqlAlarmLevel{
			value: "Critical",
		},
	}
}

func (c CreateSqlAlarmRuleRequestBodySqlAlarmLevel) Value() string {
	return c.value
}

func (c CreateSqlAlarmRuleRequestBodySqlAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlAlarmRuleRequestBodySqlAlarmLevel) UnmarshalJSON(b []byte) error {
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
