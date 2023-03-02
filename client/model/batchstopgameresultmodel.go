// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type BatchStopGameResultModel struct {
        // Paas平台游戏ID
    GameId string `json:"gameId,omitempty"`
        // 请求链路唯一标示
    RequestId string `json:"requestId,omitempty"`
        // 游戏归属的项目Id
    ProjectId string `json:"projectId,omitempty"`
        // 批量停止的回传trackInfo
    TrackInfo string `json:"trackInfo,omitempty"`
        // 返回码
    Code string `json:"code,omitempty"`
        // 返回信息
    Message string `json:"message,omitempty"`
        // 调度执行结果
    Success bool `json:"success,omitempty"`
}
