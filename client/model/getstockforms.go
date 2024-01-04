// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type GetStockForms struct {
        // Paas平台部署的游戏Id
    GameId *string `json:"gameId,omitempty"`
        // 查询库存类型
    AppKey string `json:"appKey"`
        // 通过接口获取的token
    Type_ string `json:"type"`
}
