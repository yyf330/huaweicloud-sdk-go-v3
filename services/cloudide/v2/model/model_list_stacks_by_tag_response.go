/*
    * cloudide
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// Response Object
type ListStacksByTagResponse struct {
	Stack *StacksTag `json:"stack,omitempty"`
	// 状态
	Status string `json:"status,omitempty"`
}
