// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

    // 获取结果
type InteractiveJoinPartyResultModel struct {
    PlayerList []InteractiveJoinPartyResultModelPlayerList `json:"playerList,omitempty"`
    ExtInfo InteractiveJoinPartyResultExtInfo `json:"extInfo,omitempty"`
}
