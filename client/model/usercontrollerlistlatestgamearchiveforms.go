// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontrollerListLatestGameArchiveForms struct {
        // 用户id
    AccountId string `json:"accountId"`
        // 游戏Id
    GameId string `json:"gameId"`
        // 每页数量
    PageSize *int32 `json:"pageSize,omitempty"`
        // 页码
    PageNumber *int32 `json:"pageNumber,omitempty"`
        // 打标状态，1为打标，0为非打标
    TagStatus *int32 `json:"tagStatus,omitempty"`
}
