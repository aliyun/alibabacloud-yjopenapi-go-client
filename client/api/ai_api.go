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

type AiApiService service


// BatchUpdateKbVersion
/*
 * 
 * @param varForms model.AiBatchUpdateKbVersionForms
 */
func (s *AiApiService) BatchUpdateKbVersion(
    varForms *model.AiBatchUpdateKbVersionForms,
) (model.AiBatchUpdateKbVersionResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.AiBatchUpdateKbVersionResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/ai/batchUpdateKbVersion"

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
	varFormParams.Add("roleIds", parameterToString(varForms.RoleIds, "multi"))
	varFormParams.Add("kbId", parameterToString(varForms.KbId, ""))
	varFormParams.Add("kbVersionId", parameterToString(varForms.KbVersionId, ""))
	varFormParams.Add("operator", parameterToString(varForms.Operator, ""))

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

// CreateNpcRole
/*
 * 
 * @param varForms model.AiCreateNpcRoleForms
 */
func (s *AiApiService) CreateNpcRole(
    varForms *model.AiCreateNpcRoleForms,
) (model.AiCreateNpcRoleResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.AiCreateNpcRoleResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/ai/createNpcRole"

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
	varFormParams.Add("name", parameterToString(varForms.Name, ""))
	varFormParams.Add("description", parameterToString(varForms.Description, ""))
	varFormParams.Add("figureId", parameterToString(varForms.FigureId, ""))
	varFormParams.Add("voiceId", parameterToString(varForms.VoiceId, ""))
	varFormParams.Add("characters", parameterToString(varForms.Characters, ""))
	varFormParams.Add("greetings", parameterToString(varForms.Greetings, ""))
	varFormParams.Add("guidance", parameterToString(varForms.Guidance, "multi"))
	if varForms != nil && varForms.KnowledgeBases != nil {
		varFormParams.Add("knowledgeBases", parameterToString(*varForms.KnowledgeBases, "multi"))
	}
	if varForms != nil && varForms.PluginIds != nil {
		varFormParams.Add("pluginIds", parameterToString(*varForms.PluginIds, "multi"))
	}
	varFormParams.Add("shortMemoryRound", parameterToString(varForms.ShortMemoryRound, ""))
	varFormParams.Add("tenantId", parameterToString(varForms.TenantId, ""))
	varFormParams.Add("operator", parameterToString(varForms.Operator, ""))
	varFormParams.Add("llmModelId", parameterToString(varForms.LlmModelId, ""))

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

// UpdateNpcRole
/*
 * 
 * @param varForms model.AiUpdateNpcRoleForms
 */
func (s *AiApiService) UpdateNpcRole(
    varForms *model.AiUpdateNpcRoleForms,
) (model.AiUpdateNpcRoleResult, *http.Response, error) {
	var (
		varHttpMethod = strings.ToUpper("Post")
        varReturnValue model.AiUpdateNpcRoleResult
	)

	// create path and map variables
	varPath := s.client.cfg.Scheme + "://" + s.client.cfg.Host + "/ai/updateNpcRole"

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
	varFormParams.Add("name", parameterToString(varForms.Name, ""))
	varFormParams.Add("description", parameterToString(varForms.Description, ""))
	varFormParams.Add("figureId", parameterToString(varForms.FigureId, ""))
	varFormParams.Add("voiceId", parameterToString(varForms.VoiceId, ""))
	varFormParams.Add("characters", parameterToString(varForms.Characters, ""))
	varFormParams.Add("greetings", parameterToString(varForms.Greetings, ""))
	varFormParams.Add("guidance", parameterToString(varForms.Guidance, "multi"))
	if varForms != nil && varForms.KnowledgeBases != nil {
		varFormParams.Add("knowledgeBases", parameterToString(*varForms.KnowledgeBases, "multi"))
	}
	if varForms != nil && varForms.PluginIds != nil {
		varFormParams.Add("pluginIds", parameterToString(*varForms.PluginIds, "multi"))
	}
	varFormParams.Add("shortMemoryRound", parameterToString(varForms.ShortMemoryRound, ""))
	varFormParams.Add("operator", parameterToString(varForms.Operator, ""))
	varFormParams.Add("llmModelId", parameterToString(varForms.LlmModelId, ""))

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
