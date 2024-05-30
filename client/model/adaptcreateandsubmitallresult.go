// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AdaptCreateAndSubmitAllResult struct {
    Success bool `json:"success,omitempty"`
        // 适配请求ID
    Model int64 `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
