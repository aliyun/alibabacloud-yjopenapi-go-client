// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type QueryGameHangResult struct {
    Success bool `json:"success,omitempty"`
    Model QueryGameHangResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
