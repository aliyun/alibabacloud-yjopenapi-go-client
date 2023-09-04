// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayDemoLeaveMultiPlayResult struct {
    Success bool `json:"success,omitempty"`
    Model MultiplayDemoLeaveMultiPlayResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
