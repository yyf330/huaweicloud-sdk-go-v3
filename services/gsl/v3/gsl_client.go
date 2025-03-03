package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gsl/v3/model"
)

type GslClient struct {
	HcClient *http_client.HcHttpClient
}

func NewGslClient(hcClient *http_client.HcHttpClient) *GslClient {
	return &GslClient{HcClient: hcClient}
}

func GslClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ListProPricePlans 查询套餐列表信息
//
// 查询套餐列表信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListProPricePlans(request *model.ListProPricePlansRequest) (*model.ListProPricePlansResponse, error) {
	requestDef := GenReqDefForListProPricePlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProPricePlansResponse), nil
	}
}

// ListProPricePlansInvoker 查询套餐列表信息
func (c *GslClient) ListProPricePlansInvoker(request *model.ListProPricePlansRequest) *ListProPricePlansInvoker {
	requestDef := GenReqDefForListProPricePlans()
	return &ListProPricePlansInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetAttributes 批量设置自定义属性接口
//
// 批量设置自定义属性接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) BatchSetAttributes(request *model.BatchSetAttributesRequest) (*model.BatchSetAttributesResponse, error) {
	requestDef := GenReqDefForBatchSetAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetAttributesResponse), nil
	}
}

// BatchSetAttributesInvoker 批量设置自定义属性接口
func (c *GslClient) BatchSetAttributesInvoker(request *model.BatchSetAttributesRequest) *BatchSetAttributesInvoker {
	requestDef := GenReqDefForBatchSetAttributes()
	return &BatchSetAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAttribute 用户新增自定义属性接口
//
// 用户新增自定义属性接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) CreateAttribute(request *model.CreateAttributeRequest) (*model.CreateAttributeResponse, error) {
	requestDef := GenReqDefForCreateAttribute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAttributeResponse), nil
	}
}

// CreateAttributeInvoker 用户新增自定义属性接口
func (c *GslClient) CreateAttributeInvoker(request *model.CreateAttributeRequest) *CreateAttributeInvoker {
	requestDef := GenReqDefForCreateAttribute()
	return &CreateAttributeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableAttribute 停用自定义属性接口
//
// 停用自定义属性接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) DisableAttribute(request *model.DisableAttributeRequest) (*model.DisableAttributeResponse, error) {
	requestDef := GenReqDefForDisableAttribute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableAttributeResponse), nil
	}
}

// DisableAttributeInvoker 停用自定义属性接口
func (c *GslClient) DisableAttributeInvoker(request *model.DisableAttributeRequest) *DisableAttributeInvoker {
	requestDef := GenReqDefForDisableAttribute()
	return &DisableAttributeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableAttribute 启用自定义属性接口
//
// 启用自定义属性接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) EnableAttribute(request *model.EnableAttributeRequest) (*model.EnableAttributeResponse, error) {
	requestDef := GenReqDefForEnableAttribute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableAttributeResponse), nil
	}
}

// EnableAttributeInvoker 启用自定义属性接口
func (c *GslClient) EnableAttributeInvoker(request *model.EnableAttributeRequest) *EnableAttributeInvoker {
	requestDef := GenReqDefForEnableAttribute()
	return &EnableAttributeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAttributes 查询自定义属性列表接口
//
// 查询自定义属性列表接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListAttributes(request *model.ListAttributesRequest) (*model.ListAttributesResponse, error) {
	requestDef := GenReqDefForListAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAttributesResponse), nil
	}
}

// ListAttributesInvoker 查询自定义属性列表接口
func (c *GslClient) ListAttributesInvoker(request *model.ListAttributesRequest) *ListAttributesInvoker {
	requestDef := GenReqDefForListAttributes()
	return &ListAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAttribute 修改自定义属性接口
//
// 修改自定义属性接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) UpdateAttribute(request *model.UpdateAttributeRequest) (*model.UpdateAttributeResponse, error) {
	requestDef := GenReqDefForUpdateAttribute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAttributeResponse), nil
	}
}

// UpdateAttributeInvoker 修改自定义属性接口
func (c *GslClient) UpdateAttributeInvoker(request *model.UpdateAttributeRequest) *UpdateAttributeInvoker {
	requestDef := GenReqDefForUpdateAttribute()
	return &UpdateAttributeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRealName 清除实名认证信息
//
// 清除实名认证信息，接口只支持电信卡调用
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) DeleteRealName(request *model.DeleteRealNameRequest) (*model.DeleteRealNameResponse, error) {
	requestDef := GenReqDefForDeleteRealName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRealNameResponse), nil
	}
}

// DeleteRealNameInvoker 清除实名认证信息
func (c *GslClient) DeleteRealNameInvoker(request *model.DeleteRealNameRequest) *DeleteRealNameInvoker {
	requestDef := GenReqDefForDeleteRealName()
	return &DeleteRealNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableSimCard 激活实体卡
//
// 创建激活实体卡申请，返回业务受理单号。1~2个工作日完成激活操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) EnableSimCard(request *model.EnableSimCardRequest) (*model.EnableSimCardResponse, error) {
	requestDef := GenReqDefForEnableSimCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableSimCardResponse), nil
	}
}

// EnableSimCardInvoker 激活实体卡
func (c *GslClient) EnableSimCardInvoker(request *model.EnableSimCardRequest) *EnableSimCardInvoker {
	requestDef := GenReqDefForEnableSimCard()
	return &EnableSimCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSimCards 查询SIM卡列表
//
// 查询SIM卡列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListSimCards(request *model.ListSimCardsRequest) (*model.ListSimCardsResponse, error) {
	requestDef := GenReqDefForListSimCards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimCardsResponse), nil
	}
}

// ListSimCardsInvoker 查询SIM卡列表
func (c *GslClient) ListSimCardsInvoker(request *model.ListSimCardsRequest) *ListSimCardsInvoker {
	requestDef := GenReqDefForListSimCards()
	return &ListSimCardsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterImei SIM卡机卡重绑
//
// 支持固定机卡重绑(需要上传IMEI，将SIM卡绑定到指定IMEI的设备)和普通机卡重绑(会清除之前绑定的设备,将SIM卡绑定到正在使用的设备)，单卡每月只允许重绑2次，接口只支持电信卡调用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) RegisterImei(request *model.RegisterImeiRequest) (*model.RegisterImeiResponse, error) {
	requestDef := GenReqDefForRegisterImei()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterImeiResponse), nil
	}
}

// RegisterImeiInvoker SIM卡机卡重绑
func (c *GslClient) RegisterImeiInvoker(request *model.RegisterImeiRequest) *RegisterImeiInvoker {
	requestDef := GenReqDefForRegisterImei()
	return &RegisterImeiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetSimCard SIM卡单卡复机
//
// 创建复机申请，返回业务受理单号。1~2个工作日完成复机操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ResetSimCard(request *model.ResetSimCardRequest) (*model.ResetSimCardResponse, error) {
	requestDef := GenReqDefForResetSimCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetSimCardResponse), nil
	}
}

// ResetSimCardInvoker SIM卡单卡复机
func (c *GslClient) ResetSimCardInvoker(request *model.ResetSimCardRequest) *ResetSimCardInvoker {
	requestDef := GenReqDefForResetSimCard()
	return &ResetSimCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetExceedCutNet SIM卡达量断网/恢复在用
//
// SIM卡达量断网/恢复在用,只支持电信卡。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) SetExceedCutNet(request *model.SetExceedCutNetRequest) (*model.SetExceedCutNetResponse, error) {
	requestDef := GenReqDefForSetExceedCutNet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetExceedCutNetResponse), nil
	}
}

// SetExceedCutNetInvoker SIM卡达量断网/恢复在用
func (c *GslClient) SetExceedCutNetInvoker(request *model.SetExceedCutNetRequest) *SetExceedCutNetInvoker {
	requestDef := GenReqDefForSetExceedCutNet()
	return &SetExceedCutNetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetSpeedValue 实体卡限速
//
// 实体卡限速接口,支持电信和联通实体卡调用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) SetSpeedValue(request *model.SetSpeedValueRequest) (*model.SetSpeedValueResponse, error) {
	requestDef := GenReqDefForSetSpeedValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSpeedValueResponse), nil
	}
}

// SetSpeedValueInvoker 实体卡限速
func (c *GslClient) SetSpeedValueInvoker(request *model.SetSpeedValueRequest) *SetSpeedValueInvoker {
	requestDef := GenReqDefForSetSpeedValue()
	return &SetSpeedValueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRealNamed 查询SIM卡实名认证信息
//
// 实时查询SIM卡实名认证信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ShowRealNamed(request *model.ShowRealNamedRequest) (*model.ShowRealNamedResponse, error) {
	requestDef := GenReqDefForShowRealNamed()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRealNamedResponse), nil
	}
}

// ShowRealNamedInvoker 查询SIM卡实名认证信息
func (c *GslClient) ShowRealNamedInvoker(request *model.ShowRealNamedRequest) *ShowRealNamedInvoker {
	requestDef := GenReqDefForShowRealNamed()
	return &ShowRealNamedInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSimCard 查询SIM卡详情
//
// 查询SIM卡详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ShowSimCard(request *model.ShowSimCardRequest) (*model.ShowSimCardResponse, error) {
	requestDef := GenReqDefForShowSimCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSimCardResponse), nil
	}
}

// ShowSimCardInvoker 查询SIM卡详情
func (c *GslClient) ShowSimCardInvoker(request *model.ShowSimCardRequest) *ShowSimCardInvoker {
	requestDef := GenReqDefForShowSimCard()
	return &ShowSimCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartStopNet SIM卡申请断网/恢复在用
//
// SIM卡申请断网/恢复在用,只支持电信卡。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) StartStopNet(request *model.StartStopNetRequest) (*model.StartStopNetResponse, error) {
	requestDef := GenReqDefForStartStopNet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartStopNetResponse), nil
	}
}

// StartStopNetInvoker SIM卡申请断网/恢复在用
func (c *GslClient) StartStopNetInvoker(request *model.StartStopNetRequest) *StartStopNetInvoker {
	requestDef := GenReqDefForStartStopNet()
	return &StartStopNetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopSimCard SIM卡单卡停机
//
// 创建停机申请，返回业务受理单号。1~2个工作日完成停机操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) StopSimCard(request *model.StopSimCardRequest) (*model.StopSimCardResponse, error) {
	requestDef := GenReqDefForStopSimCard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopSimCardResponse), nil
	}
}

// StopSimCardInvoker SIM卡单卡停机
func (c *GslClient) StopSimCardInvoker(request *model.StopSimCardRequest) *StopSimCardInvoker {
	requestDef := GenReqDefForStopSimCard()
	return &StopSimCardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSimPoolMembers 查询流量池成员列表
//
// 查询流量池成员列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListSimPoolMembers(request *model.ListSimPoolMembersRequest) (*model.ListSimPoolMembersResponse, error) {
	requestDef := GenReqDefForListSimPoolMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimPoolMembersResponse), nil
	}
}

// ListSimPoolMembersInvoker 查询流量池成员列表
func (c *GslClient) ListSimPoolMembersInvoker(request *model.ListSimPoolMembersRequest) *ListSimPoolMembersInvoker {
	requestDef := GenReqDefForListSimPoolMembers()
	return &ListSimPoolMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSimPools 查询流量池列表
//
// 查询流量池列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListSimPools(request *model.ListSimPoolsRequest) (*model.ListSimPoolsResponse, error) {
	requestDef := GenReqDefForListSimPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimPoolsResponse), nil
	}
}

// ListSimPoolsInvoker 查询流量池列表
func (c *GslClient) ListSimPoolsInvoker(request *model.ListSimPoolsRequest) *ListSimPoolsInvoker {
	requestDef := GenReqDefForListSimPools()
	return &ListSimPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlowBySimCards 批量查询实体卡流量
//
// 批量查询实体卡流量
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListFlowBySimCards(request *model.ListFlowBySimCardsRequest) (*model.ListFlowBySimCardsResponse, error) {
	requestDef := GenReqDefForListFlowBySimCards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlowBySimCardsResponse), nil
	}
}

// ListFlowBySimCardsInvoker 批量查询实体卡流量
func (c *GslClient) ListFlowBySimCardsInvoker(request *model.ListFlowBySimCardsRequest) *ListFlowBySimCardsInvoker {
	requestDef := GenReqDefForListFlowBySimCards()
	return &ListFlowBySimCardsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSimPricePlans sim卡套餐列表查询
//
// SIM卡套餐列表查询，实体卡只会有一个套餐，eSIM/vSIM可能会有多个套餐
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListSimPricePlans(request *model.ListSimPricePlansRequest) (*model.ListSimPricePlansResponse, error) {
	requestDef := GenReqDefForListSimPricePlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSimPricePlansResponse), nil
	}
}

// ListSimPricePlansInvoker sim卡套餐列表查询
func (c *GslClient) ListSimPricePlansInvoker(request *model.ListSimPricePlansRequest) *ListSimPricePlansInvoker {
	requestDef := GenReqDefForListSimPricePlans()
	return &ListSimPricePlansInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetTags 批量设置/取消设置标签接口
//
// 批量设置/取消设置标签接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) BatchSetTags(request *model.BatchSetTagsRequest) (*model.BatchSetTagsResponse, error) {
	requestDef := GenReqDefForBatchSetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetTagsResponse), nil
	}
}

// BatchSetTagsInvoker 批量设置/取消设置标签接口
func (c *GslClient) BatchSetTagsInvoker(request *model.BatchSetTagsRequest) *BatchSetTagsInvoker {
	requestDef := GenReqDefForBatchSetTags()
	return &BatchSetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 用户添加标签
//
// 添加标签接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 用户添加标签
func (c *GslClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除标签
//
// 删除标签
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除标签
func (c *GslClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询标签列表
//
// 查询标签列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GslClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询标签列表
func (c *GslClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
