# nexus\BlobStoreAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertBlobStoreToGroup**](BlobStoreAPI.md#ConvertBlobStoreToGroup) | **Post** /v1/blobstores/group/convert/{name}/{newNameForOriginal} | Convert a blob store to a group blob store
[**CreateBlobStore**](BlobStoreAPI.md#CreateBlobStore) | **Post** /v1/blobstores/s3 | Create an S3 blob store
[**CreateBlobStore1**](BlobStoreAPI.md#CreateBlobStore1) | **Post** /v1/blobstores/azure | Create an Azure blob store
[**CreateBlobStore2**](BlobStoreAPI.md#CreateBlobStore2) | **Post** /v1/blobstores/google | Create a Google Cloud blob store
[**CreateFileBlobStore**](BlobStoreAPI.md#CreateFileBlobStore) | **Post** /v1/blobstores/file | Create a file blob store
[**CreateGroupBlobStore**](BlobStoreAPI.md#CreateGroupBlobStore) | **Post** /v1/blobstores/group | Create a group blob store
[**DeleteBlobStore**](BlobStoreAPI.md#DeleteBlobStore) | **Delete** /v1/blobstores/{name} | Delete a blob store by name
[**GetBlobStore**](BlobStoreAPI.md#GetBlobStore) | **Get** /v1/blobstores/s3/{name} | Get a S3 blob store configuration by name
[**GetBlobStore1**](BlobStoreAPI.md#GetBlobStore1) | **Get** /v1/blobstores/azure/{name} | Get an Azure blob store configuration by name
[**GetBlobStore2**](BlobStoreAPI.md#GetBlobStore2) | **Get** /v1/blobstores/google/{name} | Get the configuration for a Google Cloud blob store
[**GetFileBlobStoreConfiguration**](BlobStoreAPI.md#GetFileBlobStoreConfiguration) | **Get** /v1/blobstores/file/{name} | Get a file blob store configuration by name
[**GetGroupBlobStoreConfiguration**](BlobStoreAPI.md#GetGroupBlobStoreConfiguration) | **Get** /v1/blobstores/group/{name} | Get a group blob store configuration by name
[**GetRegionsByProjectId**](BlobStoreAPI.md#GetRegionsByProjectId) | **Get** /v1/blobstores/google/regions/{projectId} | Get the project regions by project&#39;s id
[**ListBlobStores**](BlobStoreAPI.md#ListBlobStores) | **Get** /v1/blobstores | List the blob stores
[**QuotaStatus**](BlobStoreAPI.md#QuotaStatus) | **Get** /v1/blobstores/{name}/quota-status | Get quota status for a given blob store
[**UpdateBlobStore**](BlobStoreAPI.md#UpdateBlobStore) | **Put** /v1/blobstores/s3/{name} | Update an S3 blob store configuration by name
[**UpdateBlobStore1**](BlobStoreAPI.md#UpdateBlobStore1) | **Put** /v1/blobstores/azure/{name} | Update an Azure blob store configuration by name
[**UpdateBlobStore2**](BlobStoreAPI.md#UpdateBlobStore2) | **Put** /v1/blobstores/google/{name} | Update a Google Cloud blob store
[**UpdateFileBlobStore**](BlobStoreAPI.md#UpdateFileBlobStore) | **Put** /v1/blobstores/file/{name} | Update a file blob store configuration by name
[**UpdateGroupBlobStore**](BlobStoreAPI.md#UpdateGroupBlobStore) | **Put** /v1/blobstores/group/{name} | Update a group blob store configuration by name



## ConvertBlobStoreToGroup

> GroupBlobStoreApiModel ConvertBlobStoreToGroup(ctx, name, newNameForOriginal).Execute()

Convert a blob store to a group blob store

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "name_example" // string | The name of the group blob store
	newNameForOriginal := "newNameForOriginal_example" // string | A new name to the original blob store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.ConvertBlobStoreToGroup(context.Background(), name, newNameForOriginal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.ConvertBlobStoreToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertBlobStoreToGroup`: GroupBlobStoreApiModel
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.ConvertBlobStoreToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the group blob store | 
**newNameForOriginal** | **string** | A new name to the original blob store | 

### Other Parameters

Other parameters are passed through a pointer to a apiConvertBlobStoreToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupBlobStoreApiModel**](GroupBlobStoreApiModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBlobStore

> CreateBlobStore(ctx).Body(body).Execute()

Create an S3 blob store

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	body := *openapiclient.NewS3BlobStoreApiModel("s3", *openapiclient.NewS3BlobStoreApiBucketConfiguration(*openapiclient.NewS3BlobStoreApiBucket("DEFAULT", "Name_example", int32(3)))) // S3BlobStoreApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.CreateBlobStore(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.CreateBlobStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**S3BlobStoreApiModel**](S3BlobStoreApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBlobStore1

> CreateBlobStore1(ctx).Body(body).Execute()

Create an Azure blob store

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	body := *openapiclient.NewAzureBlobStoreApiModel("Name_example", *openapiclient.NewAzureBlobStoreApiBucketConfiguration("AccountName_example", "ContainerName_example", *openapiclient.NewAzureBlobStoreApiAuthentication("AuthenticationMethod_example"))) // AzureBlobStoreApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.CreateBlobStore1(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.CreateBlobStore1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobStore1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AzureBlobStoreApiModel**](AzureBlobStoreApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBlobStore2

> CreateBlobStore2(ctx).Body(body).Execute()

Create a Google Cloud blob store

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	body := *openapiclient.NewGoogleCloudBlobstoreApiModel("gc_storage", *openapiclient.NewGoogleCloudBlobStoreApiBucketConfiguration(*openapiclient.NewGoogleCloudBlobStoreApiBucket("Name_example", "us-central1"))) // GoogleCloudBlobstoreApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.CreateBlobStore2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.CreateBlobStore2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobStore2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GoogleCloudBlobstoreApiModel**](GoogleCloudBlobstoreApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFileBlobStore

> CreateFileBlobStore(ctx).Body(body).Execute()

Create a file blob store

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	body := *openapiclient.NewFileBlobStoreApiCreateRequest() // FileBlobStoreApiCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.CreateFileBlobStore(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.CreateFileBlobStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileBlobStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FileBlobStoreApiCreateRequest**](FileBlobStoreApiCreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroupBlobStore

> CreateGroupBlobStore(ctx).Body(body).Execute()

Create a group blob store

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	body := *openapiclient.NewGroupBlobStoreApiCreateRequest() // GroupBlobStoreApiCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.CreateGroupBlobStore(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.CreateGroupBlobStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupBlobStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GroupBlobStoreApiCreateRequest**](GroupBlobStoreApiCreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlobStore

> DeleteBlobStore(ctx, name).Execute()

Delete a blob store by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "name_example" // string | The name of the blob store to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.DeleteBlobStore(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.DeleteBlobStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the blob store to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlobStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobStore

> S3BlobStoreApiModel GetBlobStore(ctx, name).Execute()

Get a S3 blob store configuration by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "name_example" // string | Name of the blob store configuration to fetch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.GetBlobStore(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetBlobStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobStore`: S3BlobStoreApiModel
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetBlobStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the blob store configuration to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**S3BlobStoreApiModel**](S3BlobStoreApiModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobStore1

> AzureBlobStoreApiModel GetBlobStore1(ctx, name).Execute()

Get an Azure blob store configuration by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "name_example" // string | Name of the blob store configuration to fetch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.GetBlobStore1(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetBlobStore1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobStore1`: AzureBlobStoreApiModel
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetBlobStore1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the blob store configuration to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobStore1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AzureBlobStoreApiModel**](AzureBlobStoreApiModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobStore2

> GoogleCloudBlobstoreApiModel GetBlobStore2(ctx, name).Execute()

Get the configuration for a Google Cloud blob store

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "name_example" // string | the name of the blob store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.GetBlobStore2(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetBlobStore2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobStore2`: GoogleCloudBlobstoreApiModel
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetBlobStore2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | the name of the blob store | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobStore2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GoogleCloudBlobstoreApiModel**](GoogleCloudBlobstoreApiModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileBlobStoreConfiguration

> FileBlobStoreApiModel GetFileBlobStoreConfiguration(ctx, name).Execute()

Get a file blob store configuration by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "default" // string | The name of the file blob store to read

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.GetFileBlobStoreConfiguration(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetFileBlobStoreConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileBlobStoreConfiguration`: FileBlobStoreApiModel
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetFileBlobStoreConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the file blob store to read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileBlobStoreConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileBlobStoreApiModel**](FileBlobStoreApiModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupBlobStoreConfiguration

> GroupBlobStoreApiModel GetGroupBlobStoreConfiguration(ctx, name).Execute()

Get a group blob store configuration by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "name_example" // string | The name of the group blob store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.GetGroupBlobStoreConfiguration(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetGroupBlobStoreConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupBlobStoreConfiguration`: GroupBlobStoreApiModel
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetGroupBlobStoreConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the group blob store | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupBlobStoreConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupBlobStoreApiModel**](GroupBlobStoreApiModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegionsByProjectId

> []string GetRegionsByProjectId(ctx, projectId).Execute()

Get the project regions by project's id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	projectId := "projectId_example" // string | projectId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.GetRegionsByProjectId(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetRegionsByProjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegionsByProjectId`: []string
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetRegionsByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionsByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlobStores

> []GenericBlobStoreApiResponse ListBlobStores(ctx).Execute()

List the blob stores

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.ListBlobStores(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.ListBlobStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlobStores`: []GenericBlobStoreApiResponse
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.ListBlobStores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBlobStoresRequest struct via the builder pattern


### Return type

[**[]GenericBlobStoreApiResponse**](GenericBlobStoreApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotaStatus

> BlobStoreQuotaResultXO QuotaStatus(ctx, name).Execute()

Get quota status for a given blob store

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.QuotaStatus(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.QuotaStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuotaStatus`: BlobStoreQuotaResultXO
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.QuotaStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuotaStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlobStoreQuotaResultXO**](BlobStoreQuotaResultXO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlobStore

> UpdateBlobStore(ctx, name).Body(body).Execute()

Update an S3 blob store configuration by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "name_example" // string | Name of the blob store to update
	body := *openapiclient.NewS3BlobStoreApiModel("s3", *openapiclient.NewS3BlobStoreApiBucketConfiguration(*openapiclient.NewS3BlobStoreApiBucket("DEFAULT", "Name_example", int32(3)))) // S3BlobStoreApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.UpdateBlobStore(context.Background(), name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.UpdateBlobStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the blob store to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlobStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**S3BlobStoreApiModel**](S3BlobStoreApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlobStore1

> UpdateBlobStore1(ctx, name).Body(body).Execute()

Update an Azure blob store configuration by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "name_example" // string | Name of the blob store to update
	body := *openapiclient.NewAzureBlobStoreApiModel("Name_example", *openapiclient.NewAzureBlobStoreApiBucketConfiguration("AccountName_example", "ContainerName_example", *openapiclient.NewAzureBlobStoreApiAuthentication("AuthenticationMethod_example"))) // AzureBlobStoreApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.UpdateBlobStore1(context.Background(), name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.UpdateBlobStore1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the blob store to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlobStore1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AzureBlobStoreApiModel**](AzureBlobStoreApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlobStore2

> UpdateBlobStore2(ctx, name).Body(body).Execute()

Update a Google Cloud blob store

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "name_example" // string | the name of the blobstore
	body := *openapiclient.NewGoogleCloudBlobstoreApiModel("gc_storage", *openapiclient.NewGoogleCloudBlobStoreApiBucketConfiguration(*openapiclient.NewGoogleCloudBlobStoreApiBucket("Name_example", "us-central1"))) // GoogleCloudBlobstoreApiModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.UpdateBlobStore2(context.Background(), name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.UpdateBlobStore2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | the name of the blobstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlobStore2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GoogleCloudBlobstoreApiModel**](GoogleCloudBlobstoreApiModel.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFileBlobStore

> UpdateFileBlobStore(ctx, name).Body(body).Execute()

Update a file blob store configuration by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "name_example" // string | The name of the file blob store to update
	body := *openapiclient.NewFileBlobStoreApiUpdateRequest() // FileBlobStoreApiUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.UpdateFileBlobStore(context.Background(), name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.UpdateFileBlobStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the file blob store to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileBlobStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FileBlobStoreApiUpdateRequest**](FileBlobStoreApiUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupBlobStore

> UpdateGroupBlobStore(ctx, name).Body(body).Execute()

Update a group blob store configuration by name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sma-de/go-nexus"
)

func main() {
	name := "name_example" // string | The name of the blob store to update
	body := *openapiclient.NewGroupBlobStoreApiUpdateRequest() // GroupBlobStoreApiUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.UpdateGroupBlobStore(context.Background(), name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.UpdateGroupBlobStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the blob store to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupBlobStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GroupBlobStoreApiUpdateRequest**](GroupBlobStoreApiUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

