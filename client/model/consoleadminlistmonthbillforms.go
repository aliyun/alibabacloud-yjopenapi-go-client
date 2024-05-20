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
    PageNumber *int32 `json:"pageNumber,omitempty"`
    PageSize *int32 `json:"pageSize,omitempty"`
}
