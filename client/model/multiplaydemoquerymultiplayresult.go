// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayDemoQueryMultiPlayResult struct {
    Success bool `json:"success,omitempty"`
    Model MultiplayDemoQueryMultiPlayResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
