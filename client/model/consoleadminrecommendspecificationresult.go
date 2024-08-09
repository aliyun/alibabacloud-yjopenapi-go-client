// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminRecommendSpecificationResult struct {
    Success bool `json:"success,omitempty"`
    Model []ConsoleAdminRecommendSpecificationResultModel `json:"model,omitempty"`
    MsgInfo string `json:"msgInfo,omitempty"`
    MsgCode string `json:"msgCode,omitempty"`
}
