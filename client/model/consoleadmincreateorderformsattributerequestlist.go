// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminCreateOrderFormsAttributeRequestList struct {
        // 属性编码
    AttributeCode string `json:"attributeCode"`
        // 属性值(区间型专用)
    AttributeValue int32 `json:"attributeValue,omitempty"`
        // 属性值编码(枚举型，时长型 专用)
    AttributeValueCode int64 `json:"attributeValueCode,omitempty"`
        // 时长单位
    TimeUnit string `json:"timeUnit,omitempty"`
}
