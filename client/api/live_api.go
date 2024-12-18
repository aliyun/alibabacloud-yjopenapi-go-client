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

type LiveApiService service


// QueryStatus
/*
 * 
 * @param varForms model.LiveQueryStatusForms
 */
func (s *LiveApiService) QueryStatus(
    varForms *model.LiveQueryStatusForms,
) (model.LiveQueryStatusResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.LiveQueryStatusResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/live/queryStatus"

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
	varFormParams.Add("appKey", parameterToString(varForms.AppKey, ""))
	varFormParams.Add("gameSession", parameterToString(varForms.GameSession, ""))
	if varForms != nil && varForms.LiveId != nil {
		varFormParams.Add("liveId", parameterToString(*varForms.LiveId, ""))
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

// StartGameLive
/*
 * 
 * @param varForms model.LiveStartGameLiveForms
 */
func (s *LiveApiService) StartGameLive(
    varForms *model.LiveStartGameLiveForms,
) (model.LiveStartGameLiveResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.LiveStartGameLiveResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/live/startGameLive"

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
	varFormParams.Add("appKey", parameterToString(varForms.AppKey, ""))
	varFormParams.Add("gameSession", parameterToString(varForms.GameSession, ""))
	varFormParams.Add("serverUrl", parameterToString(varForms.ServerUrl, ""))
	varFormParams.Add("streamKey", parameterToString(varForms.StreamKey, ""))
	if varForms != nil && varForms.Config != nil {
		varFormParams.Add("config", parameterToString(*varForms.Config, ""))
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

// StopGameLive
/*
 * 
 * @param varForms model.LiveStopGameLiveForms
 */
func (s *LiveApiService) StopGameLive(
    varForms *model.LiveStopGameLiveForms,
) (model.LiveStopGameLiveResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.LiveStopGameLiveResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/live/stopGameLive"

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
	varFormParams.Add("appKey", parameterToString(varForms.AppKey, ""))
	varFormParams.Add("gameSession", parameterToString(varForms.GameSession, ""))
	if varForms != nil && varForms.LiveId != nil {
		varFormParams.Add("liveId", parameterToString(*varForms.LiveId, ""))
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
