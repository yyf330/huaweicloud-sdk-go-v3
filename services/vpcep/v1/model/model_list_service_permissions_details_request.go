package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListServicePermissionsDetailsRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType string `json:"Content-Type"`

	// 终端节点服务的ID。
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`

	// 权限帐号ID，格式为 “iam:domain::domain_id”。 其中“domain_id”为授权用户的 帐号ID，例如“iam:domain:: 6e9dfd51d1124e8d8498dce89492 3a0d”。 支持模糊搜索。
	Permission *string `json:"permission,omitempty"`

	// 查询返回终端节点服务的白名单数 量限制，即每页返回的个数。 取值范围：0~500，取值一般为 10，20或者50，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量。 偏移量为一个大于0小于终端节点 服务总个数的整数，表示从偏移量 后面的终端节点服务开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 查询结果中白名单列表的排序字 段，取值为create_at，表示白名单 的添加时间。
	SortKey *ListServicePermissionsDetailsRequestSortKey `json:"sort_key,omitempty"`

	// 查询结果中白名单列表的排序方 式，取值为： ● desc：降序排序 ● asc：升序排序 默认值为desc。
	SortDir *ListServicePermissionsDetailsRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListServicePermissionsDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicePermissionsDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListServicePermissionsDetailsRequest", string(data)}, " ")
}

type ListServicePermissionsDetailsRequestSortKey struct {
	value string
}

type ListServicePermissionsDetailsRequestSortKeyEnum struct {
	CREATE_AT ListServicePermissionsDetailsRequestSortKey
	UPDATE_AT ListServicePermissionsDetailsRequestSortKey
}

func GetListServicePermissionsDetailsRequestSortKeyEnum() ListServicePermissionsDetailsRequestSortKeyEnum {
	return ListServicePermissionsDetailsRequestSortKeyEnum{
		CREATE_AT: ListServicePermissionsDetailsRequestSortKey{
			value: "create_at",
		},
		UPDATE_AT: ListServicePermissionsDetailsRequestSortKey{
			value: "update_at",
		},
	}
}

func (c ListServicePermissionsDetailsRequestSortKey) Value() string {
	return c.value
}

func (c ListServicePermissionsDetailsRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServicePermissionsDetailsRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListServicePermissionsDetailsRequestSortDir struct {
	value string
}

type ListServicePermissionsDetailsRequestSortDirEnum struct {
	ASC  ListServicePermissionsDetailsRequestSortDir
	DESC ListServicePermissionsDetailsRequestSortDir
}

func GetListServicePermissionsDetailsRequestSortDirEnum() ListServicePermissionsDetailsRequestSortDirEnum {
	return ListServicePermissionsDetailsRequestSortDirEnum{
		ASC: ListServicePermissionsDetailsRequestSortDir{
			value: "asc",
		},
		DESC: ListServicePermissionsDetailsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListServicePermissionsDetailsRequestSortDir) Value() string {
	return c.value
}

func (c ListServicePermissionsDetailsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServicePermissionsDetailsRequestSortDir) UnmarshalJSON(b []byte) error {
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
