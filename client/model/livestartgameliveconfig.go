// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type LiveStartGameLiveConfig struct {
        // 分辨率
    Resolution *string `json:"resolution,omitempty"`
        // 帧率
    FrameRate *int32 `json:"frameRate,omitempty"`
        // 码率
    Bitrate *int32 `json:"bitrate,omitempty"`
}
