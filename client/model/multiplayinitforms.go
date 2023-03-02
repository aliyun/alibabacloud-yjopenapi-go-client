// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayInitForms struct {
        // 游戏会话Id
    GameSession string `json:"gameSession"`
        // 分配的项目ak
    AppKey string `json:"appKey"`
        // 联机基础配置
    Config *MultiplayInitConfig `json:"config,omitempty"`
    Tokens *[]MultiplayInitTokens `json:"tokens,omitempty"`
}
