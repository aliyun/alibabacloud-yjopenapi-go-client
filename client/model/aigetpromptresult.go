// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiGetPromptResult struct {
    Success bool `json:"success,omitempty"`
    Model AiGetPromptResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
