// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontrollerGetUserGameArchiveResultModelUserGameArchiveDto struct {
        // 游戏会话ID
    GameSessionId string `json:"gameSessionId,omitempty"`
        // 存档生成时间
    GmtCreate int64 `json:"gmtCreate,omitempty"`
}
