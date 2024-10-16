// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetGameConcurrencyResult struct {
    Success bool `json:"success,omitempty"`
    Model GetGameConcurrencyResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
