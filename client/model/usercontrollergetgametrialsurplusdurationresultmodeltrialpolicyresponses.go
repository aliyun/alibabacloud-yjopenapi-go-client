// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package model

type UsercontrollerGetGameTrialSurplusDurationResultModelTrialPolicyResponses struct {
        // 策略类型
    PolicyType string `json:"policyType,omitempty"`
        // 剩余试玩时长
    SurplusDurationInSecond int64 `json:"surplusDurationInSecond,omitempty"`
}
