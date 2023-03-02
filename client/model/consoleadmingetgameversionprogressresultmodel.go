// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminGetGameVersionProgressResultModel struct {
        // 当前所处阶段
    Event string `json:"event,omitempty"`
        // 当前状态
    Status string `json:"status,omitempty"`
        // 表示当前调用返回读取到的位置，空代表数据已经读取完毕
    Description string `json:"description,omitempty"`
        // 提供不固定的额外信息
    Extra map[string]string `json:"extra,omitempty"`
}
