// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AdaptQueryRequestByIdResult struct {
    Success bool `json:"success,omitempty"`
    Model AdaptQueryRequestByIdResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
