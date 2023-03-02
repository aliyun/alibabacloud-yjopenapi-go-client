// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListGamesForms struct {
        // 游戏名称
    NextToken *string `json:"nextToken,omitempty"`
        // 本次读取的最大数据记录数量
    MaxResults *int32 `json:"maxResults,omitempty"`
}
