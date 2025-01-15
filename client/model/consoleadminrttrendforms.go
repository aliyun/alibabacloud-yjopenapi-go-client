// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminRtTrendForms struct {
    GroupType *string `json:"groupType,omitempty"`
    IndicatorTypes *string `json:"indicatorTypes,omitempty"`
    MixGameIds *string `json:"mixGameIds,omitempty"`
    ProjectIds *string `json:"projectIds,omitempty"`
    QueryTimeStart *int64 `json:"queryTimeStart,omitempty"`
    QueryTimeEnd *int64 `json:"queryTimeEnd,omitempty"`
}
