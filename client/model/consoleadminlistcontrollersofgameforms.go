// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListControllersOfGameForms struct {
    GameId string `json:"gameId"`
    NextToken *string `json:"nextToken,omitempty"`
    MaxResults *int32 `json:"maxResults,omitempty"`
}
