# \EpasV1EpasReplicasApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EpasCreateReplica**](EpasV1EpasReplicasApiAPI.md#EpasCreateReplica) | **Post** /v1/clusters/{cluster_id}/replicas | Create Replica
[**EpasListReplicas**](EpasV1EpasReplicasApiAPI.md#EpasListReplicas) | **Get** /v1/clusters/{cluster_id}/replicas | List Replicas
[**EpasPromoteReplicaCluster**](EpasV1EpasReplicasApiAPI.md#EpasPromoteReplicaCluster) | **Post** /v1/clusters/{cluster_id}/promote | Promote Replica Cluster
[**EpasResetReplica**](EpasV1EpasReplicasApiAPI.md#EpasResetReplica) | **Post** /v1/clusters/{replica_cluster_id}/reset-replica | Reset Replica
[**EpasSyncReplicaState**](EpasV1EpasReplicasApiAPI.md#EpasSyncReplicaState) | **Post** /v1/clusters/{origin_cluster_id}/sync-replica-state | Synchronize Replica State



## EpasCreateReplica

> AsyncResponse EpasCreateReplica(ctx, clusterId).ReplicasCreateRequest(replicasCreateRequest).Execute()

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
	resp, r, err := apiClient.EpasV1EpasReplicasApiAPI.EpasCreateReplica(context.Background(), clusterId).ReplicasCreateRequest(replicasCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasReplicasApiAPI.EpasCreateReplica``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasCreateReplica`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasReplicasApiAPI.EpasCreateReplica`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasCreateReplicaRequest struct via the builder pattern


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


## EpasListReplicas

> ReplicasListResponse EpasListReplicas(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EpasV1EpasReplicasApiAPI.EpasListReplicas(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasReplicasApiAPI.EpasListReplicas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasListReplicas`: ReplicasListResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasReplicasApiAPI.EpasListReplicas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasListReplicasRequest struct via the builder pattern


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


## EpasPromoteReplicaCluster

> AsyncResponse EpasPromoteReplicaCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EpasV1EpasReplicasApiAPI.EpasPromoteReplicaCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasReplicasApiAPI.EpasPromoteReplicaCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasPromoteReplicaCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasReplicasApiAPI.EpasPromoteReplicaCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasPromoteReplicaClusterRequest struct via the builder pattern


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


## EpasResetReplica

> AsyncResponse EpasResetReplica(ctx, replicaClusterId).Execute()

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
	resp, r, err := apiClient.EpasV1EpasReplicasApiAPI.EpasResetReplica(context.Background(), replicaClusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasReplicasApiAPI.EpasResetReplica``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasResetReplica`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasReplicasApiAPI.EpasResetReplica`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaClusterId** | **string** | Replica cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasResetReplicaRequest struct via the builder pattern


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


## EpasSyncReplicaState

> AsyncResponse EpasSyncReplicaState(ctx, originClusterId).Execute()

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
	resp, r, err := apiClient.EpasV1EpasReplicasApiAPI.EpasSyncReplicaState(context.Background(), originClusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasReplicasApiAPI.EpasSyncReplicaState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSyncReplicaState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasReplicasApiAPI.EpasSyncReplicaState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originClusterId** | **string** | Origin cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSyncReplicaStateRequest struct via the builder pattern


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

