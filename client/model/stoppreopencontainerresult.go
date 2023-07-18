// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type StopPreopenContainerResult struct {
        // 服务端状态码
    MsgCode string `json:"msgCode,omitempty"`
        // 服务端描述信息
    MsgInfo string `json:"msgInfo,omitempty"`
    Success bool `json:"success,omitempty"`
}
