// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type StopGameForms struct {
        // 用户id，给到Paas平台和SDK，两者保持一致，全局唯一
    AccountId string `json:"accountId"`
        // Paas平台部署的游戏Id
    GameId string `json:"gameId"`
        // Paas平台AK(应用的AK，非服务端AK)
    AppKey string `json:"appKey"`
        // 调度成功返回的GameSession
    GameSession string `json:"gameSession"`
        // 端侧从Paas sdk获取的调度业务参数
    Reason *string `json:"reason,omitempty"`
}
