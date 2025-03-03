package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Request Object
type ListFunctionAsyncInvocationsRequest struct {

	// 函数URN
	FunctionUrn string `json:"function_urn"`

	// 需要查询的异步请求ID。如果不指定，默认查询所有异步调用记录
	RequestId *string `json:"request_id,omitempty"`

	// 本次查询最大返回的数据条数，最大值500，默认值100
	Limit *string `json:"limit,omitempty"`

	// 本次查询起始位置，默认值0
	Marker *string `json:"marker,omitempty"`

	// 本次查询指定的异步调用状态，支持5种状态，如果不指定，则查询所有状态的调用记录 WAIT: 等待 RUNNING: 执行中 SUCCESS: 执行成功 FAIL: 执行失败 DISCARD: 请求丢弃
	Status *ListFunctionAsyncInvocationsRequestStatus `json:"status,omitempty"`

	// 搜索起始时间（格式为YYYY-MM-DD'T'HH:mm:ss,UTC时间）。如果不指定默认为当前时间前1小时
	QueryBeginTime *sdktime.SdkTime `json:"query_begin_time,omitempty"`

	// 搜索结束时间（格式为YYYY-MM-DD'T'HH:mm:ss,UTC时间）。如果不指定默认为当前时间
	QueryEndTime *sdktime.SdkTime `json:"query_end_time,omitempty"`
}

func (o ListFunctionAsyncInvocationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionAsyncInvocationsRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionAsyncInvocationsRequest", string(data)}, " ")
}

type ListFunctionAsyncInvocationsRequestStatus struct {
	value string
}

type ListFunctionAsyncInvocationsRequestStatusEnum struct {
	WAIT    ListFunctionAsyncInvocationsRequestStatus
	RUNNING ListFunctionAsyncInvocationsRequestStatus
	SUCCESS ListFunctionAsyncInvocationsRequestStatus
	FAIL    ListFunctionAsyncInvocationsRequestStatus
	DISCARD ListFunctionAsyncInvocationsRequestStatus
}

func GetListFunctionAsyncInvocationsRequestStatusEnum() ListFunctionAsyncInvocationsRequestStatusEnum {
	return ListFunctionAsyncInvocationsRequestStatusEnum{
		WAIT: ListFunctionAsyncInvocationsRequestStatus{
			value: "WAIT",
		},
		RUNNING: ListFunctionAsyncInvocationsRequestStatus{
			value: "RUNNING",
		},
		SUCCESS: ListFunctionAsyncInvocationsRequestStatus{
			value: "SUCCESS",
		},
		FAIL: ListFunctionAsyncInvocationsRequestStatus{
			value: "FAIL",
		},
		DISCARD: ListFunctionAsyncInvocationsRequestStatus{
			value: "DISCARD",
		},
	}
}

func (c ListFunctionAsyncInvocationsRequestStatus) Value() string {
	return c.value
}

func (c ListFunctionAsyncInvocationsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFunctionAsyncInvocationsRequestStatus) UnmarshalJSON(b []byte) error {
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
