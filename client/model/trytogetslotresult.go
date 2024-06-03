// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type TryToGetSlotResult struct {
    Success bool `json:"success,omitempty"`
    Model TryToGetSlotResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
