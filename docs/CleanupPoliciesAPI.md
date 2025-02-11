# nexus\CleanupPoliciesAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create2**](CleanupPoliciesAPI.md#Create2) | **Post** /v1/cleanup-policies | Create a new policy
[**DeletePolicyByName**](CleanupPoliciesAPI.md#DeletePolicyByName) | **Delete** /v1/cleanup-policies/{name} | Delete cleanup policy
[**GetAll**](CleanupPoliciesAPI.md#GetAll) | **Get** /v1/cleanup-policies | Get a list of existing policies
[**GetCleanupPolicyByName**](CleanupPoliciesAPI.md#GetCleanupPolicyByName) | **Get** /v1/cleanup-policies/{name} | Get a policy by name
[**Update2**](CleanupPoliciesAPI.md#Update2) | **Put** /v1/cleanup-policies/{policyName} | Update existing policy



## Create2

> Create2(ctx).Body(body).Execute()

Create a new policy

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
	body := *openapiclient.NewCleanupPolicyResourceXO("Name_example", "Format_example") // CleanupPolicyResourceXO |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CleanupPoliciesAPI.Create2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CleanupPoliciesAPI.Create2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CleanupPolicyResourceXO**](CleanupPolicyResourceXO.md) |  | 

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


## DeletePolicyByName

> DeletePolicyByName(ctx, name).Execute()

Delete cleanup policy

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
	r, err := apiClient.CleanupPoliciesAPI.DeletePolicyByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CleanupPoliciesAPI.DeletePolicyByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyByNameRequest struct via the builder pattern


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


## GetAll

> GetAll(ctx).Execute()

Get a list of existing policies

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
	r, err := apiClient.CleanupPoliciesAPI.GetAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CleanupPoliciesAPI.GetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


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


## GetCleanupPolicyByName

> GetCleanupPolicyByName(ctx, name).Execute()

Get a policy by name

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
	r, err := apiClient.CleanupPoliciesAPI.GetCleanupPolicyByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CleanupPoliciesAPI.GetCleanupPolicyByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCleanupPolicyByNameRequest struct via the builder pattern


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


## Update2

> Update2(ctx, policyName).Body(body).Execute()

Update existing policy

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
	policyName := "policyName_example" // string | 
	body := *openapiclient.NewCleanupPolicyResourceXO("Name_example", "Format_example") // CleanupPolicyResourceXO |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CleanupPoliciesAPI.Update2(context.Background(), policyName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CleanupPoliciesAPI.Update2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CleanupPolicyResourceXO**](CleanupPolicyResourceXO.md) |  | 

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

