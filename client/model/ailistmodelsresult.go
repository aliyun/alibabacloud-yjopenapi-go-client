// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiListModelsResult struct {
    Success bool `json:"success,omitempty"`
    Model AiListModelsResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
