// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ReplaceSlotForms struct {
    ReplaceSession string `json:"replaceSession"`
    AccountId string `json:"accountId"`
    GameId string `json:"gameId"`
    AppKey string `json:"appKey"`
    BizParam *string `json:"bizParam,omitempty"`
    ClientIp *string `json:"clientIp,omitempty"`
    Tags *string `json:"tags,omitempty"`
    Codec *int32 `json:"codec,omitempty"`
    Resolution *int32 `json:"resolution,omitempty"`
    BitRate *int32 `json:"bitRate,omitempty"`
    Fps *int32 `json:"fps,omitempty"`
    GameCmdParam *string `json:"gameCmdParam,omitempty"`
    UserLevel *int32 `json:"userLevel,omitempty"`
    StartParam *ReplaceSlotFormsStartParam `json:"startParam,omitempty"`
}
