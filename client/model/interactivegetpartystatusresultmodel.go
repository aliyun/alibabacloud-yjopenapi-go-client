// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

    // 获取结果
type InteractiveGetPartyStatusResultModel struct {
        // 派对Id
    PartyId string `json:"partyId,omitempty"`
        // 游戏Id
    MixGameId string `json:"mixGameId,omitempty"`
        // 派对创建者
    Creator string `json:"creator,omitempty"`
        // 最大游戏人数
    MaxPlayerNum int32 `json:"maxPlayerNum,omitempty"`
        // 当前派对人数
    CurrentPlayerNum int32 `json:"currentPlayerNum,omitempty"`
        // 派对状态
    Status int32 `json:"status,omitempty"`
        // 派对绑定项目ID
    ProjectId string `json:"projectId,omitempty"`
    PlayerList []InteractiveGetPartyStatusResultModelPlayList `json:"playerList,omitempty"`
}
