package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudide/v2/model"
)

type CreateExtensionAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateExtensionAuthorizationInvoker) Invoke() (*model.CreateExtensionAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExtensionAuthorizationResponse), nil
	}
}

type ListProjectTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTemplatesInvoker) Invoke() (*model.ListProjectTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTemplatesResponse), nil
	}
}

type ListStacksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStacksInvoker) Invoke() (*model.ListStacksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStacksResponse), nil
	}
}

type ShowAccountStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccountStatusInvoker) Invoke() (*model.ShowAccountStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccountStatusResponse), nil
	}
}

type ShowExtensionAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExtensionAuthorizationInvoker) Invoke() (*model.ShowExtensionAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExtensionAuthorizationResponse), nil
	}
}

type ShowPriceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPriceInvoker) Invoke() (*model.ShowPriceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPriceResponse), nil
	}
}

type CheckInstanceAccessInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckInstanceAccessInvoker) Invoke() (*model.CheckInstanceAccessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckInstanceAccessResponse), nil
	}
}

type CheckNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckNameInvoker) Invoke() (*model.CheckNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckNameResponse), nil
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

type CreateInstanceBy3rdInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceBy3rdInvoker) Invoke() (*model.CreateInstanceBy3rdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceBy3rdResponse), nil
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

type ListOrgInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrgInstancesInvoker) Invoke() (*model.ListOrgInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrgInstancesResponse), nil
	}
}

type ShowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInvoker) Invoke() (*model.ShowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceResponse), nil
	}
}

type StartInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartInstanceInvoker) Invoke() (*model.StartInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartInstanceResponse), nil
	}
}

type StopInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopInstanceInvoker) Invoke() (*model.StopInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopInstanceResponse), nil
	}
}

type UpdateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceInvoker) Invoke() (*model.UpdateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceResponse), nil
	}
}

type UpdateInstanceActivityInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceActivityInvoker) Invoke() (*model.UpdateInstanceActivityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceActivityResponse), nil
	}
}
