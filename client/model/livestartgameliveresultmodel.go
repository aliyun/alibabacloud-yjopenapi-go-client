// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type LiveStartGameLiveResultModel struct {
        // 推流结果
    Data bool `json:"data,omitempty"`
        // 推流标识ID
    LiveId string `json:"liveId,omitempty"`
        // 推流状态
    Status string `json:"status,omitempty"`
}
