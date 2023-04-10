// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListActivatedInstancesResultModelDataList struct {
        // 实例ID
    CloudGameInstanceId string `json:"cloudGameInstanceId,omitempty"`
        // 实例名成
    CloudGameInstanceName string `json:"cloudGameInstanceName,omitempty"`
        // 实例路数
    ContainerCount int32 `json:"containerCount,omitempty"`
        // 最大并发数
    MaxConcurrency int32 `json:"maxConcurrency,omitempty"`
}
