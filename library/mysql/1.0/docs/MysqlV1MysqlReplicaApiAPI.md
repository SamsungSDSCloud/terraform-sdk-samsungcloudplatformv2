# \MysqlV1MysqlReplicaApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlCreateReplica**](MysqlV1MysqlReplicaApiAPI.md#MysqlCreateReplica) | **Post** /v1/clusters/{cluster_id}/replicas | Create Replica
[**MysqlListReplicas**](MysqlV1MysqlReplicaApiAPI.md#MysqlListReplicas) | **Get** /v1/clusters/{cluster_id}/replicas | List Replicas
[**MysqlPromoteReplicaCluster**](MysqlV1MysqlReplicaApiAPI.md#MysqlPromoteReplicaCluster) | **Post** /v1/clusters/{cluster_id}/promote | Promote Replica Cluster
[**MysqlResetReplica**](MysqlV1MysqlReplicaApiAPI.md#MysqlResetReplica) | **Post** /v1/clusters/{replica_cluster_id}/reset-replica | Reset Replica
[**MysqlSyncReplicaState**](MysqlV1MysqlReplicaApiAPI.md#MysqlSyncReplicaState) | **Post** /v1/clusters/{origin_cluster_id}/sync-replica-state | Synchronize Replica State



## MysqlCreateReplica

> AsyncResponse MysqlCreateReplica(ctx, clusterId).ReplicasCreateRequest(replicasCreateRequest).Execute()

Create Replica



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
	replicasCreateRequest := *openapiclient.NewReplicasCreateRequest(int32(123), []openapiclient.Replicas{*openapiclient.NewReplicas([]openapiclient.InstanceGroupRequest{*openapiclient.NewInstanceGroupRequest([]openapiclient.BlockStorageGroupRequest{*openapiclient.NewBlockStorageGroupRequest(openapiclient.BlockStorageGroupRoleType("OS"), int32(123))}, openapiclient.InstanceGroupRoleType("ACTIVE"), "ServerTypeName_example")}, "Name_example")}) // ReplicasCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlReplicaApiAPI.MysqlCreateReplica(context.Background(), clusterId).ReplicasCreateRequest(replicasCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlReplicaApiAPI.MysqlCreateReplica``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlCreateReplica`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlReplicaApiAPI.MysqlCreateReplica`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlCreateReplicaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replicasCreateRequest** | [**ReplicasCreateRequest**](ReplicasCreateRequest.md) |  | 

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


## MysqlListReplicas

> ReplicasListResponse MysqlListReplicas(ctx, clusterId).Execute()

List Replicas



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlReplicaApiAPI.MysqlListReplicas(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlReplicaApiAPI.MysqlListReplicas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlListReplicas`: ReplicasListResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlReplicaApiAPI.MysqlListReplicas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlListReplicasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReplicasListResponse**](ReplicasListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlPromoteReplicaCluster

> AsyncResponse MysqlPromoteReplicaCluster(ctx, clusterId).Execute()

Promote Replica Cluster



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlReplicaApiAPI.MysqlPromoteReplicaCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlReplicaApiAPI.MysqlPromoteReplicaCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlPromoteReplicaCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlReplicaApiAPI.MysqlPromoteReplicaCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlPromoteReplicaClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlResetReplica

> AsyncResponse MysqlResetReplica(ctx, replicaClusterId).Execute()

Reset Replica



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
	replicaClusterId := "replicaClusterId_example" // string | Replica cluster id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlReplicaApiAPI.MysqlResetReplica(context.Background(), replicaClusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlReplicaApiAPI.MysqlResetReplica``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlResetReplica`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlReplicaApiAPI.MysqlResetReplica`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaClusterId** | **string** | Replica cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlResetReplicaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlSyncReplicaState

> AsyncResponse MysqlSyncReplicaState(ctx, originClusterId).Execute()

Synchronize Replica State



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
	originClusterId := "originClusterId_example" // string | Origin cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlReplicaApiAPI.MysqlSyncReplicaState(context.Background(), originClusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlReplicaApiAPI.MysqlSyncReplicaState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSyncReplicaState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlReplicaApiAPI.MysqlSyncReplicaState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originClusterId** | **string** | Origin cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSyncReplicaStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

