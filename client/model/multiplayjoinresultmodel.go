// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayJoinResultModel struct {
        // 当前联机数据版本号
    DataVer string `json:"dataVer,omitempty"`
        // 加入人令牌Id
    TokenId string `json:"tokenId,omitempty"`
        // 加入者sessionId
    Session string `json:"session,omitempty"`
}
