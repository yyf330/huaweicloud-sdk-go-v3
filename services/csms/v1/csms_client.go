package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/csms/v1/model"
)

type CsmsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCsmsClient(hcClient *http_client.HcHttpClient) *CsmsClient {
	return &CsmsClient{HcClient: hcClient}
}

func CsmsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateOrDeleteTags 批量添加或删除凭据标签
//
// - 功能介绍：批量添加或删除凭据标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) BatchCreateOrDeleteTags(request *model.BatchCreateOrDeleteTagsRequest) (*model.BatchCreateOrDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteTagsResponse), nil
	}
}

// BatchCreateOrDeleteTagsInvoker 批量添加或删除凭据标签
func (c *CsmsClient) BatchCreateOrDeleteTagsInvoker(request *model.BatchCreateOrDeleteTagsRequest) *BatchCreateOrDeleteTagsInvoker {
	requestDef := GenReqDefForBatchCreateOrDeleteTags()
	return &BatchCreateOrDeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecret 创建凭据
//
// 创建新的凭据，并将凭据值存入凭据的初始版本。
//
// 凭据管理服务将凭据值加密后，存储在凭据对象下的版本中。每个版本可与多个凭据版本状态相关联，凭据版本状态用于标识凭据版本处于的阶段，没有版本状态标记的版本视为已弃用，可用凭据管理服务自动删除。
//
// 初始版本的状态被标记为SYSCURRENT。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) CreateSecret(request *model.CreateSecretRequest) (*model.CreateSecretResponse, error) {
	requestDef := GenReqDefForCreateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretResponse), nil
	}
}

// CreateSecretInvoker 创建凭据
func (c *CsmsClient) CreateSecretInvoker(request *model.CreateSecretRequest) *CreateSecretInvoker {
	requestDef := GenReqDefForCreateSecret()
	return &CreateSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecretTag 添加凭据标签
//
// - 功能介绍：添加凭据标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) CreateSecretTag(request *model.CreateSecretTagRequest) (*model.CreateSecretTagResponse, error) {
	requestDef := GenReqDefForCreateSecretTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretTagResponse), nil
	}
}

// CreateSecretTagInvoker 添加凭据标签
func (c *CsmsClient) CreateSecretTagInvoker(request *model.CreateSecretTagRequest) *CreateSecretTagInvoker {
	requestDef := GenReqDefForCreateSecretTag()
	return &CreateSecretTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecretVersion 创建凭据版本
//
// 在指定的凭据中，创建一个新的凭据版本，用于加密保管新的凭据值。默认情况下，新创建的凭据版本被标记为SYSCURRENT状态，而SYSCURRENT标记的前一个凭据版本被标记为SYSPREVIOUS状态。您可以通过指定VersionStage参数来覆盖默认行为。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) CreateSecretVersion(request *model.CreateSecretVersionRequest) (*model.CreateSecretVersionResponse, error) {
	requestDef := GenReqDefForCreateSecretVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretVersionResponse), nil
	}
}

// CreateSecretVersionInvoker 创建凭据版本
func (c *CsmsClient) CreateSecretVersionInvoker(request *model.CreateSecretVersionRequest) *CreateSecretVersionInvoker {
	requestDef := GenReqDefForCreateSecretVersion()
	return &CreateSecretVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecret 立即删除凭据
//
// 立即删除指定的凭据，且无法恢复。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) DeleteSecret(request *model.DeleteSecretRequest) (*model.DeleteSecretResponse, error) {
	requestDef := GenReqDefForDeleteSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretResponse), nil
	}
}

// DeleteSecretInvoker 立即删除凭据
func (c *CsmsClient) DeleteSecretInvoker(request *model.DeleteSecretRequest) *DeleteSecretInvoker {
	requestDef := GenReqDefForDeleteSecret()
	return &DeleteSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecretForSchedule 创建凭据的定时删除任务
//
// 指定延迟删除时间，创建删除凭据的定时任务，可设置7~30天的的延迟删除时间。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) DeleteSecretForSchedule(request *model.DeleteSecretForScheduleRequest) (*model.DeleteSecretForScheduleResponse, error) {
	requestDef := GenReqDefForDeleteSecretForSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretForScheduleResponse), nil
	}
}

// DeleteSecretForScheduleInvoker 创建凭据的定时删除任务
func (c *CsmsClient) DeleteSecretForScheduleInvoker(request *model.DeleteSecretForScheduleRequest) *DeleteSecretForScheduleInvoker {
	requestDef := GenReqDefForDeleteSecretForSchedule()
	return &DeleteSecretForScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecretStage 删除凭据的版本状态
//
// 删除指定的凭据版本状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) DeleteSecretStage(request *model.DeleteSecretStageRequest) (*model.DeleteSecretStageResponse, error) {
	requestDef := GenReqDefForDeleteSecretStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretStageResponse), nil
	}
}

// DeleteSecretStageInvoker 删除凭据的版本状态
func (c *CsmsClient) DeleteSecretStageInvoker(request *model.DeleteSecretStageRequest) *DeleteSecretStageInvoker {
	requestDef := GenReqDefForDeleteSecretStage()
	return &DeleteSecretStageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecretTag 删除凭据标签
//
// - 功能介绍：删除凭据标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) DeleteSecretTag(request *model.DeleteSecretTagRequest) (*model.DeleteSecretTagResponse, error) {
	requestDef := GenReqDefForDeleteSecretTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretTagResponse), nil
	}
}

// DeleteSecretTagInvoker 删除凭据标签
func (c *CsmsClient) DeleteSecretTagInvoker(request *model.DeleteSecretTagRequest) *DeleteSecretTagInvoker {
	requestDef := GenReqDefForDeleteSecretTag()
	return &DeleteSecretTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectSecretsTags 查询项目标签
//
// - 功能介绍：查询用户在指定项目下的所有凭据标签集合。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) ListProjectSecretsTags(request *model.ListProjectSecretsTagsRequest) (*model.ListProjectSecretsTagsResponse, error) {
	requestDef := GenReqDefForListProjectSecretsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectSecretsTagsResponse), nil
	}
}

// ListProjectSecretsTagsInvoker 查询项目标签
func (c *CsmsClient) ListProjectSecretsTagsInvoker(request *model.ListProjectSecretsTagsRequest) *ListProjectSecretsTagsInvoker {
	requestDef := GenReqDefForListProjectSecretsTags()
	return &ListProjectSecretsTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceInstances 查询凭据实例
//
// - 功能介绍：查询凭据实例。通过标签过滤，筛选用户凭据,返回凭据列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) ListResourceInstances(request *model.ListResourceInstancesRequest) (*model.ListResourceInstancesResponse, error) {
	requestDef := GenReqDefForListResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstancesResponse), nil
	}
}

// ListResourceInstancesInvoker 查询凭据实例
func (c *CsmsClient) ListResourceInstancesInvoker(request *model.ListResourceInstancesRequest) *ListResourceInstancesInvoker {
	requestDef := GenReqDefForListResourceInstances()
	return &ListResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecretTags 查询凭据标签
//
// - 功能介绍：查询凭据标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) ListSecretTags(request *model.ListSecretTagsRequest) (*model.ListSecretTagsResponse, error) {
	requestDef := GenReqDefForListSecretTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecretTagsResponse), nil
	}
}

// ListSecretTagsInvoker 查询凭据标签
func (c *CsmsClient) ListSecretTagsInvoker(request *model.ListSecretTagsRequest) *ListSecretTagsInvoker {
	requestDef := GenReqDefForListSecretTags()
	return &ListSecretTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecretVersions 查询凭据的版本列表
//
// 查询指定凭据下的版本列表信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) ListSecretVersions(request *model.ListSecretVersionsRequest) (*model.ListSecretVersionsResponse, error) {
	requestDef := GenReqDefForListSecretVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecretVersionsResponse), nil
	}
}

// ListSecretVersionsInvoker 查询凭据的版本列表
func (c *CsmsClient) ListSecretVersionsInvoker(request *model.ListSecretVersionsRequest) *ListSecretVersionsInvoker {
	requestDef := GenReqDefForListSecretVersions()
	return &ListSecretVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecrets 查询凭据列表
//
// 查询当前用户在本项目下创建的所有凭据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) ListSecrets(request *model.ListSecretsRequest) (*model.ListSecretsResponse, error) {
	requestDef := GenReqDefForListSecrets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecretsResponse), nil
	}
}

// ListSecretsInvoker 查询凭据列表
func (c *CsmsClient) ListSecretsInvoker(request *model.ListSecretsRequest) *ListSecretsInvoker {
	requestDef := GenReqDefForListSecrets()
	return &ListSecretsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreSecret 取消凭据的定时删除任务
//
// 取消凭据的定时删除任务，凭据对象恢复可使用状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) RestoreSecret(request *model.RestoreSecretRequest) (*model.RestoreSecretResponse, error) {
	requestDef := GenReqDefForRestoreSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreSecretResponse), nil
	}
}

// RestoreSecretInvoker 取消凭据的定时删除任务
func (c *CsmsClient) RestoreSecretInvoker(request *model.RestoreSecretRequest) *RestoreSecretInvoker {
	requestDef := GenReqDefForRestoreSecret()
	return &RestoreSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecret 查询凭据
//
// 查询指定凭据的信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) ShowSecret(request *model.ShowSecretRequest) (*model.ShowSecretResponse, error) {
	requestDef := GenReqDefForShowSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecretResponse), nil
	}
}

// ShowSecretInvoker 查询凭据
func (c *CsmsClient) ShowSecretInvoker(request *model.ShowSecretRequest) *ShowSecretInvoker {
	requestDef := GenReqDefForShowSecret()
	return &ShowSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecretStage 查询凭据的版本状态
//
// 查询指定凭据版本状态标记的版本信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) ShowSecretStage(request *model.ShowSecretStageRequest) (*model.ShowSecretStageResponse, error) {
	requestDef := GenReqDefForShowSecretStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecretStageResponse), nil
	}
}

// ShowSecretStageInvoker 查询凭据的版本状态
func (c *CsmsClient) ShowSecretStageInvoker(request *model.ShowSecretStageRequest) *ShowSecretStageInvoker {
	requestDef := GenReqDefForShowSecretStage()
	return &ShowSecretStageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecretVersion 查询凭据的版本与凭据值
//
// 查询指定凭据版本的信息和版本中的明文凭据值，只能查询ENABLED状态的凭据。
// 通过/v1/{project_id}/secrets/{secret_id}/versions/latest可访问凭据最新版本的凭据值。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) ShowSecretVersion(request *model.ShowSecretVersionRequest) (*model.ShowSecretVersionResponse, error) {
	requestDef := GenReqDefForShowSecretVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecretVersionResponse), nil
	}
}

// ShowSecretVersionInvoker 查询凭据的版本与凭据值
func (c *CsmsClient) ShowSecretVersionInvoker(request *model.ShowSecretVersionRequest) *ShowSecretVersionInvoker {
	requestDef := GenReqDefForShowSecretVersion()
	return &ShowSecretVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecret 更新凭据
//
// 更新指定凭据的元数据信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) UpdateSecret(request *model.UpdateSecretRequest) (*model.UpdateSecretResponse, error) {
	requestDef := GenReqDefForUpdateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecretResponse), nil
	}
}

// UpdateSecretInvoker 更新凭据
func (c *CsmsClient) UpdateSecretInvoker(request *model.UpdateSecretRequest) *UpdateSecretInvoker {
	requestDef := GenReqDefForUpdateSecret()
	return &UpdateSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecretStage 更新凭据的版本状态
//
// 更新凭据的版本状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CsmsClient) UpdateSecretStage(request *model.UpdateSecretStageRequest) (*model.UpdateSecretStageResponse, error) {
	requestDef := GenReqDefForUpdateSecretStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecretStageResponse), nil
	}
}

// UpdateSecretStageInvoker 更新凭据的版本状态
func (c *CsmsClient) UpdateSecretStageInvoker(request *model.UpdateSecretStageRequest) *UpdateSecretStageInvoker {
	requestDef := GenReqDefForUpdateSecretStage()
	return &UpdateSecretStageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
