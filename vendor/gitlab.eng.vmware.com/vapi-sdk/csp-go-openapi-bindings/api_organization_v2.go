/*
 * Identity APIs
 *
 * A List of all identity and account management related APIs. Note: In order to use Bearer token as authentication method, you must include Bearer as a prefix to your token in the Authorize section.   Accepted format is: Bearer <authorization_token>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiclient

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// OrganizationV2ApiService OrganizationV2Api service
type OrganizationV2ApiService service

// GetPaginatedOrgUsersInfoUsingGETOpts Optional parameters for the method 'GetPaginatedOrgUsersInfoUsingGET'
type GetPaginatedOrgUsersInfoUsingGETOpts struct {
    ServiceDefinitionId optional.String
    PageStart optional.Int32
    PageLimit optional.Int32
    ExpandProfile optional.String
}

/*
GetPaginatedOrgUsersInfoUsingGET Get Organization Users V2
Get response encapsulating organization users.&lt;br&gt;Fetched page is according to the page start and page limit passed as optional parameters.&lt;br&gt;Optionally provide \&quot;serviceDefinitionId\&quot; to filter users having access to a service. Organization Members are permitted to see only basic user information. Organization owners and read-only administrators will see also organization and service roles of the users and userProfile if expandProfile is passed. ### Access Policy: | Role | User Accounts | Service Accounts (Client Credentials Applications)  | | ----- | ----- | ---------- | | Organization Member | &amp;#x2714;&amp;#xFE0F; | &amp;#x2714;&amp;#xFE0F; | | Organization Owner | &amp;#x2714;&amp;#xFE0F; | &amp;#x2714;&amp;#xFE0F; | | Read-only Operator | &amp;#x2714;&amp;#xFE0F; | &amp;#x2714;&amp;#xFE0F; | 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization ID
 * @param optional nil or *GetPaginatedOrgUsersInfoUsingGETOpts - Optional Parameters:
 * @param "ServiceDefinitionId" (optional.String) -  Service definition id used to filter users having access to the service.
 * @param "PageStart" (optional.Int32) -  Start index, the default value is 1
 * @param "PageLimit" (optional.Int32) -  Maximum number of users to return in response, the default value is 200
 * @param "ExpandProfile" (optional.String) -  Indicates if the response should be expanded with the user profile, the value is ignored, only the existence of parameter is checked.
@return PagedResponseOfExpandedTypedUserDto
*/
func (a *OrganizationV2ApiService) GetPaginatedOrgUsersInfoUsingGET(ctx _context.Context, orgId string, localVarOptionals *GetPaginatedOrgUsersInfoUsingGETOpts) (PagedResponseOfExpandedTypedUserDto, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PagedResponseOfExpandedTypedUserDto
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/am/api/v2/orgs/{orgId}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.ServiceDefinitionId.IsSet() {
		localVarQueryParams.Add("serviceDefinitionId", parameterToString(localVarOptionals.ServiceDefinitionId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageStart.IsSet() {
		localVarQueryParams.Add("pageStart", parameterToString(localVarOptionals.PageStart.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageLimit.IsSet() {
		localVarQueryParams.Add("pageLimit", parameterToString(localVarOptionals.PageLimit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExpandProfile.IsSet() {
		localVarQueryParams.Add("expandProfile", parameterToString(localVarOptionals.ExpandProfile.Value(), ""))
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
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
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

/*
RemoveUsersFromOrgUsingDELETE Remove Users From Organization V2
Remove users from organization by user ids.&lt;br&gt;User ids will be of the format &lt;IdpId&gt;:&lt;UUID&gt; e.g. vmware.com:820e7ca5-4024-407e-8db4-f552d5d03403.&lt;br&gt;Pay attention: in case of partial success the caller must read the response to see which users have not been added successfully ### Access Policy: | Role | User Accounts | Service Accounts (Client Credentials Applications)  | | ----- | ----- | ---------- | | Organization Owner | &amp;#x2714;&amp;#xFE0F; | &amp;#x2714;&amp;#xFE0F; | 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization ID
 * @param request request
@return PartialSuccessResponseOfstring
*/
func (a *OrganizationV2ApiService) RemoveUsersFromOrgUsingDELETE(ctx _context.Context, orgId string, request RemoveUsersFromOrganizationRequestV2) (PartialSuccessResponseOfstring, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PartialSuccessResponseOfstring
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/am/api/v2/orgs/{orgId}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarPostBody = &request
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
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
