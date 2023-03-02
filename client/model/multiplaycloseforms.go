// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayCloseForms struct {
        // 联机唯一Id
    MpId string `json:"mpId"`
        // 关闭原因
    Reason *string `json:"reason,omitempty"`
}
