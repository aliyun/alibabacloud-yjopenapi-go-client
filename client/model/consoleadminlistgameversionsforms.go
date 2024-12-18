// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListGameVersionsForms struct {
    NextToken *string `json:"nextToken,omitempty"`
    MaxResults *int32 `json:"maxResults,omitempty"`
    GameId string `json:"gameId"`
}
