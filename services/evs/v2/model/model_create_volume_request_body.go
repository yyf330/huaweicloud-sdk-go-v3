/*
    * EVS
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// This is a auto create Body Object
type CreateVolumeRequestBody struct {
	BssParam *BssParamForCreateVolume `json:"bssParam,omitempty"`
	Volume *CreateVolumeOption `json:"volume"`
}
