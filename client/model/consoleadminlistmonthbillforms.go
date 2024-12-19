// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListMonthBillForms struct {
    AccountingPeriodFrom string `json:"accountingPeriodFrom"`
    AccountingPeriodTo string `json:"accountingPeriodTo"`
        // 状态
    Status *string `json:"status,omitempty"`
    PageNumber *int32 `json:"pageNumber,omitempty"`
    PageSize *int32 `json:"pageSize,omitempty"`
}
