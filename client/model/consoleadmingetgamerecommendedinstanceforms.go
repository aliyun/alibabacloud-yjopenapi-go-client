// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminGetGameRecommendedInstanceForms struct {
    Game ConsoleAdminGetGameRecommendedInstanceFormsGame `json:"game"`
    CommodityInstanceIds []string `json:"commodityInstanceIds"`
}
