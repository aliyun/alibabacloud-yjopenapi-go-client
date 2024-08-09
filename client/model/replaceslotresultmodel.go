// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ReplaceSlotResultModel struct {
    GameId string `json:"gameId,omitempty"`
    AccountId string `json:"accountId,omitempty"`
    Code string `json:"code,omitempty"`
    RegionId string `json:"regionId,omitempty"`
    Success bool `json:"success,omitempty"`
    Message string `json:"message,omitempty"`
    GameSession string `json:"gameSession,omitempty"`
}
