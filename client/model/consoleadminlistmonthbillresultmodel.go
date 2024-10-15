// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListMonthBillResultModel struct {
        // 页码
    PageNumber int32 `json:"pageNumber,omitempty"`
        // 没页数量
    PageSize int32 `json:"pageSize,omitempty"`
        // 总数量
    TotalCount int64 `json:"totalCount,omitempty"`
        // 列表
    Items []ConsoleAdminListMonthBillResultModelItems `json:"items,omitempty"`
}
