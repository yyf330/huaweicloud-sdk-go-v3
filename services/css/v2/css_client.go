package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/css/v2/model"
)

type CssClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCssClient(hcClient *http_client.HcHttpClient) *CssClient {
	return &CssClient{HcClient: hcClient}
}

func CssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateCluster 创建集群V2
//
// 该接口用于创建集群V2。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

// CreateClusterInvoker 创建集群V2
func (c *CssClient) CreateClusterInvoker(request *model.CreateClusterRequest) *CreateClusterInvoker {
	requestDef := GenReqDefForCreateCluster()
	return &CreateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartCluster 重启集群V2
//
// 该接口用于重启集群。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) RestartCluster(request *model.RestartClusterRequest) (*model.RestartClusterResponse, error) {
	requestDef := GenReqDefForRestartCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartClusterResponse), nil
	}
}

// RestartClusterInvoker 重启集群V2
func (c *CssClient) RestartClusterInvoker(request *model.RestartClusterRequest) *RestartClusterInvoker {
	requestDef := GenReqDefForRestartCluster()
	return &RestartClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RollingRestart 滚动重启
//
// 该接口用于滚动重启。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) RollingRestart(request *model.RollingRestartRequest) (*model.RollingRestartResponse, error) {
	requestDef := GenReqDefForRollingRestart()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RollingRestartResponse), nil
	}
}

// RollingRestartInvoker 滚动重启
func (c *CssClient) RollingRestartInvoker(request *model.RollingRestartRequest) *RollingRestartInvoker {
	requestDef := GenReqDefForRollingRestart()
	return &RollingRestartInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartAutoCreateSnapshots 开启自动创建快照功能
//
// 该接口用于打开自动创建快照功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StartAutoCreateSnapshots(request *model.StartAutoCreateSnapshotsRequest) (*model.StartAutoCreateSnapshotsResponse, error) {
	requestDef := GenReqDefForStartAutoCreateSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartAutoCreateSnapshotsResponse), nil
	}
}

// StartAutoCreateSnapshotsInvoker 开启自动创建快照功能
func (c *CssClient) StartAutoCreateSnapshotsInvoker(request *model.StartAutoCreateSnapshotsRequest) *StartAutoCreateSnapshotsInvoker {
	requestDef := GenReqDefForStartAutoCreateSnapshots()
	return &StartAutoCreateSnapshotsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopAutoCreateSnapshots 关闭自动创建快照功能
//
// 该接口用于关闭自动创建快照功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StopAutoCreateSnapshots(request *model.StopAutoCreateSnapshotsRequest) (*model.StopAutoCreateSnapshotsResponse, error) {
	requestDef := GenReqDefForStopAutoCreateSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopAutoCreateSnapshotsResponse), nil
	}
}

// StopAutoCreateSnapshotsInvoker 关闭自动创建快照功能
func (c *CssClient) StopAutoCreateSnapshotsInvoker(request *model.StopAutoCreateSnapshotsRequest) *StopAutoCreateSnapshotsInvoker {
	requestDef := GenReqDefForStopAutoCreateSnapshots()
	return &StopAutoCreateSnapshotsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
