// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type LiveQueryStatusResult struct {
    Success bool `json:"success,omitempty"`
    Model LiveQueryStatusResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
