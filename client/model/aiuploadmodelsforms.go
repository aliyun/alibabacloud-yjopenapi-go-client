// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiUploadModelsForms struct {
    ModelType string `json:"modelType"`
    ModelFileType string `json:"modelFileType"`
    ModelSeries string `json:"modelSeries"`
    ModelFileName string `json:"modelFileName"`
    Desc string `json:"desc"`
    Url string `json:"url"`
}
