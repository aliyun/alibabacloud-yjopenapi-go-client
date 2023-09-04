// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayDemoQueryMultiPlayResultModelConfig struct {
    MaxToken int32 `json:"maxToken,omitempty"`
    OfflineTimeout int32 `json:"offlineTimeout,omitempty"`
    ForbidJoin bool `json:"forbidJoin,omitempty"`
    HostTimeout int32 `json:"hostTimeout,omitempty"`
    ConnectTimeout int32 `json:"connectTimeout,omitempty"`
    InputTimeout int32 `json:"inputTimeout,omitempty"`
}
