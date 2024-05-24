// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type BatchStopGameForms struct {
        // Paas平台部署的游戏Id
    GameId string `json:"gameId"`
        // Paas平台AK(应用的AK，非服务端AK)
    AppKey string `json:"appKey"`
        // 通过接口获取的token
    Token string `json:"token"`
        // 踢人的原因，会透传到sdk侧
    Reason *string `json:"reason,omitempty"`
        // TrackInfo，回传消息
    TrackInfo *string `json:"trackInfo,omitempty"`
        // 支持多标签传输
    Tags *string `json:"tags,omitempty"`
}
