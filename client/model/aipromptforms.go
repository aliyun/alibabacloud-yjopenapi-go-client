// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiPromptForms struct {
    SceneType string `json:"sceneType"`
    PromptParam string `json:"promptParam"`
    Resources *[]AiPromptFormsResources `json:"resources,omitempty"`
}
