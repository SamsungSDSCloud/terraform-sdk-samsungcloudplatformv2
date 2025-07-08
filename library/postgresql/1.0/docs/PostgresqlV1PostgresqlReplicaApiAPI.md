# \PostgresqlV1PostgresqlReplicaApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostgresqlCreateReplica**](PostgresqlV1PostgresqlReplicaApiAPI.md#PostgresqlCreateReplica) | **Post** /v1/clusters/{cluster_id}/replicas | Create Replica
[**PostgresqlListReplicas**](PostgresqlV1PostgresqlReplicaApiAPI.md#PostgresqlListReplicas) | **Get** /v1/clusters/{cluster_id}/replicas | List Replicas
[**PostgresqlPromoteReplicaCluster**](PostgresqlV1PostgresqlReplicaApiAPI.md#PostgresqlPromoteReplicaCluster) | **Post** /v1/clusters/{cluster_id}/promote | Promote Replica Cluster
[**PostgresqlResetReplica**](PostgresqlV1PostgresqlReplicaApiAPI.md#PostgresqlResetReplica) | **Post** /v1/clusters/{replica_cluster_id}/reset-replica | Reset Replica
[**PostgresqlSyncReplicaState**](PostgresqlV1PostgresqlReplicaApiAPI.md#PostgresqlSyncReplicaState) | **Post** /v1/clusters/{origin_cluster_id}/sync-replica-state | Synchronize Replica State



## PostgresqlCreateReplica

> AsyncResponse PostgresqlCreateReplica(ctx, clusterId).ReplicasCreateRequest(replicasCreateRequest).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlCreateReplica(context.Background(), clusterId).ReplicasCreateRequest(replicasCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlCreateReplica``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlCreateReplica`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlCreateReplica`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlCreateReplicaRequest struct via the builder pattern


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


## PostgresqlListReplicas

> ReplicasListResponse PostgresqlListReplicas(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlListReplicas(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlListReplicas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlListReplicas`: ReplicasListResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlListReplicas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlListReplicasRequest struct via the builder pattern


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


## PostgresqlPromoteReplicaCluster

> AsyncResponse PostgresqlPromoteReplicaCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlPromoteReplicaCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlPromoteReplicaCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlPromoteReplicaCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlPromoteReplicaCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlPromoteReplicaClusterRequest struct via the builder pattern


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


## PostgresqlResetReplica

> AsyncResponse PostgresqlResetReplica(ctx, replicaClusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlResetReplica(context.Background(), replicaClusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlResetReplica``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlResetReplica`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlResetReplica`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replicaClusterId** | **string** | Replica cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlResetReplicaRequest struct via the builder pattern


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


## PostgresqlSyncReplicaState

> AsyncResponse PostgresqlSyncReplicaState(ctx, originClusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlSyncReplicaState(context.Background(), originClusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlSyncReplicaState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlSyncReplicaState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlReplicaApiAPI.PostgresqlSyncReplicaState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originClusterId** | **string** | Origin cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlSyncReplicaStateRequest struct via the builder pattern


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

