// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayInitConfig struct {
        // 联机最多能进的人数
    MaxToken *int32 `json:"maxToken,omitempty"`
        // 联机主机断连超时时间
    HostTimeout *int32 `json:"hostTimeout,omitempty"`
        // 联机从机玩家断连超时时间
    ConnectTimeout *int32 `json:"connectTimeout,omitempty"`
        // 联机从机无操作超时时间
    InputTimeout *int32 `json:"inputTimeout,omitempty"`
}
