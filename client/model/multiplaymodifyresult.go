// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayModifyResult struct {
    Success bool `json:"success,omitempty"`
    Model MultiplayModifyResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
