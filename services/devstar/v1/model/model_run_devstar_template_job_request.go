package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type RunDevstarTemplateJobRequest struct {

	// 语言类型，缺省值为“zh-cn”。  枚举值： - zh-cn：中文 - en-us：英文
	XLanguage *RunDevstarTemplateJobRequestXLanguage `json:"X-Language,omitempty"`

	Body *TemplateJobInfo `json:"body,omitempty"`
}

func (o RunDevstarTemplateJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDevstarTemplateJobRequest struct{}"
	}

	return strings.Join([]string{"RunDevstarTemplateJobRequest", string(data)}, " ")
}

type RunDevstarTemplateJobRequestXLanguage struct {
	value string
}

type RunDevstarTemplateJobRequestXLanguageEnum struct {
	ZH_CN RunDevstarTemplateJobRequestXLanguage
	EN_US RunDevstarTemplateJobRequestXLanguage
}

func GetRunDevstarTemplateJobRequestXLanguageEnum() RunDevstarTemplateJobRequestXLanguageEnum {
	return RunDevstarTemplateJobRequestXLanguageEnum{
		ZH_CN: RunDevstarTemplateJobRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: RunDevstarTemplateJobRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c RunDevstarTemplateJobRequestXLanguage) Value() string {
	return c.value
}

func (c RunDevstarTemplateJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunDevstarTemplateJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
