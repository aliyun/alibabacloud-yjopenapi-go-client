// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type MultiplayInitForms struct {
    GameSession string `json:"gameSession"`
    AppKey string `json:"appKey"`
    Config *MultiplayInitFormsConfig `json:"config,omitempty"`
    Tokens *[]MultiplayInitFormsTokens `json:"tokens,omitempty"`
}
