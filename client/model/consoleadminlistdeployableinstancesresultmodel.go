// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListDeployableInstancesResultModel struct {
    PageNumber int64 `json:"pageNumber,omitempty"`
    DataList []ConsoleAdminListDeployableInstancesResultModelDataList `json:"dataList,omitempty"`
    PageSize int64 `json:"pageSize,omitempty"`
    TotalCount int64 `json:"totalCount,omitempty"`
}
