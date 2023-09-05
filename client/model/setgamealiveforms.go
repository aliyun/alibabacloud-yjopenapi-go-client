// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type SetGameAliveForms struct {
        // 会话ID
    GameSession string `json:"gameSession"`
        // 项目应用AK
    AppKey string `json:"appKey"`
        // 游戏可运行时长
    KeepAlive string `json:"keepAlive"`
}
