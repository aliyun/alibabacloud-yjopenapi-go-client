// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListGameCcuResultModel struct {
    PageNumber int32 `json:"pageNumber,omitempty"`
    PageSize int32 `json:"pageSize,omitempty"`
    TotalCount int64 `json:"totalCount,omitempty"`
    Items []ConsoleAdminListGameCcuResultModelItems `json:"items,omitempty"`
}
