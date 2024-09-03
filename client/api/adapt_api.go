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

type AdaptApiService service


// CreateAndSubmitAll
/*
 * createAndSubmitAll
 * @param varForms model.AdaptCreateAndSubmitAllForms
 */
func (s *AdaptApiService) CreateAndSubmitAll(
    varForms *model.AdaptCreateAndSubmitAllForms,
) (model.AdaptCreateAndSubmitAllResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.AdaptCreateAndSubmitAllResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/adapt/createAndSubmitAll"

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
	if varForms != nil && varForms.GameName != nil {
		varFormParams.Add("gameName", parameterToString(*varForms.GameName, ""))
	}
	if varForms != nil && varForms.GameVersionId != nil {
		varFormParams.Add("gameVersionId", parameterToString(*varForms.GameVersionId, ""))
	}
	if varForms != nil && varForms.GameVersion != nil {
		varFormParams.Add("gameVersion", parameterToString(*varForms.GameVersion, ""))
	}
	if varForms != nil && varForms.ResolutionList != nil {
		varFormParams.Add("resolutionList", parameterToString(*varForms.ResolutionList, ""))
	}
	if varForms != nil && varForms.FrameRateList != nil {
		varFormParams.Add("frameRateList", parameterToString(*varForms.FrameRateList, ""))
	}
	varFormParams.Add("platformType", parameterToString(varForms.PlatformType, ""))
	varFormParams.Add("sourcePlatform", parameterToString(varForms.SourcePlatform, ""))
	varFormParams.Add("records", parameterToString(varForms.Records, ""))
	varFormParams.Add("mixGameVersionId", parameterToString(varForms.MixGameVersionId, ""))
	varFormParams.Add("mixGameId", parameterToString(varForms.MixGameId, ""))

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

// QueryRequestById
/*
 * queryRequestById
 * @param varForms model.AdaptQueryRequestByIdForms
 */
func (s *AdaptApiService) QueryRequestById(
    varForms *model.AdaptQueryRequestByIdForms,
) (model.AdaptQueryRequestByIdResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.AdaptQueryRequestByIdResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/adapt/queryRequestById"

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
	varFormParams.Add("id", parameterToString(varForms.Id, ""))

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
