// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UpdatePreopenStrategyForms struct {
    AppKey string `json:"appKey"`
    GameId string `json:"gameId"`
    PreStartCmd *string `json:"preStartCmd,omitempty"`
}
