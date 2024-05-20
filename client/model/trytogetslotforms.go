// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type TryToGetSlotForms struct {
        // 用户id，给到Paas平台和SDK，两者保持一致，全局唯一
    AccountId string `json:"accountId"`
        // Paas平台部署的游戏Id
    GameId string `json:"gameId"`
        // Paas平台AK(应用的AK，非服务端AK)
    AppKey string `json:"appKey"`
        // 需要调度的区域
    RegionId *string `json:"regionId,omitempty"`
        // false代表不使用断线重连，开启新游戏
    ReConnect *bool `json:"reConnect,omitempty"`
        // 端侧从Paas sdk获取的调度业务参数
    BizParam *string `json:"bizParam,omitempty"`
        // App端公网ip
    ClientIp *string `json:"clientIp,omitempty"`
        // 标签之间用半角逗号分隔
    Tags *string `json:"tags,omitempty"`
        // 调度等级
    UserLevel *int32 `json:"userLevel,omitempty"`
        // 编码
    Codec *int32 `json:"codec,omitempty"`
        // 画质
    Resolution *int32 `json:"resolution,omitempty"`
        // 码率
    BitRate *int32 `json:"bitRate,omitempty"`
        // 帧率
    Fps *int32 `json:"fps,omitempty"`
        // 启动命令，透传至ISV用于启动游戏
    GameCmdParam *string `json:"gameCmdParam,omitempty"`
        // 游戏启动设置参数
    StartParam *TryToGetSlotFormsStartParam `json:"startParam,omitempty"`
}
