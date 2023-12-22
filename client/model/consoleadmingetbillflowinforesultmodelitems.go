// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminGetBillFlowInfoResultModelItems struct {
    ConsumePeriodStartTime int64 `json:"consumePeriodStartTime,omitempty"`
    TotalUnclearAmount string `json:"totalUnclearAmount,omitempty"`
    Amount int64 `json:"amount,omitempty"`
    OrderId string `json:"orderId,omitempty"`
    BillType string `json:"billType,omitempty"`
    DiscountAmount string `json:"discountAmount,omitempty"`
    ActualTotalPayAmount int64 `json:"actualTotalPayAmount,omitempty"`
    RefundOrderId string `json:"refundOrderId,omitempty"`
    ConsumePeriodEndTime int64 `json:"consumePeriodEndTime,omitempty"`
    PayAmount int64 `json:"payAmount,omitempty"`
    TenantName string `json:"tenantName,omitempty"`
    SettlementStatus string `json:"settlementStatus,omitempty"`
    VoucherPayAmount string `json:"voucherPayAmount,omitempty"`
    AccountingPeriod string `json:"accountingPeriod,omitempty"`
    ConsumeType string `json:"consumeType,omitempty"`
    CommodityName string `json:"commodityName,omitempty"`
}
