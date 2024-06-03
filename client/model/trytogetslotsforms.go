// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type TryToGetSlotsForms struct {
    AppKey string `json:"appKey"`
    ParallelSchedule bool `json:"parallelSchedule"`
    Requests []TryToGetSlotsFormsRequests `json:"requests"`
}
