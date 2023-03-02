// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ListGameServerIpResultModel struct {
        // 分页游标
    NextToken string `json:"nextToken,omitempty"`
        // 分页大小
    PageSize int64 `json:"pageSize,omitempty"`
        // ip列表
    IpList []string `json:"ipList,omitempty"`
        // 总大小
    TotalCount int64 `json:"totalCount,omitempty"`
}
