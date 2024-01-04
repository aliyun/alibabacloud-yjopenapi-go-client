// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayInitResultModel struct {
    CreatorTokenId string `json:"creatorTokenId,omitempty"`
    Tokens []MultiplayInitResultModelTokens `json:"tokens,omitempty"`
    DataVer string `json:"dataVer,omitempty"`
    MpId string `json:"mpId,omitempty"`
}
