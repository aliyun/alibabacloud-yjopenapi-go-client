// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminSubmitDeploymentForms struct {
    GameId string `json:"gameId"`
    ProjectId string `json:"projectId"`
    VersionId string `json:"versionId"`
    CloudGameInstanceIds string `json:"cloudGameInstanceIds"`
    OperationType string `json:"operationType"`
}
