// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type QuerySessionStatusResult struct {
    Success bool `json:"success,omitempty"`
    Model QuerySessionStatusResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
