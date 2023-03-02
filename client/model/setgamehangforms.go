// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type SetGameHangForms struct {
        // 会话ID
    GameSession string `json:"gameSession"`
        // 项目应用AK
    AppKey string `json:"appKey"`
        // 挂机时长
    Duration int64 `json:"duration"`
}
