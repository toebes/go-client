/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.113
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package onshape

import (
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

// GlobalTreeNodesApiService GlobalTreeNodesApi service
type GlobalTreeNodesApiService service

type apiGlobalTreeNodesRequest struct {
	ctx _context.Context
	apiService *GlobalTreeNodesApiService
}


/*
GlobalTreeNodes Get Searchable Trees
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiGlobalTreeNodesRequest
*/
func (a *GlobalTreeNodesApiService) GlobalTreeNodes(ctx _context.Context) apiGlobalTreeNodesRequest {
	return apiGlobalTreeNodesRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request
 @return BTGlobalTreeNodesInfo
*/
func (r apiGlobalTreeNodesRequest) Execute() (BTGlobalTreeNodesInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTGlobalTreeNodesInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "GlobalTreeNodesApiService.GlobalTreeNodes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/globaltreenodes/"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v BTGlobalTreeNodesInfo
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiGlobalTreeNodesFolderRequest struct {
	ctx _context.Context
	apiService *GlobalTreeNodesApiService
	fid string
	getPathToRoot *bool
	offset *int32
	limit *int32
	sortColumn *string
	sortOrder *string
}


func (r apiGlobalTreeNodesFolderRequest) GetPathToRoot(getPathToRoot bool) apiGlobalTreeNodesFolderRequest {
	r.getPathToRoot = &getPathToRoot
	return r
}

func (r apiGlobalTreeNodesFolderRequest) Offset(offset int32) apiGlobalTreeNodesFolderRequest {
	r.offset = &offset
	return r
}

func (r apiGlobalTreeNodesFolderRequest) Limit(limit int32) apiGlobalTreeNodesFolderRequest {
	r.limit = &limit
	return r
}

func (r apiGlobalTreeNodesFolderRequest) SortColumn(sortColumn string) apiGlobalTreeNodesFolderRequest {
	r.sortColumn = &sortColumn
	return r
}

func (r apiGlobalTreeNodesFolderRequest) SortOrder(sortOrder string) apiGlobalTreeNodesFolderRequest {
	r.sortOrder = &sortOrder
	return r
}

/*
GlobalTreeNodesFolder Get Tree Node List
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param fid ID of the folder to traverse.
@return apiGlobalTreeNodesFolderRequest
*/
func (a *GlobalTreeNodesApiService) GlobalTreeNodesFolder(ctx _context.Context, fid string) apiGlobalTreeNodesFolderRequest {
	return apiGlobalTreeNodesFolderRequest{
		apiService: a,
		ctx: ctx,
		fid: fid,
	}
}

/*
Execute executes the request
 @return BTGlobalTreeNodesInfo
*/
func (r apiGlobalTreeNodesFolderRequest) Execute() (BTGlobalTreeNodesInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTGlobalTreeNodesInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "GlobalTreeNodesApiService.GlobalTreeNodesFolder")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/globaltreenodes/folder/{fid}"
	localVarPath = strings.Replace(localVarPath, "{"+"fid"+"}", _neturl.QueryEscape(parameterToString(r.fid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
					
	if r.getPathToRoot != nil {
		localVarQueryParams.Add("getPathToRoot", parameterToString(*r.getPathToRoot, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.sortColumn != nil {
		localVarQueryParams.Add("sortColumn", parameterToString(*r.sortColumn, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v BTGlobalTreeNodesInfo
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiGlobalTreeNodesMagicRequest struct {
	ctx _context.Context
	apiService *GlobalTreeNodesApiService
	mid string
	getPathToRoot *bool
	offset *int32
	limit *int32
	sortColumn *string
	sortOrder *string
}


func (r apiGlobalTreeNodesMagicRequest) GetPathToRoot(getPathToRoot bool) apiGlobalTreeNodesMagicRequest {
	r.getPathToRoot = &getPathToRoot
	return r
}

func (r apiGlobalTreeNodesMagicRequest) Offset(offset int32) apiGlobalTreeNodesMagicRequest {
	r.offset = &offset
	return r
}

func (r apiGlobalTreeNodesMagicRequest) Limit(limit int32) apiGlobalTreeNodesMagicRequest {
	r.limit = &limit
	return r
}

func (r apiGlobalTreeNodesMagicRequest) SortColumn(sortColumn string) apiGlobalTreeNodesMagicRequest {
	r.sortColumn = &sortColumn
	return r
}

func (r apiGlobalTreeNodesMagicRequest) SortOrder(sortOrder string) apiGlobalTreeNodesMagicRequest {
	r.sortOrder = &sortOrder
	return r
}

/*
GlobalTreeNodesMagic Get Tree Node List
Returns a list of all the elements in one of several globally defined lists. Known values include:
| MAGIC ID | TITLE | USAGE |
|-|-|-|
| 0 | Recently Opened | Most recently opened documents |
| 1 | My Onshape | Root folder and contents |
| 2 | Created by Me | Documents created by the logged in user |
| 3 | Public | All public documents (a very long list) |
| 4 | Trash | Trashcan for the logged in user |
| 5 | Tutorials & Samples | Desktop Tutorials |
| 6 | FeatureScript samples |  FeatureScript samples (found when you select Other documents while adding a custom feature) |
| 7 | Community spotlight | Community spotlight (found when you select Other documents while adding a custom feature) |
| 8 | Tutorials | IOS Tutorials |
| 9 | Tutorials | Android Tutorials |
| 10 | Labels | Labels created by the user  |
| 11 | Teams | Teams that the user is connected to |
| 12 | Shared with me | Documents shared with the user |
| 13 | Cloud Storage | Visual list of cloud accounts associated with the logged in user |
| 14 | Custom table samples | Custom table samples (found when you select Other documents while adding a custom table) |

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param mid Magic ID of list to return.
@return apiGlobalTreeNodesMagicRequest
*/
func (a *GlobalTreeNodesApiService) GlobalTreeNodesMagic(ctx _context.Context, mid string) apiGlobalTreeNodesMagicRequest {
	return apiGlobalTreeNodesMagicRequest{
		apiService: a,
		ctx: ctx,
		mid: mid,
	}
}

/*
Execute executes the request
 @return BTGlobalTreeNodesInfo
*/
func (r apiGlobalTreeNodesMagicRequest) Execute() (BTGlobalTreeNodesInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BTGlobalTreeNodesInfo
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "GlobalTreeNodesApiService.GlobalTreeNodesMagic")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/globaltreenodes/magic/{mid}"
	localVarPath = strings.Replace(localVarPath, "{"+"mid"+"}", _neturl.QueryEscape(parameterToString(r.mid, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
					
	if r.getPathToRoot != nil {
		localVarQueryParams.Add("getPathToRoot", parameterToString(*r.getPathToRoot, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.sortColumn != nil {
		localVarQueryParams.Add("sortColumn", parameterToString(*r.sortColumn, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v BTGlobalTreeNodesInfo
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
