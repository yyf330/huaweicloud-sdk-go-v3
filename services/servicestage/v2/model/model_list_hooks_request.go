package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHooksRequest struct {

	// 授权名称。
	XRepoAuth string `json:"X-Repo-Auth"`

	// 组织ID。
	Namespace string `json:"namespace"`

	// 仓库项目ID，如果含有“/”，需要将“/”替换为“:”。
	Project string `json:"project"`
}

func (o ListHooksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHooksRequest struct{}"
	}

	return strings.Join([]string{"ListHooksRequest", string(data)}, " ")
}
