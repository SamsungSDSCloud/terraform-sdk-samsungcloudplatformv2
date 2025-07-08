# \MysqlV1MysqlInstancesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlAddBlockStorages**](MysqlV1MysqlInstancesApiAPI.md#MysqlAddBlockStorages) | **Post** /v1/instance-groups/{instance_group_id}/block-storage-groups | Add Block Storages
[**MysqlSetBlockStorageSize**](MysqlV1MysqlInstancesApiAPI.md#MysqlSetBlockStorageSize) | **Post** /v1/block-storage-groups/{block_storage_group_id}/resize | Set Block Storage Size
[**MysqlSetSecurityGroupRules**](MysqlV1MysqlInstancesApiAPI.md#MysqlSetSecurityGroupRules) | **Post** /v1/clusters/{cluster_id}/security-group-rules | Set Security Group Rules
[**MysqlSetServerType**](MysqlV1MysqlInstancesApiAPI.md#MysqlSetServerType) | **Post** /v1/instance-groups/{instance_group_id}/resize | Set Server Type



## MysqlAddBlockStorages

> AsyncResponse MysqlAddBlockStorages(ctx, instanceGroupId).AddBlockStoragesRequest(addBlockStoragesRequest).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlInstancesApiAPI.MysqlAddBlockStorages(context.Background(), instanceGroupId).AddBlockStoragesRequest(addBlockStoragesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlInstancesApiAPI.MysqlAddBlockStorages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlAddBlockStorages`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlInstancesApiAPI.MysqlAddBlockStorages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | Instance group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlAddBlockStoragesRequest struct via the builder pattern


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


## MysqlSetBlockStorageSize

> AsyncResponse MysqlSetBlockStorageSize(ctx, blockStorageGroupId).ResizeBlockStorageGroupRequest(resizeBlockStorageGroupRequest).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlInstancesApiAPI.MysqlSetBlockStorageSize(context.Background(), blockStorageGroupId).ResizeBlockStorageGroupRequest(resizeBlockStorageGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlInstancesApiAPI.MysqlSetBlockStorageSize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSetBlockStorageSize`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlInstancesApiAPI.MysqlSetBlockStorageSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockStorageGroupId** | **string** | Block storage group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSetBlockStorageSizeRequest struct via the builder pattern


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


## MysqlSetSecurityGroupRules

> AsyncResponse MysqlSetSecurityGroupRules(ctx, clusterId).UpdateSecurityGroupRulesRequest(updateSecurityGroupRulesRequest).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlInstancesApiAPI.MysqlSetSecurityGroupRules(context.Background(), clusterId).UpdateSecurityGroupRulesRequest(updateSecurityGroupRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlInstancesApiAPI.MysqlSetSecurityGroupRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSetSecurityGroupRules`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlInstancesApiAPI.MysqlSetSecurityGroupRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSetSecurityGroupRulesRequest struct via the builder pattern


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


## MysqlSetServerType

> AsyncResponse MysqlSetServerType(ctx, instanceGroupId).InstanceGroupResizeRequest(instanceGroupResizeRequest).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlInstancesApiAPI.MysqlSetServerType(context.Background(), instanceGroupId).InstanceGroupResizeRequest(instanceGroupResizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlInstancesApiAPI.MysqlSetServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSetServerType`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlInstancesApiAPI.MysqlSetServerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | Instance group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSetServerTypeRequest struct via the builder pattern


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

