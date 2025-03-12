// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type CancelGameHangResultModel struct {
        // 错误Code
    Code string `json:"code,omitempty"`
        // 设置是否成功
    Success bool `json:"success,omitempty"`
        // 错误Message
    Message string `json:"message,omitempty"`
        // 会话ID
    GameSession string `json:"gameSession,omitempty"`
}
