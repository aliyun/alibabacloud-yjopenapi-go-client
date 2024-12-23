// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetStopGameTokenResult struct {
    Success bool `json:"success,omitempty"`
    Model GetStopGameTokenResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
