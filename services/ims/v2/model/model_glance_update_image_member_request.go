/*
    * IMS
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// Request Object
type GlanceUpdateImageMemberRequest struct {
	ImageId string `json:"image_id"`
	MemberId string `json:"member_id"`
	Body *GlanceUpdateImageMemberRequestBody `json:"body,omitempty"`
}
