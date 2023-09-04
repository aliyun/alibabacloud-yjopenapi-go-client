// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayDemoInitMultiPlayResult struct {
    Success bool `json:"success,omitempty"`
    Model MultiplayDemoInitMultiPlayResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
