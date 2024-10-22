// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiGetPromptResultModel struct {
    QueueNum int64 `json:"queueNum,omitempty"`
    Response string `json:"response,omitempty"`
    RequestId string `json:"requestId,omitempty"`
    ExecutedList string `json:"executedList,omitempty"`
    ExecutionResult string `json:"executionResult,omitempty"`
    Resources []AiGetPromptResultModelResources `json:"resources,omitempty"`
    Status string `json:"status,omitempty"`
}
