// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type InteractiveKickOutUserForms struct {
        // 派对id
    PartyId string `json:"partyId"`
        // 用户Id
    UserId string `json:"userId"`
        // 踢出原因
    KickOutReason *string `json:"kickOutReason,omitempty"`
}
