// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminGetOrderResultModel struct {
        // 订单类型
    OrderType string `json:"orderType,omitempty"`
        // 订单号
    OrderId string `json:"orderId,omitempty"`
        // 配置详情
    AttributeList []ConsoleAdminGetOrderResultModelAttributeList `json:"attributeList,omitempty"`
        // 付费方式
    ChargeType string `json:"chargeType,omitempty"`
        // 优惠金额（单位：分）
    DiscountAmount int64 `json:"discountAmount,omitempty"`
        // 信控支付金额（单位：分）
    CreditPayAmount int64 `json:"creditPayAmount,omitempty"`
        // 降配退款金额
    DowngradeRefundAmount int64 `json:"downgradeRefundAmount,omitempty"`
        // 降配信控应退款金额
    DowngradeCreditRefundAmount int64 `json:"downgradeCreditRefundAmount,omitempty"`
        // 优惠详细信息
    DiscountDetail string `json:"discountDetail,omitempty"`
        // 优惠名称
    PromotionName string `json:"promotionName,omitempty"`
        // 实例ID
    InstanceId string `json:"instanceId,omitempty"`
        // 应付金额（单位：分）
    PayAmount int64 `json:"payAmount,omitempty"`
        // 租户名称
    TenantName string `json:"tenantName,omitempty"`
        // 商品编码
    CommodityCode string `json:"commodityCode,omitempty"`
        // 是否自动续费
    AutoRenew bool `json:"autoRenew,omitempty"`
        // 币种
    Currency string `json:"currency,omitempty"`
        // 开始时间
    StartTime int64 `json:"startTime,omitempty"`
        // 降配现金应退款金额
    DowngradeBalanceRefundAmount int64 `json:"downgradeBalanceRefundAmount,omitempty"`
        // 主计费项价格类型
    PrimaryPriceType string `json:"primaryPriceType,omitempty"`
        // 订单完成时间
    FinishTime int64 `json:"finishTime,omitempty"`
        // 订单总金额（单位：分）
    Amount int64 `json:"amount,omitempty"`
        // 数量
    Quantity int32 `json:"quantity,omitempty"`
        // 发货完成时间
    DeliveryEndTime int64 `json:"deliveryEndTime,omitempty"`
        // 售中退款时间
    RefundTime int64 `json:"refundTime,omitempty"`
        // 购买时长单位
    BuyDurationUnit string `json:"buyDurationUnit,omitempty"`
        // 主计费项编码
    PrimaryChargeItemCode string `json:"primaryChargeItemCode,omitempty"`
        // 支付用户ID
    PayUserId string `json:"payUserId,omitempty"`
        // 类目编码
    CategoryCode string `json:"categoryCode,omitempty"`
        // 购买时长
    BuyDuration int64 `json:"buyDuration,omitempty"`
        // 余额支付金额（单位：分）
    BalancePayAmount int64 `json:"balancePayAmount,omitempty"`
        // 实付金额（单位：分）
    ActualPayAmount int64 `json:"actualPayAmount,omitempty"`
        // 订单创建时间
    CreateTime int64 `json:"createTime,omitempty"`
        // 租户ID
    TenantId int64 `json:"tenantId,omitempty"`
        // 支付完成时间
    PaymentEndTime int64 `json:"paymentEndTime,omitempty"`
        // 售卖类型
    BuyType string `json:"buyType,omitempty"`
        // 代金券抵扣金额（单位：分）
    VoucherPayAmount int64 `json:"voucherPayAmount,omitempty"`
        // 下单用户ID
    BuyerUserId string `json:"buyerUserId,omitempty"`
        // 截止时间
    EndTime int64 `json:"endTime,omitempty"`
        // 商品名称
    CommodityName string `json:"commodityName,omitempty"`
        // 订单状态
    Status string `json:"status,omitempty"`
}
