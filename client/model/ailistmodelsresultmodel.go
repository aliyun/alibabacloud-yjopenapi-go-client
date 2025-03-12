// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiListModelsResultModel struct {
    ModelList AiListModelsResultModelModelList `json:"modelList,omitempty"`
    RequestId string `json:"requestId,omitempty"`
}
