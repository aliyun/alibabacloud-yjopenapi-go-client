// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayDemoJoinMultiPlayResult struct {
    Success bool `json:"success,omitempty"`
    Model MultiplayDemoJoinMultiPlayResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
