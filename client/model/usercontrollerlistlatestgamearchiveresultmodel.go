// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontrollerListLatestGameArchiveResultModel struct {
    PageNumber int32 `json:"pageNumber,omitempty"`
    RequestId string `json:"requestId,omitempty"`
    PageSize int32 `json:"pageSize,omitempty"`
    TotalCount int32 `json:"totalCount,omitempty"`
    Items []UsercontrollerListLatestGameArchiveResultModelItems `json:"items,omitempty"`
}
