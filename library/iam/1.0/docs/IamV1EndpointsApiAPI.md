# \IamV1EndpointsApiAPI

All URIs are relative to *https://scp-iam-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEndpoints**](IamV1EndpointsApiAPI.md#ListEndpoints) | **Get** /v1/endpoints | Get endpoint list



## ListEndpoints

> ListEndpointsResponse ListEndpoints(ctx).Execute()

Get endpoint list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1EndpointsApiAPI.ListEndpoints(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1EndpointsApiAPI.ListEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEndpoints`: ListEndpointsResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1EndpointsApiAPI.ListEndpoints`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsRequest struct via the builder pattern


### Return type

[**ListEndpointsResponse**](ListEndpointsResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

