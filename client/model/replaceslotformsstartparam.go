// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ReplaceSlotFormsStartParam struct {
    KeepAlive *int64 `json:"keepAlive,omitempty"`
    CloudInit *string `json:"cloudInit ,omitempty"`
    DeviceResolution *string `json:"deviceResolution,omitempty"`
    DeviceDpi *int32 `json:"deviceDpi,omitempty"`
    ArchiveUrl *string `json:"archiveUrl,omitempty"`
    GameInput *string `json:"gameInput,omitempty"`
    Resolution *string `json:"resolution,omitempty"`
    KeepDisconnectAlive *int64 `json:"keepDisconnectAlive,omitempty"`
    ArchiveId *string `json:"archiveId,omitempty"`
    KeepHoldAlive *int64 `json:"keepHoldAlive ,omitempty"`
    AppChannel *string `json:"appChannel,omitempty"`
    DeviceDpr *int32 `json:"deviceDpr,omitempty"`
    ScriptId *string `json:"scriptId,omitempty"`
    GameVersionId *string `json:"gameVersionId,omitempty"`
    BitRateSelfAdaption *int32 `json:"bitRateSelfAdaption,omitempty"`
    KeepNoPlayAlive *int64 `json:"keepNoPlayAlive,omitempty"`
    Ssaid *string `json:"ssaid,omitempty"`
    ArchiveMd5 *string `json:"archiveMd5,omitempty"`
    SysDeviceId *string `json:"sysDeviceId,omitempty"`
    DeviceModel *string `json:"deviceModel,omitempty"`
    DeviceBrand *string `json:"deviceBrand,omitempty"`
}
