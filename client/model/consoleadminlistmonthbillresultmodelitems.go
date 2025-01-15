// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListMonthBillResultModelItems struct {
        // 到期还款日期
    LatestPaidTime string `json:"latestPaidTime,omitempty"`
        // 租户名称
    TenantName string `json:"tenantName,omitempty"`
        // 租户ID
    TenantId int64 `json:"tenantId,omitempty"`
        // 账单状态
    BillStatus string `json:"billStatus,omitempty"`
        // 账期
    AccountingPeriod string `json:"accountingPeriod,omitempty"`
        // 已出账的日期
    BillOutgoingTime string `json:"billOutgoingTime,omitempty"`
        // 消费金额
    CashPayAmount int64 `json:"cashPayAmount,omitempty"`
}
