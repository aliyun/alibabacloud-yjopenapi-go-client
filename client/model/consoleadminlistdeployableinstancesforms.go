// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListDeployableInstancesForms struct {
    VersionId string `json:"versionId"`
    ProjectId string `json:"projectId"`
    PageSize *int64 `json:"pageSize,omitempty"`
    PageNumber *int64 `json:"pageNumber,omitempty"`
}
