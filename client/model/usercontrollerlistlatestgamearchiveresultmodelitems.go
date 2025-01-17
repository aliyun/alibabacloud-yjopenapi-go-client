// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontrollerListLatestGameArchiveResultModelItems struct {
    GameId string `json:"gameId,omitempty"`
    AccountId string `json:"accountId,omitempty"`
    ArchiveTime int64 `json:"archiveTime,omitempty"`
    TagStatus int32 `json:"tagStatus,omitempty"`
    ArchiveId string `json:"archiveId,omitempty"`
}
