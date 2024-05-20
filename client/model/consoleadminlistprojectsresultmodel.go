// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListProjectsResultModel struct {
    MaxResults int32 `json:"maxResults,omitempty"`
    NextToken string `json:"nextToken,omitempty"`
    DataList []ConsoleAdminListProjectsResultModelDataList `json:"dataList,omitempty"`
    Count int64 `json:"count,omitempty"`
}
