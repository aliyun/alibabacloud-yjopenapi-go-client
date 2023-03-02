// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayInitResultModelTokens struct {
        // 用户联机标识Id
    TokenId string `json:"tokenId,omitempty"`
        // 控制位
    ControlId int32 `json:"controlId,omitempty"`
        // 用户Id
    AccountId string `json:"accountId,omitempty"`
        // 当前用户连接session
    Session string `json:"session,omitempty"`
}
