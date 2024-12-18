// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AdaptCreateAndSubmitAllForms struct {
    GameId *int32 `json:"gameId,omitempty"`
    GameName *string `json:"gameName,omitempty"`
    GameVersionId *int32 `json:"gameVersionId,omitempty"`
    GameVersion *string `json:"gameVersion,omitempty"`
    ResolutionList *string `json:"resolutionList,omitempty"`
    FrameRateList *string `json:"frameRateList,omitempty"`
    PlatformType int32 `json:"platformType"`
    SourcePlatform int32 `json:"sourcePlatform"`
    Records string `json:"records"`
    MixGameVersionId string `json:"mixGameVersionId"`
    MixGameId string `json:"mixGameId"`
}
