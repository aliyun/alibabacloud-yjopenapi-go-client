// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

    // 获取结果
type InteractiveModifySeatsResultModel struct {
    PlayerList []InteractiveModifySeatsResultModelPlayerList `json:"playerList,omitempty"`
        // 扩展信息
    ExtInfo string `json:"extInfo,omitempty"`
}
