// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetPairResult struct {
        // 临时token
    Token string `json:"token,omitempty"`
        // 临时secretKey
    AccessSecret string `json:"accessSecret,omitempty"`
        // token失效时间戳,单位:秒
    Expired string `json:"expired,omitempty"`
        // 返回码
    Code string `json:"code,omitempty"`
        // 返回信息
    Message string `json:"message,omitempty"`
}
