// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type ListGameServerIpResultModel struct {
    NextToken string `json:"nextToken,omitempty"`
    PageSize int64 `json:"pageSize,omitempty"`
    TotalCount int64 `json:"totalCount,omitempty"`
    IpList []string `json:"ipList,omitempty"`
}
