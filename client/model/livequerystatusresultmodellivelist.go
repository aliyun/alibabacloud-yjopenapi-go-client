// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type LiveQueryStatusResultModelLiveList struct {
        // 推流标识ID
    LiveId string `json:"liveId,omitempty"`
        // 推流服务地址
    ServerUrl string `json:"serverUrl,omitempty"`
        // 推流鉴权参数
    StreamKey string `json:"streamKey,omitempty"`
        // 推流状态
    Status string `json:"status,omitempty"`
        // 推流状态消息
    Message string `json:"message,omitempty"`
    Config LiveQueryStatusResultModelLiveListConfig `json:"config,omitempty"`
}
