// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiUpdateModelsResult struct {
    Success bool `json:"success,omitempty"`
    Model AiUpdateModelsResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
