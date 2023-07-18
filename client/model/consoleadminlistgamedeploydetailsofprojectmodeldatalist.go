// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListGameDeployDetailsOfProjectModelDataList struct {
        // 版本ID
    VersionId string `json:"versionId,omitempty"`
        // 版本名称
    VersionName string `json:"versionName,omitempty"`
        // 版本部署状态
    DeployStatus string `json:"deployStatus,omitempty"`
        // 版本待自动激活状态
    AutoActiveStatus string `json:"autoActiveStatus,omitempty"`
}
