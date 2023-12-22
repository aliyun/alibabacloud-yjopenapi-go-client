// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminGetBillFlowInfoForms struct {
    AccountingPeriodFrom string `json:"accountingPeriodFrom"`
    AccountingPeriodTo string `json:"accountingPeriodTo"`
    CommodityCode *string `json:"commodityCode,omitempty"`
    TenantId string `json:"tenantId"`
    OrderId *string `json:"orderId,omitempty"`
    Status *string `json:"status,omitempty"`
    ConsumeType *string `json:"consumeType,omitempty"`
    BillType *string `json:"billType,omitempty"`
    PageNumber *int32 `json:"pageNumber,omitempty"`
    PageSize *int32 `json:"pageSize,omitempty"`
}
