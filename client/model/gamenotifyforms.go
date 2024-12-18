// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GameNotifyForms struct {
        // appKey
    AppKey string `json:"appKey"`
    GameSession string `json:"gameSession"`
    Type_ string `json:"type"`
    Value *string `json:"value,omitempty"`
}
