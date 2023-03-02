// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type SetGameHangResultModel struct {
        // 会话ID
    GameSession string `json:"gameSession,omitempty"`
        // 设置是否成功
    Success bool `json:"success,omitempty"`
        // 错误Code
    Code string `json:"code,omitempty"`
        // 错误Message
    Message string `json:"message,omitempty"`
        // 开始挂机毫秒时间戳
    StartHangTimestamp int64 `json:"startHangTimestamp,omitempty"`
        // 挂机时长，单位：秒
    Duration int64 `json:"duration,omitempty"`
}
