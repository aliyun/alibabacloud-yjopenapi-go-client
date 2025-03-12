// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiGetQueueResult struct {
    Success bool `json:"success,omitempty"`
    Model AiGetQueueResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
