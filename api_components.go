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
	"strings"
	"os"
)


// ComponentsAPIService ComponentsAPI service
type ComponentsAPIService service

type ApiDeleteComponentRequest struct {
	ctx context.Context
	ApiService *ComponentsAPIService
	id string
}

func (r ApiDeleteComponentRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteComponentExecute(r)
}

/*
DeleteComponent Delete a single component

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the component to delete
 @return ApiDeleteComponentRequest
*/
func (a *ComponentsAPIService) DeleteComponent(ctx context.Context, id string) ApiDeleteComponentRequest {
	return ApiDeleteComponentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ComponentsAPIService) DeleteComponentExecute(r ApiDeleteComponentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentsAPIService.DeleteComponent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/components/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetComponentByIdRequest struct {
	ctx context.Context
	ApiService *ComponentsAPIService
	id string
}

func (r ApiGetComponentByIdRequest) Execute() (*ComponentXO, *http.Response, error) {
	return r.ApiService.GetComponentByIdExecute(r)
}

/*
GetComponentById Get a single component

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the component to retrieve
 @return ApiGetComponentByIdRequest
*/
func (a *ComponentsAPIService) GetComponentById(ctx context.Context, id string) ApiGetComponentByIdRequest {
	return ApiGetComponentByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ComponentXO
func (a *ComponentsAPIService) GetComponentByIdExecute(r ApiGetComponentByIdRequest) (*ComponentXO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ComponentXO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentsAPIService.GetComponentById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/components/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetComponentsRequest struct {
	ctx context.Context
	ApiService *ComponentsAPIService
	repository *string
	continuationToken *string
}

// Repository from which you would like to retrieve components
func (r ApiGetComponentsRequest) Repository(repository string) ApiGetComponentsRequest {
	r.repository = &repository
	return r
}

// A token returned by a prior request. If present, the next page of results are returned
func (r ApiGetComponentsRequest) ContinuationToken(continuationToken string) ApiGetComponentsRequest {
	r.continuationToken = &continuationToken
	return r
}

func (r ApiGetComponentsRequest) Execute() (*PageComponentXO, *http.Response, error) {
	return r.ApiService.GetComponentsExecute(r)
}

/*
GetComponents List components

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetComponentsRequest
*/
func (a *ComponentsAPIService) GetComponents(ctx context.Context) ApiGetComponentsRequest {
	return ApiGetComponentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PageComponentXO
func (a *ComponentsAPIService) GetComponentsExecute(r ApiGetComponentsRequest) (*PageComponentXO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PageComponentXO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentsAPIService.GetComponents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/components"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repository == nil {
		return localVarReturnValue, nil, reportError("repository is required and must be specified")
	}

	if r.continuationToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "continuationToken", r.continuationToken, "", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "", "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUploadComponentRequest struct {
	ctx context.Context
	ApiService *ComponentsAPIService
	repository *string
	helmTag *string
	helmAsset *os.File
	rTag *string
	rAsset *os.File
	rAssetPathId *string
	pypiTag *string
	pypiAsset *os.File
	dockerTag *string
	dockerAsset *os.File
	yumDirectory *string
	yumTag *string
	yumAsset *os.File
	yumAssetFilename *string
	rubygemsTag *string
	rubygemsAsset *os.File
	nugetTag *string
	nugetAsset *os.File
	npmTag *string
	npmAsset *os.File
	rawDirectory *string
	rawTag *string
	rawAsset1 *os.File
	rawAsset1Filename *string
	rawAsset2 *os.File
	rawAsset2Filename *string
	rawAsset3 *os.File
	rawAsset3Filename *string
	aptTag *string
	aptAsset *os.File
	maven2GroupId *string
	maven2ArtifactId *string
	maven2Version *string
	maven2GeneratePom *bool
	maven2Packaging *string
	maven2Tag *string
	maven2Asset1 *os.File
	maven2Asset1Classifier *string
	maven2Asset1Extension *string
	maven2Asset2 *os.File
	maven2Asset2Classifier *string
	maven2Asset2Extension *string
	maven2Asset3 *os.File
	maven2Asset3Classifier *string
	maven2Asset3Extension *string
}

// Name of the repository to which you would like to upload the component
func (r ApiUploadComponentRequest) Repository(repository string) ApiUploadComponentRequest {
	r.repository = &repository
	return r
}

// helm Tag
func (r ApiUploadComponentRequest) HelmTag(helmTag string) ApiUploadComponentRequest {
	r.helmTag = &helmTag
	return r
}

// helm Asset 
func (r ApiUploadComponentRequest) HelmAsset(helmAsset *os.File) ApiUploadComponentRequest {
	r.helmAsset = helmAsset
	return r
}

// r Tag
func (r ApiUploadComponentRequest) RTag(rTag string) ApiUploadComponentRequest {
	r.rTag = &rTag
	return r
}

// r Asset 
func (r ApiUploadComponentRequest) RAsset(rAsset *os.File) ApiUploadComponentRequest {
	r.rAsset = rAsset
	return r
}

// r Asset  Package Path
func (r ApiUploadComponentRequest) RAssetPathId(rAssetPathId string) ApiUploadComponentRequest {
	r.rAssetPathId = &rAssetPathId
	return r
}

// pypi Tag
func (r ApiUploadComponentRequest) PypiTag(pypiTag string) ApiUploadComponentRequest {
	r.pypiTag = &pypiTag
	return r
}

// pypi Asset 
func (r ApiUploadComponentRequest) PypiAsset(pypiAsset *os.File) ApiUploadComponentRequest {
	r.pypiAsset = pypiAsset
	return r
}

// docker Tag
func (r ApiUploadComponentRequest) DockerTag(dockerTag string) ApiUploadComponentRequest {
	r.dockerTag = &dockerTag
	return r
}

// docker Asset 
func (r ApiUploadComponentRequest) DockerAsset(dockerAsset *os.File) ApiUploadComponentRequest {
	r.dockerAsset = dockerAsset
	return r
}

// yum Directory
func (r ApiUploadComponentRequest) YumDirectory(yumDirectory string) ApiUploadComponentRequest {
	r.yumDirectory = &yumDirectory
	return r
}

// yum Tag
func (r ApiUploadComponentRequest) YumTag(yumTag string) ApiUploadComponentRequest {
	r.yumTag = &yumTag
	return r
}

// yum Asset 
func (r ApiUploadComponentRequest) YumAsset(yumAsset *os.File) ApiUploadComponentRequest {
	r.yumAsset = yumAsset
	return r
}

// yum Asset  Filename
func (r ApiUploadComponentRequest) YumAssetFilename(yumAssetFilename string) ApiUploadComponentRequest {
	r.yumAssetFilename = &yumAssetFilename
	return r
}

// rubygems Tag
func (r ApiUploadComponentRequest) RubygemsTag(rubygemsTag string) ApiUploadComponentRequest {
	r.rubygemsTag = &rubygemsTag
	return r
}

// rubygems Asset 
func (r ApiUploadComponentRequest) RubygemsAsset(rubygemsAsset *os.File) ApiUploadComponentRequest {
	r.rubygemsAsset = rubygemsAsset
	return r
}

// nuget Tag
func (r ApiUploadComponentRequest) NugetTag(nugetTag string) ApiUploadComponentRequest {
	r.nugetTag = &nugetTag
	return r
}

// nuget Asset 
func (r ApiUploadComponentRequest) NugetAsset(nugetAsset *os.File) ApiUploadComponentRequest {
	r.nugetAsset = nugetAsset
	return r
}

// npm Tag
func (r ApiUploadComponentRequest) NpmTag(npmTag string) ApiUploadComponentRequest {
	r.npmTag = &npmTag
	return r
}

// npm Asset 
func (r ApiUploadComponentRequest) NpmAsset(npmAsset *os.File) ApiUploadComponentRequest {
	r.npmAsset = npmAsset
	return r
}

// raw Directory
func (r ApiUploadComponentRequest) RawDirectory(rawDirectory string) ApiUploadComponentRequest {
	r.rawDirectory = &rawDirectory
	return r
}

// raw Tag
func (r ApiUploadComponentRequest) RawTag(rawTag string) ApiUploadComponentRequest {
	r.rawTag = &rawTag
	return r
}

// raw Asset 1
func (r ApiUploadComponentRequest) RawAsset1(rawAsset1 *os.File) ApiUploadComponentRequest {
	r.rawAsset1 = rawAsset1
	return r
}

// raw Asset 1 Filename
func (r ApiUploadComponentRequest) RawAsset1Filename(rawAsset1Filename string) ApiUploadComponentRequest {
	r.rawAsset1Filename = &rawAsset1Filename
	return r
}

// raw Asset 2
func (r ApiUploadComponentRequest) RawAsset2(rawAsset2 *os.File) ApiUploadComponentRequest {
	r.rawAsset2 = rawAsset2
	return r
}

// raw Asset 2 Filename
func (r ApiUploadComponentRequest) RawAsset2Filename(rawAsset2Filename string) ApiUploadComponentRequest {
	r.rawAsset2Filename = &rawAsset2Filename
	return r
}

// raw Asset 3
func (r ApiUploadComponentRequest) RawAsset3(rawAsset3 *os.File) ApiUploadComponentRequest {
	r.rawAsset3 = rawAsset3
	return r
}

// raw Asset 3 Filename
func (r ApiUploadComponentRequest) RawAsset3Filename(rawAsset3Filename string) ApiUploadComponentRequest {
	r.rawAsset3Filename = &rawAsset3Filename
	return r
}

// apt Tag
func (r ApiUploadComponentRequest) AptTag(aptTag string) ApiUploadComponentRequest {
	r.aptTag = &aptTag
	return r
}

// apt Asset 
func (r ApiUploadComponentRequest) AptAsset(aptAsset *os.File) ApiUploadComponentRequest {
	r.aptAsset = aptAsset
	return r
}

// maven2 Group ID
func (r ApiUploadComponentRequest) Maven2GroupId(maven2GroupId string) ApiUploadComponentRequest {
	r.maven2GroupId = &maven2GroupId
	return r
}

// maven2 Artifact ID
func (r ApiUploadComponentRequest) Maven2ArtifactId(maven2ArtifactId string) ApiUploadComponentRequest {
	r.maven2ArtifactId = &maven2ArtifactId
	return r
}

// maven2 Version
func (r ApiUploadComponentRequest) Maven2Version(maven2Version string) ApiUploadComponentRequest {
	r.maven2Version = &maven2Version
	return r
}

// maven2 Generate a POM file with these coordinates
func (r ApiUploadComponentRequest) Maven2GeneratePom(maven2GeneratePom bool) ApiUploadComponentRequest {
	r.maven2GeneratePom = &maven2GeneratePom
	return r
}

// maven2 Packaging
func (r ApiUploadComponentRequest) Maven2Packaging(maven2Packaging string) ApiUploadComponentRequest {
	r.maven2Packaging = &maven2Packaging
	return r
}

// maven2 Tag
func (r ApiUploadComponentRequest) Maven2Tag(maven2Tag string) ApiUploadComponentRequest {
	r.maven2Tag = &maven2Tag
	return r
}

// maven2 Asset 1
func (r ApiUploadComponentRequest) Maven2Asset1(maven2Asset1 *os.File) ApiUploadComponentRequest {
	r.maven2Asset1 = maven2Asset1
	return r
}

// maven2 Asset 1 Classifier
func (r ApiUploadComponentRequest) Maven2Asset1Classifier(maven2Asset1Classifier string) ApiUploadComponentRequest {
	r.maven2Asset1Classifier = &maven2Asset1Classifier
	return r
}

// maven2 Asset 1 Extension
func (r ApiUploadComponentRequest) Maven2Asset1Extension(maven2Asset1Extension string) ApiUploadComponentRequest {
	r.maven2Asset1Extension = &maven2Asset1Extension
	return r
}

// maven2 Asset 2
func (r ApiUploadComponentRequest) Maven2Asset2(maven2Asset2 *os.File) ApiUploadComponentRequest {
	r.maven2Asset2 = maven2Asset2
	return r
}

// maven2 Asset 2 Classifier
func (r ApiUploadComponentRequest) Maven2Asset2Classifier(maven2Asset2Classifier string) ApiUploadComponentRequest {
	r.maven2Asset2Classifier = &maven2Asset2Classifier
	return r
}

// maven2 Asset 2 Extension
func (r ApiUploadComponentRequest) Maven2Asset2Extension(maven2Asset2Extension string) ApiUploadComponentRequest {
	r.maven2Asset2Extension = &maven2Asset2Extension
	return r
}

// maven2 Asset 3
func (r ApiUploadComponentRequest) Maven2Asset3(maven2Asset3 *os.File) ApiUploadComponentRequest {
	r.maven2Asset3 = maven2Asset3
	return r
}

// maven2 Asset 3 Classifier
func (r ApiUploadComponentRequest) Maven2Asset3Classifier(maven2Asset3Classifier string) ApiUploadComponentRequest {
	r.maven2Asset3Classifier = &maven2Asset3Classifier
	return r
}

// maven2 Asset 3 Extension
func (r ApiUploadComponentRequest) Maven2Asset3Extension(maven2Asset3Extension string) ApiUploadComponentRequest {
	r.maven2Asset3Extension = &maven2Asset3Extension
	return r
}

func (r ApiUploadComponentRequest) Execute() (*http.Response, error) {
	return r.ApiService.UploadComponentExecute(r)
}

/*
UploadComponent Upload a single component

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadComponentRequest
*/
func (a *ComponentsAPIService) UploadComponent(ctx context.Context) ApiUploadComponentRequest {
	return ApiUploadComponentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ComponentsAPIService) UploadComponentExecute(r ApiUploadComponentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentsAPIService.UploadComponent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/components"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repository == nil {
		return nil, reportError("repository is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	if r.helmTag != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "helm.tag", r.helmTag, "", "")
	}
	var helmAssetLocalVarFormFileName string
	var helmAssetLocalVarFileName     string
	var helmAssetLocalVarFileBytes    []byte

	helmAssetLocalVarFormFileName = "helm.asset"
	helmAssetLocalVarFile := r.helmAsset

	if helmAssetLocalVarFile != nil {
		fbs, _ := io.ReadAll(helmAssetLocalVarFile)

		helmAssetLocalVarFileBytes = fbs
		helmAssetLocalVarFileName = helmAssetLocalVarFile.Name()
		helmAssetLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: helmAssetLocalVarFileBytes, fileName: helmAssetLocalVarFileName, formFileName: helmAssetLocalVarFormFileName})
	}
	if r.rTag != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "r.tag", r.rTag, "", "")
	}
	var rAssetLocalVarFormFileName string
	var rAssetLocalVarFileName     string
	var rAssetLocalVarFileBytes    []byte

	rAssetLocalVarFormFileName = "r.asset"
	rAssetLocalVarFile := r.rAsset

	if rAssetLocalVarFile != nil {
		fbs, _ := io.ReadAll(rAssetLocalVarFile)

		rAssetLocalVarFileBytes = fbs
		rAssetLocalVarFileName = rAssetLocalVarFile.Name()
		rAssetLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: rAssetLocalVarFileBytes, fileName: rAssetLocalVarFileName, formFileName: rAssetLocalVarFormFileName})
	}
	if r.rAssetPathId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "r.asset.pathId", r.rAssetPathId, "", "")
	}
	if r.pypiTag != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pypi.tag", r.pypiTag, "", "")
	}
	var pypiAssetLocalVarFormFileName string
	var pypiAssetLocalVarFileName     string
	var pypiAssetLocalVarFileBytes    []byte

	pypiAssetLocalVarFormFileName = "pypi.asset"
	pypiAssetLocalVarFile := r.pypiAsset

	if pypiAssetLocalVarFile != nil {
		fbs, _ := io.ReadAll(pypiAssetLocalVarFile)

		pypiAssetLocalVarFileBytes = fbs
		pypiAssetLocalVarFileName = pypiAssetLocalVarFile.Name()
		pypiAssetLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: pypiAssetLocalVarFileBytes, fileName: pypiAssetLocalVarFileName, formFileName: pypiAssetLocalVarFormFileName})
	}
	if r.dockerTag != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "docker.tag", r.dockerTag, "", "")
	}
	var dockerAssetLocalVarFormFileName string
	var dockerAssetLocalVarFileName     string
	var dockerAssetLocalVarFileBytes    []byte

	dockerAssetLocalVarFormFileName = "docker.asset"
	dockerAssetLocalVarFile := r.dockerAsset

	if dockerAssetLocalVarFile != nil {
		fbs, _ := io.ReadAll(dockerAssetLocalVarFile)

		dockerAssetLocalVarFileBytes = fbs
		dockerAssetLocalVarFileName = dockerAssetLocalVarFile.Name()
		dockerAssetLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: dockerAssetLocalVarFileBytes, fileName: dockerAssetLocalVarFileName, formFileName: dockerAssetLocalVarFormFileName})
	}
	if r.yumDirectory != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "yum.directory", r.yumDirectory, "", "")
	}
	if r.yumTag != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "yum.tag", r.yumTag, "", "")
	}
	var yumAssetLocalVarFormFileName string
	var yumAssetLocalVarFileName     string
	var yumAssetLocalVarFileBytes    []byte

	yumAssetLocalVarFormFileName = "yum.asset"
	yumAssetLocalVarFile := r.yumAsset

	if yumAssetLocalVarFile != nil {
		fbs, _ := io.ReadAll(yumAssetLocalVarFile)

		yumAssetLocalVarFileBytes = fbs
		yumAssetLocalVarFileName = yumAssetLocalVarFile.Name()
		yumAssetLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: yumAssetLocalVarFileBytes, fileName: yumAssetLocalVarFileName, formFileName: yumAssetLocalVarFormFileName})
	}
	if r.yumAssetFilename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "yum.asset.filename", r.yumAssetFilename, "", "")
	}
	if r.rubygemsTag != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rubygems.tag", r.rubygemsTag, "", "")
	}
	var rubygemsAssetLocalVarFormFileName string
	var rubygemsAssetLocalVarFileName     string
	var rubygemsAssetLocalVarFileBytes    []byte

	rubygemsAssetLocalVarFormFileName = "rubygems.asset"
	rubygemsAssetLocalVarFile := r.rubygemsAsset

	if rubygemsAssetLocalVarFile != nil {
		fbs, _ := io.ReadAll(rubygemsAssetLocalVarFile)

		rubygemsAssetLocalVarFileBytes = fbs
		rubygemsAssetLocalVarFileName = rubygemsAssetLocalVarFile.Name()
		rubygemsAssetLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: rubygemsAssetLocalVarFileBytes, fileName: rubygemsAssetLocalVarFileName, formFileName: rubygemsAssetLocalVarFormFileName})
	}
	if r.nugetTag != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "nuget.tag", r.nugetTag, "", "")
	}
	var nugetAssetLocalVarFormFileName string
	var nugetAssetLocalVarFileName     string
	var nugetAssetLocalVarFileBytes    []byte

	nugetAssetLocalVarFormFileName = "nuget.asset"
	nugetAssetLocalVarFile := r.nugetAsset

	if nugetAssetLocalVarFile != nil {
		fbs, _ := io.ReadAll(nugetAssetLocalVarFile)

		nugetAssetLocalVarFileBytes = fbs
		nugetAssetLocalVarFileName = nugetAssetLocalVarFile.Name()
		nugetAssetLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: nugetAssetLocalVarFileBytes, fileName: nugetAssetLocalVarFileName, formFileName: nugetAssetLocalVarFormFileName})
	}
	if r.npmTag != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "npm.tag", r.npmTag, "", "")
	}
	var npmAssetLocalVarFormFileName string
	var npmAssetLocalVarFileName     string
	var npmAssetLocalVarFileBytes    []byte

	npmAssetLocalVarFormFileName = "npm.asset"
	npmAssetLocalVarFile := r.npmAsset

	if npmAssetLocalVarFile != nil {
		fbs, _ := io.ReadAll(npmAssetLocalVarFile)

		npmAssetLocalVarFileBytes = fbs
		npmAssetLocalVarFileName = npmAssetLocalVarFile.Name()
		npmAssetLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: npmAssetLocalVarFileBytes, fileName: npmAssetLocalVarFileName, formFileName: npmAssetLocalVarFormFileName})
	}
	if r.rawDirectory != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "raw.directory", r.rawDirectory, "", "")
	}
	if r.rawTag != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "raw.tag", r.rawTag, "", "")
	}
	var rawAsset1LocalVarFormFileName string
	var rawAsset1LocalVarFileName     string
	var rawAsset1LocalVarFileBytes    []byte

	rawAsset1LocalVarFormFileName = "raw.asset1"
	rawAsset1LocalVarFile := r.rawAsset1

	if rawAsset1LocalVarFile != nil {
		fbs, _ := io.ReadAll(rawAsset1LocalVarFile)

		rawAsset1LocalVarFileBytes = fbs
		rawAsset1LocalVarFileName = rawAsset1LocalVarFile.Name()
		rawAsset1LocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: rawAsset1LocalVarFileBytes, fileName: rawAsset1LocalVarFileName, formFileName: rawAsset1LocalVarFormFileName})
	}
	if r.rawAsset1Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "raw.asset1.filename", r.rawAsset1Filename, "", "")
	}
	var rawAsset2LocalVarFormFileName string
	var rawAsset2LocalVarFileName     string
	var rawAsset2LocalVarFileBytes    []byte

	rawAsset2LocalVarFormFileName = "raw.asset2"
	rawAsset2LocalVarFile := r.rawAsset2

	if rawAsset2LocalVarFile != nil {
		fbs, _ := io.ReadAll(rawAsset2LocalVarFile)

		rawAsset2LocalVarFileBytes = fbs
		rawAsset2LocalVarFileName = rawAsset2LocalVarFile.Name()
		rawAsset2LocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: rawAsset2LocalVarFileBytes, fileName: rawAsset2LocalVarFileName, formFileName: rawAsset2LocalVarFormFileName})
	}
	if r.rawAsset2Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "raw.asset2.filename", r.rawAsset2Filename, "", "")
	}
	var rawAsset3LocalVarFormFileName string
	var rawAsset3LocalVarFileName     string
	var rawAsset3LocalVarFileBytes    []byte

	rawAsset3LocalVarFormFileName = "raw.asset3"
	rawAsset3LocalVarFile := r.rawAsset3

	if rawAsset3LocalVarFile != nil {
		fbs, _ := io.ReadAll(rawAsset3LocalVarFile)

		rawAsset3LocalVarFileBytes = fbs
		rawAsset3LocalVarFileName = rawAsset3LocalVarFile.Name()
		rawAsset3LocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: rawAsset3LocalVarFileBytes, fileName: rawAsset3LocalVarFileName, formFileName: rawAsset3LocalVarFormFileName})
	}
	if r.rawAsset3Filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "raw.asset3.filename", r.rawAsset3Filename, "", "")
	}
	if r.aptTag != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "apt.tag", r.aptTag, "", "")
	}
	var aptAssetLocalVarFormFileName string
	var aptAssetLocalVarFileName     string
	var aptAssetLocalVarFileBytes    []byte

	aptAssetLocalVarFormFileName = "apt.asset"
	aptAssetLocalVarFile := r.aptAsset

	if aptAssetLocalVarFile != nil {
		fbs, _ := io.ReadAll(aptAssetLocalVarFile)

		aptAssetLocalVarFileBytes = fbs
		aptAssetLocalVarFileName = aptAssetLocalVarFile.Name()
		aptAssetLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: aptAssetLocalVarFileBytes, fileName: aptAssetLocalVarFileName, formFileName: aptAssetLocalVarFormFileName})
	}
	if r.maven2GroupId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "maven2.groupId", r.maven2GroupId, "", "")
	}
	if r.maven2ArtifactId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "maven2.artifactId", r.maven2ArtifactId, "", "")
	}
	if r.maven2Version != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "maven2.version", r.maven2Version, "", "")
	}
	if r.maven2GeneratePom != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "maven2.generate-pom", r.maven2GeneratePom, "", "")
	}
	if r.maven2Packaging != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "maven2.packaging", r.maven2Packaging, "", "")
	}
	if r.maven2Tag != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "maven2.tag", r.maven2Tag, "", "")
	}
	var maven2Asset1LocalVarFormFileName string
	var maven2Asset1LocalVarFileName     string
	var maven2Asset1LocalVarFileBytes    []byte

	maven2Asset1LocalVarFormFileName = "maven2.asset1"
	maven2Asset1LocalVarFile := r.maven2Asset1

	if maven2Asset1LocalVarFile != nil {
		fbs, _ := io.ReadAll(maven2Asset1LocalVarFile)

		maven2Asset1LocalVarFileBytes = fbs
		maven2Asset1LocalVarFileName = maven2Asset1LocalVarFile.Name()
		maven2Asset1LocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: maven2Asset1LocalVarFileBytes, fileName: maven2Asset1LocalVarFileName, formFileName: maven2Asset1LocalVarFormFileName})
	}
	if r.maven2Asset1Classifier != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "maven2.asset1.classifier", r.maven2Asset1Classifier, "", "")
	}
	if r.maven2Asset1Extension != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "maven2.asset1.extension", r.maven2Asset1Extension, "", "")
	}
	var maven2Asset2LocalVarFormFileName string
	var maven2Asset2LocalVarFileName     string
	var maven2Asset2LocalVarFileBytes    []byte

	maven2Asset2LocalVarFormFileName = "maven2.asset2"
	maven2Asset2LocalVarFile := r.maven2Asset2

	if maven2Asset2LocalVarFile != nil {
		fbs, _ := io.ReadAll(maven2Asset2LocalVarFile)

		maven2Asset2LocalVarFileBytes = fbs
		maven2Asset2LocalVarFileName = maven2Asset2LocalVarFile.Name()
		maven2Asset2LocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: maven2Asset2LocalVarFileBytes, fileName: maven2Asset2LocalVarFileName, formFileName: maven2Asset2LocalVarFormFileName})
	}
	if r.maven2Asset2Classifier != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "maven2.asset2.classifier", r.maven2Asset2Classifier, "", "")
	}
	if r.maven2Asset2Extension != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "maven2.asset2.extension", r.maven2Asset2Extension, "", "")
	}
	var maven2Asset3LocalVarFormFileName string
	var maven2Asset3LocalVarFileName     string
	var maven2Asset3LocalVarFileBytes    []byte

	maven2Asset3LocalVarFormFileName = "maven2.asset3"
	maven2Asset3LocalVarFile := r.maven2Asset3

	if maven2Asset3LocalVarFile != nil {
		fbs, _ := io.ReadAll(maven2Asset3LocalVarFile)

		maven2Asset3LocalVarFileBytes = fbs
		maven2Asset3LocalVarFileName = maven2Asset3LocalVarFile.Name()
		maven2Asset3LocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: maven2Asset3LocalVarFileBytes, fileName: maven2Asset3LocalVarFileName, formFileName: maven2Asset3LocalVarFormFileName})
	}
	if r.maven2Asset3Classifier != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "maven2.asset3.classifier", r.maven2Asset3Classifier, "", "")
	}
	if r.maven2Asset3Extension != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "maven2.asset3.extension", r.maven2Asset3Extension, "", "")
	}
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
