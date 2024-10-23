// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetGameConcurrencyResultModelPlayingRegionList struct {
        // 当前游戏指定region游戏中人数
    Ccu int64 `json:"ccu,omitempty"`
        // 当前统计的regionId
    RegionId string `json:"regionId,omitempty"`
}
