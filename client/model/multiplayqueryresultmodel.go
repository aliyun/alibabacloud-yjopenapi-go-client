// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayQueryResultModel struct {
    CreatorTokenId string `json:"creatorTokenId,omitempty"`
    EndTs int64 `json:"endTs,omitempty"`
    Tokens []MultiplayQueryResultModelTokens `json:"tokens,omitempty"`
    DataVer string `json:"dataVer,omitempty"`
    MpId string `json:"mpId,omitempty"`
    Config MultiplayQueryResultModelConfig `json:"config,omitempty"`
}
