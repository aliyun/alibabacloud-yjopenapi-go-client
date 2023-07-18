// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListVersionDeployInstancesForms struct {
        // 项目ID
    ProjectId *string `json:"projectId,omitempty"`
        // 游戏ID
    GameId *string `json:"gameId,omitempty"`
        // 版本ID
    VersionId *string `json:"versionId,omitempty"`
        // 查询的部署状态
    DeployStatus *string `json:"deployStatus,omitempty"`
}
