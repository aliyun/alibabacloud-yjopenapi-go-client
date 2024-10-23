// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListInstancesOfProjectResultModel struct {
    MaxResult int32 `json:"maxResult,omitempty"`
    NextToken string `json:"nextToken,omitempty"`
    DataList []ConsoleAdminListInstancesOfProjectResultModelDataList `json:"dataList,omitempty"`
    Count int32 `json:"count,omitempty"`
}
