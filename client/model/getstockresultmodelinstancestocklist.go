// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetStockResultModelInstanceStockList struct {
        // 实例id
    InstanceId string `json:"instanceId,omitempty"`
        // 实例游戏当前可使用路数
    InstanceGameAvailableTotal int32 `json:"instanceGameAvailableTotal,omitempty"`
        // 实例总路数
    InstanceTotal int32 `json:"instanceTotal,omitempty"`
        // 实例可用路数
    InstanceAvailableTotal int32 `json:"instanceAvailableTotal,omitempty"`
        // 实例配置游戏总路数
    InstanceGameTotal int32 `json:"instanceGameTotal,omitempty"`
        // 实例已用路数
    InstanceUsedTotal int32 `json:"instanceUsedTotal,omitempty"`
        // 实例游戏当前已使用路数
    InstanceGameUsedTotal int32 `json:"instanceGameUsedTotal,omitempty"`
        // 实例大区ID
    InstanceRegionId string `json:"instanceRegionId,omitempty"`
        // 实例调度等级
    InstanceUserLevel int32 `json:"instanceUserLevel,omitempty"`
}
