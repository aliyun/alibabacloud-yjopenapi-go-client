// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontrollerListLatestGameArchiveResultModel struct {
    PageNumber string `json:"pageNumber,omitempty"`
    RequestId string `json:"requestId,omitempty"`
    PageSize string `json:"pageSize,omitempty"`
    TotalCount string `json:"totalCount,omitempty"`
    Items []UsercontrollerListLatestGameArchiveResultModelItems `json:"items,omitempty"`
}
