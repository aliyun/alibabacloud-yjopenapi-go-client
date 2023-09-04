// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayDemoModifyMultiPlayResult struct {
    Success bool `json:"success,omitempty"`
    Model MultiplayDemoModifyMultiPlayResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
