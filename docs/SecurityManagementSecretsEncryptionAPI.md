# nexus\SecurityManagementSecretsEncryptionAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReEncrypt**](SecurityManagementSecretsEncryptionAPI.md#ReEncrypt) | **Put** /v1/secrets/encryption/re-encrypt | Re-encrypt secrets using the specified key



## ReEncrypt

> ReEncrypt(ctx).Body(body).Execute()

Re-encrypt secrets using the specified key



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
	body := *openapiclient.NewReEncryptionRequestApiXO() // ReEncryptionRequestApiXO |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementSecretsEncryptionAPI.ReEncrypt(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSecretsEncryptionAPI.ReEncrypt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReEncryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ReEncryptionRequestApiXO**](ReEncryptionRequestApiXO.md) |  | 

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

