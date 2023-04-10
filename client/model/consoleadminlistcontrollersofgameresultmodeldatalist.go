// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListControllersOfGameResultModelDataList struct {
        // 控制器ID
    Id string `json:"id,omitempty"`
        // 名称
    Name string `json:"name,omitempty"`
        // 背景图
    BgPic string `json:"bgPic,omitempty"`
        // 配置
    Config string `json:"config,omitempty"`
        // 方案类型
    Type_ int32 `json:"type,omitempty"`
        // 控制器类型
    Ctltype string `json:"ctltype,omitempty"`
        // 优先级
    Priority int32 `json:"priority,omitempty"`
}
