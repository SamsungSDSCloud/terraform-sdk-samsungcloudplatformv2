# \SearchengineV1SearchEngineInstancesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchengineAddBlockStorages**](SearchengineV1SearchEngineInstancesApiAPI.md#SearchengineAddBlockStorages) | **Post** /v1/instance-groups/{instance_group_id}/block-storage-groups | Add Block Storages
[**SearchengineSetBlockStorageSize**](SearchengineV1SearchEngineInstancesApiAPI.md#SearchengineSetBlockStorageSize) | **Post** /v1/block-storage-groups/{block_storage_group_id}/resize | Set Block Storage Size
[**SearchengineSetSecurityGroupRules**](SearchengineV1SearchEngineInstancesApiAPI.md#SearchengineSetSecurityGroupRules) | **Post** /v1/clusters/{cluster_id}/security-group-rules | Set Security Group Rules
[**SearchengineSetServerType**](SearchengineV1SearchEngineInstancesApiAPI.md#SearchengineSetServerType) | **Post** /v1/instance-groups/{instance_group_id}/resize | Set Server Type



## SearchengineAddBlockStorages

> AsyncResponse SearchengineAddBlockStorages(ctx, instanceGroupId).AddBlockStoragesRequest(addBlockStoragesRequest).Execute()

Add Block Storages



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
	addBlockStoragesRequest := *openapiclient.NewAddBlockStoragesRequest(openapiclient.BlockStorageGroupRoleType("OS"), int32(123)) // AddBlockStoragesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchengineV1SearchEngineInstancesApiAPI.SearchengineAddBlockStorages(context.Background(), instanceGroupId).AddBlockStoragesRequest(addBlockStoragesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineInstancesApiAPI.SearchengineAddBlockStorages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineAddBlockStorages`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineInstancesApiAPI.SearchengineAddBlockStorages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | Instance group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineAddBlockStoragesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addBlockStoragesRequest** | [**AddBlockStoragesRequest**](AddBlockStoragesRequest.md) |  | 

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


## SearchengineSetBlockStorageSize

> AsyncResponse SearchengineSetBlockStorageSize(ctx, blockStorageGroupId).ResizeBlockStorageGroupRequest(resizeBlockStorageGroupRequest).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineInstancesApiAPI.SearchengineSetBlockStorageSize(context.Background(), blockStorageGroupId).ResizeBlockStorageGroupRequest(resizeBlockStorageGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineInstancesApiAPI.SearchengineSetBlockStorageSize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineSetBlockStorageSize`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineInstancesApiAPI.SearchengineSetBlockStorageSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockStorageGroupId** | **string** | Block storage group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineSetBlockStorageSizeRequest struct via the builder pattern


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


## SearchengineSetSecurityGroupRules

> AsyncResponse SearchengineSetSecurityGroupRules(ctx, clusterId).UpdateSecurityGroupRulesRequest(updateSecurityGroupRulesRequest).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineInstancesApiAPI.SearchengineSetSecurityGroupRules(context.Background(), clusterId).UpdateSecurityGroupRulesRequest(updateSecurityGroupRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineInstancesApiAPI.SearchengineSetSecurityGroupRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineSetSecurityGroupRules`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineInstancesApiAPI.SearchengineSetSecurityGroupRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineSetSecurityGroupRulesRequest struct via the builder pattern


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


## SearchengineSetServerType

> AsyncResponse SearchengineSetServerType(ctx, instanceGroupId).InstanceGroupResizeRequest(instanceGroupResizeRequest).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineInstancesApiAPI.SearchengineSetServerType(context.Background(), instanceGroupId).InstanceGroupResizeRequest(instanceGroupResizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineInstancesApiAPI.SearchengineSetServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineSetServerType`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineInstancesApiAPI.SearchengineSetServerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | Instance group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineSetServerTypeRequest struct via the builder pattern


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

