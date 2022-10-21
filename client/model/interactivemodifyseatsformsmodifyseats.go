// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type InteractiveModifySeatsFormsModifySeats struct {
        // 被操作用户Id
    UserId string `json:"userId"`
        // 派对坐次
    SeatId int32 `json:"seatId"`
        // 控制位
    ControlId int32 `json:"controlId"`
}
