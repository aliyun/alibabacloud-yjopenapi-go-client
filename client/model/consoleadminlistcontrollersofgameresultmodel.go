// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListControllersOfGameResultModel struct {
    NextToken string `json:"nextToken,omitempty"`
    MaxResults int32 `json:"maxResults,omitempty"`
    DataList []ConsoleAdminListControllersOfGameResultModelDataList `json:"dataList,omitempty"`
    Count int64 `json:"count,omitempty"`
}
