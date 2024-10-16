// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiUpdateNpcRoleForms struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    FigureId int64 `json:"figureId"`
    VoiceId int64 `json:"voiceId"`
    Characters string `json:"characters"`
    Greetings string `json:"greetings"`
    Guidance []string `json:"guidance"`
    KnowledgeBases *[]AiUpdateNpcRoleFormsKnowledgeBases `json:"knowledgeBases,omitempty"`
    PluginIds *[]int64 `json:"pluginIds,omitempty"`
    ShortMemoryRound int32 `json:"shortMemoryRound"`
    Operator string `json:"operator"`
    LlmModelId int64 `json:"llmModelId"`
}
