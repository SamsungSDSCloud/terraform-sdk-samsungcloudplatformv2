# \EventstreamsV1EventStreamsInstancesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventstreamsSetBlockStorageSize**](EventstreamsV1EventStreamsInstancesApiAPI.md#EventstreamsSetBlockStorageSize) | **Post** /v1/block-storage-groups/{block_storage_group_id}/resize | Set Block Storage Size
[**EventstreamsSetSecurityGroupRules**](EventstreamsV1EventStreamsInstancesApiAPI.md#EventstreamsSetSecurityGroupRules) | **Post** /v1/clusters/{cluster_id}/security-group-rules | Set Security Group Rules
[**EventstreamsSetServerType**](EventstreamsV1EventStreamsInstancesApiAPI.md#EventstreamsSetServerType) | **Post** /v1/instance-groups/{instance_group_id}/resize | Set Server Type



## EventstreamsSetBlockStorageSize

> AsyncResponse EventstreamsSetBlockStorageSize(ctx, blockStorageGroupId).ResizeBlockStorageGroupRequest(resizeBlockStorageGroupRequest).Execute()

Set Block Storage Size



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
	blockStorageGroupId := "afc33891202643ba8ba1a5298a1affff" // string | Block storage group ID
	resizeBlockStorageGroupRequest := *openapiclient.NewResizeBlockStorageGroupRequest(int32(16)) // ResizeBlockStorageGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventstreamsV1EventStreamsInstancesApiAPI.EventstreamsSetBlockStorageSize(context.Background(), blockStorageGroupId).ResizeBlockStorageGroupRequest(resizeBlockStorageGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsInstancesApiAPI.EventstreamsSetBlockStorageSize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsSetBlockStorageSize`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsInstancesApiAPI.EventstreamsSetBlockStorageSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockStorageGroupId** | **string** | Block storage group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsSetBlockStorageSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resizeBlockStorageGroupRequest** | [**ResizeBlockStorageGroupRequest**](ResizeBlockStorageGroupRequest.md) |  | 

### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventstreamsSetSecurityGroupRules

> AsyncResponse EventstreamsSetSecurityGroupRules(ctx, clusterId).UpdateSecurityGroupRulesRequest(updateSecurityGroupRulesRequest).Execute()

Set Security Group Rules



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
	clusterId := "clusterId_example" // string | Cluster ID
	updateSecurityGroupRulesRequest := *openapiclient.NewUpdateSecurityGroupRulesRequest() // UpdateSecurityGroupRulesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventstreamsV1EventStreamsInstancesApiAPI.EventstreamsSetSecurityGroupRules(context.Background(), clusterId).UpdateSecurityGroupRulesRequest(updateSecurityGroupRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsInstancesApiAPI.EventstreamsSetSecurityGroupRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsSetSecurityGroupRules`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsInstancesApiAPI.EventstreamsSetSecurityGroupRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsSetSecurityGroupRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSecurityGroupRulesRequest** | [**UpdateSecurityGroupRulesRequest**](UpdateSecurityGroupRulesRequest.md) |  | 

### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventstreamsSetServerType

> AsyncResponse EventstreamsSetServerType(ctx, instanceGroupId).InstanceGroupResizeRequest(instanceGroupResizeRequest).Execute()

Set Server Type



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
	instanceGroupId := "instanceGroupId_example" // string | Instance group ID
	instanceGroupResizeRequest := *openapiclient.NewInstanceGroupResizeRequest("ServerTypeName_example") // InstanceGroupResizeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventstreamsV1EventStreamsInstancesApiAPI.EventstreamsSetServerType(context.Background(), instanceGroupId).InstanceGroupResizeRequest(instanceGroupResizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsInstancesApiAPI.EventstreamsSetServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsSetServerType`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsInstancesApiAPI.EventstreamsSetServerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | Instance group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsSetServerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instanceGroupResizeRequest** | [**InstanceGroupResizeRequest**](InstanceGroupResizeRequest.md) |  | 

### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

