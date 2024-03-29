/*
Statuspage API

# Code of Conduct Please don't abuse the API, and please report all feature requests and issues to https://support.atlassian.com/contact  # Rate Limiting Each API token is limited to 1 request / second as measured on a 60 second rolling window. To get this limit increased, please contact us at https://support.atlassian.com/contact  Error codes 420 or 429 indicate that you have exceeded the rate limit and the request has been rejected.  # Basics  ## HTTPS It's required  ## URL Prefix In order to maintain version integrity into the future, the API is versioned. All calls currently begin with the following prefix:    https://api.statuspage.io/v1/  ## RESTful Interface Wherever possible, the API seeks to implement repeatable patterns with logical, representative URLs and descriptive HTTP verbs. Below are some examples and conventions you will see throughout the documentation.  * Collections are buckets: https://api.statuspage.io/v1/pages/asdf123/incidents.json * Elements have unique IDs: https://api.statuspage.io/v1/pages/asdf123/incidents/jklm456.json * GET will retrieve information about a collection/element * POST will create an element in a collection * PATCH will update a single element * PUT will replace a single element in a collection (rarely used) * DELETE will destroy a single element  ## Sending Data Information can be sent in the body as form urlencoded or JSON, but make sure the Content-Type header matches the body structure or the server gremlins will be angry.  All examples are provided in JSON format, however they can easily be converted to form encoding if required.  Some examples of how to convert things are below:      // JSON     {       \"incident\": {         \"name\": \"test incident\",         \"components\": [\"8kbf7d35c070\", \"vtnh60py4yd7\"]       }     }      // Form Encoded (using curl as an example):     curl -X POST https://api.statuspage.io/v1/example \\       -d \"incident[name]=test incident\" \\       -d \"incident[components][]=8kbf7d35c070\" \\       -d \"incident[components][]=vtnh60py4yd7\"  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// PageAccessGroupsApiService PageAccessGroupsApi service
type PageAccessGroupsApiService service

type ApiDeletePagesPageIdPageAccessGroupsPageAccessGroupIdRequest struct {
	ctx _context.Context
	ApiService *PageAccessGroupsApiService
	pageId string
	pageAccessGroupId string
}


func (r ApiDeletePagesPageIdPageAccessGroupsPageAccessGroupIdRequest) Execute() (PageAccessGroup, *_nethttp.Response, error) {
	return r.ApiService.DeletePagesPageIdPageAccessGroupsPageAccessGroupIdExecute(r)
}

/*
DeletePagesPageIdPageAccessGroupsPageAccessGroupId Remove a page access group

Remove a page access group

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pageId Page identifier
 @param pageAccessGroupId Page Access Group Identifier
 @return ApiDeletePagesPageIdPageAccessGroupsPageAccessGroupIdRequest
*/
func (a *PageAccessGroupsApiService) DeletePagesPageIdPageAccessGroupsPageAccessGroupId(ctx _context.Context, pageId string, pageAccessGroupId string) ApiDeletePagesPageIdPageAccessGroupsPageAccessGroupIdRequest {
	return ApiDeletePagesPageIdPageAccessGroupsPageAccessGroupIdRequest{
		ApiService: a,
		ctx: ctx,
		pageId: pageId,
		pageAccessGroupId: pageAccessGroupId,
	}
}

// Execute executes the request
//  @return PageAccessGroup
func (a *PageAccessGroupsApiService) DeletePagesPageIdPageAccessGroupsPageAccessGroupIdExecute(r ApiDeletePagesPageIdPageAccessGroupsPageAccessGroupIdRequest) (PageAccessGroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PageAccessGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PageAccessGroupsApiService.DeletePagesPageIdPageAccessGroupsPageAccessGroupId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pages/{page_id}/page_access_groups/{page_access_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"page_id"+"}", _neturl.PathEscape(parameterToString(r.pageId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"page_access_group_id"+"}", _neturl.PathEscape(parameterToString(r.pageAccessGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPagesPageIdPageAccessGroupsRequest struct {
	ctx _context.Context
	ApiService *PageAccessGroupsApiService
	pageId string
	page *int32
	perPage *int32
}

// Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided.
func (r ApiGetPagesPageIdPageAccessGroupsRequest) Page(page int32) ApiGetPagesPageIdPageAccessGroupsRequest {
	r.page = &page
	return r
}
// Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided.
func (r ApiGetPagesPageIdPageAccessGroupsRequest) PerPage(perPage int32) ApiGetPagesPageIdPageAccessGroupsRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetPagesPageIdPageAccessGroupsRequest) Execute() ([]PageAccessGroup, *_nethttp.Response, error) {
	return r.ApiService.GetPagesPageIdPageAccessGroupsExecute(r)
}

/*
GetPagesPageIdPageAccessGroups Get a list of page access groups

Get a list of page access groups

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pageId Page identifier
 @return ApiGetPagesPageIdPageAccessGroupsRequest
*/
func (a *PageAccessGroupsApiService) GetPagesPageIdPageAccessGroups(ctx _context.Context, pageId string) ApiGetPagesPageIdPageAccessGroupsRequest {
	return ApiGetPagesPageIdPageAccessGroupsRequest{
		ApiService: a,
		ctx: ctx,
		pageId: pageId,
	}
}

// Execute executes the request
//  @return []PageAccessGroup
func (a *PageAccessGroupsApiService) GetPagesPageIdPageAccessGroupsExecute(r ApiGetPagesPageIdPageAccessGroupsRequest) ([]PageAccessGroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []PageAccessGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PageAccessGroupsApiService.GetPagesPageIdPageAccessGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pages/{page_id}/page_access_groups"
	localVarPath = strings.Replace(localVarPath, "{"+"page_id"+"}", _neturl.PathEscape(parameterToString(r.pageId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPagesPageIdPageAccessGroupsPageAccessGroupIdRequest struct {
	ctx _context.Context
	ApiService *PageAccessGroupsApiService
	pageId string
	pageAccessGroupId string
}


func (r ApiGetPagesPageIdPageAccessGroupsPageAccessGroupIdRequest) Execute() (PageAccessGroup, *_nethttp.Response, error) {
	return r.ApiService.GetPagesPageIdPageAccessGroupsPageAccessGroupIdExecute(r)
}

/*
GetPagesPageIdPageAccessGroupsPageAccessGroupId Get a page access group

Get a page access group

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pageId Page identifier
 @param pageAccessGroupId Page Access Group Identifier
 @return ApiGetPagesPageIdPageAccessGroupsPageAccessGroupIdRequest
*/
func (a *PageAccessGroupsApiService) GetPagesPageIdPageAccessGroupsPageAccessGroupId(ctx _context.Context, pageId string, pageAccessGroupId string) ApiGetPagesPageIdPageAccessGroupsPageAccessGroupIdRequest {
	return ApiGetPagesPageIdPageAccessGroupsPageAccessGroupIdRequest{
		ApiService: a,
		ctx: ctx,
		pageId: pageId,
		pageAccessGroupId: pageAccessGroupId,
	}
}

// Execute executes the request
//  @return PageAccessGroup
func (a *PageAccessGroupsApiService) GetPagesPageIdPageAccessGroupsPageAccessGroupIdExecute(r ApiGetPagesPageIdPageAccessGroupsPageAccessGroupIdRequest) (PageAccessGroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PageAccessGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PageAccessGroupsApiService.GetPagesPageIdPageAccessGroupsPageAccessGroupId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pages/{page_id}/page_access_groups/{page_access_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"page_id"+"}", _neturl.PathEscape(parameterToString(r.pageId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"page_access_group_id"+"}", _neturl.PathEscape(parameterToString(r.pageAccessGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchPagesPageIdPageAccessGroupsPageAccessGroupIdRequest struct {
	ctx _context.Context
	ApiService *PageAccessGroupsApiService
	pageId string
	pageAccessGroupId string
	patchPagesPageIdPageAccessGroups *PatchPagesPageIdPageAccessGroups
}

func (r ApiPatchPagesPageIdPageAccessGroupsPageAccessGroupIdRequest) PatchPagesPageIdPageAccessGroups(patchPagesPageIdPageAccessGroups PatchPagesPageIdPageAccessGroups) ApiPatchPagesPageIdPageAccessGroupsPageAccessGroupIdRequest {
	r.patchPagesPageIdPageAccessGroups = &patchPagesPageIdPageAccessGroups
	return r
}

func (r ApiPatchPagesPageIdPageAccessGroupsPageAccessGroupIdRequest) Execute() (PageAccessGroup, *_nethttp.Response, error) {
	return r.ApiService.PatchPagesPageIdPageAccessGroupsPageAccessGroupIdExecute(r)
}

/*
PatchPagesPageIdPageAccessGroupsPageAccessGroupId Update a page access group

Update a page access group

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pageId Page identifier
 @param pageAccessGroupId Page Access Group Identifier
 @return ApiPatchPagesPageIdPageAccessGroupsPageAccessGroupIdRequest
*/
func (a *PageAccessGroupsApiService) PatchPagesPageIdPageAccessGroupsPageAccessGroupId(ctx _context.Context, pageId string, pageAccessGroupId string) ApiPatchPagesPageIdPageAccessGroupsPageAccessGroupIdRequest {
	return ApiPatchPagesPageIdPageAccessGroupsPageAccessGroupIdRequest{
		ApiService: a,
		ctx: ctx,
		pageId: pageId,
		pageAccessGroupId: pageAccessGroupId,
	}
}

// Execute executes the request
//  @return PageAccessGroup
func (a *PageAccessGroupsApiService) PatchPagesPageIdPageAccessGroupsPageAccessGroupIdExecute(r ApiPatchPagesPageIdPageAccessGroupsPageAccessGroupIdRequest) (PageAccessGroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PageAccessGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PageAccessGroupsApiService.PatchPagesPageIdPageAccessGroupsPageAccessGroupId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pages/{page_id}/page_access_groups/{page_access_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"page_id"+"}", _neturl.PathEscape(parameterToString(r.pageId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"page_access_group_id"+"}", _neturl.PathEscape(parameterToString(r.pageAccessGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.patchPagesPageIdPageAccessGroups == nil {
		return localVarReturnValue, nil, reportError("patchPagesPageIdPageAccessGroups is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.patchPagesPageIdPageAccessGroups
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostPagesPageIdPageAccessGroupsRequest struct {
	ctx _context.Context
	ApiService *PageAccessGroupsApiService
	pageId string
	postPagesPageIdPageAccessGroups *PostPagesPageIdPageAccessGroups
}

func (r ApiPostPagesPageIdPageAccessGroupsRequest) PostPagesPageIdPageAccessGroups(postPagesPageIdPageAccessGroups PostPagesPageIdPageAccessGroups) ApiPostPagesPageIdPageAccessGroupsRequest {
	r.postPagesPageIdPageAccessGroups = &postPagesPageIdPageAccessGroups
	return r
}

func (r ApiPostPagesPageIdPageAccessGroupsRequest) Execute() (PageAccessGroup, *_nethttp.Response, error) {
	return r.ApiService.PostPagesPageIdPageAccessGroupsExecute(r)
}

/*
PostPagesPageIdPageAccessGroups Create a page access group

Create a page access group

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pageId Page identifier
 @return ApiPostPagesPageIdPageAccessGroupsRequest
*/
func (a *PageAccessGroupsApiService) PostPagesPageIdPageAccessGroups(ctx _context.Context, pageId string) ApiPostPagesPageIdPageAccessGroupsRequest {
	return ApiPostPagesPageIdPageAccessGroupsRequest{
		ApiService: a,
		ctx: ctx,
		pageId: pageId,
	}
}

// Execute executes the request
//  @return PageAccessGroup
func (a *PageAccessGroupsApiService) PostPagesPageIdPageAccessGroupsExecute(r ApiPostPagesPageIdPageAccessGroupsRequest) (PageAccessGroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PageAccessGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PageAccessGroupsApiService.PostPagesPageIdPageAccessGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pages/{page_id}/page_access_groups"
	localVarPath = strings.Replace(localVarPath, "{"+"page_id"+"}", _neturl.PathEscape(parameterToString(r.pageId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.postPagesPageIdPageAccessGroups == nil {
		return localVarReturnValue, nil, reportError("postPagesPageIdPageAccessGroups is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.postPagesPageIdPageAccessGroups
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutPagesPageIdPageAccessGroupsPageAccessGroupIdRequest struct {
	ctx _context.Context
	ApiService *PageAccessGroupsApiService
	pageId string
	pageAccessGroupId string
	putPagesPageIdPageAccessGroups *PutPagesPageIdPageAccessGroups
}

func (r ApiPutPagesPageIdPageAccessGroupsPageAccessGroupIdRequest) PutPagesPageIdPageAccessGroups(putPagesPageIdPageAccessGroups PutPagesPageIdPageAccessGroups) ApiPutPagesPageIdPageAccessGroupsPageAccessGroupIdRequest {
	r.putPagesPageIdPageAccessGroups = &putPagesPageIdPageAccessGroups
	return r
}

func (r ApiPutPagesPageIdPageAccessGroupsPageAccessGroupIdRequest) Execute() (PageAccessGroup, *_nethttp.Response, error) {
	return r.ApiService.PutPagesPageIdPageAccessGroupsPageAccessGroupIdExecute(r)
}

/*
PutPagesPageIdPageAccessGroupsPageAccessGroupId Update a page access group

Update a page access group

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pageId Page identifier
 @param pageAccessGroupId Page Access Group Identifier
 @return ApiPutPagesPageIdPageAccessGroupsPageAccessGroupIdRequest
*/
func (a *PageAccessGroupsApiService) PutPagesPageIdPageAccessGroupsPageAccessGroupId(ctx _context.Context, pageId string, pageAccessGroupId string) ApiPutPagesPageIdPageAccessGroupsPageAccessGroupIdRequest {
	return ApiPutPagesPageIdPageAccessGroupsPageAccessGroupIdRequest{
		ApiService: a,
		ctx: ctx,
		pageId: pageId,
		pageAccessGroupId: pageAccessGroupId,
	}
}

// Execute executes the request
//  @return PageAccessGroup
func (a *PageAccessGroupsApiService) PutPagesPageIdPageAccessGroupsPageAccessGroupIdExecute(r ApiPutPagesPageIdPageAccessGroupsPageAccessGroupIdRequest) (PageAccessGroup, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PageAccessGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PageAccessGroupsApiService.PutPagesPageIdPageAccessGroupsPageAccessGroupId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pages/{page_id}/page_access_groups/{page_access_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"page_id"+"}", _neturl.PathEscape(parameterToString(r.pageId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"page_access_group_id"+"}", _neturl.PathEscape(parameterToString(r.pageAccessGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.putPagesPageIdPageAccessGroups == nil {
		return localVarReturnValue, nil, reportError("putPagesPageIdPageAccessGroups is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.putPagesPageIdPageAccessGroups
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
