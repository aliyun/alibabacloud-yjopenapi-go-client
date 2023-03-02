// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListDeployableInstancesResultModel struct {
        // 本次请求条件下的数据总量
    TotalCount int64 `json:"totalCount,omitempty"`
        // 表示当前调用返回读取到的位置
    PageNumber int64 `json:"pageNumber,omitempty"`
        // 本次请求所返回的最大记录条数
    PageSize int64 `json:"pageSize,omitempty"`
        // 数据列表
    DataList []ConsoleAdminListDeployableInstancesResultModelDataList `json:"dataList,omitempty"`
}
