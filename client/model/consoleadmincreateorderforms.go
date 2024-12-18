// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminCreateOrderForms struct {
        // 类目编码
    CategoryCode string `json:"categoryCode"`
        // 商品编码
    CommodityCode string `json:"commodityCode"`
        // 实例ID
    InstanceId *string `json:"instanceId,omitempty"`
        // 主计费项编码
    PrimaryChargeItemCode string `json:"primaryChargeItemCode"`
    AttributeRequestList []ConsoleAdminCreateOrderFormsAttributeRequestList `json:"attributeRequestList"`
        // 订单类型
    OrderType string `json:"orderType"`
        // 是否开启自动续费
    AutoRenew *bool `json:"autoRenew,omitempty"`
        // 下单扩展信息
    CreateOrderExtParams *string `json:"createOrderExtParams,omitempty"`
}
