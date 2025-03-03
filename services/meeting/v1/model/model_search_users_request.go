package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type SearchUsersRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN, 英文为en-US
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 查询偏移量,若超过最大数量，则返回最后一页
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量 默认值：0
	Limit *int32 `json:"limit,omitempty"`

	// 搜索条件，支持姓名、手机、邮箱、账号、第三方账号模糊搜索。
	SearchKey *string `json:"searchKey,omitempty"`

	// 排序字段名称  约束 - 长度范围为0到64个字符  支持的取值 - userType - adminType - ldapAccount - deptCode - status - sortLevel
	SortField *string `json:"sortField,omitempty"`

	// 是否按升序排序
	IsAsc *bool `json:"isAsc,omitempty"`

	// 部门编码，不带则查询所有。
	DeptCode *string `json:"deptCode,omitempty"`

	// 是否查询子部门 默认值: true
	EnableSubDept *bool `json:"enableSubDept,omitempty"`

	// 根据管理员类型查询 * 1：普通管理员； * 2：非管理员。
	AdminType *SearchUsersRequestAdminType `json:"adminType,omitempty"`

	// 是否开启智能协同白板功能功能位，不带则搜索所有。
	EnableRoom *bool `json:"enableRoom,omitempty"`

	// 用户类型 * 2：普通用户； * 12：智慧屏用户； * 13：ideaHub用户； * 若不携带，则默认查询普通用户。
	UserType *[]int32 `json:"userType,omitempty"`

	// 用户状态 * 0：正常； * 1：停用。不带则查询所有
	Status *int32 `json:"status,omitempty"`

	// 是否查询未激活的终端 默认值: false
	ContainsUnActive *bool `json:"containsUnActive,omitempty"`
}

func (o SearchUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchUsersRequest struct{}"
	}

	return strings.Join([]string{"SearchUsersRequest", string(data)}, " ")
}

type SearchUsersRequestAdminType struct {
	value int32
}

type SearchUsersRequestAdminTypeEnum struct {
	E_1 SearchUsersRequestAdminType
	E_2 SearchUsersRequestAdminType
}

func GetSearchUsersRequestAdminTypeEnum() SearchUsersRequestAdminTypeEnum {
	return SearchUsersRequestAdminTypeEnum{
		E_1: SearchUsersRequestAdminType{
			value: 1,
		}, E_2: SearchUsersRequestAdminType{
			value: 2,
		},
	}
}

func (c SearchUsersRequestAdminType) Value() int32 {
	return c.value
}

func (c SearchUsersRequestAdminType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchUsersRequestAdminType) UnmarshalJSON(b []byte) error {
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
