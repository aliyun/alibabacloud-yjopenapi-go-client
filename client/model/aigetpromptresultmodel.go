// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiGetPromptResultModel struct {
    PromptResult string `json:"promptResult,omitempty"`
    RequestId string `json:"requestId,omitempty"`
    PromptResponse string `json:"promptResponse,omitempty"`
    Resources []AiGetPromptResultModelResources `json:"resources,omitempty"`
}
