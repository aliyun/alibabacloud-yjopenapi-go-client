// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListGameVersionsForms struct {
        // 标记当前开始读取的位置
    NextToken *string `json:"nextToken,omitempty"`
        // 本次读取的最大数据记录数量
    MaxResults *int32 `json:"maxResults,omitempty"`
        // 游戏ID
    GameId string `json:"gameId"`
}
