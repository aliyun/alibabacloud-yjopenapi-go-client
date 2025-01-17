// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type CancelGameHangResult struct {
    Success bool `json:"success,omitempty"`
    Model CancelGameHangResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
