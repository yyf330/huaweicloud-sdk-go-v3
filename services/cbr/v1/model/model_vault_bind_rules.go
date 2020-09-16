/*
 * Cbr
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type VaultBindRules struct {
	// 按tags过滤自动绑定的资源
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o VaultBindRules) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"VaultBindRules", string(data)}, " ")
}
