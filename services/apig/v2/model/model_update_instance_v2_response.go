package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateInstanceV2Response struct {

	// 实例ID
	Id *string `json:"id,omitempty"`

	// 实例所属租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例状态： - Creating：创建中 - CreateSuccess：创建成功 - CreateFail：创建失败 - Initing：初始化中 - Registering：注册中 - Running：运行中 - InitingFailed：初始化失败 - RegisterFailed：注册失败 - Installing：安装中 - InstallFailed：安装失败 - Updating：升级中 - UpdateFailed：升级失败 - Rollbacking：回滚中 - RollbackSuccess：回滚成功 - RollbackFailed：回滚失败 - Deleting：删除中 - DeleteFailed：删除失败 - Unregistering：注销中 - UnRegisterFailed：注销失败 - CreateTimeout：创建超时 - InitTimeout：初始化超时 - RegisterTimeout：注册超时 - InstallTimeout：安装超时 - UpdateTimeout：升级超时 - RollbackTimeout：回滚超时 - DeleteTimeout：删除超时 - UnregisterTimeout：注销超时 - Starting：启动中 - Freezing：冻结中 - Frozen：已冻结 - Restarting：重启中 - RestartFail：重启失败 - Unhealthy：实例异常 - RestartTimeout：重启超时
	Status *UpdateInstanceV2ResponseStatus `json:"status,omitempty"`

	// 实例状态对应编号 - 1：创建中 - 2：创建成功 - 3：创建失败 - 4：初始化中 - 5：注册中 - 6：运行中 - 7：初始化失败 - 8：注册失败 - 10：安装中 - 11：安装失败 - 12：升级中 - 13：升级失败 - 20：回滚中 - 21：回滚成功 - 22：回滚失败 - 23：删除中 - 24：删除失败 - 25：注销中 - 26：注销失败 - 27：创建超时 - 28：初始化超时 - 29：注册超时 - 30：安装超时 - 31：升级超时 - 32：回滚超时 - 33：删除超时 - 34：注销超时 - 35：启动中 - 36：冻结中 - 37：已冻结 - 38：重启中 - 39：重启失败 - 40：实例异常 - 41：重启超时
	InstanceStatus *UpdateInstanceV2ResponseInstanceStatus `json:"instance_status,omitempty"`

	// 实例类型  默认apig
	Type *string `json:"type,omitempty"`

	// 实例规格： - BASIC：基础版实例 - PROFESSIONAL：专业版实例 - ENTERPRISE：企业版实例 - PLATINUM：铂金版实例 - BASIC_IPV6：基础版IPV6实例 - PROFESSIONAL_IPV6：专业版IPV6实例 - ENTERPRISE_IPV6：企业版IPV6实例 - PLATINUM_IPV6：铂金版IPV6实例
	Spec *UpdateInstanceV2ResponseSpec `json:"spec,omitempty"`

	// 实例创建时间。unix时间戳格式。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 企业项目ID，企业帐号必填
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例绑定的弹性IP地址
	EipAddress *string `json:"eip_address,omitempty"`

	// 实例计费方式： - 0：按需计费 - 1：包周期计费
	ChargingMode *UpdateInstanceV2ResponseChargingMode `json:"charging_mode,omitempty"`

	// 包周期计费订单编号
	CbcMetadata *string `json:"cbc_metadata,omitempty"`

	// 实例使用的负载均衡器类型 - lvs Linux虚拟服务器 - elb 弹性负载均衡，elb仅部分region支持
	LoadbalancerProvider *UpdateInstanceV2ResponseLoadbalancerProvider `json:"loadbalancer_provider,omitempty"`

	// 实例描述
	Description *string `json:"description,omitempty"`

	// 虚拟私有云ID。  获取方法如下：   - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。   - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询VPC列表”章节。
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网的网络ID。  获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询子网列表”章节。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 指定实例所属的安全组。  获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询安全组列表”章节。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// '维护时间窗开始时间。时间格式为 xx:00:00，xx取值为02,06,10,14,18,22。'  '在这个时间段内，运维人员可以对该实例的节点进行维护操作。维护期间，业务可以正常使用，可能会发生闪断。维护操作通常几个月一次。'
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// '维护时间窗结束时间。时间格式为 xx:00:00，与维护时间窗开始时间相差4个小时。'  '在这个时间段内，运维人员可以对该实例的节点进行维护操作。维护期间，业务可以正常使用，可能会发生闪断。维护操作通常几个月一次'。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// 实例入口，虚拟私有云访问地址
	IngressIp *string `json:"ingress_ip,omitempty"`

	// 实例所属用户ID
	UserId *string `json:"user_id,omitempty"`

	// 出公网网段 (IPv6)  。  当前仅部分region部分可用区支持IPv6
	NatEipIpv6Cidr *string `json:"nat_eip_ipv6_cidr,omitempty"`

	// 弹性IP地址(IPv6)。  当前仅部分region部分可用区支持IPv6
	EipIpv6Address *string `json:"eip_ipv6_address,omitempty"`

	// 实例出公网IP
	NatEipAddress *string `json:"nat_eip_address,omitempty"`

	// 出公网带宽
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 可用区
	AvailableZoneIds *string `json:"available_zone_ids,omitempty"`

	// 实例版本编号
	InstanceVersion *string `json:"instance_version,omitempty"`

	// 子网的网络ID。  暂不支持
	VirsubnetId *string `json:"virsubnet_id,omitempty"`

	// roma弹性公网IP。  暂不支持
	RomaEipAddress *string `json:"roma_eip_address,omitempty"`

	// 监听信息  暂不支持
	Listeners *interface{} `json:"listeners,omitempty"`

	// 实例支持的特性列表
	SupportedFeatures *[]string `json:"supported_features,omitempty"`

	EndpointService *EndpointService `json:"endpoint_service,omitempty"`

	// 终端节点服务列表
	EndpointServices *[]EndpointService `json:"endpoint_services,omitempty"`

	NodeIps *NodeIps `json:"node_ips,omitempty"`

	// 公网入口地址列表
	Publicips *[]IpDetails `json:"publicips,omitempty"`

	// 私网入口地址列表
	Privateips     *[]IpDetails `json:"privateips,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateInstanceV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceV2Response struct{}"
	}

	return strings.Join([]string{"UpdateInstanceV2Response", string(data)}, " ")
}

type UpdateInstanceV2ResponseStatus struct {
	value string
}

type UpdateInstanceV2ResponseStatusEnum struct {
	CREATING           UpdateInstanceV2ResponseStatus
	CREATE_SUCCESS     UpdateInstanceV2ResponseStatus
	CREATE_FAIL        UpdateInstanceV2ResponseStatus
	INITING            UpdateInstanceV2ResponseStatus
	REGISTERING        UpdateInstanceV2ResponseStatus
	RUNNING            UpdateInstanceV2ResponseStatus
	INITING_FAILED     UpdateInstanceV2ResponseStatus
	REGISTER_FAILED    UpdateInstanceV2ResponseStatus
	INSTALLING         UpdateInstanceV2ResponseStatus
	INSTALL_FAILED     UpdateInstanceV2ResponseStatus
	UPDATING           UpdateInstanceV2ResponseStatus
	UPDATE_FAILED      UpdateInstanceV2ResponseStatus
	ROLLBACKING        UpdateInstanceV2ResponseStatus
	ROLLBACK_SUCCESS   UpdateInstanceV2ResponseStatus
	ROLLBACK_FAILED    UpdateInstanceV2ResponseStatus
	DELETING           UpdateInstanceV2ResponseStatus
	DELETE_FAILED      UpdateInstanceV2ResponseStatus
	UNREGISTERING      UpdateInstanceV2ResponseStatus
	UN_REGISTER_FAILED UpdateInstanceV2ResponseStatus
	CREATE_TIMEOUT     UpdateInstanceV2ResponseStatus
	INIT_TIMEOUT       UpdateInstanceV2ResponseStatus
	REGISTER_TIMEOUT   UpdateInstanceV2ResponseStatus
	INSTALL_TIMEOUT    UpdateInstanceV2ResponseStatus
	UPDATE_TIMEOUT     UpdateInstanceV2ResponseStatus
	ROLLBACK_TIMEOUT   UpdateInstanceV2ResponseStatus
	DELETE_TIMEOUT     UpdateInstanceV2ResponseStatus
	UNREGISTER_TIMEOUT UpdateInstanceV2ResponseStatus
	STARTING           UpdateInstanceV2ResponseStatus
	FREEZING           UpdateInstanceV2ResponseStatus
	FROZEN             UpdateInstanceV2ResponseStatus
	RESTARTING         UpdateInstanceV2ResponseStatus
	RESTART_FAIL       UpdateInstanceV2ResponseStatus
	UNHEALTHY          UpdateInstanceV2ResponseStatus
	RESTART_TIMEOUT    UpdateInstanceV2ResponseStatus
}

func GetUpdateInstanceV2ResponseStatusEnum() UpdateInstanceV2ResponseStatusEnum {
	return UpdateInstanceV2ResponseStatusEnum{
		CREATING: UpdateInstanceV2ResponseStatus{
			value: "Creating",
		},
		CREATE_SUCCESS: UpdateInstanceV2ResponseStatus{
			value: "CreateSuccess",
		},
		CREATE_FAIL: UpdateInstanceV2ResponseStatus{
			value: "CreateFail",
		},
		INITING: UpdateInstanceV2ResponseStatus{
			value: "Initing",
		},
		REGISTERING: UpdateInstanceV2ResponseStatus{
			value: "Registering",
		},
		RUNNING: UpdateInstanceV2ResponseStatus{
			value: "Running",
		},
		INITING_FAILED: UpdateInstanceV2ResponseStatus{
			value: "InitingFailed",
		},
		REGISTER_FAILED: UpdateInstanceV2ResponseStatus{
			value: "RegisterFailed",
		},
		INSTALLING: UpdateInstanceV2ResponseStatus{
			value: "Installing",
		},
		INSTALL_FAILED: UpdateInstanceV2ResponseStatus{
			value: "InstallFailed",
		},
		UPDATING: UpdateInstanceV2ResponseStatus{
			value: "Updating",
		},
		UPDATE_FAILED: UpdateInstanceV2ResponseStatus{
			value: "UpdateFailed",
		},
		ROLLBACKING: UpdateInstanceV2ResponseStatus{
			value: "Rollbacking",
		},
		ROLLBACK_SUCCESS: UpdateInstanceV2ResponseStatus{
			value: "RollbackSuccess",
		},
		ROLLBACK_FAILED: UpdateInstanceV2ResponseStatus{
			value: "RollbackFailed",
		},
		DELETING: UpdateInstanceV2ResponseStatus{
			value: "Deleting",
		},
		DELETE_FAILED: UpdateInstanceV2ResponseStatus{
			value: "DeleteFailed",
		},
		UNREGISTERING: UpdateInstanceV2ResponseStatus{
			value: "Unregistering",
		},
		UN_REGISTER_FAILED: UpdateInstanceV2ResponseStatus{
			value: "UnRegisterFailed",
		},
		CREATE_TIMEOUT: UpdateInstanceV2ResponseStatus{
			value: "CreateTimeout",
		},
		INIT_TIMEOUT: UpdateInstanceV2ResponseStatus{
			value: "InitTimeout",
		},
		REGISTER_TIMEOUT: UpdateInstanceV2ResponseStatus{
			value: "RegisterTimeout",
		},
		INSTALL_TIMEOUT: UpdateInstanceV2ResponseStatus{
			value: "InstallTimeout",
		},
		UPDATE_TIMEOUT: UpdateInstanceV2ResponseStatus{
			value: "UpdateTimeout",
		},
		ROLLBACK_TIMEOUT: UpdateInstanceV2ResponseStatus{
			value: "RollbackTimeout",
		},
		DELETE_TIMEOUT: UpdateInstanceV2ResponseStatus{
			value: "DeleteTimeout",
		},
		UNREGISTER_TIMEOUT: UpdateInstanceV2ResponseStatus{
			value: "UnregisterTimeout",
		},
		STARTING: UpdateInstanceV2ResponseStatus{
			value: "Starting",
		},
		FREEZING: UpdateInstanceV2ResponseStatus{
			value: "Freezing",
		},
		FROZEN: UpdateInstanceV2ResponseStatus{
			value: "Frozen",
		},
		RESTARTING: UpdateInstanceV2ResponseStatus{
			value: "Restarting",
		},
		RESTART_FAIL: UpdateInstanceV2ResponseStatus{
			value: "RestartFail",
		},
		UNHEALTHY: UpdateInstanceV2ResponseStatus{
			value: "Unhealthy",
		},
		RESTART_TIMEOUT: UpdateInstanceV2ResponseStatus{
			value: "RestartTimeout",
		},
	}
}

func (c UpdateInstanceV2ResponseStatus) Value() string {
	return c.value
}

func (c UpdateInstanceV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInstanceV2ResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateInstanceV2ResponseInstanceStatus struct {
	value int32
}

type UpdateInstanceV2ResponseInstanceStatusEnum struct {
	E_1  UpdateInstanceV2ResponseInstanceStatus
	E_2  UpdateInstanceV2ResponseInstanceStatus
	E_3  UpdateInstanceV2ResponseInstanceStatus
	E_4  UpdateInstanceV2ResponseInstanceStatus
	E_5  UpdateInstanceV2ResponseInstanceStatus
	E_6  UpdateInstanceV2ResponseInstanceStatus
	E_7  UpdateInstanceV2ResponseInstanceStatus
	E_8  UpdateInstanceV2ResponseInstanceStatus
	E_10 UpdateInstanceV2ResponseInstanceStatus
	E_11 UpdateInstanceV2ResponseInstanceStatus
	E_12 UpdateInstanceV2ResponseInstanceStatus
	E_13 UpdateInstanceV2ResponseInstanceStatus
	E_20 UpdateInstanceV2ResponseInstanceStatus
	E_21 UpdateInstanceV2ResponseInstanceStatus
	E_22 UpdateInstanceV2ResponseInstanceStatus
	E_23 UpdateInstanceV2ResponseInstanceStatus
	E_24 UpdateInstanceV2ResponseInstanceStatus
	E_25 UpdateInstanceV2ResponseInstanceStatus
	E_26 UpdateInstanceV2ResponseInstanceStatus
	E_27 UpdateInstanceV2ResponseInstanceStatus
	E_28 UpdateInstanceV2ResponseInstanceStatus
	E_29 UpdateInstanceV2ResponseInstanceStatus
	E_30 UpdateInstanceV2ResponseInstanceStatus
	E_31 UpdateInstanceV2ResponseInstanceStatus
	E_32 UpdateInstanceV2ResponseInstanceStatus
	E_33 UpdateInstanceV2ResponseInstanceStatus
	E_34 UpdateInstanceV2ResponseInstanceStatus
	E_35 UpdateInstanceV2ResponseInstanceStatus
	E_36 UpdateInstanceV2ResponseInstanceStatus
	E_37 UpdateInstanceV2ResponseInstanceStatus
	E_38 UpdateInstanceV2ResponseInstanceStatus
	E_39 UpdateInstanceV2ResponseInstanceStatus
	E_40 UpdateInstanceV2ResponseInstanceStatus
	E_41 UpdateInstanceV2ResponseInstanceStatus
}

func GetUpdateInstanceV2ResponseInstanceStatusEnum() UpdateInstanceV2ResponseInstanceStatusEnum {
	return UpdateInstanceV2ResponseInstanceStatusEnum{
		E_1: UpdateInstanceV2ResponseInstanceStatus{
			value: 1,
		}, E_2: UpdateInstanceV2ResponseInstanceStatus{
			value: 2,
		}, E_3: UpdateInstanceV2ResponseInstanceStatus{
			value: 3,
		}, E_4: UpdateInstanceV2ResponseInstanceStatus{
			value: 4,
		}, E_5: UpdateInstanceV2ResponseInstanceStatus{
			value: 5,
		}, E_6: UpdateInstanceV2ResponseInstanceStatus{
			value: 6,
		}, E_7: UpdateInstanceV2ResponseInstanceStatus{
			value: 7,
		}, E_8: UpdateInstanceV2ResponseInstanceStatus{
			value: 8,
		}, E_10: UpdateInstanceV2ResponseInstanceStatus{
			value: 10,
		}, E_11: UpdateInstanceV2ResponseInstanceStatus{
			value: 11,
		}, E_12: UpdateInstanceV2ResponseInstanceStatus{
			value: 12,
		}, E_13: UpdateInstanceV2ResponseInstanceStatus{
			value: 13,
		}, E_20: UpdateInstanceV2ResponseInstanceStatus{
			value: 20,
		}, E_21: UpdateInstanceV2ResponseInstanceStatus{
			value: 21,
		}, E_22: UpdateInstanceV2ResponseInstanceStatus{
			value: 22,
		}, E_23: UpdateInstanceV2ResponseInstanceStatus{
			value: 23,
		}, E_24: UpdateInstanceV2ResponseInstanceStatus{
			value: 24,
		}, E_25: UpdateInstanceV2ResponseInstanceStatus{
			value: 25,
		}, E_26: UpdateInstanceV2ResponseInstanceStatus{
			value: 26,
		}, E_27: UpdateInstanceV2ResponseInstanceStatus{
			value: 27,
		}, E_28: UpdateInstanceV2ResponseInstanceStatus{
			value: 28,
		}, E_29: UpdateInstanceV2ResponseInstanceStatus{
			value: 29,
		}, E_30: UpdateInstanceV2ResponseInstanceStatus{
			value: 30,
		}, E_31: UpdateInstanceV2ResponseInstanceStatus{
			value: 31,
		}, E_32: UpdateInstanceV2ResponseInstanceStatus{
			value: 32,
		}, E_33: UpdateInstanceV2ResponseInstanceStatus{
			value: 33,
		}, E_34: UpdateInstanceV2ResponseInstanceStatus{
			value: 34,
		}, E_35: UpdateInstanceV2ResponseInstanceStatus{
			value: 35,
		}, E_36: UpdateInstanceV2ResponseInstanceStatus{
			value: 36,
		}, E_37: UpdateInstanceV2ResponseInstanceStatus{
			value: 37,
		}, E_38: UpdateInstanceV2ResponseInstanceStatus{
			value: 38,
		}, E_39: UpdateInstanceV2ResponseInstanceStatus{
			value: 39,
		}, E_40: UpdateInstanceV2ResponseInstanceStatus{
			value: 40,
		}, E_41: UpdateInstanceV2ResponseInstanceStatus{
			value: 41,
		},
	}
}

func (c UpdateInstanceV2ResponseInstanceStatus) Value() int32 {
	return c.value
}

func (c UpdateInstanceV2ResponseInstanceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInstanceV2ResponseInstanceStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type UpdateInstanceV2ResponseSpec struct {
	value string
}

type UpdateInstanceV2ResponseSpecEnum struct {
	BASIC             UpdateInstanceV2ResponseSpec
	PROFESSIONAL      UpdateInstanceV2ResponseSpec
	ENTERPRISE        UpdateInstanceV2ResponseSpec
	PLATINUM          UpdateInstanceV2ResponseSpec
	BASIC_IPV6        UpdateInstanceV2ResponseSpec
	PROFESSIONAL_IPV6 UpdateInstanceV2ResponseSpec
	ENTERPRISE_IPV6   UpdateInstanceV2ResponseSpec
	PLATINUM_IPV6     UpdateInstanceV2ResponseSpec
}

func GetUpdateInstanceV2ResponseSpecEnum() UpdateInstanceV2ResponseSpecEnum {
	return UpdateInstanceV2ResponseSpecEnum{
		BASIC: UpdateInstanceV2ResponseSpec{
			value: "BASIC",
		},
		PROFESSIONAL: UpdateInstanceV2ResponseSpec{
			value: "PROFESSIONAL",
		},
		ENTERPRISE: UpdateInstanceV2ResponseSpec{
			value: "ENTERPRISE",
		},
		PLATINUM: UpdateInstanceV2ResponseSpec{
			value: "PLATINUM",
		},
		BASIC_IPV6: UpdateInstanceV2ResponseSpec{
			value: "BASIC_IPV6",
		},
		PROFESSIONAL_IPV6: UpdateInstanceV2ResponseSpec{
			value: "PROFESSIONAL_IPV6",
		},
		ENTERPRISE_IPV6: UpdateInstanceV2ResponseSpec{
			value: "ENTERPRISE_IPV6",
		},
		PLATINUM_IPV6: UpdateInstanceV2ResponseSpec{
			value: "PLATINUM_IPV6",
		},
	}
}

func (c UpdateInstanceV2ResponseSpec) Value() string {
	return c.value
}

func (c UpdateInstanceV2ResponseSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInstanceV2ResponseSpec) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateInstanceV2ResponseChargingMode struct {
	value int32
}

type UpdateInstanceV2ResponseChargingModeEnum struct {
	E_0 UpdateInstanceV2ResponseChargingMode
	E_1 UpdateInstanceV2ResponseChargingMode
}

func GetUpdateInstanceV2ResponseChargingModeEnum() UpdateInstanceV2ResponseChargingModeEnum {
	return UpdateInstanceV2ResponseChargingModeEnum{
		E_0: UpdateInstanceV2ResponseChargingMode{
			value: 0,
		}, E_1: UpdateInstanceV2ResponseChargingMode{
			value: 1,
		},
	}
}

func (c UpdateInstanceV2ResponseChargingMode) Value() int32 {
	return c.value
}

func (c UpdateInstanceV2ResponseChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInstanceV2ResponseChargingMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type UpdateInstanceV2ResponseLoadbalancerProvider struct {
	value string
}

type UpdateInstanceV2ResponseLoadbalancerProviderEnum struct {
	LVS UpdateInstanceV2ResponseLoadbalancerProvider
	ELB UpdateInstanceV2ResponseLoadbalancerProvider
}

func GetUpdateInstanceV2ResponseLoadbalancerProviderEnum() UpdateInstanceV2ResponseLoadbalancerProviderEnum {
	return UpdateInstanceV2ResponseLoadbalancerProviderEnum{
		LVS: UpdateInstanceV2ResponseLoadbalancerProvider{
			value: "lvs",
		},
		ELB: UpdateInstanceV2ResponseLoadbalancerProvider{
			value: "elb",
		},
	}
}

func (c UpdateInstanceV2ResponseLoadbalancerProvider) Value() string {
	return c.value
}

func (c UpdateInstanceV2ResponseLoadbalancerProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInstanceV2ResponseLoadbalancerProvider) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
