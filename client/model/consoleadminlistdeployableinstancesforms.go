// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListDeployableInstancesForms struct {
        // 版本id
    VersionId string `json:"versionId"`
        // 项目id
    ProjectId string `json:"projectId"`
        // 每页大小
    PageSize *int64 `json:"pageSize,omitempty"`
        // 页码
    PageNumber *int64 `json:"pageNumber,omitempty"`
}
