// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminActivateDeploymentForms struct {
    GameId string `json:"gameId"`
    ProjectId string `json:"projectId"`
    VersionId string `json:"versionId"`
    MaxConcurrency *int32 `json:"maxConcurrency,omitempty"`
}
