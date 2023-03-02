// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminSubmitDeploymentForms struct {
        // 游戏id
    GameId string `json:"gameId"`
        // 项目id
    ProjectId string `json:"projectId"`
        // 版本id
    VersionId string `json:"versionId"`
        // 实例id列表
    CloudGameInstanceIds string `json:"cloudGameInstanceIds"`
        // 操作类型
    OperationType string `json:"operationType"`
}
