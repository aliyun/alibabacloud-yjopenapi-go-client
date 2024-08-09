// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminQueryAdaptResultByVersionIdResult struct {
    Success bool `json:"success,omitempty"`
    Model []ConsoleAdminQueryAdaptResultByVersionIdResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
