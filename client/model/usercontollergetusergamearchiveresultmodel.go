// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontollerGetUserGameArchiveResultModel struct {
        // 下载链接
    DownloadUrl string `json:"downloadUrl,omitempty"`
        // 存档MD5
    ObjectMD5 string `json:"objectMD5,omitempty"`
    UserGameArchiveDTO UsercontollerGetUserGameArchiveResultModelUserGameArchiveDto `json:"userGameArchiveDTO,omitempty"`
}
