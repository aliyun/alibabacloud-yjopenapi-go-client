// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListActivatedInstancesResultModel struct {
        // 总记录数
    Count int64 `json:"count,omitempty"`
        // 数据列表
    DataList []ConsoleAdminListActivatedInstancesResultModelDataList `json:"dataList,omitempty"`
}
