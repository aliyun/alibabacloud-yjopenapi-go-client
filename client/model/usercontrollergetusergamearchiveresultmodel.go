// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontrollerGetUserGameArchiveResultModel struct {
    UserGameArchiveDTO UsercontrollerGetUserGameArchiveResultModelUserGameArchiveDto `json:"userGameArchiveDTO,omitempty"`
    DownloadUrl string `json:"downloadUrl,omitempty"`
    ObjectMD5 string `json:"objectMD5,omitempty"`
}
