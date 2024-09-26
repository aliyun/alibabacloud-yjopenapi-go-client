// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiUploadModelsResult struct {
    Success bool `json:"success,omitempty"`
    Model AiUploadModelsResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
