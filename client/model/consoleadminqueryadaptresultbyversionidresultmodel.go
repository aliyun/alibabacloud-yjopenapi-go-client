// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminQueryAdaptResultByVersionIdResultModel struct {
        // -1：已下线 0:待适配 1:适配中 2:适配完成 3:适配失败
    State int32 `json:"state,omitempty"`
    MaxConcurrency int32 `json:"maxConcurrency,omitempty"`
}
