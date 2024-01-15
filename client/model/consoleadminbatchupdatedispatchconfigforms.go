// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminBatchUpdateDispatchConfigForms struct {
        // 项目id(混淆)
    MixProjectId string `json:"mixProjectId"`
        // 实例id
    InstanceId string `json:"instanceId"`
        // 接入方唯一标识
    AppName string `json:"appName"`
    ConfigList []ConsoleAdminBatchUpdateDispatchConfigFormsConfigList `json:"configList"`
}
