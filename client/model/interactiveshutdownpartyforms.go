// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type InteractiveShutDownPartyForms struct {
        // 派对id
    PartyId string `json:"partyId"`
        // 关闭原因
    ShutDownReason *string `json:"shutDownReason,omitempty"`
}
