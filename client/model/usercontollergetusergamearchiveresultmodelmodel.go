// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontollerGetUserGameArchiveResultModelModel struct {
        // 游戏ID
    GameId int64 `json:"gameId,omitempty"`
        // 游戏会话ID
    GameSessionId string `json:"gameSessionId,omitempty"`
        // 租户ID
    TenantId int64 `json:"tenantId,omitempty"`
        // 项目ID
    ProjectId int64 `json:"projectId,omitempty"`
        // 存档生成时间
    GmtCreate int64 `json:"gmtCreate,omitempty"`
}
