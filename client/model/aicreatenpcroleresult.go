// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiCreateNpcRoleResult struct {
    Success bool `json:"success,omitempty"`
    Model map[string]string `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
