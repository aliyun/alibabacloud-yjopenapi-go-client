// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiListModelsForms struct {
    ModelType *string `json:"modelType,omitempty"`
    ModelFileType *string `json:"modelFileType,omitempty"`
    ModelSeries *string `json:"modelSeries,omitempty"`
    ModelFileName *string `json:"modelFileName,omitempty"`
    PageNumber int32 `json:"pageNumber"`
    PageSize int32 `json:"pageSize"`
}
