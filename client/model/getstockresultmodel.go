// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetStockResultModel struct {
        // 请求链路唯一标示
    RequestId string `json:"requestId,omitempty"`
        // 返回码
    Code string `json:"code,omitempty"`
        // 返回信息
    Message string `json:"message,omitempty"`
        // 调度执行结果
    Success bool `json:"success,omitempty"`
        // 游戏归属的项目Id
    ProjectId string `json:"projectId,omitempty"`
        // Paas平台游戏ID
    GameId string `json:"gameId,omitempty"`
        // 当前毫秒时间戳
    CurrentTime int64 `json:"currentTime,omitempty"`
        // 总路数
    QuotaTotal int32 `json:"quotaTotal,omitempty"`
        // 已使用路数
    UsedTotal int32 `json:"usedTotal,omitempty"`
        // 可用剩余路数
    AvailableTotal int32 `json:"availableTotal,omitempty"`
    InstanceStockList []GetStockResultInstanceStockList `json:"instanceStockList,omitempty"`
}
