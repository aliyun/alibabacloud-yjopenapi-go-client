// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetStopGameTokenForms struct {
        // Paas平台部署的游戏Id
    GameId string `json:"gameId"`
        // Paas平台AK(应用的AK，非服务端AK)
    AppKey string `json:"appKey"`
}
