// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ReplaceSlotResult struct {
    Success bool `json:"success,omitempty"`
    Model ReplaceSlotResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
