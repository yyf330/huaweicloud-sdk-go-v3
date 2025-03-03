package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配置项。
type Configs struct {

	// 回源请求头配置
	OriginRequestHeader *[]OriginRequestHeader `json:"origin_request_header,omitempty"`

	// http header配置
	HttpResponseHeader *[]HttpResponseHeader `json:"http_response_header,omitempty"`

	UrlAuth *UrlAuth `json:"url_auth,omitempty"`

	Https *HttpPutBody `json:"https,omitempty"`

	// 源站配置。
	Sources *[]SourcesConfig `json:"sources,omitempty"`

	// 回源协议（follow：协议跟随回源，http：HTTP回源(默认)，https：https回源）。
	OriginProtocol *string `json:"origin_protocol,omitempty"`

	ForceRedirect *ForceRedirectConfig `json:"force_redirect,omitempty"`

	Compress *Compress `json:"compress,omitempty"`
}

func (o Configs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Configs struct{}"
	}

	return strings.Join([]string{"Configs", string(data)}, " ")
}
