// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiBatchUpdateKbVersionForms struct {
    RoleIds []int64 `json:"roleIds"`
    KbId int64 `json:"kbId"`
    KbVersionId int64 `json:"kbVersionId"`
    Operator string `json:"operator"`
}
