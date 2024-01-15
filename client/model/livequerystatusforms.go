// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type LiveQueryStatusForms struct {
    AppKey string `json:"appKey"`
    GameSession string `json:"gameSession"`
    LiveId *string `json:"liveId,omitempty"`
}
