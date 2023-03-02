// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayQueryResultModelConfig struct {
        // 联机最多能进的人数
    MaxToken int32 `json:"maxToken,omitempty"`
        // 联机主机断连超时时间，超出则关闭联机
    HostTimeout int32 `json:"hostTimeout,omitempty"`
        // 联机从机玩家断连超时时间，超出则退出联机
    ConnectTimeout int32 `json:"connectTimeout,omitempty"`
        // 联机无操作超时时间，超出则退出联机
    InputTimeout int32 `json:"inputTimeout,omitempty"`
}
