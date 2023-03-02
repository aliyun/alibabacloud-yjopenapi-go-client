// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayQueryResultModel struct {
        // 当前联机数据版本号
    DataVer string `json:"dataVer,omitempty"`
        // 联机Id
    MpId string `json:"mpId,omitempty"`
    Config MultiplayQueryResultModelConfig `json:"config,omitempty"`
        // 创建者tokenId
    CreatorTokenId string `json:"creatorTokenId,omitempty"`
        // 联机结束时间
    EndTs int64 `json:"endTs,omitempty"`
        // 当前联机玩家列表
    Tokens []MultiplayQueryResultModelTokens `json:"tokens,omitempty"`
}
