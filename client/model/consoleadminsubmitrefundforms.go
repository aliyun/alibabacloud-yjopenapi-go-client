// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminSubmitRefundForms struct {
        // 商品编码
    CommodityCode string `json:"commodityCode"`
        // 实例ID
    InstanceId string `json:"instanceId"`
        // 退订类型
    RefundType string `json:"refundType"`
        // 退订理由
    RefundReason *string `json:"refundReason,omitempty"`
        // 退订理由类型
    RefundReasonType *string `json:"refundReasonType,omitempty"`
        // 退订后的过期时间 退续费场景使用
    TargetExpireTime *string `json:"targetExpireTime,omitempty"`
}
