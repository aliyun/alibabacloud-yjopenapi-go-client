// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontollerListLatestGameArchiveResultModelItems struct {
        // 用户ID
    AccountId string `json:"accountId,omitempty"`
        // 游戏ID
    GameId string `json:"gameId,omitempty"`
        // 存档ID
    ArchiveId string `json:"archiveId,omitempty"`
        // 存档时间，时间戳
    ArchiveTime int64 `json:"archiveTime,omitempty"`
        // 打标状态
    TagStatus int32 `json:"tagStatus,omitempty"`
}
