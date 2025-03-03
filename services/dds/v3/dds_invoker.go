package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dds/v3/model"
)

type AddShardingNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddShardingNodeInvoker) Invoke() (*model.AddShardingNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddShardingNodeResponse), nil
	}
}

type AttachEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachEipInvoker) Invoke() (*model.AttachEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachEipResponse), nil
	}
}

type AttachInternalIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachInternalIpInvoker) Invoke() (*model.AttachInternalIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachInternalIpResponse), nil
	}
}

type BatchTagActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchTagActionInvoker) Invoke() (*model.BatchTagActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchTagActionResponse), nil
	}
}

type CancelEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelEipInvoker) Invoke() (*model.CancelEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelEipResponse), nil
	}
}

type CheckPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckPasswordInvoker) Invoke() (*model.CheckPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckPasswordResponse), nil
	}
}

type CreateDatabaseRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseRoleInvoker) Invoke() (*model.CreateDatabaseRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseRoleResponse), nil
	}
}

type CreateDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseUserInvoker) Invoke() (*model.CreateDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseUserResponse), nil
	}
}

type CreateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceInvoker) Invoke() (*model.CreateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceResponse), nil
	}
}

type CreateIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIpInvoker) Invoke() (*model.CreateIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIpResponse), nil
	}
}

type CreateManualBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateManualBackupInvoker) Invoke() (*model.CreateManualBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateManualBackupResponse), nil
	}
}

type DeleteDatabaseRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabaseRoleInvoker) Invoke() (*model.DeleteDatabaseRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseRoleResponse), nil
	}
}

type DeleteDatabaseUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabaseUserInvoker) Invoke() (*model.DeleteDatabaseUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseUserResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
	}
}

type DeleteManualBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteManualBackupInvoker) Invoke() (*model.DeleteManualBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteManualBackupResponse), nil
	}
}

type DeleteSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSessionInvoker) Invoke() (*model.DeleteSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSessionResponse), nil
	}
}

type DownloadErrorlogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadErrorlogInvoker) Invoke() (*model.DownloadErrorlogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadErrorlogResponse), nil
	}
}

type DownloadSlowlogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadSlowlogInvoker) Invoke() (*model.DownloadSlowlogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadSlowlogResponse), nil
	}
}

type ListAuditlogLinksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditlogLinksInvoker) Invoke() (*model.ListAuditlogLinksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditlogLinksResponse), nil
	}
}

type ListAuditlogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuditlogsInvoker) Invoke() (*model.ListAuditlogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuditlogsResponse), nil
	}
}

type ListAz2MigrateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAz2MigrateInvoker) Invoke() (*model.ListAz2MigrateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAz2MigrateResponse), nil
	}
}

type ListBackupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupsInvoker) Invoke() (*model.ListBackupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupsResponse), nil
	}
}

type ListDatabaseRolesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseRolesInvoker) Invoke() (*model.ListDatabaseRolesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseRolesResponse), nil
	}
}

type ListDatabaseUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseUsersInvoker) Invoke() (*model.ListDatabaseUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseUsersResponse), nil
	}
}

type ListDatastoreVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatastoreVersionsInvoker) Invoke() (*model.ListDatastoreVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatastoreVersionsResponse), nil
	}
}

type ListErrorLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListErrorLogsInvoker) Invoke() (*model.ListErrorLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListErrorLogsResponse), nil
	}
}

type ListFlavorInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorInfosInvoker) Invoke() (*model.ListFlavorInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorInfosResponse), nil
	}
}

type ListFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorsInvoker) Invoke() (*model.ListFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorsResponse), nil
	}
}

type ListInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceTagsInvoker) Invoke() (*model.ListInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceTagsResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type ListInstancesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesByTagsInvoker) Invoke() (*model.ListInstancesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesByTagsResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ListRestoreCollectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestoreCollectionsInvoker) Invoke() (*model.ListRestoreCollectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestoreCollectionsResponse), nil
	}
}

type ListRestoreDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestoreDatabasesInvoker) Invoke() (*model.ListRestoreDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestoreDatabasesResponse), nil
	}
}

type ListRestoreTimesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRestoreTimesInvoker) Invoke() (*model.ListRestoreTimesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRestoreTimesResponse), nil
	}
}

type ListSessionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSessionsInvoker) Invoke() (*model.ListSessionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSessionsResponse), nil
	}
}

type ListSlowLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSlowLogsInvoker) Invoke() (*model.ListSlowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSlowLogsResponse), nil
	}
}

type ListStorageTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStorageTypeInvoker) Invoke() (*model.ListStorageTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStorageTypeResponse), nil
	}
}

type MigrateAzInvoker struct {
	*invoker.BaseInvoker
}

func (i *MigrateAzInvoker) Invoke() (*model.MigrateAzResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrateAzResponse), nil
	}
}

type ResetPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPasswordInvoker) Invoke() (*model.ResetPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPasswordResponse), nil
	}
}

type ResizeInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeInstanceInvoker) Invoke() (*model.ResizeInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeInstanceResponse), nil
	}
}

type ResizeInstanceVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeInstanceVolumeInvoker) Invoke() (*model.ResizeInstanceVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeInstanceVolumeResponse), nil
	}
}

type RestartInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartInstanceInvoker) Invoke() (*model.RestartInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartInstanceResponse), nil
	}
}

type RestoreInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreInstanceInvoker) Invoke() (*model.RestoreInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreInstanceResponse), nil
	}
}

type RestoreInstanceFromCollectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreInstanceFromCollectionInvoker) Invoke() (*model.RestoreInstanceFromCollectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreInstanceFromCollectionResponse), nil
	}
}

type RestoreNewInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreNewInstanceInvoker) Invoke() (*model.RestoreNewInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreNewInstanceResponse), nil
	}
}

type SetAuditlogPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetAuditlogPolicyInvoker) Invoke() (*model.SetAuditlogPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAuditlogPolicyResponse), nil
	}
}

type SetBackupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetBackupPolicyInvoker) Invoke() (*model.SetBackupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetBackupPolicyResponse), nil
	}
}

type SetBalancerSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetBalancerSwitchInvoker) Invoke() (*model.SetBalancerSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetBalancerSwitchResponse), nil
	}
}

type SetBalancerWindowInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetBalancerWindowInvoker) Invoke() (*model.SetBalancerWindowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetBalancerWindowResponse), nil
	}
}

type ShowAuditlogPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditlogPolicyInvoker) Invoke() (*model.ShowAuditlogPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditlogPolicyResponse), nil
	}
}

type ShowBackupDownloadLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupDownloadLinkInvoker) Invoke() (*model.ShowBackupDownloadLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupDownloadLinkResponse), nil
	}
}

type ShowBackupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupPolicyInvoker) Invoke() (*model.ShowBackupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupPolicyResponse), nil
	}
}

type ShowConnectionStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConnectionStatisticsInvoker) Invoke() (*model.ShowConnectionStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConnectionStatisticsResponse), nil
	}
}

type ShowJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobDetailInvoker) Invoke() (*model.ShowJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobDetailResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
	}
}

type ShowShardingBalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowShardingBalancerInvoker) Invoke() (*model.ShowShardingBalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowShardingBalancerResponse), nil
	}
}

type SwitchSlowlogDesensitizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchSlowlogDesensitizationInvoker) Invoke() (*model.SwitchSlowlogDesensitizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchSlowlogDesensitizationResponse), nil
	}
}

type SwitchSslInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchSslInvoker) Invoke() (*model.SwitchSslResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchSslResponse), nil
	}
}

type SwitchoverReplicaSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchoverReplicaSetInvoker) Invoke() (*model.SwitchoverReplicaSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchoverReplicaSetResponse), nil
	}
}

type UpdateClientNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClientNetworkInvoker) Invoke() (*model.UpdateClientNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClientNetworkResponse), nil
	}
}

type UpdateInstanceNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceNameInvoker) Invoke() (*model.UpdateInstanceNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceNameResponse), nil
	}
}

type UpdateInstancePortInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstancePortInvoker) Invoke() (*model.UpdateInstancePortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstancePortResponse), nil
	}
}

type UpdateInstanceRemarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceRemarkInvoker) Invoke() (*model.UpdateInstanceRemarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceRemarkResponse), nil
	}
}

type UpdateSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityGroupInvoker) Invoke() (*model.UpdateSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityGroupResponse), nil
	}
}

type ListApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionInvoker) Invoke() (*model.ListApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionResponse), nil
	}
}

type ShowApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiVersionInvoker) Invoke() (*model.ShowApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiVersionResponse), nil
	}
}
