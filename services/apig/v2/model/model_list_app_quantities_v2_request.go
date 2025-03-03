package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAppQuantitiesV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListAppQuantitiesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppQuantitiesV2Request struct{}"
	}

	return strings.Join([]string{"ListAppQuantitiesV2Request", string(data)}, " ")
}
