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

type ConsoleAdminApiService service


// ActivateDeployment
/*
 * 激活已部署成功的游戏版本的部署
 * @param varForms model.ConsoleAdminActivateDeploymentForms
 */
func (s *ConsoleAdminApiService) ActivateDeployment(
    varForms *model.ConsoleAdminActivateDeploymentForms,
) (model.ConsoleAdminActivateDeploymentResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminActivateDeploymentResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/activateDeployment"

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
	varFormParams.Add("gameId", parameterToString(varForms.GameId, ""))
	varFormParams.Add("projectId", parameterToString(varForms.ProjectId, ""))
	varFormParams.Add("versionId", parameterToString(varForms.VersionId, ""))
	if varForms != nil && varForms.MaxConcurrency != nil {
		varFormParams.Add("maxConcurrency", parameterToString(*varForms.MaxConcurrency, ""))
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

// AdaptGameVersion
/*
 * 发起游戏版本适配
 * @param varForms model.ConsoleAdminAdaptGameVersionForms
 */
func (s *ConsoleAdminApiService) AdaptGameVersion(
    varForms *model.ConsoleAdminAdaptGameVersionForms,
) (model.ConsoleAdminAdaptGameVersionResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminAdaptGameVersionResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/adaptGameVersion"

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
	varFormParams.Add("versionId", parameterToString(varForms.VersionId, ""))

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

// AddGameToProject
/*
 * 将游戏添加到项目
 * @param varForms model.ConsoleAdminAddGameToProjectForms
 */
func (s *ConsoleAdminApiService) AddGameToProject(
    varForms *model.ConsoleAdminAddGameToProjectForms,
) (model.ConsoleAdminAddGameToProjectResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminAddGameToProjectResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/addGameToProject"

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
	varFormParams.Add("projectId", parameterToString(varForms.ProjectId, ""))
	varFormParams.Add("gameId", parameterToString(varForms.GameId, ""))

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

// CreateGame
/*
 * 创建游戏
 * @param varForms model.ConsoleAdminCreateGameForms
 */
func (s *ConsoleAdminApiService) CreateGame(
    varForms *model.ConsoleAdminCreateGameForms,
) (model.ConsoleAdminCreateGameResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminCreateGameResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/createGame"

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
	varFormParams.Add("gameName", parameterToString(varForms.GameName, ""))
	varFormParams.Add("platformType", parameterToString(varForms.PlatformType, ""))

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

// CreateProject
/*
 * 创建项目
 * @param varForms model.ConsoleAdminCreateProjectForms
 */
func (s *ConsoleAdminApiService) CreateProject(
    varForms *model.ConsoleAdminCreateProjectForms,
) (model.ConsoleAdminCreateProjectResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminCreateProjectResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/createProject"

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
	varFormParams.Add("projectName", parameterToString(varForms.ProjectName, ""))

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

// DeleteGame
/*
 * 删除指定的游戏
 * @param varForms model.ConsoleAdminDeleteGameForms
 */
func (s *ConsoleAdminApiService) DeleteGame(
    varForms *model.ConsoleAdminDeleteGameForms,
) (model.ConsoleAdminDeleteGameResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminDeleteGameResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/deleteGame"

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
	varFormParams.Add("gameId", parameterToString(varForms.GameId, ""))

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

// DeleteGameVersion
/*
 * 发起游戏版本适配
 * @param varForms model.ConsoleAdminDeleteGameVersionForms
 */
func (s *ConsoleAdminApiService) DeleteGameVersion(
    varForms *model.ConsoleAdminDeleteGameVersionForms,
) (model.ConsoleAdminDeleteGameVersionResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminDeleteGameVersionResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/deleteGameVersion"

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
	varFormParams.Add("versionId", parameterToString(varForms.VersionId, ""))

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

// DeleteProject
/*
 * 删除指定的项目
 * @param varForms model.ConsoleAdminDeleteProjectForms
 */
func (s *ConsoleAdminApiService) DeleteProject(
    varForms *model.ConsoleAdminDeleteProjectForms,
) (model.ConsoleAdminDeleteProjectResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminDeleteProjectResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/deleteProject"

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
	varFormParams.Add("projectId", parameterToString(varForms.ProjectId, ""))

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

// GetGameVersion
/*
 * 获取单个游戏版本信息
 * @param varForms model.ConsoleAdminGetGameVersionForms
 */
func (s *ConsoleAdminApiService) GetGameVersion(
    varForms *model.ConsoleAdminGetGameVersionForms,
) (model.ConsoleAdminGetGameVersionResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminGetGameVersionResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/getGameVersion"

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
	varFormParams.Add("versionId", parameterToString(varForms.VersionId, ""))

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

// GetGameVersionProgress
/*
 * 查询版本处理进度（包含上传、适配、部署）
 * @param varForms model.ConsoleAdminGetGameVersionProgressForms
 */
func (s *ConsoleAdminApiService) GetGameVersionProgress(
    varForms *model.ConsoleAdminGetGameVersionProgressForms,
) (model.ConsoleAdminGetGameVersionProgressResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminGetGameVersionProgressResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/getGameVersionProgress"

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
	varFormParams.Add("taskId", parameterToString(varForms.TaskId, ""))

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

// ListActivateableInstances
/*
 * 指定项目和游戏版本，获取可激活且可调度的实例及调度配置
 * @param varForms model.ConsoleAdminListActivateableInstancesForms
 */
func (s *ConsoleAdminApiService) ListActivateableInstances(
    varForms *model.ConsoleAdminListActivateableInstancesForms,
) (model.ConsoleAdminListActivateableInstancesResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminListActivateableInstancesResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/listActivateableInstances"

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
	varFormParams.Add("projectId", parameterToString(varForms.ProjectId, ""))
	varFormParams.Add("versionId", parameterToString(varForms.VersionId, ""))

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

// ListActivatedInstances
/*
 * 指定项目和游戏，获取已激活版本的可调度实例及调度配置
 * @param varForms model.ConsoleAdminListActivatedInstancesForms
 */
func (s *ConsoleAdminApiService) ListActivatedInstances(
    varForms *model.ConsoleAdminListActivatedInstancesForms,
) (model.ConsoleAdminListActivatedInstancesResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminListActivatedInstancesResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/listActivatedInstances"

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
	varFormParams.Add("gameId", parameterToString(varForms.GameId, ""))
	varFormParams.Add("projectId", parameterToString(varForms.ProjectId, ""))

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

// ListControllersOfGame
/*
 * 获取单个游戏关联的控制器列表
 * @param varForms model.ConsoleAdminListControllersOfGameForms
 */
func (s *ConsoleAdminApiService) ListControllersOfGame(
    varForms *model.ConsoleAdminListControllersOfGameForms,
) (model.ConsoleAdminListControllersOfGameResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminListControllersOfGameResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/listControllersOfGame"

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
	varFormParams.Add("gameId", parameterToString(varForms.GameId, ""))
	if varForms != nil && varForms.NextToken != nil {
		varFormParams.Add("nextToken", parameterToString(*varForms.NextToken, ""))
	}
	if varForms != nil && varForms.MaxResults != nil {
		varFormParams.Add("maxResults", parameterToString(*varForms.MaxResults, ""))
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

// ListDeployableInstances
/*
 * 指定项目和游戏版本，获取可以部署的实例
 * @param varForms model.ConsoleAdminListDeployableInstancesForms
 */
func (s *ConsoleAdminApiService) ListDeployableInstances(
    varForms *model.ConsoleAdminListDeployableInstancesForms,
) (model.ConsoleAdminListDeployableInstancesResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminListDeployableInstancesResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/listDeployableInstances"

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
	varFormParams.Add("versionId", parameterToString(varForms.VersionId, ""))
	varFormParams.Add("projectId", parameterToString(varForms.ProjectId, ""))
	if varForms != nil && varForms.PageSize != nil {
		varFormParams.Add("pageSize", parameterToString(*varForms.PageSize, ""))
	}
	if varForms != nil && varForms.PageNumber != nil {
		varFormParams.Add("pageNumber", parameterToString(*varForms.PageNumber, ""))
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

// ListGameDeployDetailsOfProject
/*
 * 获取项目下游戏部署版本信息
 * @param varForms model.ConsoleAdminListGameDeployDetailsOfProjectForms
 */
func (s *ConsoleAdminApiService) ListGameDeployDetailsOfProject(
    varForms *model.ConsoleAdminListGameDeployDetailsOfProjectForms,
) (model.ConsoleAdminListGameDeployDetailsOfProjectResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminListGameDeployDetailsOfProjectResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/listGameDeployDetailsOfProject"

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
	if varForms != nil && varForms.ProjectId != nil {
		varFormParams.Add("projectId", parameterToString(*varForms.ProjectId, ""))
	}
	if varForms != nil && varForms.GameId != nil {
		varFormParams.Add("gameId", parameterToString(*varForms.GameId, ""))
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

// ListGameVersions
/*
 * 分页获取游戏版本列表
 * @param varForms model.ConsoleAdminListGameVersionsForms
 */
func (s *ConsoleAdminApiService) ListGameVersions(
    varForms *model.ConsoleAdminListGameVersionsForms,
) (model.ConsoleAdminListGameVersionsResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminListGameVersionsResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/listGameVersions"

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
	if varForms != nil && varForms.NextToken != nil {
		varFormParams.Add("nextToken", parameterToString(*varForms.NextToken, ""))
	}
	if varForms != nil && varForms.MaxResults != nil {
		varFormParams.Add("maxResults", parameterToString(*varForms.MaxResults, ""))
	}
	varFormParams.Add("gameId", parameterToString(varForms.GameId, ""))

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

// ListGames
/*
 * 分页获取游戏列表
 * @param varForms model.ConsoleAdminListGamesForms
 */
func (s *ConsoleAdminApiService) ListGames(
    varForms *model.ConsoleAdminListGamesForms,
) (model.ConsoleAdminListGamesResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminListGamesResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/listGames"

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
	if varForms != nil && varForms.NextToken != nil {
		varFormParams.Add("nextToken", parameterToString(*varForms.NextToken, ""))
	}
	if varForms != nil && varForms.MaxResults != nil {
		varFormParams.Add("maxResults", parameterToString(*varForms.MaxResults, ""))
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

// ListInstancesOfProject
/*
 * 分页获取项目中的实例
 * @param varForms model.ConsoleAdminListInstancesOfProjectForms
 */
func (s *ConsoleAdminApiService) ListInstancesOfProject(
    varForms *model.ConsoleAdminListInstancesOfProjectForms,
) (model.ConsoleAdminListInstancesOfProjectResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminListInstancesOfProjectResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/listInstancesOfProject"

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
	if varForms != nil && varForms.NextToken != nil {
		varFormParams.Add("nextToken", parameterToString(*varForms.NextToken, ""))
	}
	if varForms != nil && varForms.MaxResult != nil {
		varFormParams.Add("maxResult", parameterToString(*varForms.MaxResult, ""))
	}
	varFormParams.Add("projectId", parameterToString(varForms.ProjectId, ""))

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

// ListProjects
/*
 * 分页获取项目列表
 * @param varForms model.ConsoleAdminListProjectsForms
 */
func (s *ConsoleAdminApiService) ListProjects(
    varForms *model.ConsoleAdminListProjectsForms,
) (model.ConsoleAdminListProjectsResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminListProjectsResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/listProjects"

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
	if varForms != nil && varForms.NextToken != nil {
		varFormParams.Add("nextToken", parameterToString(*varForms.NextToken, ""))
	}
	if varForms != nil && varForms.MaxResults != nil {
		varFormParams.Add("maxResults", parameterToString(*varForms.MaxResults, ""))
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

// ListVersionDeployInstances
/*
 * 获取项目下游戏版本的部署实例信息
 * @param varForms model.ConsoleAdminListVersionDeployInstancesForms
 */
func (s *ConsoleAdminApiService) ListVersionDeployInstances(
    varForms *model.ConsoleAdminListVersionDeployInstancesForms,
) (model.ConsoleAdminListVersionDeployInstancesResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminListVersionDeployInstancesResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/listVersionDeployInstances"

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
	if varForms != nil && varForms.ProjectId != nil {
		varFormParams.Add("projectId", parameterToString(*varForms.ProjectId, ""))
	}
	if varForms != nil && varForms.GameId != nil {
		varFormParams.Add("gameId", parameterToString(*varForms.GameId, ""))
	}
	if varForms != nil && varForms.VersionId != nil {
		varFormParams.Add("versionId", parameterToString(*varForms.VersionId, ""))
	}
	if varForms != nil && varForms.DeployStatus != nil {
		varFormParams.Add("deployStatus", parameterToString(*varForms.DeployStatus, ""))
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

// RemoveGameFromProject
/*
 * 将游戏移出项目
 * @param varForms model.ConsoleAdminRemoveGameFromProjectForms
 */
func (s *ConsoleAdminApiService) RemoveGameFromProject(
    varForms *model.ConsoleAdminRemoveGameFromProjectForms,
) (model.ConsoleAdminRemoveGameFromProjectResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminRemoveGameFromProjectResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/removeGameFromProject"

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
	varFormParams.Add("projectId", parameterToString(varForms.ProjectId, ""))
	varFormParams.Add("gameId", parameterToString(varForms.GameId, ""))

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

// SubmitDeployment
/*
 * 提交游戏版本的部署请求
 * @param varForms model.ConsoleAdminSubmitDeploymentForms
 */
func (s *ConsoleAdminApiService) SubmitDeployment(
    varForms *model.ConsoleAdminSubmitDeploymentForms,
) (model.ConsoleAdminSubmitDeploymentResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminSubmitDeploymentResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/submitDeployment"

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
	varFormParams.Add("gameId", parameterToString(varForms.GameId, ""))
	varFormParams.Add("projectId", parameterToString(varForms.ProjectId, ""))
	varFormParams.Add("versionId", parameterToString(varForms.VersionId, ""))
	varFormParams.Add("cloudGameInstanceIds", parameterToString(varForms.CloudGameInstanceIds, ""))
	varFormParams.Add("operationType", parameterToString(varForms.OperationType, ""))

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

// UploadGameVersionByDownload
/*
 * 一键上传：文件上传接口，用远程下载的方式生成新版本
 * @param varForms model.ConsoleAdminUploadGameVersionByDownloadForms
 */
func (s *ConsoleAdminApiService) UploadGameVersionByDownload(
    varForms *model.ConsoleAdminUploadGameVersionByDownloadForms,
) (model.ConsoleAdminUploadGameVersionByDownloadResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.ConsoleAdminUploadGameVersionByDownloadResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/consoleAdmin/uploadGameVersionByDownload"

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
	varFormParams.Add("hash", parameterToString(varForms.Hash, ""))
	varFormParams.Add("gameId", parameterToString(varForms.GameId, ""))
	varFormParams.Add("downloadType", parameterToString(varForms.DownloadType, ""))
	varFormParams.Add("versionName", parameterToString(varForms.VersionName, ""))

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
