// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiUploadModelsForms struct {
    ModelType string `json:"modelType"`
    ModelName string `json:"modelName"`
    ModelFilePath string `json:"modelFilePath"`
    ModelFiles []AiUploadModelsFormsModelFiles `json:"modelFiles"`
    Desc *string `json:"desc,omitempty"`
    ModelFileType string `json:"modelFileType"`
}
