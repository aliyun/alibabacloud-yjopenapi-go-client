// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetGameConcurrencyResultModel struct {
        // Paas平台游戏ID
    GameId string `json:"gameId,omitempty"`
        // 当前毫秒时间戳
    CurrentTime int64 `json:"currentTime,omitempty"`
        // 当前游戏region维度游戏中人数统计
    PlayingRegionList []GetGameConcurrencyResultModelPlayingRegionList `json:"playingRegionList,omitempty"`
        // 错误码
    Code string `json:"code,omitempty"`
        // 当前游戏中人数
    PlayingTotal int64 `json:"playingTotal,omitempty"`
        // 查询结果
    Success bool `json:"success,omitempty"`
        // 当前排队人数
    QueuingTotal int64 `json:"queuingTotal,omitempty"`
        // 错误信息
    Message string `json:"message,omitempty"`
        // 游戏归属的项目Id
    ProjectId string `json:"projectId,omitempty"`
}
