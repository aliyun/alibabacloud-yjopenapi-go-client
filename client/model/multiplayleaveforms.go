// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayLeaveForms struct {
    MpId string `json:"mpId"`
    KickOut bool `json:"kickOut"`
    Reason *string `json:"reason,omitempty"`
    TokenIds []string `json:"tokenIds"`
}
