// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListActivateableInstancesResultModelDataList struct {
    CloudGameInstanceId string `json:"cloudGameInstanceId,omitempty"`
    CloudGameInstanceName string `json:"cloudGameInstanceName,omitempty"`
    ContainerCount int32 `json:"containerCount,omitempty"`
    MaxConcurrency int32 `json:"maxConcurrency,omitempty"`
}
