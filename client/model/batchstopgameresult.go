// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type BatchStopGameResult struct {
    Success bool `json:"success,omitempty"`
    Model BatchStopGameResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
