// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListVersionDeployInstancesModel struct {
        // 总记录数
    Count int64 `json:"count,omitempty"`
        // 数据列表
    DataList []ConsoleAdminListVersionDeployInstancesModelDataList `json:"dataList,omitempty"`
}
