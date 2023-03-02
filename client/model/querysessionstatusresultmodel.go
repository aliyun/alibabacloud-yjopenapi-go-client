// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type QuerySessionStatusResultModel struct {
        // 返回码
    Code string `json:"code,omitempty"`
        // 返回信息
    Message string `json:"message,omitempty"`
        // 是否成功
    Success bool `json:"success,omitempty"`
        // 会话归属的租户ID
    TenantId int64 `json:"tenantId,omitempty"`
        // 会话归属的项目ID
    ProjectId string `json:"projectId,omitempty"`
        // 会话归属的游戏ID
    GameId string `json:"gameId,omitempty"`
        // 会话ID
    GameSession string `json:"gameSession,omitempty"`
        // 会话当前状态：STARTED: 运行中，STOPPED: 已停止
    Status string `json:"status,omitempty"`
        // 会话所属用户ID
    AccountId string `json:"accountId,omitempty"`
        // 会话调度用户等级
    UserLevel int32 `json:"userLevel,omitempty"`
        // 会话调度大区ID
    RegionId string `json:"regionId,omitempty"`
        // 会话调度毫秒时间戳
    DispatchTime int64 `json:"dispatchTime,omitempty"`
        // 会话停止毫秒时间戳
    StopTime int64 `json:"stopTime,omitempty"`
        // 用户开始游戏毫秒时间戳
    PlayTime int64 `json:"playTime,omitempty"`
}
