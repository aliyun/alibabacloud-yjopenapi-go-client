// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayDemoQueryMultiPlayResultModel struct {
    NextTokenId string `json:"nextTokenId,omitempty"`
    CreatorTokenId string `json:"creatorTokenId,omitempty"`
    EndTs int64 `json:"endTs,omitempty"`
    Tokens []MultiplayDemoQueryMultiPlayResultModelTokens `json:"tokens,omitempty"`
    DataVer int64 `json:"dataVer,omitempty"`
    MpId string `json:"mpId,omitempty"`
    Config MultiplayDemoQueryMultiPlayResultModelConfig `json:"config,omitempty"`
}
