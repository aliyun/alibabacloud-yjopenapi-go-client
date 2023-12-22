// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontrollerListLatestGameArchiveForms struct {
    AccountId string `json:"accountId"`
    GameId string `json:"gameId"`
    PageSize *int32 `json:"pageSize,omitempty"`
    PageNumber *int32 `json:"pageNumber,omitempty"`
    TagStatus *int32 `json:"tagStatus,omitempty"`
}
