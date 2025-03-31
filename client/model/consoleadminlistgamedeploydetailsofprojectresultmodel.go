// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListGameDeployDetailsOfProjectResultModel struct {
    NextToken string `json:"nextToken,omitempty"`
    MaxResults int32 `json:"maxResults,omitempty"`
    DataList []ConsoleAdminListGameDeployDetailsOfProjectResultModelDataList `json:"dataList,omitempty"`
    Count int64 `json:"count,omitempty"`
}
