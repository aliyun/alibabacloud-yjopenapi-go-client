// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayModifyForms struct {
        // 联机唯一Id
    MpId string `json:"mpId"`
    Tokens *[]MultiplayModifyTokens `json:"tokens,omitempty"`
}
