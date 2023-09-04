// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminUploadGameVersionByDownloadForms struct {
    Hash string `json:"hash"`
    GameId string `json:"gameId"`
    DownloadType ConsoleAdminUploadGameVersionByDownloadFormsDownloadType `json:"downloadType"`
    VersionName string `json:"versionName"`
}
