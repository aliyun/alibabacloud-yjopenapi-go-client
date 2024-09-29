// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiCancelQueueResult struct {
    Success bool `json:"success,omitempty"`
    Model AiCancelQueueResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
