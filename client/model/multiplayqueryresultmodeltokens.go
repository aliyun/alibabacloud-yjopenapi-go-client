// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayQueryResultModelTokens struct {
        // 联机令牌Id
    TokenId string `json:"tokenId,omitempty"`
        // 玩家控制位
    ControlId int32 `json:"controlId,omitempty"`
        // 玩家用户Id
    AccountId string `json:"accountId,omitempty"`
        // 玩家sessionId
    Session string `json:"session,omitempty"`
}
