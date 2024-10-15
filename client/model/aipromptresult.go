// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiPromptResult struct {
    Success bool `json:"success,omitempty"`
    Model AiPromptResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
