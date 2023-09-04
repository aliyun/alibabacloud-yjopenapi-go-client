// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayDemoInitMultiPlayResultModelTokens struct {
    AccountId string `json:"accountId,omitempty"`
    TokenId string `json:"tokenId,omitempty"`
    OfflineTs int64 `json:"offlineTs,omitempty"`
    Session string `json:"session,omitempty"`
    Admin bool `json:"admin,omitempty"`
    ControlId int32 `json:"controlId,omitempty"`
}
