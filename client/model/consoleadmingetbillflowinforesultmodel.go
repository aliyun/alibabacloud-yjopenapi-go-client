// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminGetBillFlowInfoResultModel struct {
    PageNumber int32 `json:"pageNumber,omitempty"`
    Pagesize int32 `json:"pagesize,omitempty"`
    TotalCount int32 `json:"totalCount,omitempty"`
    Items []ConsoleAdminGetBillFlowInfoResultModelItems `json:"items,omitempty"`
}
