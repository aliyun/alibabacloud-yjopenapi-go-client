// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ClientNotifyResult struct {
    Success bool `json:"success,omitempty"`
    Model ClientNotifyResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
