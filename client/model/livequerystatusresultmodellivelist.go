// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type LiveQueryStatusResultModelLiveList struct {
    StreamKey string `json:"streamKey,omitempty"`
    ServerUrl string `json:"serverUrl,omitempty"`
    Message string `json:"message,omitempty"`
    LiveId string `json:"liveId,omitempty"`
    Config LiveQueryStatusResultModelLiveListConfig `json:"config,omitempty"`
    Status string `json:"status,omitempty"`
}
