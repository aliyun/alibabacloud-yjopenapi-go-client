// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type TryToGetSlotsFormsRequests struct {
    GameId string `json:"gameId,omitempty"`
    Fps int32 `json:"fps,omitempty"`
    ReConnect bool `json:"reConnect,omitempty"`
    Resolution int32 `json:"resolution,omitempty"`
    Tags string `json:"tags,omitempty"`
    GameCmdParam string `json:"gameCmdParam,omitempty"`
    AccountId string `json:"accountId,omitempty"`
    Codec int32 `json:"codec,omitempty"`
    UserLevel int32 `json:"userLevel,omitempty"`
    StartParam TryToGetSlotsFormsRequestsStartParam `json:"startParam,omitempty"`
    RegionId string `json:"regionId,omitempty"`
    BitRate int32 `json:"bitRate,omitempty"`
    ClientIp string `json:"clientIp,omitempty"`
    BizParam string `json:"bizParam,omitempty"`
}
