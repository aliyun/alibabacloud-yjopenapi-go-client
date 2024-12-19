// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetGameConcurrencyForms struct {
        // 游戏ID
    GameId string `json:"gameId"`
        // 云游戏项目应用AK
    AppKey string `json:"appKey"`
        // 是否查询PaaS当前排队人数
    QueryQueueConcurrency *bool `json:"queryQueueConcurrency,omitempty"`
        // 指定userLevel查询当前userLevel排队人数，不指定时，查询当前所有排队人数
    QueueUserLevel *int32 `json:"queueUserLevel,omitempty"`
}
