package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/model"
)

type ElbClient struct {
	HcClient *http_client.HcHttpClient
}

func NewElbClient(hcClient *http_client.HcHttpClient) *ElbClient {
	return &ElbClient{HcClient: hcClient}
}

func ElbClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateMembers 批量创建后端服务器
//
// 在指定pool下批量创建后端服务器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) BatchCreateMembers(request *model.BatchCreateMembersRequest) (*model.BatchCreateMembersResponse, error) {
	requestDef := GenReqDefForBatchCreateMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateMembersResponse), nil
	}
}

// BatchCreateMembersInvoker 批量创建后端服务器
func (c *ElbClient) BatchCreateMembersInvoker(request *model.BatchCreateMembersRequest) *BatchCreateMembersInvoker {
	requestDef := GenReqDefForBatchCreateMembers()
	return &BatchCreateMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteMembers 批量删除后端服务器
//
// 在指定pool下批量删除后端服务器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) BatchDeleteMembers(request *model.BatchDeleteMembersRequest) (*model.BatchDeleteMembersResponse, error) {
	requestDef := GenReqDefForBatchDeleteMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteMembersResponse), nil
	}
}

// BatchDeleteMembersInvoker 批量删除后端服务器
func (c *ElbClient) BatchDeleteMembersInvoker(request *model.BatchDeleteMembersRequest) *BatchDeleteMembersInvoker {
	requestDef := GenReqDefForBatchDeleteMembers()
	return &BatchDeleteMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdatePoliciesPriority 批量更新转发策略优先级
//
// 批量更新转发策略的优先级。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) BatchUpdatePoliciesPriority(request *model.BatchUpdatePoliciesPriorityRequest) (*model.BatchUpdatePoliciesPriorityResponse, error) {
	requestDef := GenReqDefForBatchUpdatePoliciesPriority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdatePoliciesPriorityResponse), nil
	}
}

// BatchUpdatePoliciesPriorityInvoker 批量更新转发策略优先级
func (c *ElbClient) BatchUpdatePoliciesPriorityInvoker(request *model.BatchUpdatePoliciesPriorityRequest) *BatchUpdatePoliciesPriorityInvoker {
	requestDef := GenReqDefForBatchUpdatePoliciesPriority()
	return &BatchUpdatePoliciesPriorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeLoadbalancerChargeMode 负载均衡器计费模式变更
//
// 负载均衡器计费模式变更，当前只支持按需计费转包周期计费。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ChangeLoadbalancerChargeMode(request *model.ChangeLoadbalancerChargeModeRequest) (*model.ChangeLoadbalancerChargeModeResponse, error) {
	requestDef := GenReqDefForChangeLoadbalancerChargeMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeLoadbalancerChargeModeResponse), nil
	}
}

// ChangeLoadbalancerChargeModeInvoker 负载均衡器计费模式变更
func (c *ElbClient) ChangeLoadbalancerChargeModeInvoker(request *model.ChangeLoadbalancerChargeModeRequest) *ChangeLoadbalancerChargeModeInvoker {
	requestDef := GenReqDefForChangeLoadbalancerChargeMode()
	return &ChangeLoadbalancerChargeModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCertificate 创建证书
//
// 创建证书。用于HTTPS协议监听器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := GenReqDefForCreateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateResponse), nil
	}
}

// CreateCertificateInvoker 创建证书
func (c *ElbClient) CreateCertificateInvoker(request *model.CreateCertificateRequest) *CreateCertificateInvoker {
	requestDef := GenReqDefForCreateCertificate()
	return &CreateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHealthMonitor 创建健康检查
//
// 创建健康检查。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateHealthMonitor(request *model.CreateHealthMonitorRequest) (*model.CreateHealthMonitorResponse, error) {
	requestDef := GenReqDefForCreateHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHealthMonitorResponse), nil
	}
}

// CreateHealthMonitorInvoker 创建健康检查
func (c *ElbClient) CreateHealthMonitorInvoker(request *model.CreateHealthMonitorRequest) *CreateHealthMonitorInvoker {
	requestDef := GenReqDefForCreateHealthMonitor()
	return &CreateHealthMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateL7Policy 创建转发策略
//
// 创建七层转发策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateL7Policy(request *model.CreateL7PolicyRequest) (*model.CreateL7PolicyResponse, error) {
	requestDef := GenReqDefForCreateL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateL7PolicyResponse), nil
	}
}

// CreateL7PolicyInvoker 创建转发策略
func (c *ElbClient) CreateL7PolicyInvoker(request *model.CreateL7PolicyRequest) *CreateL7PolicyInvoker {
	requestDef := GenReqDefForCreateL7Policy()
	return &CreateL7PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateL7Rule 创建转发规则
//
// 创建七层转发规则。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateL7Rule(request *model.CreateL7RuleRequest) (*model.CreateL7RuleResponse, error) {
	requestDef := GenReqDefForCreateL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateL7RuleResponse), nil
	}
}

// CreateL7RuleInvoker 创建转发规则
func (c *ElbClient) CreateL7RuleInvoker(request *model.CreateL7RuleRequest) *CreateL7RuleInvoker {
	requestDef := GenReqDefForCreateL7Rule()
	return &CreateL7RuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateListener 创建监听器
//
// 创建监听器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateListener(request *model.CreateListenerRequest) (*model.CreateListenerResponse, error) {
	requestDef := GenReqDefForCreateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateListenerResponse), nil
	}
}

// CreateListenerInvoker 创建监听器
func (c *ElbClient) CreateListenerInvoker(request *model.CreateListenerRequest) *CreateListenerInvoker {
	requestDef := GenReqDefForCreateListener()
	return &CreateListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLoadBalancer 创建负载均衡器
//
// 创建负载均衡器。
//
// 1.若要创建内网IPv4负载均衡器，则需要设置vip_subnet_cidr_id。
//
// 2.若要创建公网IPv4负载均衡器，则需要设置publicip，以及设置vpc_id和vip_subnet_cidr_id这两个参数中的一个。
//
// 3.若要绑定有已有公网IPv4地址，需要设置publicip_ids，以及设置vpc_id和vip_subnet_cidr_id这两个参数中的一个。
//
// 4.若要创建内网双栈负载均衡器，则需要设置ipv6_vip_virsubnet_id。
//
// 5.若要创建公网双栈负载均衡器，则需要设置ipv6_vip_virsubnet_id和ipv6_bandwidth。
//
// 6.不支持绑定已有未使用的内网IPv4、内网IPv6或公网IPv6地址。
//
// [&gt;不支持创建IPv6地址负载均衡器](tag:dt,dt_test)
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateLoadBalancer(request *model.CreateLoadBalancerRequest) (*model.CreateLoadBalancerResponse, error) {
	requestDef := GenReqDefForCreateLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLoadBalancerResponse), nil
	}
}

// CreateLoadBalancerInvoker 创建负载均衡器
func (c *ElbClient) CreateLoadBalancerInvoker(request *model.CreateLoadBalancerRequest) *CreateLoadBalancerInvoker {
	requestDef := GenReqDefForCreateLoadBalancer()
	return &CreateLoadBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLogtank 创建云日志
//
// 创建云日志
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateLogtank(request *model.CreateLogtankRequest) (*model.CreateLogtankResponse, error) {
	requestDef := GenReqDefForCreateLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogtankResponse), nil
	}
}

// CreateLogtankInvoker 创建云日志
func (c *ElbClient) CreateLogtankInvoker(request *model.CreateLogtankRequest) *CreateLogtankInvoker {
	requestDef := GenReqDefForCreateLogtank()
	return &CreateLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMember 创建后端服务器
//
// 创建后端服务器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateMember(request *model.CreateMemberRequest) (*model.CreateMemberResponse, error) {
	requestDef := GenReqDefForCreateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMemberResponse), nil
	}
}

// CreateMemberInvoker 创建后端服务器
func (c *ElbClient) CreateMemberInvoker(request *model.CreateMemberRequest) *CreateMemberInvoker {
	requestDef := GenReqDefForCreateMember()
	return &CreateMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePool 创建后端服务器组
//
// 创建后端服务器组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreatePool(request *model.CreatePoolRequest) (*model.CreatePoolResponse, error) {
	requestDef := GenReqDefForCreatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePoolResponse), nil
	}
}

// CreatePoolInvoker 创建后端服务器组
func (c *ElbClient) CreatePoolInvoker(request *model.CreatePoolRequest) *CreatePoolInvoker {
	requestDef := GenReqDefForCreatePool()
	return &CreatePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityPolicy 创建自定义安全策略
//
// 创建自定义安全策略。用于在创建HTTPS监听器时，请求参数中指定security_policy_id来设置监听器的自定义安全策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateSecurityPolicy(request *model.CreateSecurityPolicyRequest) (*model.CreateSecurityPolicyResponse, error) {
	requestDef := GenReqDefForCreateSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityPolicyResponse), nil
	}
}

// CreateSecurityPolicyInvoker 创建自定义安全策略
func (c *ElbClient) CreateSecurityPolicyInvoker(request *model.CreateSecurityPolicyRequest) *CreateSecurityPolicyInvoker {
	requestDef := GenReqDefForCreateSecurityPolicy()
	return &CreateSecurityPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCertificate 删除证书
//
// 删除证书。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

// DeleteCertificateInvoker 删除证书
func (c *ElbClient) DeleteCertificateInvoker(request *model.DeleteCertificateRequest) *DeleteCertificateInvoker {
	requestDef := GenReqDefForDeleteCertificate()
	return &DeleteCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHealthMonitor 删除健康检查
//
// 删除健康检查。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteHealthMonitor(request *model.DeleteHealthMonitorRequest) (*model.DeleteHealthMonitorResponse, error) {
	requestDef := GenReqDefForDeleteHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHealthMonitorResponse), nil
	}
}

// DeleteHealthMonitorInvoker 删除健康检查
func (c *ElbClient) DeleteHealthMonitorInvoker(request *model.DeleteHealthMonitorRequest) *DeleteHealthMonitorInvoker {
	requestDef := GenReqDefForDeleteHealthMonitor()
	return &DeleteHealthMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteL7Policy 删除转发策略
//
// 删除七层转发策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteL7Policy(request *model.DeleteL7PolicyRequest) (*model.DeleteL7PolicyResponse, error) {
	requestDef := GenReqDefForDeleteL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteL7PolicyResponse), nil
	}
}

// DeleteL7PolicyInvoker 删除转发策略
func (c *ElbClient) DeleteL7PolicyInvoker(request *model.DeleteL7PolicyRequest) *DeleteL7PolicyInvoker {
	requestDef := GenReqDefForDeleteL7Policy()
	return &DeleteL7PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteL7Rule 删除转发规则
//
// 删除七层转发规则。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteL7Rule(request *model.DeleteL7RuleRequest) (*model.DeleteL7RuleResponse, error) {
	requestDef := GenReqDefForDeleteL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteL7RuleResponse), nil
	}
}

// DeleteL7RuleInvoker 删除转发规则
func (c *ElbClient) DeleteL7RuleInvoker(request *model.DeleteL7RuleRequest) *DeleteL7RuleInvoker {
	requestDef := GenReqDefForDeleteL7Rule()
	return &DeleteL7RuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteListener 删除监听器
//
// 删除监听器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteListener(request *model.DeleteListenerRequest) (*model.DeleteListenerResponse, error) {
	requestDef := GenReqDefForDeleteListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteListenerResponse), nil
	}
}

// DeleteListenerInvoker 删除监听器
func (c *ElbClient) DeleteListenerInvoker(request *model.DeleteListenerRequest) *DeleteListenerInvoker {
	requestDef := GenReqDefForDeleteListener()
	return &DeleteListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLoadBalancer 删除负载均衡器
//
// 删除负载均衡器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteLoadBalancer(request *model.DeleteLoadBalancerRequest) (*model.DeleteLoadBalancerResponse, error) {
	requestDef := GenReqDefForDeleteLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLoadBalancerResponse), nil
	}
}

// DeleteLoadBalancerInvoker 删除负载均衡器
func (c *ElbClient) DeleteLoadBalancerInvoker(request *model.DeleteLoadBalancerRequest) *DeleteLoadBalancerInvoker {
	requestDef := GenReqDefForDeleteLoadBalancer()
	return &DeleteLoadBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLogtank 删除云日志
//
// 删除云日志。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteLogtank(request *model.DeleteLogtankRequest) (*model.DeleteLogtankResponse, error) {
	requestDef := GenReqDefForDeleteLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogtankResponse), nil
	}
}

// DeleteLogtankInvoker 删除云日志
func (c *ElbClient) DeleteLogtankInvoker(request *model.DeleteLogtankRequest) *DeleteLogtankInvoker {
	requestDef := GenReqDefForDeleteLogtank()
	return &DeleteLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMember 删除后端服务器
//
// 删除后端服务器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteMember(request *model.DeleteMemberRequest) (*model.DeleteMemberResponse, error) {
	requestDef := GenReqDefForDeleteMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberResponse), nil
	}
}

// DeleteMemberInvoker 删除后端服务器
func (c *ElbClient) DeleteMemberInvoker(request *model.DeleteMemberRequest) *DeleteMemberInvoker {
	requestDef := GenReqDefForDeleteMember()
	return &DeleteMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePool 删除后端服务器组
//
// 删除后端服务器组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeletePool(request *model.DeletePoolRequest) (*model.DeletePoolResponse, error) {
	requestDef := GenReqDefForDeletePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePoolResponse), nil
	}
}

// DeletePoolInvoker 删除后端服务器组
func (c *ElbClient) DeletePoolInvoker(request *model.DeletePoolRequest) *DeletePoolInvoker {
	requestDef := GenReqDefForDeletePool()
	return &DeletePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecurityPolicy 删除自定义安全策略
//
// 删除自定义安全策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteSecurityPolicy(request *model.DeleteSecurityPolicyRequest) (*model.DeleteSecurityPolicyResponse, error) {
	requestDef := GenReqDefForDeleteSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityPolicyResponse), nil
	}
}

// DeleteSecurityPolicyInvoker 删除自定义安全策略
func (c *ElbClient) DeleteSecurityPolicyInvoker(request *model.DeleteSecurityPolicyRequest) *DeleteSecurityPolicyInvoker {
	requestDef := GenReqDefForDeleteSecurityPolicy()
	return &DeleteSecurityPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllMembers 后端服务器全局列表
//
// 查询当前租户下的后端服务器列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListAllMembers(request *model.ListAllMembersRequest) (*model.ListAllMembersResponse, error) {
	requestDef := GenReqDefForListAllMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllMembersResponse), nil
	}
}

// ListAllMembersInvoker 后端服务器全局列表
func (c *ElbClient) ListAllMembersInvoker(request *model.ListAllMembersRequest) *ListAllMembersInvoker {
	requestDef := GenReqDefForListAllMembers()
	return &ListAllMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailabilityZones 查询可用区列表
//
// 返回租户创建LB时可使用的可用区集合列表情况。
//
// 默认情况下，会返回一个可用区集合。在（如创建LB）设置可用区时，填写的可用区必须包含在可用区集合中、为这个可用区集合的子集。
//
// [特殊场景下，部分客户要求负载均衡只能创建在指定可用区集合中，此时会返回客户定制的可用区集合。返回可用区集合可能为一个也可能为多个，比如列表有两个可用区集合[az1,az2],
// [az2,az3]。在创建负载均衡器时，可以选择创建在多个可用区，但所选的多个可用区必须同属于其中一个可用区集合，如可以选az2和az3，但不能选择az1和az3。你可以选择多个可用区，只要这些可用区在一个子集中](tag:hws,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42)
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListAvailabilityZones(request *model.ListAvailabilityZonesRequest) (*model.ListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZonesResponse), nil
	}
}

// ListAvailabilityZonesInvoker 查询可用区列表
func (c *ElbClient) ListAvailabilityZonesInvoker(request *model.ListAvailabilityZonesRequest) *ListAvailabilityZonesInvoker {
	requestDef := GenReqDefForListAvailabilityZones()
	return &ListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCertificates 证书列表
//
// 查询证书列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

// ListCertificatesInvoker 证书列表
func (c *ElbClient) ListCertificatesInvoker(request *model.ListCertificatesRequest) *ListCertificatesInvoker {
	requestDef := GenReqDefForListCertificates()
	return &ListCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询规格列表
//
// 查询租户在当前region下可用的负载均衡规格列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询规格列表
func (c *ElbClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHealthMonitors 查询健康检查列表
//
// 健康检查列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListHealthMonitors(request *model.ListHealthMonitorsRequest) (*model.ListHealthMonitorsResponse, error) {
	requestDef := GenReqDefForListHealthMonitors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHealthMonitorsResponse), nil
	}
}

// ListHealthMonitorsInvoker 查询健康检查列表
func (c *ElbClient) ListHealthMonitorsInvoker(request *model.ListHealthMonitorsRequest) *ListHealthMonitorsInvoker {
	requestDef := GenReqDefForListHealthMonitors()
	return &ListHealthMonitorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListL7Policies 查询转发策略列表
//
// 查询七层转发策略列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListL7Policies(request *model.ListL7PoliciesRequest) (*model.ListL7PoliciesResponse, error) {
	requestDef := GenReqDefForListL7Policies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListL7PoliciesResponse), nil
	}
}

// ListL7PoliciesInvoker 查询转发策略列表
func (c *ElbClient) ListL7PoliciesInvoker(request *model.ListL7PoliciesRequest) *ListL7PoliciesInvoker {
	requestDef := GenReqDefForListL7Policies()
	return &ListL7PoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListL7Rules 查询转发规则列表
//
// 查询转发规则列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListL7Rules(request *model.ListL7RulesRequest) (*model.ListL7RulesResponse, error) {
	requestDef := GenReqDefForListL7Rules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListL7RulesResponse), nil
	}
}

// ListL7RulesInvoker 查询转发规则列表
func (c *ElbClient) ListL7RulesInvoker(request *model.ListL7RulesRequest) *ListL7RulesInvoker {
	requestDef := GenReqDefForListL7Rules()
	return &ListL7RulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListListeners 查询监听器列表
//
// 查询监听器列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListListeners(request *model.ListListenersRequest) (*model.ListListenersResponse, error) {
	requestDef := GenReqDefForListListeners()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListListenersResponse), nil
	}
}

// ListListenersInvoker 查询监听器列表
func (c *ElbClient) ListListenersInvoker(request *model.ListListenersRequest) *ListListenersInvoker {
	requestDef := GenReqDefForListListeners()
	return &ListListenersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLoadBalancers 查询负载均衡器列表
//
// 查询负载均衡器列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListLoadBalancers(request *model.ListLoadBalancersRequest) (*model.ListLoadBalancersResponse, error) {
	requestDef := GenReqDefForListLoadBalancers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoadBalancersResponse), nil
	}
}

// ListLoadBalancersInvoker 查询负载均衡器列表
func (c *ElbClient) ListLoadBalancersInvoker(request *model.ListLoadBalancersRequest) *ListLoadBalancersInvoker {
	requestDef := GenReqDefForListLoadBalancers()
	return &ListLoadBalancersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogtanks 云日志列表
//
// 云日志列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListLogtanks(request *model.ListLogtanksRequest) (*model.ListLogtanksResponse, error) {
	requestDef := GenReqDefForListLogtanks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogtanksResponse), nil
	}
}

// ListLogtanksInvoker 云日志列表
func (c *ElbClient) ListLogtanksInvoker(request *model.ListLogtanksRequest) *ListLogtanksInvoker {
	requestDef := GenReqDefForListLogtanks()
	return &ListLogtanksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMembers 后端服务器列表
//
// Pool下的后端服务器列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListMembers(request *model.ListMembersRequest) (*model.ListMembersResponse, error) {
	requestDef := GenReqDefForListMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMembersResponse), nil
	}
}

// ListMembersInvoker 后端服务器列表
func (c *ElbClient) ListMembersInvoker(request *model.ListMembersRequest) *ListMembersInvoker {
	requestDef := GenReqDefForListMembers()
	return &ListMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPools 查询后端服务器组列表
//
// 后端服务器组列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListPools(request *model.ListPoolsRequest) (*model.ListPoolsResponse, error) {
	requestDef := GenReqDefForListPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoolsResponse), nil
	}
}

// ListPoolsInvoker 查询后端服务器组列表
func (c *ElbClient) ListPoolsInvoker(request *model.ListPoolsRequest) *ListPoolsInvoker {
	requestDef := GenReqDefForListPools()
	return &ListPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotaDetails 查询配额使用详情
//
// 查询指定项目中负载均衡相关的各类资源的当前配额和已使用配额信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListQuotaDetails(request *model.ListQuotaDetailsRequest) (*model.ListQuotaDetailsResponse, error) {
	requestDef := GenReqDefForListQuotaDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaDetailsResponse), nil
	}
}

// ListQuotaDetailsInvoker 查询配额使用详情
func (c *ElbClient) ListQuotaDetailsInvoker(request *model.ListQuotaDetailsRequest) *ListQuotaDetailsInvoker {
	requestDef := GenReqDefForListQuotaDetails()
	return &ListQuotaDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityPolicies 查询自定义安全策略列表
//
// 查询自定义安全策略列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListSecurityPolicies(request *model.ListSecurityPoliciesRequest) (*model.ListSecurityPoliciesResponse, error) {
	requestDef := GenReqDefForListSecurityPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityPoliciesResponse), nil
	}
}

// ListSecurityPoliciesInvoker 查询自定义安全策略列表
func (c *ElbClient) ListSecurityPoliciesInvoker(request *model.ListSecurityPoliciesRequest) *ListSecurityPoliciesInvoker {
	requestDef := GenReqDefForListSecurityPolicies()
	return &ListSecurityPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSystemSecurityPolicies 查询系统安全策略列表
//
// 查询系统安全策略列表。
//
// 系统安全策略为预置的所有租户通用的安全策略，租户不可新增或修改。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListSystemSecurityPolicies(request *model.ListSystemSecurityPoliciesRequest) (*model.ListSystemSecurityPoliciesResponse, error) {
	requestDef := GenReqDefForListSystemSecurityPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSystemSecurityPoliciesResponse), nil
	}
}

// ListSystemSecurityPoliciesInvoker 查询系统安全策略列表
func (c *ElbClient) ListSystemSecurityPoliciesInvoker(request *model.ListSystemSecurityPoliciesRequest) *ListSystemSecurityPoliciesInvoker {
	requestDef := GenReqDefForListSystemSecurityPolicies()
	return &ListSystemSecurityPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificate 证书详情
//
// 查询证书详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}

// ShowCertificateInvoker 证书详情
func (c *ElbClient) ShowCertificateInvoker(request *model.ShowCertificateRequest) *ShowCertificateInvoker {
	requestDef := GenReqDefForShowCertificate()
	return &ShowCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlavor 查询规格详情
//
// 查询规格的详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowFlavor(request *model.ShowFlavorRequest) (*model.ShowFlavorResponse, error) {
	requestDef := GenReqDefForShowFlavor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlavorResponse), nil
	}
}

// ShowFlavorInvoker 查询规格详情
func (c *ElbClient) ShowFlavorInvoker(request *model.ShowFlavorRequest) *ShowFlavorInvoker {
	requestDef := GenReqDefForShowFlavor()
	return &ShowFlavorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHealthMonitor 查询健康检查详情
//
// 查询健康检查详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowHealthMonitor(request *model.ShowHealthMonitorRequest) (*model.ShowHealthMonitorResponse, error) {
	requestDef := GenReqDefForShowHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHealthMonitorResponse), nil
	}
}

// ShowHealthMonitorInvoker 查询健康检查详情
func (c *ElbClient) ShowHealthMonitorInvoker(request *model.ShowHealthMonitorRequest) *ShowHealthMonitorInvoker {
	requestDef := GenReqDefForShowHealthMonitor()
	return &ShowHealthMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowL7Policy 查询转发策略详情
//
// 查询七层转发策略详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowL7Policy(request *model.ShowL7PolicyRequest) (*model.ShowL7PolicyResponse, error) {
	requestDef := GenReqDefForShowL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowL7PolicyResponse), nil
	}
}

// ShowL7PolicyInvoker 查询转发策略详情
func (c *ElbClient) ShowL7PolicyInvoker(request *model.ShowL7PolicyRequest) *ShowL7PolicyInvoker {
	requestDef := GenReqDefForShowL7Policy()
	return &ShowL7PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowL7Rule 查询转发规则详情
//
// 查询七层转发规则详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowL7Rule(request *model.ShowL7RuleRequest) (*model.ShowL7RuleResponse, error) {
	requestDef := GenReqDefForShowL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowL7RuleResponse), nil
	}
}

// ShowL7RuleInvoker 查询转发规则详情
func (c *ElbClient) ShowL7RuleInvoker(request *model.ShowL7RuleRequest) *ShowL7RuleInvoker {
	requestDef := GenReqDefForShowL7Rule()
	return &ShowL7RuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowListener 查询监听器详情
//
// 监听器详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowListener(request *model.ShowListenerRequest) (*model.ShowListenerResponse, error) {
	requestDef := GenReqDefForShowListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListenerResponse), nil
	}
}

// ShowListenerInvoker 查询监听器详情
func (c *ElbClient) ShowListenerInvoker(request *model.ShowListenerRequest) *ShowListenerInvoker {
	requestDef := GenReqDefForShowListener()
	return &ShowListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLoadBalancer 查询负载均衡器详情
//
// 查询负载均衡器详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowLoadBalancer(request *model.ShowLoadBalancerRequest) (*model.ShowLoadBalancerResponse, error) {
	requestDef := GenReqDefForShowLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadBalancerResponse), nil
	}
}

// ShowLoadBalancerInvoker 查询负载均衡器详情
func (c *ElbClient) ShowLoadBalancerInvoker(request *model.ShowLoadBalancerRequest) *ShowLoadBalancerInvoker {
	requestDef := GenReqDefForShowLoadBalancer()
	return &ShowLoadBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLoadBalancerStatus 查询负载均衡器状态树
//
// 查询负载均衡器状态树，包括负载均衡器及其关联的子资源的状态信息。
//
// 注意：该接口中的operating_status不一定与对应资源的operating_status相同。如：当Member的admin_state_up&#x3D;false且operating_status&#x3D;OFFLINE时，该接口返回member的operating_status&#x3D;DISABLE。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowLoadBalancerStatus(request *model.ShowLoadBalancerStatusRequest) (*model.ShowLoadBalancerStatusResponse, error) {
	requestDef := GenReqDefForShowLoadBalancerStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadBalancerStatusResponse), nil
	}
}

// ShowLoadBalancerStatusInvoker 查询负载均衡器状态树
func (c *ElbClient) ShowLoadBalancerStatusInvoker(request *model.ShowLoadBalancerStatusRequest) *ShowLoadBalancerStatusInvoker {
	requestDef := GenReqDefForShowLoadBalancerStatus()
	return &ShowLoadBalancerStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogtank 云日志配置详情
//
// 云日志详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowLogtank(request *model.ShowLogtankRequest) (*model.ShowLogtankResponse, error) {
	requestDef := GenReqDefForShowLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogtankResponse), nil
	}
}

// ShowLogtankInvoker 云日志配置详情
func (c *ElbClient) ShowLogtankInvoker(request *model.ShowLogtankRequest) *ShowLogtankInvoker {
	requestDef := GenReqDefForShowLogtank()
	return &ShowLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMember 后端服务器详情
//
// 后端服务器详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowMember(request *model.ShowMemberRequest) (*model.ShowMemberResponse, error) {
	requestDef := GenReqDefForShowMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMemberResponse), nil
	}
}

// ShowMemberInvoker 后端服务器详情
func (c *ElbClient) ShowMemberInvoker(request *model.ShowMemberRequest) *ShowMemberInvoker {
	requestDef := GenReqDefForShowMember()
	return &ShowMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPool 查询后端服务器组详情
//
// 后端服务器组详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowPool(request *model.ShowPoolRequest) (*model.ShowPoolResponse, error) {
	requestDef := GenReqDefForShowPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolResponse), nil
	}
}

// ShowPoolInvoker 查询后端服务器组详情
func (c *ElbClient) ShowPoolInvoker(request *model.ShowPoolRequest) *ShowPoolInvoker {
	requestDef := GenReqDefForShowPool()
	return &ShowPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuota 查询配额
//
// 查询指定项目中负载均衡相关的各类资源的当前配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

// ShowQuotaInvoker 查询配额
func (c *ElbClient) ShowQuotaInvoker(request *model.ShowQuotaRequest) *ShowQuotaInvoker {
	requestDef := GenReqDefForShowQuota()
	return &ShowQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecurityPolicy 查询自定义安全策略详情
//
// 查询自定义安全策略详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowSecurityPolicy(request *model.ShowSecurityPolicyRequest) (*model.ShowSecurityPolicyResponse, error) {
	requestDef := GenReqDefForShowSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityPolicyResponse), nil
	}
}

// ShowSecurityPolicyInvoker 查询自定义安全策略详情
func (c *ElbClient) ShowSecurityPolicyInvoker(request *model.ShowSecurityPolicyRequest) *ShowSecurityPolicyInvoker {
	requestDef := GenReqDefForShowSecurityPolicy()
	return &ShowSecurityPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCertificate 更新证书
//
// 更新证书。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateCertificate(request *model.UpdateCertificateRequest) (*model.UpdateCertificateResponse, error) {
	requestDef := GenReqDefForUpdateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCertificateResponse), nil
	}
}

// UpdateCertificateInvoker 更新证书
func (c *ElbClient) UpdateCertificateInvoker(request *model.UpdateCertificateRequest) *UpdateCertificateInvoker {
	requestDef := GenReqDefForUpdateCertificate()
	return &UpdateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHealthMonitor 更新健康检查
//
// 更新健康检查。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateHealthMonitor(request *model.UpdateHealthMonitorRequest) (*model.UpdateHealthMonitorResponse, error) {
	requestDef := GenReqDefForUpdateHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHealthMonitorResponse), nil
	}
}

// UpdateHealthMonitorInvoker 更新健康检查
func (c *ElbClient) UpdateHealthMonitorInvoker(request *model.UpdateHealthMonitorRequest) *UpdateHealthMonitorInvoker {
	requestDef := GenReqDefForUpdateHealthMonitor()
	return &UpdateHealthMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateL7Policy 更新转发策略
//
// 更新七层转发策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateL7Policy(request *model.UpdateL7PolicyRequest) (*model.UpdateL7PolicyResponse, error) {
	requestDef := GenReqDefForUpdateL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateL7PolicyResponse), nil
	}
}

// UpdateL7PolicyInvoker 更新转发策略
func (c *ElbClient) UpdateL7PolicyInvoker(request *model.UpdateL7PolicyRequest) *UpdateL7PolicyInvoker {
	requestDef := GenReqDefForUpdateL7Policy()
	return &UpdateL7PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateL7Rule 更新转发规则
//
// 更新七层转发规则。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateL7Rule(request *model.UpdateL7RuleRequest) (*model.UpdateL7RuleResponse, error) {
	requestDef := GenReqDefForUpdateL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateL7RuleResponse), nil
	}
}

// UpdateL7RuleInvoker 更新转发规则
func (c *ElbClient) UpdateL7RuleInvoker(request *model.UpdateL7RuleRequest) *UpdateL7RuleInvoker {
	requestDef := GenReqDefForUpdateL7Rule()
	return &UpdateL7RuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateListener 更新监听器
//
// 更新监听器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateListener(request *model.UpdateListenerRequest) (*model.UpdateListenerResponse, error) {
	requestDef := GenReqDefForUpdateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateListenerResponse), nil
	}
}

// UpdateListenerInvoker 更新监听器
func (c *ElbClient) UpdateListenerInvoker(request *model.UpdateListenerRequest) *UpdateListenerInvoker {
	requestDef := GenReqDefForUpdateListener()
	return &UpdateListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLoadBalancer 更新负载均衡器
//
// 更新负载均衡器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateLoadBalancer(request *model.UpdateLoadBalancerRequest) (*model.UpdateLoadBalancerResponse, error) {
	requestDef := GenReqDefForUpdateLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLoadBalancerResponse), nil
	}
}

// UpdateLoadBalancerInvoker 更新负载均衡器
func (c *ElbClient) UpdateLoadBalancerInvoker(request *model.UpdateLoadBalancerRequest) *UpdateLoadBalancerInvoker {
	requestDef := GenReqDefForUpdateLoadBalancer()
	return &UpdateLoadBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLogtank 更新云日志
//
// 更新云日志
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateLogtank(request *model.UpdateLogtankRequest) (*model.UpdateLogtankResponse, error) {
	requestDef := GenReqDefForUpdateLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLogtankResponse), nil
	}
}

// UpdateLogtankInvoker 更新云日志
func (c *ElbClient) UpdateLogtankInvoker(request *model.UpdateLogtankRequest) *UpdateLogtankInvoker {
	requestDef := GenReqDefForUpdateLogtank()
	return &UpdateLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMember 更新后端服务器
//
// 更新后端服务器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateMember(request *model.UpdateMemberRequest) (*model.UpdateMemberResponse, error) {
	requestDef := GenReqDefForUpdateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberResponse), nil
	}
}

// UpdateMemberInvoker 更新后端服务器
func (c *ElbClient) UpdateMemberInvoker(request *model.UpdateMemberRequest) *UpdateMemberInvoker {
	requestDef := GenReqDefForUpdateMember()
	return &UpdateMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePool 更新后端服务器组
//
// 更新后端服务器组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdatePool(request *model.UpdatePoolRequest) (*model.UpdatePoolResponse, error) {
	requestDef := GenReqDefForUpdatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePoolResponse), nil
	}
}

// UpdatePoolInvoker 更新后端服务器组
func (c *ElbClient) UpdatePoolInvoker(request *model.UpdatePoolRequest) *UpdatePoolInvoker {
	requestDef := GenReqDefForUpdatePool()
	return &UpdatePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityPolicy 更新自定义安全策略
//
// 更新自定义安全策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateSecurityPolicy(request *model.UpdateSecurityPolicyRequest) (*model.UpdateSecurityPolicyResponse, error) {
	requestDef := GenReqDefForUpdateSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityPolicyResponse), nil
	}
}

// UpdateSecurityPolicyInvoker 更新自定义安全策略
func (c *ElbClient) UpdateSecurityPolicyInvoker(request *model.UpdateSecurityPolicyRequest) *UpdateSecurityPolicyInvoker {
	requestDef := GenReqDefForUpdateSecurityPolicy()
	return &UpdateSecurityPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersions 查询API版本列表信息
//
// 返回ELB当前所有可用的API版本。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// ListApiVersionsInvoker 查询API版本列表信息
func (c *ElbClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteIpList 删除IP地址组的IP列表项
//
// 批量删除IP地址组的IP列表项。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) BatchDeleteIpList(request *model.BatchDeleteIpListRequest) (*model.BatchDeleteIpListResponse, error) {
	requestDef := GenReqDefForBatchDeleteIpList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteIpListResponse), nil
	}
}

// BatchDeleteIpListInvoker 删除IP地址组的IP列表项
func (c *ElbClient) BatchDeleteIpListInvoker(request *model.BatchDeleteIpListRequest) *BatchDeleteIpListInvoker {
	requestDef := GenReqDefForBatchDeleteIpList()
	return &BatchDeleteIpListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountPreoccupyIpNum 计算预占IP数
//
// 计算以下几种场景的预占用IP数量：
// - 计算创建LB的第一个七层监听器后总占用IP数量：传入loadbalancer_id、l7_flavor_id为空、ip_target_enable不传或为false。
// - 计算LB规格变更或开启跨VPC后总占用IP数量：传入参数loadbalancer_id，及l7_flavor_id不为空或ip_target_enable为true。
// - 计算创建LB所需IP数量：传入参数availability_zone_id，及可选参数l7_flavor_id、ip_target_enable、ip_version，不能传loadbalancer_id。
// &gt; 查询出的预占IP数大于等于最终实际占用的IP数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CountPreoccupyIpNum(request *model.CountPreoccupyIpNumRequest) (*model.CountPreoccupyIpNumResponse, error) {
	requestDef := GenReqDefForCountPreoccupyIpNum()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountPreoccupyIpNumResponse), nil
	}
}

// CountPreoccupyIpNumInvoker 计算预占IP数
func (c *ElbClient) CountPreoccupyIpNumInvoker(request *model.CountPreoccupyIpNumRequest) *CountPreoccupyIpNumInvoker {
	requestDef := GenReqDefForCountPreoccupyIpNum()
	return &CountPreoccupyIpNumInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIpGroup 创建IP地址组
//
// 创建IP地址组。输入的ip可为ip地址或者CIDR子网，支持IPV4和IPV6。需要注意，0.0.0.0与0.0.0.0/32视为重复，0:0:0:0:0:0:0:1与::1与::1/128视为重复，会只保留其中一个写入。
// [不支持IPv6，请勿传入IPv6地址。](tag:dt,dt_test)
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateIpGroup(request *model.CreateIpGroupRequest) (*model.CreateIpGroupResponse, error) {
	requestDef := GenReqDefForCreateIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIpGroupResponse), nil
	}
}

// CreateIpGroupInvoker 创建IP地址组
func (c *ElbClient) CreateIpGroupInvoker(request *model.CreateIpGroupRequest) *CreateIpGroupInvoker {
	requestDef := GenReqDefForCreateIpGroup()
	return &CreateIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIpGroup 删除IP地址组
//
// 删除ip地址组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteIpGroup(request *model.DeleteIpGroupRequest) (*model.DeleteIpGroupResponse, error) {
	requestDef := GenReqDefForDeleteIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIpGroupResponse), nil
	}
}

// DeleteIpGroupInvoker 删除IP地址组
func (c *ElbClient) DeleteIpGroupInvoker(request *model.DeleteIpGroupRequest) *DeleteIpGroupInvoker {
	requestDef := GenReqDefForDeleteIpGroup()
	return &DeleteIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpGroups 查询IP地址组列表
//
// 查询IP地址组列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListIpGroups(request *model.ListIpGroupsRequest) (*model.ListIpGroupsResponse, error) {
	requestDef := GenReqDefForListIpGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpGroupsResponse), nil
	}
}

// ListIpGroupsInvoker 查询IP地址组列表
func (c *ElbClient) ListIpGroupsInvoker(request *model.ListIpGroupsRequest) *ListIpGroupsInvoker {
	requestDef := GenReqDefForListIpGroups()
	return &ListIpGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIpGroup 查询IP地址组详情
//
// 获取IP地址组详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowIpGroup(request *model.ShowIpGroupRequest) (*model.ShowIpGroupResponse, error) {
	requestDef := GenReqDefForShowIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpGroupResponse), nil
	}
}

// ShowIpGroupInvoker 查询IP地址组详情
func (c *ElbClient) ShowIpGroupInvoker(request *model.ShowIpGroupRequest) *ShowIpGroupInvoker {
	requestDef := GenReqDefForShowIpGroup()
	return &ShowIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIpGroup 更新IP地址组
//
// 更新IP地址组，只支持全量更新IP。即IP地址组中的ip_list将被全量覆盖，不在请求参数中的IP地址将被移除。输入的ip可为ip地址或者CIDR子网，支持IPV4和IPV6。需要注意，0.0.0.0与0.0.0.0/32视为重复，0:0:0:0:0:0:0:1与::1与::1/128视为重复，会只保留其中一个写入。
// [不支持IPv6，请勿传入IPv6地址。](tag:dt,dt_test)
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateIpGroup(request *model.UpdateIpGroupRequest) (*model.UpdateIpGroupResponse, error) {
	requestDef := GenReqDefForUpdateIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpGroupResponse), nil
	}
}

// UpdateIpGroupInvoker 更新IP地址组
func (c *ElbClient) UpdateIpGroupInvoker(request *model.UpdateIpGroupRequest) *UpdateIpGroupInvoker {
	requestDef := GenReqDefForUpdateIpGroup()
	return &UpdateIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIpList 更新IP地址组的IP列表项
//
// 更新IP地址组的IP列表信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateIpList(request *model.UpdateIpListRequest) (*model.UpdateIpListResponse, error) {
	requestDef := GenReqDefForUpdateIpList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpListResponse), nil
	}
}

// UpdateIpListInvoker 更新IP地址组的IP列表项
func (c *ElbClient) UpdateIpListInvoker(request *model.UpdateIpListRequest) *UpdateIpListInvoker {
	requestDef := GenReqDefForUpdateIpList()
	return &UpdateIpListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
