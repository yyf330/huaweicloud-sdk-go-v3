/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMetricDataRequest struct {
	Body *[]MetricDataItem `json:"body,omitempty"`
}

func (o CreateMetricDataRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateMetricDataRequest", string(data)}, " ")
}
