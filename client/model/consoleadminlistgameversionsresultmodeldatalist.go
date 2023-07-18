// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListGameVersionsResultModelDataList struct {
        // 版本id
    VersionId string `json:"versionId,omitempty"`
        // 版本名称
    VersionName string `json:"versionName,omitempty"`
        // 适配评测状态
    AdaptState string `json:"adaptState,omitempty"`
        // 适配评测完成时间戳
    AdaptFinishTime int64 `json:"adaptFinishTime,omitempty"`
}
