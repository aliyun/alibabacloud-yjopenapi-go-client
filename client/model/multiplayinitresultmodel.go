// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayInitResultModel struct {
        // 联机唯一Id
    MpId string `json:"mpId,omitempty"`
        // 当前联机数据版本号
    DataVer string `json:"dataVer,omitempty"`
        // 创建者令牌Id
    CreatorTokenId string `json:"creatorTokenId,omitempty"`
        // 联机所有人员信息
    Tokens []MultiplayInitResultModelTokens `json:"tokens,omitempty"`
}
