# nexus\SecurityManagementLDAPAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeOrder**](SecurityManagementLDAPAPI.md#ChangeOrder) | **Post** /v1/security/ldap/change-order | Change LDAP server order
[**CreateLdapServer**](SecurityManagementLDAPAPI.md#CreateLdapServer) | **Post** /v1/security/ldap | Create LDAP server
[**DeleteLdapServer**](SecurityManagementLDAPAPI.md#DeleteLdapServer) | **Delete** /v1/security/ldap/{name} | Delete LDAP server
[**GetLdapServer**](SecurityManagementLDAPAPI.md#GetLdapServer) | **Get** /v1/security/ldap/{name} | Get LDAP server
[**GetLdapServers**](SecurityManagementLDAPAPI.md#GetLdapServers) | **Get** /v1/security/ldap | List LDAP servers
[**UpdateLdapServer**](SecurityManagementLDAPAPI.md#UpdateLdapServer) | **Put** /v1/security/ldap/{name} | Update LDAP server



## ChangeOrder

> ChangeOrder(ctx).Body(body).Execute()

Change LDAP server order

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
	body := []string{"Property_example"} // []string | Ordered list of LDAP server names (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.ChangeOrder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.ChangeOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | Ordered list of LDAP server names | 

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


## CreateLdapServer

> CreateLdapServer(ctx).Body(body).Execute()

Create LDAP server

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
	body := *openapiclient.NewCreateLdapServerXo("Name_example", "Protocol_example", "Host_example", int32(636), "dc=example,dc=com", "AuthScheme_example", int32(1), int32(123), int32(123), "GroupType_example", "AuthPassword_example") // CreateLdapServerXo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.CreateLdapServer(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.CreateLdapServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLdapServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateLdapServerXo**](CreateLdapServerXo.md) |  | 

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


## DeleteLdapServer

> DeleteLdapServer(ctx, name).Execute()

Delete LDAP server

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
	name := "name_example" // string | Name of the LDAP server to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.DeleteLdapServer(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.DeleteLdapServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP server to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLdapServerRequest struct via the builder pattern


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


## GetLdapServer

> GetLdapServer(ctx, name).Execute()

Get LDAP server

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
	name := "name_example" // string | Name of the LDAP server to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.GetLdapServer(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.GetLdapServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP server to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapServerRequest struct via the builder pattern


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


## GetLdapServers

> GetLdapServers(ctx).Execute()

List LDAP servers

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
	r, err := apiClient.SecurityManagementLDAPAPI.GetLdapServers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.GetLdapServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapServersRequest struct via the builder pattern


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


## UpdateLdapServer

> UpdateLdapServer(ctx, name).Body(body).Execute()

Update LDAP server

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
	name := "name_example" // string | Name of the LDAP server to update
	body := *openapiclient.NewUpdateLdapServerXo("Name_example", "Protocol_example", "Host_example", int32(636), "dc=example,dc=com", "AuthScheme_example", int32(1), int32(123), int32(123), "GroupType_example", "AuthPassword_example") // UpdateLdapServerXo | Updated values of LDAP server (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.UpdateLdapServer(context.Background(), name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.UpdateLdapServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP server to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateLdapServerXo**](UpdateLdapServerXo.md) | Updated values of LDAP server | 

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

