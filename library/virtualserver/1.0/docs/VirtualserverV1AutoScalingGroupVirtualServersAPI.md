# \VirtualserverV1AutoScalingGroupVirtualServersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAutoScalingGroupVirtualServers**](VirtualserverV1AutoScalingGroupVirtualServersAPI.md#ListAutoScalingGroupVirtualServers) | **Get** /v1/auto-scaling-groups/{auto_scaling_group_id}/virtual-servers | List Auto-Scaling Group Virtual Servers



## ListAutoScalingGroupVirtualServers

> AutoScalingGroupVirtualServerListResponse ListAutoScalingGroupVirtualServers(ctx, autoScalingGroupId).Marker(marker).Limit(limit).Sort(sort).Execute()

List Auto-Scaling Group Virtual Servers



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
	autoScalingGroupId := "52613bd852b04b39adcb15a8364d856d" // string | Auto-Scaling Group ID
	marker := "2a9be312-5d4b-4bc8-b2ae-35100fa9241f" // string | Marker (optional)
	limit := int32(20) // int32 | Limit (optional)
	sort := "created_at:desc" // string | Sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupVirtualServersAPI.ListAutoScalingGroupVirtualServers(context.Background(), autoScalingGroupId).Marker(marker).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupVirtualServersAPI.ListAutoScalingGroupVirtualServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAutoScalingGroupVirtualServers`: AutoScalingGroupVirtualServerListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupVirtualServersAPI.ListAutoScalingGroupVirtualServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAutoScalingGroupVirtualServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marker** | **string** | Marker | 
 **limit** | **int32** | Limit | 
 **sort** | **string** | Sort | 

### Return type

[**AutoScalingGroupVirtualServerListResponse**](AutoScalingGroupVirtualServerListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

