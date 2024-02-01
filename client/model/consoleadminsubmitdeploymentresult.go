// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminSubmitDeploymentResult struct {
    Success bool `json:"success,omitempty"`
    Model ConsoleAdminSubmitDeploymentResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
