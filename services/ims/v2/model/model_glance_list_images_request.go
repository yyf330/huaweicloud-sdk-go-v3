/*
    * IMS
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// Request Object
type GlanceListImagesRequest struct {
	Imagetype string `json:"__imagetype,omitempty"`
	Isregistered bool `json:"__isregistered,omitempty"`
	OsBit string `json:"__os_bit,omitempty"`
	OsType string `json:"__os_type,omitempty"`
	Platform string `json:"__platform,omitempty"`
	SupportDiskintensive string `json:"__support_diskintensive,omitempty"`
	SupportHighperformance string `json:"__support_highperformance,omitempty"`
	SupportKvm string `json:"__support_kvm,omitempty"`
	SupportKvmGpuType string `json:"__support_kvm_gpu_type,omitempty"`
	SupportKvmInfiniband string `json:"__support_kvm_infiniband,omitempty"`
	SupportLargememory string `json:"__support_largememory,omitempty"`
	SupportXen string `json:"__support_xen,omitempty"`
	SupportXenGpuType string `json:"__support_xen_gpu_type,omitempty"`
	SupportXenHana string `json:"__support_xen_hana,omitempty"`
	ContainerFormat string `json:"container_format,omitempty"`
	DiskFormat string `json:"disk_format,omitempty"`
	Id string `json:"id,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	Marker string `json:"marker,omitempty"`
	MemberStatus string `json:"member_status,omitempty"`
	MinDisk int32 `json:"min_disk,omitempty"`
	MinRam int32 `json:"min_ram,omitempty"`
	Name string `json:"name,omitempty"`
	Owner string `json:"owner,omitempty"`
	Protected bool `json:"protected,omitempty"`
	SortDir string `json:"sort_dir,omitempty"`
	SortKey string `json:"sort_key,omitempty"`
	Status string `json:"status,omitempty"`
	Tag string `json:"tag,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
