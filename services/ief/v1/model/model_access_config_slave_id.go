package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Modbus设备访问id(0-255)
type AccessConfigSlaveId struct {

	// value 最大长度512， value允许英文字母、数字、下划线、中划线、点、逗号、@、#、+、\\、/、？、^、=、%、&、:、~
	Value string `json:"value"`

	// 标识属性是否可选，默认为true
	Optional *bool `json:"optional,omitempty"`

	Metadata *ValueInPropertyVisitorsRegisterTypeMetadata `json:"metadata,omitempty"`
}

func (o AccessConfigSlaveId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigSlaveId struct{}"
	}

	return strings.Join([]string{"AccessConfigSlaveId", string(data)}, " ")
}
