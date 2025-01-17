// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetStockResult struct {
    Success bool `json:"success,omitempty"`
    Model GetStockResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
