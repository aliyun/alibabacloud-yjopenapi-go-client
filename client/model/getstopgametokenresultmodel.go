// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetStopGameTokenResultModel struct {
        // 当前token失效时间
    ExpireTime int64 `json:"expireTime,omitempty"`
        // 返回码
    Code string `json:"code,omitempty"`
        // 请求链路唯一标示
    RequestId string `json:"requestId,omitempty"`
        // 调度执行结果
    Success bool `json:"success,omitempty"`
        // 返回信息
    Message string `json:"message,omitempty"`
        // token
    Token string `json:"token,omitempty"`
}
