// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminGetGameVersionProgressResultModel struct {
        // 提供不固定的额外信息
    Extra map[string]string `json:"extra,omitempty"`
    Description string `json:"description,omitempty"`
    Event string `json:"event,omitempty"`
    Status string `json:"status,omitempty"`
}
