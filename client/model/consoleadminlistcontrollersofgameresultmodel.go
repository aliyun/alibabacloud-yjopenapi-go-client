// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ConsoleAdminListControllersOfGameResultModel struct {
        // 总记录数
    Count int64 `json:"count,omitempty"`
        // 表示当前调用返回读取到的位置，空代表数据已经读取完毕
    NextToken string `json:"nextToken,omitempty"`
        // 本次请求所返回的最大记录条数
    MaxResults int32 `json:"maxResults,omitempty"`
        // 数据列表
    DataList []ConsoleAdminListControllersOfGameResultModelDataList `json:"dataList,omitempty"`
}
