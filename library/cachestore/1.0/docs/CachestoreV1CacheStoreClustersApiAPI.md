# \CachestoreV1CacheStoreClustersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CachestoreCreateCluster**](CachestoreV1CacheStoreClustersApiAPI.md#CachestoreCreateCluster) | **Post** /v1/clusters | Create Cluster
[**CachestoreCreateRestore**](CachestoreV1CacheStoreClustersApiAPI.md#CachestoreCreateRestore) | **Post** /v1/clusters/{cluster_id}/restore | Create Restore Cluster
[**CachestoreDeleteCluster**](CachestoreV1CacheStoreClustersApiAPI.md#CachestoreDeleteCluster) | **Delete** /v1/clusters/{cluster_id} | Terminate Cluster
[**CachestoreListClusters**](CachestoreV1CacheStoreClustersApiAPI.md#CachestoreListClusters) | **Get** /v1/clusters | List Clusters
[**CachestoreRestartCluster**](CachestoreV1CacheStoreClustersApiAPI.md#CachestoreRestartCluster) | **Post** /v1/clusters/{cluster_id}/restart | Restart Cluster
[**CachestoreShowCluster**](CachestoreV1CacheStoreClustersApiAPI.md#CachestoreShowCluster) | **Get** /v1/clusters/{cluster_id} | Show Cluster
[**CachestoreStartCluster**](CachestoreV1CacheStoreClustersApiAPI.md#CachestoreStartCluster) | **Post** /v1/clusters/{cluster_id}/start | Start Cluster
[**CachestoreStopCluster**](CachestoreV1CacheStoreClustersApiAPI.md#CachestoreStopCluster) | **Post** /v1/clusters/{cluster_id}/stop | Stop Cluster
[**CachestoreSwitchoverCluster**](CachestoreV1CacheStoreClustersApiAPI.md#CachestoreSwitchoverCluster) | **Post** /v1/clusters/{cluster_id}/switchover | Switchover Cluster
[**CachestoreSyncClusterState**](CachestoreV1CacheStoreClustersApiAPI.md#CachestoreSyncClusterState) | **Post** /v1/clusters/{cluster_id}/sync-state | Synchronize Cluster State



## CachestoreCreateCluster

> AsyncResponse CachestoreCreateCluster(ctx).RedisClusterCreateRequest(redisClusterCreateRequest).Execute()

Create Cluster



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
	redisClusterCreateRequest := *openapiclient.NewRedisClusterCreateRequest("DbaasEngineVersionId_example", *openapiclient.NewRedisInitConfigOption("DatabaseUserPassword_example"), []openapiclient.RedisInstanceGroupRequest{*openapiclient.NewRedisInstanceGroupRequest([]openapiclient.RedisBlockStorageGroupRequest{*openapiclient.NewRedisBlockStorageGroupRequest(openapiclient.BlockStorageGroupRoleType("OS"), int32(123))}, []openapiclient.InstanceRequest{*openapiclient.NewInstanceRequest(openapiclient.InstanceRoleType("ACTIVE"))}, openapiclient.InstanceGroupRoleType("ACTIVE"), "ServerTypeName_example")}, "InstanceNamePrefix_example", "Name_example", "SubnetId_example", "Timezone_example") // RedisClusterCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreClustersApiAPI.CachestoreCreateCluster(context.Background()).RedisClusterCreateRequest(redisClusterCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreClustersApiAPI.CachestoreCreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreCreateCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreClustersApiAPI.CachestoreCreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **redisClusterCreateRequest** | [**RedisClusterCreateRequest**](RedisClusterCreateRequest.md) |  | 

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


## CachestoreCreateRestore

> AsyncResponse CachestoreCreateRestore(ctx, clusterId).ClusterRestoreRequest(clusterRestoreRequest).Execute()

Create Restore Cluster



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
	clusterRestoreRequest := *openapiclient.NewClusterRestoreRequest("InstanceNamePrefix_example", "Name_example", "ServerTypeName_example") // ClusterRestoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreClustersApiAPI.CachestoreCreateRestore(context.Background(), clusterId).ClusterRestoreRequest(clusterRestoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreClustersApiAPI.CachestoreCreateRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreCreateRestore`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreClustersApiAPI.CachestoreCreateRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreCreateRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterRestoreRequest** | [**ClusterRestoreRequest**](ClusterRestoreRequest.md) |  | 

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


## CachestoreDeleteCluster

> AsyncResponse CachestoreDeleteCluster(ctx, clusterId).Execute()

Terminate Cluster



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
	resp, r, err := apiClient.CachestoreV1CacheStoreClustersApiAPI.CachestoreDeleteCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreClustersApiAPI.CachestoreDeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreDeleteCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreClustersApiAPI.CachestoreDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreDeleteClusterRequest struct via the builder pattern


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


## CachestoreListClusters

> ClusterPageResponse CachestoreListClusters(ctx).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).Execute()

List Clusters



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
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	name := "name_example" // string | Cluster name (optional)
	serviceState := "serviceState_example" // string | Service state (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreClustersApiAPI.CachestoreListClusters(context.Background()).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreClustersApiAPI.CachestoreListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreListClusters`: ClusterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreClustersApiAPI.CachestoreListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Cluster name | 
 **serviceState** | **string** | Service state | 

### Return type

[**ClusterPageResponse**](ClusterPageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CachestoreRestartCluster

> AsyncResponse CachestoreRestartCluster(ctx, clusterId).Execute()

Restart Cluster



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
	resp, r, err := apiClient.CachestoreV1CacheStoreClustersApiAPI.CachestoreRestartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreClustersApiAPI.CachestoreRestartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreRestartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreClustersApiAPI.CachestoreRestartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreRestartClusterRequest struct via the builder pattern


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


## CachestoreShowCluster

> RedisClusterDetailResponse CachestoreShowCluster(ctx, clusterId).Execute()

Show Cluster



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
	resp, r, err := apiClient.CachestoreV1CacheStoreClustersApiAPI.CachestoreShowCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreClustersApiAPI.CachestoreShowCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreShowCluster`: RedisClusterDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreClustersApiAPI.CachestoreShowCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreShowClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RedisClusterDetailResponse**](RedisClusterDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CachestoreStartCluster

> AsyncResponse CachestoreStartCluster(ctx, clusterId).Execute()

Start Cluster



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
	resp, r, err := apiClient.CachestoreV1CacheStoreClustersApiAPI.CachestoreStartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreClustersApiAPI.CachestoreStartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreStartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreClustersApiAPI.CachestoreStartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreStartClusterRequest struct via the builder pattern


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


## CachestoreStopCluster

> AsyncResponse CachestoreStopCluster(ctx, clusterId).Execute()

Stop Cluster



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
	resp, r, err := apiClient.CachestoreV1CacheStoreClustersApiAPI.CachestoreStopCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreClustersApiAPI.CachestoreStopCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreStopCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreClustersApiAPI.CachestoreStopCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreStopClusterRequest struct via the builder pattern


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


## CachestoreSwitchoverCluster

> AsyncResponse CachestoreSwitchoverCluster(ctx, clusterId).RedisSwitchOverTargetRequest(redisSwitchOverTargetRequest).Execute()

Switchover Cluster



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
	redisSwitchOverTargetRequest := *openapiclient.NewRedisSwitchOverTargetRequest("SwitchHostName_example") // RedisSwitchOverTargetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreClustersApiAPI.CachestoreSwitchoverCluster(context.Background(), clusterId).RedisSwitchOverTargetRequest(redisSwitchOverTargetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreClustersApiAPI.CachestoreSwitchoverCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreSwitchoverCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreClustersApiAPI.CachestoreSwitchoverCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreSwitchoverClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redisSwitchOverTargetRequest** | [**RedisSwitchOverTargetRequest**](RedisSwitchOverTargetRequest.md) |  | 

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


## CachestoreSyncClusterState

> AsyncResponse CachestoreSyncClusterState(ctx, clusterId).Execute()

Synchronize Cluster State



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
	resp, r, err := apiClient.CachestoreV1CacheStoreClustersApiAPI.CachestoreSyncClusterState(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreClustersApiAPI.CachestoreSyncClusterState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreSyncClusterState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreClustersApiAPI.CachestoreSyncClusterState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreSyncClusterStateRequest struct via the builder pattern


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

