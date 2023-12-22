// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type LiveStartGameLiveForms struct {
    AppKey string `json:"appKey"`
    GameSession string `json:"gameSession"`
    ServerUrl string `json:"serverUrl"`
    StreamKey string `json:"streamKey"`
    Config *LiveStartGameLiveFormsConfig `json:"config,omitempty"`
}
