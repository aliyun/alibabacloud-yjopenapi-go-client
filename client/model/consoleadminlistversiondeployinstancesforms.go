// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListVersionDeployInstancesForms struct {
    ProjectId *string `json:"projectId,omitempty"`
    GameId *string `json:"gameId,omitempty"`
    VersionId *string `json:"versionId,omitempty"`
    DeployStatus *string `json:"deployStatus,omitempty"`
}
