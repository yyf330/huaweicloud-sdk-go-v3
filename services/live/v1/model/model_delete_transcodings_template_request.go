/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Request Object
type DeleteTranscodingsTemplateRequest struct {
	Domain string `json:"domain"`
	AppName string `json:"app_name"`
}
