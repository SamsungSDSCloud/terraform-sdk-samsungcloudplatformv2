# \MariadbV1MariadbReplicasApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariadbCreateReplica**](MariadbV1MariadbReplicasApiAPI.md#MariadbCreateReplica) | **Post** /v1/clusters/{cluster_id}/replicas | Create Replica
[**MariadbListReplicas**](MariadbV1MariadbReplicasApiAPI.md#MariadbListReplicas) | **Get** /v1/clusters/{cluster_id}/replicas | List Replicas
[**MariadbPromoteReplicaCluster**](MariadbV1MariadbReplicasApiAPI.md#MariadbPromoteReplicaCluster) | **Post** /v1/clusters/{cluster_id}/promote | Promote Replica Cluster
[**MariadbResetReplica**](MariadbV1MariadbReplicasApiAPI.md#MariadbResetReplica) | **Post** /v1/clusters/{replica_cluster_id}/reset-replica | Reset Replica
[**MariadbSyncReplicaState**](MariadbV1MariadbReplicasApiAPI.md#MariadbSyncReplicaState) | **Post** /v1/clusters/{origin_cluster_id}/sync-replica-state | Synchronize Replica State



## MariadbCreateReplica

> AsyncResponse MariadbCreateReplica(ctx, clusterId).ReplicasCreateRequest(replicasCreateRequest).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbReplicasApiAPI.MariadbCreateReplica(context.Background(), clusterId).ReplicasCreateRequest(replicasCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbReplicasApiAPI.MariadbCreateReplica``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbCreateReplica`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbReplicasApiAPI.MariadbCreateReplica`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbCreateReplicaRequest struct via the builder pattern


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


## MariadbListReplicas

> ReplicasListResponse MariadbListReplicas(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbReplicasApiAPI.MariadbListReplicas(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbReplicasApiAPI.MariadbListReplicas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbListReplicas`: ReplicasListResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbReplicasApiAPI.MariadbListReplicas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbListReplicasRequest struct via the builder pattern


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


## MariadbPromoteReplicaCluster

> AsyncResponse MariadbPromoteReplicaCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbReplicasApiAPI.MariadbPromoteReplicaCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbReplicasApiAPI.MariadbPromoteReplicaCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbPromoteReplicaCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbReplicasApiAPI.MariadbPromoteReplicaCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbPromoteReplicaClusterRequest struct via the builder pattern


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


## MariadbResetReplica

> AsyncResponse MariadbResetReplica(ctx, replicaClusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbReplicasApiAPI.MariadbResetReplica(context.Background(), replicaClusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbReplicasApiAPI.MariadbResetReplica``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbResetReplica`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbReplicasApiAPI.MariadbResetReplica`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaClusterId** | **string** | Replica cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbResetReplicaRequest struct via the builder pattern


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


## MariadbSyncReplicaState

> AsyncResponse MariadbSyncReplicaState(ctx, originClusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbReplicasApiAPI.MariadbSyncReplicaState(context.Background(), originClusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbReplicasApiAPI.MariadbSyncReplicaState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbSyncReplicaState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbReplicasApiAPI.MariadbSyncReplicaState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originClusterId** | **string** | Origin cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSyncReplicaStateRequest struct via the builder pattern


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

