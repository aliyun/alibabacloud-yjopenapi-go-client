// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListProjectsResult struct {
        // 服务端状态码
    MsgCode string `json:"msgCode,omitempty"`
        // 服务端描述信息
    MsgInfo string `json:"msgInfo,omitempty"`
        // 是否成功
    Success bool `json:"success,omitempty"`
    Model ConsoleAdminListProjectsResultModel `json:"model,omitempty"`
}
