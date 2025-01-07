/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.75.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// SecurityManagementSecretsEncryptionAPIService SecurityManagementSecretsEncryptionAPI service
type SecurityManagementSecretsEncryptionAPIService service

type ApiReEncryptRequest struct {
	ctx context.Context
	ApiService *SecurityManagementSecretsEncryptionAPIService
	body *ReEncryptionRequestApiXO
}

func (r ApiReEncryptRequest) Body(body ReEncryptionRequestApiXO) ApiReEncryptRequest {
	r.body = &body
	return r
}

func (r ApiReEncryptRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReEncryptExecute(r)
}

/*
ReEncrypt Re-encrypt secrets using the specified key

Ensure all nodes have access to the key, and they use the same key

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReEncryptRequest
*/
func (a *SecurityManagementSecretsEncryptionAPIService) ReEncrypt(ctx context.Context) ApiReEncryptRequest {
	return ApiReEncryptRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SecurityManagementSecretsEncryptionAPIService) ReEncryptExecute(r ApiReEncryptRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityManagementSecretsEncryptionAPIService.ReEncrypt")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/secrets/encryption/re-encrypt"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
