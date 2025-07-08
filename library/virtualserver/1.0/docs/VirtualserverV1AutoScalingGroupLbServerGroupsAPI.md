# \VirtualserverV1AutoScalingGroupLbServerGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAutoScalingGroupLbServerGroups**](VirtualserverV1AutoScalingGroupLbServerGroupsAPI.md#ListAutoScalingGroupLbServerGroups) | **Get** /v1/auto-scaling-groups/{auto_scaling_group_id}/lb-server-groups | List Auto-Scaling Group LB Server Groups
[**UpdateAutoScalingGroupLbServerGroups**](VirtualserverV1AutoScalingGroupLbServerGroupsAPI.md#UpdateAutoScalingGroupLbServerGroups) | **Put** /v1/auto-scaling-groups/{auto_scaling_group_id}/lb-server-groups | Update Auto-Scaling Group LB Server Groups



## ListAutoScalingGroupLbServerGroups

> AutoScalingGroupLbServerGroupListResponse ListAutoScalingGroupLbServerGroups(ctx, autoScalingGroupId).Offset(offset).Limit(limit).Sort(sort).Execute()

List Auto-Scaling Group LB Server Groups



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
	offset := int32(0) // int32 | Offset (optional)
	limit := int32(20) // int32 | Limit (optional)
	sort := "created_at:desc" // string | Sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupLbServerGroupsAPI.ListAutoScalingGroupLbServerGroups(context.Background(), autoScalingGroupId).Offset(offset).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupLbServerGroupsAPI.ListAutoScalingGroupLbServerGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAutoScalingGroupLbServerGroups`: AutoScalingGroupLbServerGroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupLbServerGroupsAPI.ListAutoScalingGroupLbServerGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAutoScalingGroupLbServerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Offset | 
 **limit** | **int32** | Limit | 
 **sort** | **string** | Sort | 

### Return type

[**AutoScalingGroupLbServerGroupListResponse**](AutoScalingGroupLbServerGroupListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoScalingGroupLbServerGroups

> UpdateAutoScalingGroupLbServerGroups(ctx, autoScalingGroupId).AutoScalingGroupLbServerGroupSetRequest(autoScalingGroupLbServerGroupSetRequest).Execute()

Update Auto-Scaling Group LB Server Groups



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
	autoScalingGroupLbServerGroupSetRequest := *openapiclient.NewAutoScalingGroupLbServerGroupSetRequest([]openapiclient.AutoScalingGroupLbServerGroup{*openapiclient.NewAutoScalingGroupLbServerGroup("d06e87d3-ca9a-461b-8e88-077a542a7335", int32(8080))}) // AutoScalingGroupLbServerGroupSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1AutoScalingGroupLbServerGroupsAPI.UpdateAutoScalingGroupLbServerGroups(context.Background(), autoScalingGroupId).AutoScalingGroupLbServerGroupSetRequest(autoScalingGroupLbServerGroupSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupLbServerGroupsAPI.UpdateAutoScalingGroupLbServerGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoScalingGroupLbServerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoScalingGroupLbServerGroupSetRequest** | [**AutoScalingGroupLbServerGroupSetRequest**](AutoScalingGroupLbServerGroupSetRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

