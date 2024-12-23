// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminQueryAdaptResultByVersionIdResultModel struct {
    AdaptFileSourceName string `json:"adaptFileSourceName,omitempty"`
    DictFrameRate string `json:"dictFrameRate,omitempty"`
    DictMachineTypeVm string `json:"dictMachineTypeVm,omitempty"`
    ContainerType int32 `json:"containerType,omitempty"`
    Sandbox int32 `json:"sandbox,omitempty"`
    PlatformType int32 `json:"platformType,omitempty"`
    Priority int32 `json:"priority,omitempty"`
    Resolution string `json:"resolution,omitempty"`
    AdaptFileSource string `json:"adaptFileSource,omitempty"`
    Required int32 `json:"required,omitempty"`
    MaxConcurrency int32 `json:"maxConcurrency,omitempty"`
    State int32 `json:"state,omitempty"`
    PipelineType int32 `json:"pipelineType,omitempty"`
    DictPicQuality string `json:"dictPicQuality,omitempty"`
    DictIsvName string `json:"dictIsvName,omitempty"`
    CalculationEvaluation map[string]string `json:"calculationEvaluation,omitempty"`
}
