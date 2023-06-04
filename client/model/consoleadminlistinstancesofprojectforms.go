// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListInstancesOfProjectForms struct {
    NextToken *string `json:"nextToken,omitempty"`
    MaxResult *int64 `json:"maxResult,omitempty"`
    ProjectId string `json:"projectId"`
}
