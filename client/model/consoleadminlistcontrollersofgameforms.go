// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListControllersOfGameForms struct {
        // 游戏id
    GameId string `json:"gameId"`
        // 标记当前开始读取的位置
    NextToken *string `json:"nextToken,omitempty"`
        // 本次读取的最大数据记录数量
    MaxResults *string `json:"maxResults,omitempty"`
}
