// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayJoinForms struct {
    MpId string `json:"mpId"`
    AccountId string `json:"accountId"`
    ControlId *int32 `json:"controlId,omitempty"`
}
