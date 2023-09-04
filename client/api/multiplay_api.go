// Package client
/*
 * YuanJing OpenAPI SDK for Go
 *
 *
 */
package api

import (
    "github.com/aliyun/alibabacloud-yjopenapi-go-client/client/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type MultiplayApiService service


// Close
/*
 * 
 * @param varForms model.MultiplayCloseForms
 */
func (s *MultiplayApiService) Close(
    varForms *model.MultiplayCloseForms,
) (model.MultiplayCloseResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.MultiplayCloseResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/multiplay/close/v2"

	varHeaderParams := make(map[string]string)
	varQueryParams := url.Values{}
	varFormParams := url.Values{}

	// to determine the Content-Type header
	varHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	varHttpContentType := selectHeaderContentType(varHttpContentTypes)
	if varHttpContentType != "" {
		varHeaderParams["Content-Type"] = varHttpContentType
	}

	// to determine the Accept header
	varHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	varHttpHeaderAccept := selectHeaderAccept(varHttpHeaderAccepts)
	if varHttpHeaderAccept != "" {
		varHeaderParams["Accept"] = varHttpHeaderAccept
	}
	varFormParams.Add("mpId", parameterToString(varForms.MpId, ""))
	if varForms != nil && varForms.Reason != nil {
		varFormParams.Add("reason", parameterToString(*varForms.Reason, ""))
	}

	r, err := s.client.prepareRequest(varPath, varHttpMethod, varHeaderParams, varQueryParams, varFormParams)
	if err != nil {
		return varReturnValue, nil, err
	}

	varHttpResponse, err := s.client.callAPI(r)
	if err != nil || varHttpResponse == nil {
		return varReturnValue, varHttpResponse, err
	}

    defer varHttpResponse.Body.Close()
	varBody, err := ioutil.ReadAll(varHttpResponse.Body)
	if err != nil {
		return varReturnValue, varHttpResponse, err
	}

	if varHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = s.client.decode(&varReturnValue, varBody, varHttpResponse.Header.Get("Content-Type"))
		if err == nil { 
			return varReturnValue, varHttpResponse, err
		}
	}

	if varHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body: varBody,
			error: varHttpResponse.Status,
		}
		return varReturnValue, varHttpResponse, newErr
	}

	return varReturnValue, varHttpResponse, nil
}

// DemoCloseMultiPlay
/*
 * 联机demo关闭联机
 * @param varForms model.MultiplayDemoCloseMultiPlayForms
 */
func (s *MultiplayApiService) DemoCloseMultiPlay(
    varForms *model.MultiplayDemoCloseMultiPlayForms,
) (model.MultiplayDemoCloseMultiPlayResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.MultiplayDemoCloseMultiPlayResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/multiplay/demoCloseMultiPlay"

	varHeaderParams := make(map[string]string)
	varQueryParams := url.Values{}
	varFormParams := url.Values{}

	// to determine the Content-Type header
	varHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	varHttpContentType := selectHeaderContentType(varHttpContentTypes)
	if varHttpContentType != "" {
		varHeaderParams["Content-Type"] = varHttpContentType
	}

	// to determine the Accept header
	varHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	varHttpHeaderAccept := selectHeaderAccept(varHttpHeaderAccepts)
	if varHttpHeaderAccept != "" {
		varHeaderParams["Accept"] = varHttpHeaderAccept
	}
	if varForms != nil && varForms.MpId != nil {
		varFormParams.Add("mpId", parameterToString(*varForms.MpId, ""))
	}
	if varForms != nil && varForms.Reason != nil {
		varFormParams.Add("reason", parameterToString(*varForms.Reason, ""))
	}
	if varForms != nil && varForms.OpenAk != nil {
		varFormParams.Add("openAk", parameterToString(*varForms.OpenAk, ""))
	}
	if varForms != nil && varForms.OpenSk != nil {
		varFormParams.Add("openSk", parameterToString(*varForms.OpenSk, ""))
	}

	r, err := s.client.prepareRequest(varPath, varHttpMethod, varHeaderParams, varQueryParams, varFormParams)
	if err != nil {
		return varReturnValue, nil, err
	}

	varHttpResponse, err := s.client.callAPI(r)
	if err != nil || varHttpResponse == nil {
		return varReturnValue, varHttpResponse, err
	}

    defer varHttpResponse.Body.Close()
	varBody, err := ioutil.ReadAll(varHttpResponse.Body)
	if err != nil {
		return varReturnValue, varHttpResponse, err
	}

	if varHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = s.client.decode(&varReturnValue, varBody, varHttpResponse.Header.Get("Content-Type"))
		if err == nil { 
			return varReturnValue, varHttpResponse, err
		}
	}

	if varHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body: varBody,
			error: varHttpResponse.Status,
		}
		return varReturnValue, varHttpResponse, newErr
	}

	return varReturnValue, varHttpResponse, nil
}

// DemoInitMultiPlay
/*
 * 联机demo初始化联机
 * @param varForms model.MultiplayDemoInitMultiPlayForms
 */
func (s *MultiplayApiService) DemoInitMultiPlay(
    varForms *model.MultiplayDemoInitMultiPlayForms,
) (model.MultiplayDemoInitMultiPlayResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.MultiplayDemoInitMultiPlayResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/multiplay/demoInitMultiPlay"

	varHeaderParams := make(map[string]string)
	varQueryParams := url.Values{}
	varFormParams := url.Values{}

	// to determine the Content-Type header
	varHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	varHttpContentType := selectHeaderContentType(varHttpContentTypes)
	if varHttpContentType != "" {
		varHeaderParams["Content-Type"] = varHttpContentType
	}

	// to determine the Accept header
	varHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	varHttpHeaderAccept := selectHeaderAccept(varHttpHeaderAccepts)
	if varHttpHeaderAccept != "" {
		varHeaderParams["Accept"] = varHttpHeaderAccept
	}
	if varForms != nil && varForms.GameId != nil {
		varFormParams.Add("gameId", parameterToString(*varForms.GameId, ""))
	}
	if varForms != nil && varForms.GameSession != nil {
		varFormParams.Add("gameSession", parameterToString(*varForms.GameSession, ""))
	}
	if varForms != nil && varForms.AppKey != nil {
		varFormParams.Add("appKey", parameterToString(*varForms.AppKey, ""))
	}
	if varForms != nil && varForms.Config != nil {
		varFormParams.Add("config", parameterToString(*varForms.Config, ""))
	}
	if varForms != nil && varForms.SchedulerConfig != nil {
		varFormParams.Add("schedulerConfig", parameterToString(*varForms.SchedulerConfig, ""))
	}
	if varForms != nil && varForms.Tokens != nil {
		varFormParams.Add("tokens", parameterToString(*varForms.Tokens, ""))
	}
	if varForms != nil && varForms.OpenAk != nil {
		varFormParams.Add("openAk", parameterToString(*varForms.OpenAk, ""))
	}
	if varForms != nil && varForms.OpenSk != nil {
		varFormParams.Add("openSk", parameterToString(*varForms.OpenSk, ""))
	}
	if varForms != nil && varForms.AccountId != nil {
		varFormParams.Add("accountId", parameterToString(*varForms.AccountId, ""))
	}

	r, err := s.client.prepareRequest(varPath, varHttpMethod, varHeaderParams, varQueryParams, varFormParams)
	if err != nil {
		return varReturnValue, nil, err
	}

	varHttpResponse, err := s.client.callAPI(r)
	if err != nil || varHttpResponse == nil {
		return varReturnValue, varHttpResponse, err
	}

    defer varHttpResponse.Body.Close()
	varBody, err := ioutil.ReadAll(varHttpResponse.Body)
	if err != nil {
		return varReturnValue, varHttpResponse, err
	}

	if varHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = s.client.decode(&varReturnValue, varBody, varHttpResponse.Header.Get("Content-Type"))
		if err == nil { 
			return varReturnValue, varHttpResponse, err
		}
	}

	if varHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body: varBody,
			error: varHttpResponse.Status,
		}
		return varReturnValue, varHttpResponse, newErr
	}

	return varReturnValue, varHttpResponse, nil
}

// DemoJoinMultiPlay
/*
 * 联机Demo加入联机接口
 * @param varForms model.MultiplayDemoJoinMultiPlayForms
 */
func (s *MultiplayApiService) DemoJoinMultiPlay(
    varForms *model.MultiplayDemoJoinMultiPlayForms,
) (model.MultiplayDemoJoinMultiPlayResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.MultiplayDemoJoinMultiPlayResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/multiplay/demoJoinMultiPlay"

	varHeaderParams := make(map[string]string)
	varQueryParams := url.Values{}
	varFormParams := url.Values{}

	// to determine the Content-Type header
	varHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	varHttpContentType := selectHeaderContentType(varHttpContentTypes)
	if varHttpContentType != "" {
		varHeaderParams["Content-Type"] = varHttpContentType
	}

	// to determine the Accept header
	varHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	varHttpHeaderAccept := selectHeaderAccept(varHttpHeaderAccepts)
	if varHttpHeaderAccept != "" {
		varHeaderParams["Accept"] = varHttpHeaderAccept
	}
	if varForms != nil && varForms.MpId != nil {
		varFormParams.Add("mpId", parameterToString(*varForms.MpId, ""))
	}
	if varForms != nil && varForms.ControlId != nil {
		varFormParams.Add("controlId", parameterToString(*varForms.ControlId, ""))
	}
	if varForms != nil && varForms.OpenAk != nil {
		varFormParams.Add("openAk", parameterToString(*varForms.OpenAk, ""))
	}
	if varForms != nil && varForms.OpenSk != nil {
		varFormParams.Add("openSk", parameterToString(*varForms.OpenSk, ""))
	}
	if varForms != nil && varForms.AccountId != nil {
		varFormParams.Add("accountId", parameterToString(*varForms.AccountId, ""))
	}

	r, err := s.client.prepareRequest(varPath, varHttpMethod, varHeaderParams, varQueryParams, varFormParams)
	if err != nil {
		return varReturnValue, nil, err
	}

	varHttpResponse, err := s.client.callAPI(r)
	if err != nil || varHttpResponse == nil {
		return varReturnValue, varHttpResponse, err
	}

    defer varHttpResponse.Body.Close()
	varBody, err := ioutil.ReadAll(varHttpResponse.Body)
	if err != nil {
		return varReturnValue, varHttpResponse, err
	}

	if varHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = s.client.decode(&varReturnValue, varBody, varHttpResponse.Header.Get("Content-Type"))
		if err == nil { 
			return varReturnValue, varHttpResponse, err
		}
	}

	if varHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body: varBody,
			error: varHttpResponse.Status,
		}
		return varReturnValue, varHttpResponse, newErr
	}

	return varReturnValue, varHttpResponse, nil
}

// DemoLeaveMultiPlay
/*
 * 联机demo离开联机
 * @param varForms model.MultiplayDemoLeaveMultiPlayForms
 */
func (s *MultiplayApiService) DemoLeaveMultiPlay(
    varForms *model.MultiplayDemoLeaveMultiPlayForms,
) (model.MultiplayDemoLeaveMultiPlayResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.MultiplayDemoLeaveMultiPlayResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/multiplay/demoLeaveMultiPlay"

	varHeaderParams := make(map[string]string)
	varQueryParams := url.Values{}
	varFormParams := url.Values{}

	// to determine the Content-Type header
	varHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	varHttpContentType := selectHeaderContentType(varHttpContentTypes)
	if varHttpContentType != "" {
		varHeaderParams["Content-Type"] = varHttpContentType
	}

	// to determine the Accept header
	varHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	varHttpHeaderAccept := selectHeaderAccept(varHttpHeaderAccepts)
	if varHttpHeaderAccept != "" {
		varHeaderParams["Accept"] = varHttpHeaderAccept
	}
	if varForms != nil && varForms.MpId != nil {
		varFormParams.Add("mpId", parameterToString(*varForms.MpId, ""))
	}
	if varForms != nil && varForms.TokenIds != nil {
		varFormParams.Add("tokenIds", parameterToString(*varForms.TokenIds, ""))
	}
	if varForms != nil && varForms.KickOut != nil {
		varFormParams.Add("kickOut", parameterToString(*varForms.KickOut, ""))
	}
	if varForms != nil && varForms.Reason != nil {
		varFormParams.Add("reason", parameterToString(*varForms.Reason, ""))
	}
	if varForms != nil && varForms.OpenAk != nil {
		varFormParams.Add("openAk", parameterToString(*varForms.OpenAk, ""))
	}
	if varForms != nil && varForms.OpenSk != nil {
		varFormParams.Add("openSk", parameterToString(*varForms.OpenSk, ""))
	}

	r, err := s.client.prepareRequest(varPath, varHttpMethod, varHeaderParams, varQueryParams, varFormParams)
	if err != nil {
		return varReturnValue, nil, err
	}

	varHttpResponse, err := s.client.callAPI(r)
	if err != nil || varHttpResponse == nil {
		return varReturnValue, varHttpResponse, err
	}

    defer varHttpResponse.Body.Close()
	varBody, err := ioutil.ReadAll(varHttpResponse.Body)
	if err != nil {
		return varReturnValue, varHttpResponse, err
	}

	if varHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = s.client.decode(&varReturnValue, varBody, varHttpResponse.Header.Get("Content-Type"))
		if err == nil { 
			return varReturnValue, varHttpResponse, err
		}
	}

	if varHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body: varBody,
			error: varHttpResponse.Status,
		}
		return varReturnValue, varHttpResponse, newErr
	}

	return varReturnValue, varHttpResponse, nil
}

// DemoModifyMultiPlay
/*
 * 联机demo修改联机
 * @param varForms model.MultiplayDemoModifyMultiPlayForms
 */
func (s *MultiplayApiService) DemoModifyMultiPlay(
    varForms *model.MultiplayDemoModifyMultiPlayForms,
) (model.MultiplayDemoModifyMultiPlayResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.MultiplayDemoModifyMultiPlayResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/multiplay/demoModifyMultiPlay"

	varHeaderParams := make(map[string]string)
	varQueryParams := url.Values{}
	varFormParams := url.Values{}

	// to determine the Content-Type header
	varHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	varHttpContentType := selectHeaderContentType(varHttpContentTypes)
	if varHttpContentType != "" {
		varHeaderParams["Content-Type"] = varHttpContentType
	}

	// to determine the Accept header
	varHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	varHttpHeaderAccept := selectHeaderAccept(varHttpHeaderAccepts)
	if varHttpHeaderAccept != "" {
		varHeaderParams["Accept"] = varHttpHeaderAccept
	}
	if varForms != nil && varForms.MpId != nil {
		varFormParams.Add("mpId", parameterToString(*varForms.MpId, ""))
	}
	if varForms != nil && varForms.Tokens != nil {
		varFormParams.Add("tokens", parameterToString(*varForms.Tokens, ""))
	}
	if varForms != nil && varForms.OpenAk != nil {
		varFormParams.Add("openAk", parameterToString(*varForms.OpenAk, ""))
	}
	if varForms != nil && varForms.OpenSk != nil {
		varFormParams.Add("openSk", parameterToString(*varForms.OpenSk, ""))
	}

	r, err := s.client.prepareRequest(varPath, varHttpMethod, varHeaderParams, varQueryParams, varFormParams)
	if err != nil {
		return varReturnValue, nil, err
	}

	varHttpResponse, err := s.client.callAPI(r)
	if err != nil || varHttpResponse == nil {
		return varReturnValue, varHttpResponse, err
	}

    defer varHttpResponse.Body.Close()
	varBody, err := ioutil.ReadAll(varHttpResponse.Body)
	if err != nil {
		return varReturnValue, varHttpResponse, err
	}

	if varHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = s.client.decode(&varReturnValue, varBody, varHttpResponse.Header.Get("Content-Type"))
		if err == nil { 
			return varReturnValue, varHttpResponse, err
		}
	}

	if varHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body: varBody,
			error: varHttpResponse.Status,
		}
		return varReturnValue, varHttpResponse, newErr
	}

	return varReturnValue, varHttpResponse, nil
}

// DemoQueryMultiPlay
/*
 * 联机demo查询联机接口
 * @param varForms model.MultiplayDemoQueryMultiPlayForms
 */
func (s *MultiplayApiService) DemoQueryMultiPlay(
    varForms *model.MultiplayDemoQueryMultiPlayForms,
) (model.MultiplayDemoQueryMultiPlayResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.MultiplayDemoQueryMultiPlayResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/multiplay/demoQueryMultiPlay"

	varHeaderParams := make(map[string]string)
	varQueryParams := url.Values{}
	varFormParams := url.Values{}

	// to determine the Content-Type header
	varHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	varHttpContentType := selectHeaderContentType(varHttpContentTypes)
	if varHttpContentType != "" {
		varHeaderParams["Content-Type"] = varHttpContentType
	}

	// to determine the Accept header
	varHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	varHttpHeaderAccept := selectHeaderAccept(varHttpHeaderAccepts)
	if varHttpHeaderAccept != "" {
		varHeaderParams["Accept"] = varHttpHeaderAccept
	}
	if varForms != nil && varForms.MpId != nil {
		varFormParams.Add("mpId", parameterToString(*varForms.MpId, ""))
	}
	if varForms != nil && varForms.OpenAk != nil {
		varFormParams.Add("openAk", parameterToString(*varForms.OpenAk, ""))
	}
	if varForms != nil && varForms.OpenSk != nil {
		varFormParams.Add("openSk", parameterToString(*varForms.OpenSk, ""))
	}

	r, err := s.client.prepareRequest(varPath, varHttpMethod, varHeaderParams, varQueryParams, varFormParams)
	if err != nil {
		return varReturnValue, nil, err
	}

	varHttpResponse, err := s.client.callAPI(r)
	if err != nil || varHttpResponse == nil {
		return varReturnValue, varHttpResponse, err
	}

    defer varHttpResponse.Body.Close()
	varBody, err := ioutil.ReadAll(varHttpResponse.Body)
	if err != nil {
		return varReturnValue, varHttpResponse, err
	}

	if varHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = s.client.decode(&varReturnValue, varBody, varHttpResponse.Header.Get("Content-Type"))
		if err == nil { 
			return varReturnValue, varHttpResponse, err
		}
	}

	if varHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body: varBody,
			error: varHttpResponse.Status,
		}
		return varReturnValue, varHttpResponse, newErr
	}

	return varReturnValue, varHttpResponse, nil
}

// Init
/*
 * 
 * @param varForms model.MultiplayInitForms
 */
func (s *MultiplayApiService) Init(
    varForms *model.MultiplayInitForms,
) (model.MultiplayInitResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.MultiplayInitResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/multiplay/init/v2"

	varHeaderParams := make(map[string]string)
	varQueryParams := url.Values{}
	varFormParams := url.Values{}

	// to determine the Content-Type header
	varHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	varHttpContentType := selectHeaderContentType(varHttpContentTypes)
	if varHttpContentType != "" {
		varHeaderParams["Content-Type"] = varHttpContentType
	}

	// to determine the Accept header
	varHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	varHttpHeaderAccept := selectHeaderAccept(varHttpHeaderAccepts)
	if varHttpHeaderAccept != "" {
		varHeaderParams["Accept"] = varHttpHeaderAccept
	}
	varFormParams.Add("gameSession", parameterToString(varForms.GameSession, ""))
	varFormParams.Add("appKey", parameterToString(varForms.AppKey, ""))
	if varForms != nil && varForms.Config != nil {
		varFormParams.Add("config", parameterToString(*varForms.Config, ""))
	}
	if varForms != nil && varForms.Tokens != nil {
		varFormParams.Add("tokens", parameterToString(*varForms.Tokens, "multi"))
	}

	r, err := s.client.prepareRequest(varPath, varHttpMethod, varHeaderParams, varQueryParams, varFormParams)
	if err != nil {
		return varReturnValue, nil, err
	}

	varHttpResponse, err := s.client.callAPI(r)
	if err != nil || varHttpResponse == nil {
		return varReturnValue, varHttpResponse, err
	}

    defer varHttpResponse.Body.Close()
	varBody, err := ioutil.ReadAll(varHttpResponse.Body)
	if err != nil {
		return varReturnValue, varHttpResponse, err
	}

	if varHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = s.client.decode(&varReturnValue, varBody, varHttpResponse.Header.Get("Content-Type"))
		if err == nil { 
			return varReturnValue, varHttpResponse, err
		}
	}

	if varHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body: varBody,
			error: varHttpResponse.Status,
		}
		return varReturnValue, varHttpResponse, newErr
	}

	return varReturnValue, varHttpResponse, nil
}

// Join
/*
 * 
 * @param varForms model.MultiplayJoinForms
 */
func (s *MultiplayApiService) Join(
    varForms *model.MultiplayJoinForms,
) (model.MultiplayJoinResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.MultiplayJoinResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/multiplay/join/v2"

	varHeaderParams := make(map[string]string)
	varQueryParams := url.Values{}
	varFormParams := url.Values{}

	// to determine the Content-Type header
	varHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	varHttpContentType := selectHeaderContentType(varHttpContentTypes)
	if varHttpContentType != "" {
		varHeaderParams["Content-Type"] = varHttpContentType
	}

	// to determine the Accept header
	varHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	varHttpHeaderAccept := selectHeaderAccept(varHttpHeaderAccepts)
	if varHttpHeaderAccept != "" {
		varHeaderParams["Accept"] = varHttpHeaderAccept
	}
	varFormParams.Add("mpId", parameterToString(varForms.MpId, ""))
	varFormParams.Add("accountId", parameterToString(varForms.AccountId, ""))
	if varForms != nil && varForms.ControlId != nil {
		varFormParams.Add("controlId", parameterToString(*varForms.ControlId, ""))
	}

	r, err := s.client.prepareRequest(varPath, varHttpMethod, varHeaderParams, varQueryParams, varFormParams)
	if err != nil {
		return varReturnValue, nil, err
	}

	varHttpResponse, err := s.client.callAPI(r)
	if err != nil || varHttpResponse == nil {
		return varReturnValue, varHttpResponse, err
	}

    defer varHttpResponse.Body.Close()
	varBody, err := ioutil.ReadAll(varHttpResponse.Body)
	if err != nil {
		return varReturnValue, varHttpResponse, err
	}

	if varHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = s.client.decode(&varReturnValue, varBody, varHttpResponse.Header.Get("Content-Type"))
		if err == nil { 
			return varReturnValue, varHttpResponse, err
		}
	}

	if varHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body: varBody,
			error: varHttpResponse.Status,
		}
		return varReturnValue, varHttpResponse, newErr
	}

	return varReturnValue, varHttpResponse, nil
}

// Leave
/*
 * 
 * @param varForms model.MultiplayLeaveForms
 */
func (s *MultiplayApiService) Leave(
    varForms *model.MultiplayLeaveForms,
) (model.MultiplayLeaveResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.MultiplayLeaveResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/multiplay/leave/v2"

	varHeaderParams := make(map[string]string)
	varQueryParams := url.Values{}
	varFormParams := url.Values{}

	// to determine the Content-Type header
	varHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	varHttpContentType := selectHeaderContentType(varHttpContentTypes)
	if varHttpContentType != "" {
		varHeaderParams["Content-Type"] = varHttpContentType
	}

	// to determine the Accept header
	varHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	varHttpHeaderAccept := selectHeaderAccept(varHttpHeaderAccepts)
	if varHttpHeaderAccept != "" {
		varHeaderParams["Accept"] = varHttpHeaderAccept
	}
	varFormParams.Add("mpId", parameterToString(varForms.MpId, ""))
	varFormParams.Add("kickOut", parameterToString(varForms.KickOut, ""))
	if varForms != nil && varForms.Reason != nil {
		varFormParams.Add("reason", parameterToString(*varForms.Reason, ""))
	}
	varFormParams.Add("tokenIds", parameterToString(varForms.TokenIds, "multi"))

	r, err := s.client.prepareRequest(varPath, varHttpMethod, varHeaderParams, varQueryParams, varFormParams)
	if err != nil {
		return varReturnValue, nil, err
	}

	varHttpResponse, err := s.client.callAPI(r)
	if err != nil || varHttpResponse == nil {
		return varReturnValue, varHttpResponse, err
	}

    defer varHttpResponse.Body.Close()
	varBody, err := ioutil.ReadAll(varHttpResponse.Body)
	if err != nil {
		return varReturnValue, varHttpResponse, err
	}

	if varHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = s.client.decode(&varReturnValue, varBody, varHttpResponse.Header.Get("Content-Type"))
		if err == nil { 
			return varReturnValue, varHttpResponse, err
		}
	}

	if varHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body: varBody,
			error: varHttpResponse.Status,
		}
		return varReturnValue, varHttpResponse, newErr
	}

	return varReturnValue, varHttpResponse, nil
}

// Modify
/*
 * 
 * @param varForms model.MultiplayModifyForms
 */
func (s *MultiplayApiService) Modify(
    varForms *model.MultiplayModifyForms,
) (model.MultiplayModifyResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.MultiplayModifyResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/multiplay/modify/v2"

	varHeaderParams := make(map[string]string)
	varQueryParams := url.Values{}
	varFormParams := url.Values{}

	// to determine the Content-Type header
	varHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	varHttpContentType := selectHeaderContentType(varHttpContentTypes)
	if varHttpContentType != "" {
		varHeaderParams["Content-Type"] = varHttpContentType
	}

	// to determine the Accept header
	varHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	varHttpHeaderAccept := selectHeaderAccept(varHttpHeaderAccepts)
	if varHttpHeaderAccept != "" {
		varHeaderParams["Accept"] = varHttpHeaderAccept
	}
	varFormParams.Add("mpId", parameterToString(varForms.MpId, ""))
	if varForms != nil && varForms.Tokens != nil {
		varFormParams.Add("tokens", parameterToString(*varForms.Tokens, "multi"))
	}

	r, err := s.client.prepareRequest(varPath, varHttpMethod, varHeaderParams, varQueryParams, varFormParams)
	if err != nil {
		return varReturnValue, nil, err
	}

	varHttpResponse, err := s.client.callAPI(r)
	if err != nil || varHttpResponse == nil {
		return varReturnValue, varHttpResponse, err
	}

    defer varHttpResponse.Body.Close()
	varBody, err := ioutil.ReadAll(varHttpResponse.Body)
	if err != nil {
		return varReturnValue, varHttpResponse, err
	}

	if varHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = s.client.decode(&varReturnValue, varBody, varHttpResponse.Header.Get("Content-Type"))
		if err == nil { 
			return varReturnValue, varHttpResponse, err
		}
	}

	if varHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body: varBody,
			error: varHttpResponse.Status,
		}
		return varReturnValue, varHttpResponse, newErr
	}

	return varReturnValue, varHttpResponse, nil
}

// Query
/*
 * 
 * @param varForms model.MultiplayQueryForms
 */
func (s *MultiplayApiService) Query(
    varForms *model.MultiplayQueryForms,
) (model.MultiplayQueryResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.MultiplayQueryResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/multiplay/query/v2"

	varHeaderParams := make(map[string]string)
	varQueryParams := url.Values{}
	varFormParams := url.Values{}

	// to determine the Content-Type header
	varHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	varHttpContentType := selectHeaderContentType(varHttpContentTypes)
	if varHttpContentType != "" {
		varHeaderParams["Content-Type"] = varHttpContentType
	}

	// to determine the Accept header
	varHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	varHttpHeaderAccept := selectHeaderAccept(varHttpHeaderAccepts)
	if varHttpHeaderAccept != "" {
		varHeaderParams["Accept"] = varHttpHeaderAccept
	}
	varFormParams.Add("mpId", parameterToString(varForms.MpId, ""))

	r, err := s.client.prepareRequest(varPath, varHttpMethod, varHeaderParams, varQueryParams, varFormParams)
	if err != nil {
		return varReturnValue, nil, err
	}

	varHttpResponse, err := s.client.callAPI(r)
	if err != nil || varHttpResponse == nil {
		return varReturnValue, varHttpResponse, err
	}

    defer varHttpResponse.Body.Close()
	varBody, err := ioutil.ReadAll(varHttpResponse.Body)
	if err != nil {
		return varReturnValue, varHttpResponse, err
	}

	if varHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = s.client.decode(&varReturnValue, varBody, varHttpResponse.Header.Get("Content-Type"))
		if err == nil { 
			return varReturnValue, varHttpResponse, err
		}
	}

	if varHttpResponse.StatusCode >= 300 {
		newErr := GenericError{
			body: varBody,
			error: varHttpResponse.Status,
		}
		return varReturnValue, varHttpResponse, newErr
	}

	return varReturnValue, varHttpResponse, nil
}
