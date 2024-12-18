// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListGamesResultModel struct {
    MaxResults int32 `json:"maxResults,omitempty"`
    NextToken string `json:"nextToken,omitempty"`
    DataList []ConsoleAdminListGamesResultModelDataList `json:"dataList,omitempty"`
    Count string `json:"count,omitempty"`
}
