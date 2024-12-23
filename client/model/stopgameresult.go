// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type StopGameResult struct {
        // 是否成功
    Success bool `json:"success,omitempty"`
    Model StopGameResultModel `json:"model,omitempty"`
        // 错误信息
    MsgInfo string `json:"msgInfo,omitempty"`
        // 错误码
    MsgCode string `json:"msgCode,omitempty"`
}
