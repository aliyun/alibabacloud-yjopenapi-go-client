// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayDemoInitMultiPlayResultModel struct {
    CreatorTokenId string `json:"creatorTokenId,omitempty"`
    Tokens []MultiplayDemoInitMultiPlayResultModelTokens `json:"tokens,omitempty"`
    DataVer int64 `json:"dataVer,omitempty"`
    MpId string `json:"mpId,omitempty"`
}
