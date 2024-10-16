// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListGameCcuForms struct {
    MixGameIds *string `json:"mixGameIds,omitempty"`
    ProjectIds *string `json:"projectIds,omitempty"`
    QueryTimeStart *int64 `json:"queryTimeStart,omitempty"`
    QueryTimeEnd *int64 `json:"queryTimeEnd,omitempty"`
    PageNumber *int32 `json:"pageNumber,omitempty"`
    PageSize *int32 `json:"pageSize,omitempty"`
}
