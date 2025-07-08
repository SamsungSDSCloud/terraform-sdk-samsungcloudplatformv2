# \PostgresqlV1PostgresqlInstancesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostgresqlAddBlockStorages**](PostgresqlV1PostgresqlInstancesApiAPI.md#PostgresqlAddBlockStorages) | **Post** /v1/instance-groups/{instance_group_id}/block-storage-groups | Add Block Storages
[**PostgresqlSetBlockStorageSize**](PostgresqlV1PostgresqlInstancesApiAPI.md#PostgresqlSetBlockStorageSize) | **Post** /v1/block-storage-groups/{block_storage_group_id}/resize | Set Block Storage Size
[**PostgresqlSetSecurityGroupRules**](PostgresqlV1PostgresqlInstancesApiAPI.md#PostgresqlSetSecurityGroupRules) | **Post** /v1/clusters/{cluster_id}/security-group-rules | Set Security Group Rules
[**PostgresqlSetServerType**](PostgresqlV1PostgresqlInstancesApiAPI.md#PostgresqlSetServerType) | **Post** /v1/instance-groups/{instance_group_id}/resize | Set Server Type



## PostgresqlAddBlockStorages

> AsyncResponse PostgresqlAddBlockStorages(ctx, instanceGroupId).AddBlockStoragesRequest(addBlockStoragesRequest).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlInstancesApiAPI.PostgresqlAddBlockStorages(context.Background(), instanceGroupId).AddBlockStoragesRequest(addBlockStoragesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlInstancesApiAPI.PostgresqlAddBlockStorages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlAddBlockStorages`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlInstancesApiAPI.PostgresqlAddBlockStorages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | Instance group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlAddBlockStoragesRequest struct via the builder pattern


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


## PostgresqlSetBlockStorageSize

> AsyncResponse PostgresqlSetBlockStorageSize(ctx, blockStorageGroupId).ResizeBlockStorageGroupRequest(resizeBlockStorageGroupRequest).Execute()

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
	blockStorageGroupId := "blockStorageGroupId_example" // string | Block storage group ID
	resizeBlockStorageGroupRequest := *openapiclient.NewResizeBlockStorageGroupRequest(int32(123)) // ResizeBlockStorageGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostgresqlV1PostgresqlInstancesApiAPI.PostgresqlSetBlockStorageSize(context.Background(), blockStorageGroupId).ResizeBlockStorageGroupRequest(resizeBlockStorageGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlInstancesApiAPI.PostgresqlSetBlockStorageSize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlSetBlockStorageSize`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlInstancesApiAPI.PostgresqlSetBlockStorageSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockStorageGroupId** | **string** | Block storage group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlSetBlockStorageSizeRequest struct via the builder pattern


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


## PostgresqlSetSecurityGroupRules

> AsyncResponse PostgresqlSetSecurityGroupRules(ctx, clusterId).UpdateSecurityGroupRulesRequest(updateSecurityGroupRulesRequest).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlInstancesApiAPI.PostgresqlSetSecurityGroupRules(context.Background(), clusterId).UpdateSecurityGroupRulesRequest(updateSecurityGroupRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlInstancesApiAPI.PostgresqlSetSecurityGroupRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlSetSecurityGroupRules`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlInstancesApiAPI.PostgresqlSetSecurityGroupRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlSetSecurityGroupRulesRequest struct via the builder pattern


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


## PostgresqlSetServerType

> AsyncResponse PostgresqlSetServerType(ctx, instanceGroupId).InstanceGroupResizeRequest(instanceGroupResizeRequest).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlInstancesApiAPI.PostgresqlSetServerType(context.Background(), instanceGroupId).InstanceGroupResizeRequest(instanceGroupResizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlInstancesApiAPI.PostgresqlSetServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlSetServerType`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlInstancesApiAPI.PostgresqlSetServerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | Instance group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlSetServerTypeRequest struct via the builder pattern


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

