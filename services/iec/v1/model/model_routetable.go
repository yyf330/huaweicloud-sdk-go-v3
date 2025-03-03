package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 路由表详情
type Routetable struct {

	// 路由表ID  取值范围：标准UUID
	Id *string `json:"id,omitempty"`

	// 路由表名称  取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// 路由表所关联的子网  约束：只能关联路由表所属VPC下的子网
	Subnets *[]BaseId `json:"subnets,omitempty"`

	// 路由表所在的虚拟私有云ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 帐号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 路由表描述信息  取值范围：0-255个字符，不能包含“<”和“>”
	Description *string `json:"description,omitempty"`

	// 是否为默认路由表  取值范围：true表示默认路由表；false表示自定义路由表
	Default *bool `json:"default,omitempty"`
}

func (o Routetable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Routetable struct{}"
	}

	return strings.Join([]string{"Routetable", string(data)}, " ")
}
