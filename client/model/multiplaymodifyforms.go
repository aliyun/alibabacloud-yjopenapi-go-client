// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayModifyForms struct {
    MpId string `json:"mpId"`
    Tokens *[]MultiplayModifyFormsTokens `json:"tokens,omitempty"`
}
