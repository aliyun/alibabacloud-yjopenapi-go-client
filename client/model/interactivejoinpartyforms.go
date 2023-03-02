// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type InteractiveJoinPartyForms struct {
        // 派对id
    PartyId string `json:"partyId"`
        // 用户ID
    UserId string `json:"userId"`
        // 派对坐次
    SeatId int32 `json:"seatId"`
        // 项目
    ControlId int32 `json:"controlId"`
    ReConnect int32 `json:"reConnect"`
}
