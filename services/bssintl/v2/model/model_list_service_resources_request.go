package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListServiceResourcesRequest struct {

	// 语言。中文：zh_CN英文：en_US缺省为zh_CN。
	XLanguage *string `json:"X-Language,omitempty"`

	// 云服务类型编码。例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。
	ServiceTypeCode string `json:"service_type_code"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的数量，默认值为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListServiceResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListServiceResourcesRequest", string(data)}, " ")
}
