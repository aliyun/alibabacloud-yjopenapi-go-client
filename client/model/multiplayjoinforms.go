// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayJoinForms struct {
        // 联机唯一Id
    MpId string `json:"mpId"`
        // 用户Id
    AccountId string `json:"accountId"`
        // 控制位
    ControlId *int32 `json:"controlId,omitempty"`
}
