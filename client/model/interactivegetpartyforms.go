// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type InteractiveGetPartyForms struct {
        // 游戏ID
    MixGameId string `json:"mixGameId"`
        // 用户ID
    UserId string `json:"userId"`
    ReConnect int32 `json:"reConnect"`
        // 项目
    ProjectId string `json:"projectId"`
        // 派对配置
    Config *InteractiveGetPartyFormsConfig `json:"config,omitempty"`
}
