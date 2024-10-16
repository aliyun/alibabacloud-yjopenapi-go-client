// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminBatchUpdateDispatchConfigFormsConfigList struct {
        // 编码游戏id
    MixGameId string `json:"mixGameId"`
        // 独占并发数
    ExclusiveConcurrency int32 `json:"exclusiveConcurrency,omitempty"`
        // 最大并发百分比-Double类型[0.0,1.0]
    MaxConcurrencyPercent string `json:"maxConcurrencyPercent"`
        // 独占并发百分比--Double类型[0.0,1.0]
    ExclusiveConcurrencyPercent string `json:"exclusiveConcurrencyPercent"`
        // 调度等级
    Priority int32 `json:"priority"`
        // 最大并发数
    MaxConcurrency int32 `json:"maxConcurrency,omitempty"`
}
