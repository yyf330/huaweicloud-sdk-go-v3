package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kms/v1/model"
)

type KmsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewKmsClient(hcClient *http_client.HcHttpClient) *KmsClient {
	return &KmsClient{HcClient: hcClient}
}

func KmsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateKmsTags 批量添加删除密钥标签
//
// - 功能介绍：批量添加删除密钥标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) BatchCreateKmsTags(request *model.BatchCreateKmsTagsRequest) (*model.BatchCreateKmsTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateKmsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateKmsTagsResponse), nil
	}
}

// BatchCreateKmsTagsInvoker 批量添加删除密钥标签
func (c *KmsClient) BatchCreateKmsTagsInvoker(request *model.BatchCreateKmsTagsRequest) *BatchCreateKmsTagsInvoker {
	requestDef := GenReqDefForBatchCreateKmsTags()
	return &BatchCreateKmsTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelGrant 撤销授权
//
// - 功能介绍：撤销授权，授权用户撤销被授权用户操作密钥的权限。
// - 说明：
//    - 创建密钥的用户才能撤销该密钥授权。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CancelGrant(request *model.CancelGrantRequest) (*model.CancelGrantResponse, error) {
	requestDef := GenReqDefForCancelGrant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelGrantResponse), nil
	}
}

// CancelGrantInvoker 撤销授权
func (c *KmsClient) CancelGrantInvoker(request *model.CancelGrantRequest) *CancelGrantInvoker {
	requestDef := GenReqDefForCancelGrant()
	return &CancelGrantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelKeyDeletion 取消计划删除密钥
//
// - 功能介绍：取消计划删除密钥。
//
// - 说明：密钥处于“计划删除”状态才能取消计划删除密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CancelKeyDeletion(request *model.CancelKeyDeletionRequest) (*model.CancelKeyDeletionResponse, error) {
	requestDef := GenReqDefForCancelKeyDeletion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelKeyDeletionResponse), nil
	}
}

// CancelKeyDeletionInvoker 取消计划删除密钥
func (c *KmsClient) CancelKeyDeletionInvoker(request *model.CancelKeyDeletionRequest) *CancelKeyDeletionInvoker {
	requestDef := GenReqDefForCancelKeyDeletion()
	return &CancelKeyDeletionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelSelfGrant 退役授权
//
// - 功能介绍：退役授权，表示被授权用户不再具有授权密钥的操作权。
//   例如：用户A授权用户B可以操作密钥A/key，同时授权用户C可以撤销该授权，
//   那么用户A、B、C均可退役该授权，退役授权后，用户B不再可以使用A/key。
// - 须知：
//      可执行退役授权的主体包括：
//    - 创建授权的用户；
//    - 授权中retiring_principal指向的用户；
//    - 当授权的操作列表中包含retire-grant时，grantee_principal指向的用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CancelSelfGrant(request *model.CancelSelfGrantRequest) (*model.CancelSelfGrantResponse, error) {
	requestDef := GenReqDefForCancelSelfGrant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSelfGrantResponse), nil
	}
}

// CancelSelfGrantInvoker 退役授权
func (c *KmsClient) CancelSelfGrantInvoker(request *model.CancelSelfGrantRequest) *CancelSelfGrantInvoker {
	requestDef := GenReqDefForCancelSelfGrant()
	return &CancelSelfGrantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatakey 创建数据密钥
//
// - 功能介绍：创建数据密钥，返回结果包含明文和密文。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateDatakey(request *model.CreateDatakeyRequest) (*model.CreateDatakeyResponse, error) {
	requestDef := GenReqDefForCreateDatakey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatakeyResponse), nil
	}
}

// CreateDatakeyInvoker 创建数据密钥
func (c *KmsClient) CreateDatakeyInvoker(request *model.CreateDatakeyRequest) *CreateDatakeyInvoker {
	requestDef := GenReqDefForCreateDatakey()
	return &CreateDatakeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatakeyWithoutPlaintext 创建不含明文数据密钥
//
// - 功能介绍：创建数据密钥，返回结果只包含密文。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateDatakeyWithoutPlaintext(request *model.CreateDatakeyWithoutPlaintextRequest) (*model.CreateDatakeyWithoutPlaintextResponse, error) {
	requestDef := GenReqDefForCreateDatakeyWithoutPlaintext()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatakeyWithoutPlaintextResponse), nil
	}
}

// CreateDatakeyWithoutPlaintextInvoker 创建不含明文数据密钥
func (c *KmsClient) CreateDatakeyWithoutPlaintextInvoker(request *model.CreateDatakeyWithoutPlaintextRequest) *CreateDatakeyWithoutPlaintextInvoker {
	requestDef := GenReqDefForCreateDatakeyWithoutPlaintext()
	return &CreateDatakeyWithoutPlaintextInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGrant 创建授权
//
// - 功能介绍：创建授权，被授权用户可以对授权密钥进行操作。
// - 说明：
//    - 服务默认主密钥（密钥别名后缀为“/default”）不可以授权。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateGrant(request *model.CreateGrantRequest) (*model.CreateGrantResponse, error) {
	requestDef := GenReqDefForCreateGrant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGrantResponse), nil
	}
}

// CreateGrantInvoker 创建授权
func (c *KmsClient) CreateGrantInvoker(request *model.CreateGrantRequest) *CreateGrantInvoker {
	requestDef := GenReqDefForCreateGrant()
	return &CreateGrantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateKey 创建密钥
//
// 创建用户主密钥，用户主密钥可以为对称密钥或非对称密钥。
// - 对称密钥为长度为256位AES密钥，可用于小量数据的加密或者用于加密数据密钥。
// - 非对称密钥可以为RSA密钥对或者ECC密钥对，可用于数字签名及验签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateKey(request *model.CreateKeyRequest) (*model.CreateKeyResponse, error) {
	requestDef := GenReqDefForCreateKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKeyResponse), nil
	}
}

// CreateKeyInvoker 创建密钥
func (c *KmsClient) CreateKeyInvoker(request *model.CreateKeyRequest) *CreateKeyInvoker {
	requestDef := GenReqDefForCreateKey()
	return &CreateKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateKmsTag 添加密钥标签
//
// - 功能介绍：添加密钥标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateKmsTag(request *model.CreateKmsTagRequest) (*model.CreateKmsTagResponse, error) {
	requestDef := GenReqDefForCreateKmsTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKmsTagResponse), nil
	}
}

// CreateKmsTagInvoker 添加密钥标签
func (c *KmsClient) CreateKmsTagInvoker(request *model.CreateKmsTagRequest) *CreateKmsTagInvoker {
	requestDef := GenReqDefForCreateKmsTag()
	return &CreateKmsTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateParametersForImport 获取密钥导入参数
//
// - 功能介绍：获取导入密钥的必要参数，包括密钥导入令牌和密钥加密公钥。
// - 说明：返回的公钥类型默认为RSA_2048。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateParametersForImport(request *model.CreateParametersForImportRequest) (*model.CreateParametersForImportResponse, error) {
	requestDef := GenReqDefForCreateParametersForImport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateParametersForImportResponse), nil
	}
}

// CreateParametersForImportInvoker 获取密钥导入参数
func (c *KmsClient) CreateParametersForImportInvoker(request *model.CreateParametersForImportRequest) *CreateParametersForImportInvoker {
	requestDef := GenReqDefForCreateParametersForImport()
	return &CreateParametersForImportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRandom 创建随机数
//
// - 功能介绍：
//   生成8~8192bit范围内的随机数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateRandom(request *model.CreateRandomRequest) (*model.CreateRandomResponse, error) {
	requestDef := GenReqDefForCreateRandom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRandomResponse), nil
	}
}

// CreateRandomInvoker 创建随机数
func (c *KmsClient) CreateRandomInvoker(request *model.CreateRandomRequest) *CreateRandomInvoker {
	requestDef := GenReqDefForCreateRandom()
	return &CreateRandomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecret 创建凭据
//
// 创建新的凭据，并将凭据值存入凭据的初始版本。
//
//
// 凭据管理服务将凭据值加密后，存储在凭据对象下的版本中。每个版本可与多个凭据版本状态相关联，凭据版本状态用于标识凭据版本处于的阶段，没有版本状态标记的版本视为已弃用，可用凭据管理服务自动删除。
//
// 初始版本的状态被标记为SYSCURRENT。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateSecret(request *model.CreateSecretRequest) (*model.CreateSecretResponse, error) {
	requestDef := GenReqDefForCreateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretResponse), nil
	}
}

// CreateSecretInvoker 创建凭据
func (c *KmsClient) CreateSecretInvoker(request *model.CreateSecretRequest) *CreateSecretInvoker {
	requestDef := GenReqDefForCreateSecret()
	return &CreateSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecretVersion 创建凭据版本
//
// 在指定的凭据中，创建一个新的凭据版本，用于加密保管新的凭据值。默认情况下，新创建的凭据版本被标记为SYSCURRENT状态，而SYSCURRENT标记的前一个凭据版本被标记为SYSPREVIOUS状态。您可以通过指定VersionStage参数来覆盖默认行为。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateSecretVersion(request *model.CreateSecretVersionRequest) (*model.CreateSecretVersionResponse, error) {
	requestDef := GenReqDefForCreateSecretVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretVersionResponse), nil
	}
}

// CreateSecretVersionInvoker 创建凭据版本
func (c *KmsClient) CreateSecretVersionInvoker(request *model.CreateSecretVersionRequest) *CreateSecretVersionInvoker {
	requestDef := GenReqDefForCreateSecretVersion()
	return &CreateSecretVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DecryptData 解密数据
//
// - 功能介绍：解密数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DecryptData(request *model.DecryptDataRequest) (*model.DecryptDataResponse, error) {
	requestDef := GenReqDefForDecryptData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DecryptDataResponse), nil
	}
}

// DecryptDataInvoker 解密数据
func (c *KmsClient) DecryptDataInvoker(request *model.DecryptDataRequest) *DecryptDataInvoker {
	requestDef := GenReqDefForDecryptData()
	return &DecryptDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DecryptDatakey 解密数据密钥
//
// - 功能介绍：解密数据密钥，用指定的主密钥解密数据密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DecryptDatakey(request *model.DecryptDatakeyRequest) (*model.DecryptDatakeyResponse, error) {
	requestDef := GenReqDefForDecryptDatakey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DecryptDatakeyResponse), nil
	}
}

// DecryptDatakeyInvoker 解密数据密钥
func (c *KmsClient) DecryptDatakeyInvoker(request *model.DecryptDatakeyRequest) *DecryptDatakeyInvoker {
	requestDef := GenReqDefForDecryptDatakey()
	return &DecryptDatakeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteImportedKeyMaterial 删除密钥材料
//
// - 功能介绍：删除密钥材料信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DeleteImportedKeyMaterial(request *model.DeleteImportedKeyMaterialRequest) (*model.DeleteImportedKeyMaterialResponse, error) {
	requestDef := GenReqDefForDeleteImportedKeyMaterial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImportedKeyMaterialResponse), nil
	}
}

// DeleteImportedKeyMaterialInvoker 删除密钥材料
func (c *KmsClient) DeleteImportedKeyMaterialInvoker(request *model.DeleteImportedKeyMaterialRequest) *DeleteImportedKeyMaterialInvoker {
	requestDef := GenReqDefForDeleteImportedKeyMaterial()
	return &DeleteImportedKeyMaterialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteKey 计划删除密钥
//
// - 功能介绍：计划多少天后删除密钥，可设置7天～1096天内删除密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DeleteKey(request *model.DeleteKeyRequest) (*model.DeleteKeyResponse, error) {
	requestDef := GenReqDefForDeleteKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKeyResponse), nil
	}
}

// DeleteKeyInvoker 计划删除密钥
func (c *KmsClient) DeleteKeyInvoker(request *model.DeleteKeyRequest) *DeleteKeyInvoker {
	requestDef := GenReqDefForDeleteKey()
	return &DeleteKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecret 立即删除凭据
//
// 立即删除指定的凭据，且无法恢复。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DeleteSecret(request *model.DeleteSecretRequest) (*model.DeleteSecretResponse, error) {
	requestDef := GenReqDefForDeleteSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretResponse), nil
	}
}

// DeleteSecretInvoker 立即删除凭据
func (c *KmsClient) DeleteSecretInvoker(request *model.DeleteSecretRequest) *DeleteSecretInvoker {
	requestDef := GenReqDefForDeleteSecret()
	return &DeleteSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecretForSchedule 创建凭据的定时删除任务
//
// 指定延迟删除时间，创建删除凭据的定时任务，可设置7~30天的的延迟删除时间。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DeleteSecretForSchedule(request *model.DeleteSecretForScheduleRequest) (*model.DeleteSecretForScheduleResponse, error) {
	requestDef := GenReqDefForDeleteSecretForSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretForScheduleResponse), nil
	}
}

// DeleteSecretForScheduleInvoker 创建凭据的定时删除任务
func (c *KmsClient) DeleteSecretForScheduleInvoker(request *model.DeleteSecretForScheduleRequest) *DeleteSecretForScheduleInvoker {
	requestDef := GenReqDefForDeleteSecretForSchedule()
	return &DeleteSecretForScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecretStage 删除凭据的版本状态
//
// 删除指定的凭据版本状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DeleteSecretStage(request *model.DeleteSecretStageRequest) (*model.DeleteSecretStageResponse, error) {
	requestDef := GenReqDefForDeleteSecretStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretStageResponse), nil
	}
}

// DeleteSecretStageInvoker 删除凭据的版本状态
func (c *KmsClient) DeleteSecretStageInvoker(request *model.DeleteSecretStageRequest) *DeleteSecretStageInvoker {
	requestDef := GenReqDefForDeleteSecretStage()
	return &DeleteSecretStageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除密钥标签
//
// - 功能介绍：删除密钥标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除密钥标签
func (c *KmsClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableKey 禁用密钥
//
// - 功能介绍：禁用密钥，密钥禁用后不可以使用。
//
// - 说明：密钥为启用状态才能禁用密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DisableKey(request *model.DisableKeyRequest) (*model.DisableKeyResponse, error) {
	requestDef := GenReqDefForDisableKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableKeyResponse), nil
	}
}

// DisableKeyInvoker 禁用密钥
func (c *KmsClient) DisableKeyInvoker(request *model.DisableKeyRequest) *DisableKeyInvoker {
	requestDef := GenReqDefForDisableKey()
	return &DisableKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableKeyRotation 关闭密钥轮换
//
// - 功能介绍：关闭用户主密钥轮换。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DisableKeyRotation(request *model.DisableKeyRotationRequest) (*model.DisableKeyRotationResponse, error) {
	requestDef := GenReqDefForDisableKeyRotation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableKeyRotationResponse), nil
	}
}

// DisableKeyRotationInvoker 关闭密钥轮换
func (c *KmsClient) DisableKeyRotationInvoker(request *model.DisableKeyRotationRequest) *DisableKeyRotationInvoker {
	requestDef := GenReqDefForDisableKeyRotation()
	return &DisableKeyRotationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableKey 启用密钥
//
// - 功能介绍：启用密钥，密钥启用后才可以使用。
//
// - 说明：密钥为禁用状态才能启用密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) EnableKey(request *model.EnableKeyRequest) (*model.EnableKeyResponse, error) {
	requestDef := GenReqDefForEnableKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableKeyResponse), nil
	}
}

// EnableKeyInvoker 启用密钥
func (c *KmsClient) EnableKeyInvoker(request *model.EnableKeyRequest) *EnableKeyInvoker {
	requestDef := GenReqDefForEnableKey()
	return &EnableKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableKeyRotation 开启密钥轮换
//
// - 功能介绍：开启用户主密钥轮换。
// - 说明：
//   - 开启密钥轮换后，默认轮询间隔时间为365天。
//   - 默认主密钥及外部导入密钥不支持轮换操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) EnableKeyRotation(request *model.EnableKeyRotationRequest) (*model.EnableKeyRotationResponse, error) {
	requestDef := GenReqDefForEnableKeyRotation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableKeyRotationResponse), nil
	}
}

// EnableKeyRotationInvoker 开启密钥轮换
func (c *KmsClient) EnableKeyRotationInvoker(request *model.EnableKeyRotationRequest) *EnableKeyRotationInvoker {
	requestDef := GenReqDefForEnableKeyRotation()
	return &EnableKeyRotationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EncryptData 加密数据
//
// - 功能介绍：加密数据，用指定的用户主密钥加密数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) EncryptData(request *model.EncryptDataRequest) (*model.EncryptDataResponse, error) {
	requestDef := GenReqDefForEncryptData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EncryptDataResponse), nil
	}
}

// EncryptDataInvoker 加密数据
func (c *KmsClient) EncryptDataInvoker(request *model.EncryptDataRequest) *EncryptDataInvoker {
	requestDef := GenReqDefForEncryptData()
	return &EncryptDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EncryptDatakey 加密数据密钥
//
// - 功能介绍：加密数据密钥，用指定的主密钥加密数据密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) EncryptDatakey(request *model.EncryptDatakeyRequest) (*model.EncryptDatakeyResponse, error) {
	requestDef := GenReqDefForEncryptDatakey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EncryptDatakeyResponse), nil
	}
}

// EncryptDatakeyInvoker 加密数据密钥
func (c *KmsClient) EncryptDatakeyInvoker(request *model.EncryptDatakeyRequest) *EncryptDatakeyInvoker {
	requestDef := GenReqDefForEncryptDatakey()
	return &EncryptDatakeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportKeyMaterial 导入密钥材料
//
// - 功能介绍：导入密钥材料。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ImportKeyMaterial(request *model.ImportKeyMaterialRequest) (*model.ImportKeyMaterialResponse, error) {
	requestDef := GenReqDefForImportKeyMaterial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportKeyMaterialResponse), nil
	}
}

// ImportKeyMaterialInvoker 导入密钥材料
func (c *KmsClient) ImportKeyMaterialInvoker(request *model.ImportKeyMaterialRequest) *ImportKeyMaterialInvoker {
	requestDef := GenReqDefForImportKeyMaterial()
	return &ImportKeyMaterialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGrants 查询授权列表
//
// - 功能介绍：查询密钥的授权列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListGrants(request *model.ListGrantsRequest) (*model.ListGrantsResponse, error) {
	requestDef := GenReqDefForListGrants()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGrantsResponse), nil
	}
}

// ListGrantsInvoker 查询授权列表
func (c *KmsClient) ListGrantsInvoker(request *model.ListGrantsRequest) *ListGrantsInvoker {
	requestDef := GenReqDefForListGrants()
	return &ListGrantsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKeyDetail 查询密钥信息
//
// - 功能介绍：查询密钥详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListKeyDetail(request *model.ListKeyDetailRequest) (*model.ListKeyDetailResponse, error) {
	requestDef := GenReqDefForListKeyDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeyDetailResponse), nil
	}
}

// ListKeyDetailInvoker 查询密钥信息
func (c *KmsClient) ListKeyDetailInvoker(request *model.ListKeyDetailRequest) *ListKeyDetailInvoker {
	requestDef := GenReqDefForListKeyDetail()
	return &ListKeyDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKeys 查询密钥列表
//
// - 功能介绍：查询用户所有密钥列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListKeys(request *model.ListKeysRequest) (*model.ListKeysResponse, error) {
	requestDef := GenReqDefForListKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeysResponse), nil
	}
}

// ListKeysInvoker 查询密钥列表
func (c *KmsClient) ListKeysInvoker(request *model.ListKeysRequest) *ListKeysInvoker {
	requestDef := GenReqDefForListKeys()
	return &ListKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKmsByTags 查询密钥实例
//
// - 功能介绍：查询密钥实例。通过标签过滤，查询指定用户主密钥的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListKmsByTags(request *model.ListKmsByTagsRequest) (*model.ListKmsByTagsResponse, error) {
	requestDef := GenReqDefForListKmsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKmsByTagsResponse), nil
	}
}

// ListKmsByTagsInvoker 查询密钥实例
func (c *KmsClient) ListKmsByTagsInvoker(request *model.ListKmsByTagsRequest) *ListKmsByTagsInvoker {
	requestDef := GenReqDefForListKmsByTags()
	return &ListKmsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKmsTags 查询项目标签
//
// - 功能介绍：查询用户在指定项目下的所有标签集合。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListKmsTags(request *model.ListKmsTagsRequest) (*model.ListKmsTagsResponse, error) {
	requestDef := GenReqDefForListKmsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKmsTagsResponse), nil
	}
}

// ListKmsTagsInvoker 查询项目标签
func (c *KmsClient) ListKmsTagsInvoker(request *model.ListKmsTagsRequest) *ListKmsTagsInvoker {
	requestDef := GenReqDefForListKmsTags()
	return &ListKmsTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRetirableGrants 查询可退役授权列表
//
// - 功能介绍：查询用户可以退役的授权列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListRetirableGrants(request *model.ListRetirableGrantsRequest) (*model.ListRetirableGrantsResponse, error) {
	requestDef := GenReqDefForListRetirableGrants()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRetirableGrantsResponse), nil
	}
}

// ListRetirableGrantsInvoker 查询可退役授权列表
func (c *KmsClient) ListRetirableGrantsInvoker(request *model.ListRetirableGrantsRequest) *ListRetirableGrantsInvoker {
	requestDef := GenReqDefForListRetirableGrants()
	return &ListRetirableGrantsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecretStage 查询凭据的版本状态
//
// 查询指定凭据版本状态标记的版本信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListSecretStage(request *model.ListSecretStageRequest) (*model.ListSecretStageResponse, error) {
	requestDef := GenReqDefForListSecretStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecretStageResponse), nil
	}
}

// ListSecretStageInvoker 查询凭据的版本状态
func (c *KmsClient) ListSecretStageInvoker(request *model.ListSecretStageRequest) *ListSecretStageInvoker {
	requestDef := GenReqDefForListSecretStage()
	return &ListSecretStageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecretVersions 查询凭据的版本列表
//
// 查询指定凭据下的版本列表信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListSecretVersions(request *model.ListSecretVersionsRequest) (*model.ListSecretVersionsResponse, error) {
	requestDef := GenReqDefForListSecretVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecretVersionsResponse), nil
	}
}

// ListSecretVersionsInvoker 查询凭据的版本列表
func (c *KmsClient) ListSecretVersionsInvoker(request *model.ListSecretVersionsRequest) *ListSecretVersionsInvoker {
	requestDef := GenReqDefForListSecretVersions()
	return &ListSecretVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecrets 查询凭据列表
//
// 查询当前用户在本项目下创建的所有凭据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListSecrets(request *model.ListSecretsRequest) (*model.ListSecretsResponse, error) {
	requestDef := GenReqDefForListSecrets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecretsResponse), nil
	}
}

// ListSecretsInvoker 查询凭据列表
func (c *KmsClient) ListSecretsInvoker(request *model.ListSecretsRequest) *ListSecretsInvoker {
	requestDef := GenReqDefForListSecrets()
	return &ListSecretsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreSecret 取消凭据的定时删除任务
//
// 取消凭据的定时删除任务，凭据对象恢复可使用状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) RestoreSecret(request *model.RestoreSecretRequest) (*model.RestoreSecretResponse, error) {
	requestDef := GenReqDefForRestoreSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreSecretResponse), nil
	}
}

// RestoreSecretInvoker 取消凭据的定时删除任务
func (c *KmsClient) RestoreSecretInvoker(request *model.RestoreSecretRequest) *RestoreSecretInvoker {
	requestDef := GenReqDefForRestoreSecret()
	return &RestoreSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKeyRotationStatus 查询密钥轮换状态
//
// - 功能介绍：查询用户主密钥轮换状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowKeyRotationStatus(request *model.ShowKeyRotationStatusRequest) (*model.ShowKeyRotationStatusResponse, error) {
	requestDef := GenReqDefForShowKeyRotationStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKeyRotationStatusResponse), nil
	}
}

// ShowKeyRotationStatusInvoker 查询密钥轮换状态
func (c *KmsClient) ShowKeyRotationStatusInvoker(request *model.ShowKeyRotationStatusRequest) *ShowKeyRotationStatusInvoker {
	requestDef := GenReqDefForShowKeyRotationStatus()
	return &ShowKeyRotationStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKmsTags 查询密钥标签
//
// - 功能介绍：查询密钥标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowKmsTags(request *model.ShowKmsTagsRequest) (*model.ShowKmsTagsResponse, error) {
	requestDef := GenReqDefForShowKmsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKmsTagsResponse), nil
	}
}

// ShowKmsTagsInvoker 查询密钥标签
func (c *KmsClient) ShowKmsTagsInvoker(request *model.ShowKmsTagsRequest) *ShowKmsTagsInvoker {
	requestDef := GenReqDefForShowKmsTags()
	return &ShowKmsTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicKey 查询公钥信息
//
// - 查询用户指定非对称密钥的公钥信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowPublicKey(request *model.ShowPublicKeyRequest) (*model.ShowPublicKeyResponse, error) {
	requestDef := GenReqDefForShowPublicKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicKeyResponse), nil
	}
}

// ShowPublicKeyInvoker 查询公钥信息
func (c *KmsClient) ShowPublicKeyInvoker(request *model.ShowPublicKeyRequest) *ShowPublicKeyInvoker {
	requestDef := GenReqDefForShowPublicKey()
	return &ShowPublicKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecret 查询凭据
//
// 查询指定凭据的信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowSecret(request *model.ShowSecretRequest) (*model.ShowSecretResponse, error) {
	requestDef := GenReqDefForShowSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecretResponse), nil
	}
}

// ShowSecretInvoker 查询凭据
func (c *KmsClient) ShowSecretInvoker(request *model.ShowSecretRequest) *ShowSecretInvoker {
	requestDef := GenReqDefForShowSecret()
	return &ShowSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecretVersion 查询凭据的版本与凭据值
//
// 查询指定凭据版本的信息和版本中的明文凭据值，只能查询ENABLED状态的凭据。
// 通过/v1/{project_id}/secrets/{secret_id}/versions/latest可访问凭据最新版本的凭据值。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowSecretVersion(request *model.ShowSecretVersionRequest) (*model.ShowSecretVersionResponse, error) {
	requestDef := GenReqDefForShowSecretVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecretVersionResponse), nil
	}
}

// ShowSecretVersionInvoker 查询凭据的版本与凭据值
func (c *KmsClient) ShowSecretVersionInvoker(request *model.ShowSecretVersionRequest) *ShowSecretVersionInvoker {
	requestDef := GenReqDefForShowSecretVersion()
	return &ShowSecretVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserInstances 查询实例数
//
// - 功能介绍：查询实例数，获取用户已经创建的用户主密钥数量。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowUserInstances(request *model.ShowUserInstancesRequest) (*model.ShowUserInstancesResponse, error) {
	requestDef := GenReqDefForShowUserInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserInstancesResponse), nil
	}
}

// ShowUserInstancesInvoker 查询实例数
func (c *KmsClient) ShowUserInstancesInvoker(request *model.ShowUserInstancesRequest) *ShowUserInstancesInvoker {
	requestDef := GenReqDefForShowUserInstances()
	return &ShowUserInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserQuotas 查询配额
//
// - 功能介绍：查询配额，查询用户可以创建的用户主密钥配额总数及当前使用量信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowUserQuotas(request *model.ShowUserQuotasRequest) (*model.ShowUserQuotasResponse, error) {
	requestDef := GenReqDefForShowUserQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserQuotasResponse), nil
	}
}

// ShowUserQuotasInvoker 查询配额
func (c *KmsClient) ShowUserQuotasInvoker(request *model.ShowUserQuotasRequest) *ShowUserQuotasInvoker {
	requestDef := GenReqDefForShowUserQuotas()
	return &ShowUserQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Sign 签名数据
//
// - 功能介绍：使用非对称密钥的私钥对消息或消息摘要进行数字签名。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) Sign(request *model.SignRequest) (*model.SignResponse, error) {
	requestDef := GenReqDefForSign()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SignResponse), nil
	}
}

// SignInvoker 签名数据
func (c *KmsClient) SignInvoker(request *model.SignRequest) *SignInvoker {
	requestDef := GenReqDefForSign()
	return &SignInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateKeyAlias 修改密钥别名
//
// - 功能介绍：修改用户主密钥别名。
// - 说明：
//    - 服务默认主密钥（密钥别名后缀为“/default”）不可以修改。
//    - 密钥处于“计划删除”状态，密钥别名不可以修改。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) UpdateKeyAlias(request *model.UpdateKeyAliasRequest) (*model.UpdateKeyAliasResponse, error) {
	requestDef := GenReqDefForUpdateKeyAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeyAliasResponse), nil
	}
}

// UpdateKeyAliasInvoker 修改密钥别名
func (c *KmsClient) UpdateKeyAliasInvoker(request *model.UpdateKeyAliasRequest) *UpdateKeyAliasInvoker {
	requestDef := GenReqDefForUpdateKeyAlias()
	return &UpdateKeyAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateKeyDescription 修改密钥描述
//
// - 功能介绍：修改用户主密钥描述信息。
// - 说明：
//    - 服务默认主密钥（密钥别名后缀为“/default”）不可以修改。
//    - 密钥处于“计划删除”状态，密钥描述不可以修改。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) UpdateKeyDescription(request *model.UpdateKeyDescriptionRequest) (*model.UpdateKeyDescriptionResponse, error) {
	requestDef := GenReqDefForUpdateKeyDescription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeyDescriptionResponse), nil
	}
}

// UpdateKeyDescriptionInvoker 修改密钥描述
func (c *KmsClient) UpdateKeyDescriptionInvoker(request *model.UpdateKeyDescriptionRequest) *UpdateKeyDescriptionInvoker {
	requestDef := GenReqDefForUpdateKeyDescription()
	return &UpdateKeyDescriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateKeyRotationInterval 修改密钥轮换周期
//
// - 功能介绍：修改用户主密钥轮换周期。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) UpdateKeyRotationInterval(request *model.UpdateKeyRotationIntervalRequest) (*model.UpdateKeyRotationIntervalResponse, error) {
	requestDef := GenReqDefForUpdateKeyRotationInterval()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeyRotationIntervalResponse), nil
	}
}

// UpdateKeyRotationIntervalInvoker 修改密钥轮换周期
func (c *KmsClient) UpdateKeyRotationIntervalInvoker(request *model.UpdateKeyRotationIntervalRequest) *UpdateKeyRotationIntervalInvoker {
	requestDef := GenReqDefForUpdateKeyRotationInterval()
	return &UpdateKeyRotationIntervalInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecret 更新凭据
//
// 更新指定凭据的元数据信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) UpdateSecret(request *model.UpdateSecretRequest) (*model.UpdateSecretResponse, error) {
	requestDef := GenReqDefForUpdateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecretResponse), nil
	}
}

// UpdateSecretInvoker 更新凭据
func (c *KmsClient) UpdateSecretInvoker(request *model.UpdateSecretRequest) *UpdateSecretInvoker {
	requestDef := GenReqDefForUpdateSecret()
	return &UpdateSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecretStage 更新凭据的版本状态
//
// 更新凭据的版本状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) UpdateSecretStage(request *model.UpdateSecretStageRequest) (*model.UpdateSecretStageResponse, error) {
	requestDef := GenReqDefForUpdateSecretStage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecretStageResponse), nil
	}
}

// UpdateSecretStageInvoker 更新凭据的版本状态
func (c *KmsClient) UpdateSecretStageInvoker(request *model.UpdateSecretStageRequest) *UpdateSecretStageInvoker {
	requestDef := GenReqDefForUpdateSecretStage()
	return &UpdateSecretStageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateSignature 验证签名
//
// - 功能介绍：使用非对称密钥的私钥对消息或消息摘要进行数字签名。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ValidateSignature(request *model.ValidateSignatureRequest) (*model.ValidateSignatureResponse, error) {
	requestDef := GenReqDefForValidateSignature()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateSignatureResponse), nil
	}
}

// ValidateSignatureInvoker 验证签名
func (c *KmsClient) ValidateSignatureInvoker(request *model.ValidateSignatureRequest) *ValidateSignatureInvoker {
	requestDef := GenReqDefForValidateSignature()
	return &ValidateSignatureInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVersion 查询指定版本信息
//
// - 功能介绍：查指定API版本信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowVersion(request *model.ShowVersionRequest) (*model.ShowVersionResponse, error) {
	requestDef := GenReqDefForShowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionResponse), nil
	}
}

// ShowVersionInvoker 查询指定版本信息
func (c *KmsClient) ShowVersionInvoker(request *model.ShowVersionRequest) *ShowVersionInvoker {
	requestDef := GenReqDefForShowVersion()
	return &ShowVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVersions 查询版本信息列表
//
// - 功能介绍：查询API版本信息列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowVersions(request *model.ShowVersionsRequest) (*model.ShowVersionsResponse, error) {
	requestDef := GenReqDefForShowVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionsResponse), nil
	}
}

// ShowVersionsInvoker 查询版本信息列表
func (c *KmsClient) ShowVersionsInvoker(request *model.ShowVersionsRequest) *ShowVersionsInvoker {
	requestDef := GenReqDefForShowVersions()
	return &ShowVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
