// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminUploadGameVersionByDownloadForms struct {
        // 文件MD5校验值
    Hash string `json:"hash"`
        // 游戏id
    GameId string `json:"gameId"`
        // 下载方式
    DownloadType ConsoleAdminUploadGameVersionByDownloadDownloadType `json:"downloadType"`
        // 自定义版本名称
    VersionName string `json:"versionName"`
}
