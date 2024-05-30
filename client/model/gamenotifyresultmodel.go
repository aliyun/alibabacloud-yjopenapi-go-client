// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GameNotifyResultModel struct {
    Code string `json:"code,omitempty"`
    Success bool `json:"success,omitempty"`
    Message string `json:"message,omitempty"`
    GameSession string `json:"gameSession,omitempty"`
}
