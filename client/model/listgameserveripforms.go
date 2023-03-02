// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ListGameServerIpForms struct {
        // 分页大小
    PageSize *int64 `json:"pageSize,omitempty"`
        // 分页标识
    NextToken *string `json:"nextToken,omitempty"`
}
