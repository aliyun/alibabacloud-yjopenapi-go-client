// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type LiveStopGameLiveForms struct {
        // 用户appKey
    AppKey string `json:"appKey"`
        // 游戏会话Id
    GameSession string `json:"gameSession"`
        // 推流标识ID
    LiveId *string `json:"liveId,omitempty"`
}
