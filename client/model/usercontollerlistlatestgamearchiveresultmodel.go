// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontollerListLatestGameArchiveResultModel struct {
        // 总数
    TotalCount int32 `json:"totalCount,omitempty"`
        // 分页数量
    PageSize int32 `json:"pageSize,omitempty"`
        // 当前页码
    PageNumber int32 `json:"pageNumber,omitempty"`
        // id of the request
    RequestId string `json:"requestId,omitempty"`
        // 存档数据列表
    Items []UsercontollerListLatestGameArchiveResultModelItems `json:"items,omitempty"`
}
