// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminGetGameInstanceContainerRatioResult struct {
    Success bool `json:"success,omitempty"`
    Model []ConsoleAdminGetGameInstanceContainerRatioResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
