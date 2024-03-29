// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type TryToGetSlotFormsStartParam struct {
        // 启动待缩放分辨率
    Resolution *string `json:"resolution,omitempty"`
        // 控制游戏运行时长
    KeepAlive *int64 `json:"keepAlive,omitempty"`
        // 游戏中设备掉线后会话保活时长
    KeepDisconnectAlive *int64 `json:"keepDisconnectAlive,omitempty"`
        // 游戏中无操作时会话保活时长
    KeepNoPlayAlive *int64 `json:"keepNoPlayAlive,omitempty"`
        // 设备品牌
    DeviceBrand *string `json:"deviceBrand,omitempty"`
        // 设备型号
    DeviceModel *string `json:"deviceModel,omitempty"`
        // 设备DPI
    DeviceDpi *int32 `json:"deviceDpi,omitempty"`
        // 设备DPR
    DeviceDpr *int32 `json:"deviceDpr,omitempty"`
        // 设备分辨率
    DeviceResolution *string `json:"deviceResolution,omitempty"`
        // 手游输入法控制
    GameInput *string `json:"gameInput,omitempty"`
        // 游戏脚本id
    ScriptId *int32 `json:"scriptId,omitempty"`
        // 码率自适应开关
    BitRateSelfAdaption *int32 `json:"bitRateSelfAdaption,omitempty"`
        // 游戏调度实例等级
    ScheduleUserLevels *string `json:"scheduleUserLevels,omitempty"`
        // 设备号Android
    Ssaid *string `json:"ssaid,omitempty"`
        // 应用渠道
    AppChannel *string `json:"appChannel,omitempty"`
        // 游戏初始化路径
    CloudInit *string `json:"cloudInit,omitempty"`
        // 设备id
    SysDeviceId *string `json:"sysDeviceId,omitempty"`
    ArchiveUrl *string `json:"archiveUrl,omitempty"`
    ArchiveMd5 *string `json:"archiveMd5,omitempty"`
}
