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
    PageSize *string `json:"pageSize,omitempty"`
    PageNumber *string `json:"pageNumber,omitempty"`
    TagStatus *int32 `json:"tagStatus,omitempty"`
}
