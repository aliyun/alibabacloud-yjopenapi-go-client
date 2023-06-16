// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type TryToGetSlotResultModel struct {
        // Paas平台游戏ID
    GameId string `json:"gameId,omitempty"`
        // 会话标识
    GameSession string `json:"gameSession,omitempty"`
        // 分配的region
    RegionId string `json:"regionId,omitempty"`
        // 用户id，给到Paas平台和SDK，两者保持一致，全局唯一
    AccountId string `json:"accountId,omitempty"`
        // 返回码
    Code string `json:"code,omitempty"`
        // 返回信息
    Message string `json:"message,omitempty"`
        // 调度执行结果
    Success bool `json:"success,omitempty"`
    SlotData string `json:"slotData,omitempty"`
}
