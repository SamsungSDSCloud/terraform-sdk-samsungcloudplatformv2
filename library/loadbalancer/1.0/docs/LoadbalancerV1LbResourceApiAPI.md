# \LoadbalancerV1LbResourceApiAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowResourceUsageLimit**](LoadbalancerV1LbResourceApiAPI.md#ShowResourceUsageLimit) | **Get** /v1/resources/{resource_type}/usage-limit | ShowResourceUsageLimit



## ShowResourceUsageLimit

> ResourceUsageLimitResponse ShowResourceUsageLimit(ctx, resourceType).Execute()

ShowResourceUsageLimit



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
	resourceType := openapiclient.ResourceUsageLimitType("LOADBALANCER") // ResourceUsageLimitType | resource type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LbResourceApiAPI.ShowResourceUsageLimit(context.Background(), resourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LbResourceApiAPI.ShowResourceUsageLimit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowResourceUsageLimit`: ResourceUsageLimitResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LbResourceApiAPI.ShowResourceUsageLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | [**ResourceUsageLimitType**](.md) | resource type | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowResourceUsageLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceUsageLimitResponse**](ResourceUsageLimitResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

