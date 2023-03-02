// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayLeaveForms struct {
        // 联机唯一Id
    MpId string `json:"mpId"`
        // 是否踢出
    KickOut bool `json:"kickOut"`
        // 踢出原因
    Reason *string `json:"reason,omitempty"`
    TokenIds []string `json:"tokenIds"`
}
