# nexus\MaliciousRiskVisualizerAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMaliciousRiskSummary**](MaliciousRiskVisualizerAPI.md#GetMaliciousRiskSummary) | **Get** /v1/malicious-risk/summary | Get Dashboard Data Summary



## GetMaliciousRiskSummary

> GetMaliciousRiskSummary(ctx).Execute()

Get Dashboard Data Summary

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
	r, err := apiClient.MaliciousRiskVisualizerAPI.GetMaliciousRiskSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaliciousRiskVisualizerAPI.GetMaliciousRiskSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaliciousRiskSummaryRequest struct via the builder pattern


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

