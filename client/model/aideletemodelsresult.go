// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiDeleteModelsResult struct {
    Success bool `json:"success,omitempty"`
    Model AiDeleteModelsResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
