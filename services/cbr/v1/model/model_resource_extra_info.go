/*
    * Cbr
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

type ResourceExtraInfo struct {
	// 需要排除备份的卷id。仅虚拟机磁盘级备份有效
	ExcludeVolumes []string `json:"exclude_volumes,omitempty"`
	// 要备份的卷
	IncludeVolumes []ResourceExtraInfoIncludeVolumes `json:"include_volumes,omitempty"`
}
