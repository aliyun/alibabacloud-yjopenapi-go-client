// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type StopPreopenContainerForms struct {
        // 云游戏项目appKey
    AppKey string `json:"appKey"`
        // 云游戏平台游戏ID
    GameId string `json:"gameId"`
    NumberOfBatches *int32 `json:"numberOfBatches,omitempty"`
}
