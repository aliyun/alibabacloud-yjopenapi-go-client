// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type LiveStartGameLiveForms struct {
        // 用户appKey
    AppKey string `json:"appKey"`
        // 游戏会话Id
    GameSession string `json:"gameSession"`
        // 推流服务器地址
    ServerUrl string `json:"serverUrl"`
        // 鉴权参数
    StreamKey string `json:"streamKey"`
        // 直播推流配置
    Config *LiveStartGameLiveConfig `json:"config,omitempty"`
}
