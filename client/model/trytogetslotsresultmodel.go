// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type TryToGetSlotsResultModel struct {
    Code string `json:"code,omitempty"`
    Success bool `json:"success,omitempty"`
    Responses []TryToGetSlotsResultModelResponses `json:"responses,omitempty"`
    Message string `json:"message,omitempty"`
}
