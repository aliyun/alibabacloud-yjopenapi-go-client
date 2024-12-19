// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiListModelsResultModelModelList struct {
    PageNumber string `json:"pageNumber,omitempty"`
    PageSize string `json:"pageSize,omitempty"`
    TotalCount int64 `json:"totalCount,omitempty"`
    Items []AiListModelsResultModelModelListItems `json:"items,omitempty"`
}
