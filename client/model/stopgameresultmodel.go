// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type StopGameResultModel struct {
        // Paas平台游戏ID
    GameId string `json:"gameId,omitempty"`
        // 会话标识
    GameSession string `json:"gameSession,omitempty"`
        // 返回码
    Code string `json:"code,omitempty"`
        // 返回信息
    Message string `json:"message,omitempty"`
        // 调度执行结果
    Success bool `json:"success,omitempty"`
}
