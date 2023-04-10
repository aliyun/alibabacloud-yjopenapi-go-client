// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminActivateDeploymentForms struct {
        // 游戏id
    GameId string `json:"gameId"`
        // 项目id
    ProjectId string `json:"projectId"`
        // 版本id
    VersionId string `json:"versionId"`
        // 期望的激活后可调度的总最大并发数不小于该参数
    MaxConcurrency *int32 `json:"maxConcurrency,omitempty"`
}
