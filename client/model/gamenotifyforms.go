// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GameNotifyForms struct {
        // 会话ID
    GameSession string `json:"gameSession"`
        // 项目应用AK
    AppKey string `json:"appKey"`
        // 通知类型
    Type_ string `json:"type"`
        // 通知值
    Value *string `json:"value,omitempty"`
}
