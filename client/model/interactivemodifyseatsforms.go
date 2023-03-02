// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type InteractiveModifySeatsForms struct {
        // 派对id
    PartyId string `json:"partyId"`
        // 操作者
    Operator string `json:"operator"`
    ModifySeats []InteractiveModifySeatsFormsModifySeats `json:"modifySeats"`
}
