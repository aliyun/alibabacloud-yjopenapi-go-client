// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetTripleResult struct {
        // token失效时间戳,单位:秒
    Expired string `json:"expired,omitempty"`
        // 返回码
    Code string `json:"code,omitempty"`
        // 临时accessKey
    AccessKey string `json:"accessKey,omitempty"`
        // 返回信息
    Message string `json:"message,omitempty"`
        // 临时token
    Token string `json:"token,omitempty"`
        // 临时secretKey
    AccessSecret string `json:"accessSecret,omitempty"`
}
