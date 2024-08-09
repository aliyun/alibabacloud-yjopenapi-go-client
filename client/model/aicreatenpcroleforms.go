// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type AiCreateNpcRoleForms struct {
    Name string `json:"name"`
    Description string `json:"description"`
    FigureId int64 `json:"figureId"`
    VoiceId int64 `json:"voiceId"`
    Characters string `json:"characters"`
    Greetings string `json:"greetings"`
    Guidance []string `json:"guidance"`
    KnowledgeBases []AiCreateNpcRoleFormsKnowledgeBases `json:"knowledgeBases"`
    PluginIds *[]int64 `json:"pluginIds,omitempty"`
    ShortMemoryRound int32 `json:"shortMemoryRound"`
    TenantId int64 `json:"tenantId"`
    Operator string `json:"operator"`
}
