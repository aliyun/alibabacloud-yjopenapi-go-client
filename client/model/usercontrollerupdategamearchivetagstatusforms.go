// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontrollerUpdateGameArchiveTagStatusForms struct {
        // 用户id
    AccountId string `json:"accountId"`
        // 游戏Id
    GameId string `json:"gameId"`
        // 存档ID
    ArchiveId string `json:"archiveId"`
        // 打标状态，1为打标，0为非打标
    TagStatus int32 `json:"tagStatus"`
}
