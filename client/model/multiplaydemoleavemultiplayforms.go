// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayDemoLeaveMultiPlayForms struct {
    MpId *string `json:"mpId,omitempty"`
    TokenIds *string `json:"tokenIds,omitempty"`
    KickOut *bool `json:"kickOut,omitempty"`
    Reason *string `json:"reason,omitempty"`
    OpenAk *string `json:"openAk,omitempty"`
    OpenSk *string `json:"openSk,omitempty"`
}
