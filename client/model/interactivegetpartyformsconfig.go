// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type InteractiveGetPartyFormsConfig struct {
        // 派对创建者断线保活时间
    OwnerDisconnectKeepAliveTime *int32 `json:"ownerDisconnectKeepAliveTime,omitempty"`
        // 加入派对玩家断线保活时间
    JoinerDisconnectKeepAliveTime *int32 `json:"joinerDisconnectKeepAliveTime,omitempty"`
        // 无操作保活时间
    NoActionKeepAliveTime *int32 `json:"noActionKeepAliveTime,omitempty"`
}
