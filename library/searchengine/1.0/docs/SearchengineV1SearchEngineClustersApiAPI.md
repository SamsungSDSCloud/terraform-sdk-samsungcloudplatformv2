# \SearchengineV1SearchEngineClustersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchengineAddInstances**](SearchengineV1SearchEngineClustersApiAPI.md#SearchengineAddInstances) | **Post** /v1/clusters/{cluster_id}/add-instances | Add Instances
[**SearchengineCreateCluster**](SearchengineV1SearchEngineClustersApiAPI.md#SearchengineCreateCluster) | **Post** /v1/clusters | Create Cluster
[**SearchengineCreateRestore**](SearchengineV1SearchEngineClustersApiAPI.md#SearchengineCreateRestore) | **Post** /v1/clusters/{cluster_id}/restore | Create Restore Cluster
[**SearchengineDeleteCluster**](SearchengineV1SearchEngineClustersApiAPI.md#SearchengineDeleteCluster) | **Delete** /v1/clusters/{cluster_id} | Terminate Cluster
[**SearchengineListClusters**](SearchengineV1SearchEngineClustersApiAPI.md#SearchengineListClusters) | **Get** /v1/clusters | List Clusters
[**SearchengineRestartCluster**](SearchengineV1SearchEngineClustersApiAPI.md#SearchengineRestartCluster) | **Post** /v1/clusters/{cluster_id}/restart | Restart Cluster
[**SearchengineShowCluster**](SearchengineV1SearchEngineClustersApiAPI.md#SearchengineShowCluster) | **Get** /v1/clusters/{cluster_id} | Show Cluster
[**SearchengineStartCluster**](SearchengineV1SearchEngineClustersApiAPI.md#SearchengineStartCluster) | **Post** /v1/clusters/{cluster_id}/start | Start Cluster
[**SearchengineStopCluster**](SearchengineV1SearchEngineClustersApiAPI.md#SearchengineStopCluster) | **Post** /v1/clusters/{cluster_id}/stop | Stop Cluster
[**SearchengineSyncClusterState**](SearchengineV1SearchEngineClustersApiAPI.md#SearchengineSyncClusterState) | **Post** /v1/clusters/{cluster_id}/sync-state | Synchronize Cluster State



## SearchengineAddInstances

> AsyncResponse SearchengineAddInstances(ctx, clusterId).SearchEngineClusterAddInstancesRequest(searchEngineClusterAddInstancesRequest).Execute()

Add Instances



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
	searchEngineClusterAddInstancesRequest := *openapiclient.NewSearchEngineClusterAddInstancesRequest(int32(123), []string{"ServiceIpAddresses_example"}) // SearchEngineClusterAddInstancesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchengineV1SearchEngineClustersApiAPI.SearchengineAddInstances(context.Background(), clusterId).SearchEngineClusterAddInstancesRequest(searchEngineClusterAddInstancesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineClustersApiAPI.SearchengineAddInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineAddInstances`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineClustersApiAPI.SearchengineAddInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineAddInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchEngineClusterAddInstancesRequest** | [**SearchEngineClusterAddInstancesRequest**](SearchEngineClusterAddInstancesRequest.md) |  | 

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


## SearchengineCreateCluster

> AsyncResponse SearchengineCreateCluster(ctx).SearchEngineClusterCreateRequest(searchEngineClusterCreateRequest).Execute()

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
	searchEngineClusterCreateRequest := *openapiclient.NewSearchEngineClusterCreateRequest("DbaasEngineVersionId_example", *openapiclient.NewSearchEngineInitConfigOptionRequest("DatabaseUserName_example", "DatabaseUserPassword_example"), []openapiclient.InstanceGroupRequest{*openapiclient.NewInstanceGroupRequest([]openapiclient.BlockStorageGroupRequest{*openapiclient.NewBlockStorageGroupRequest(openapiclient.BlockStorageGroupRoleType("OS"), int32(123))}, openapiclient.InstanceGroupRoleType("ACTIVE"), "ServerTypeName_example")}, "InstanceNamePrefix_example", "Name_example", "SubnetId_example", "Timezone_example") // SearchEngineClusterCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchengineV1SearchEngineClustersApiAPI.SearchengineCreateCluster(context.Background()).SearchEngineClusterCreateRequest(searchEngineClusterCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineClustersApiAPI.SearchengineCreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineCreateCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineClustersApiAPI.SearchengineCreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchEngineClusterCreateRequest** | [**SearchEngineClusterCreateRequest**](SearchEngineClusterCreateRequest.md) |  | 

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


## SearchengineCreateRestore

> AsyncResponse SearchengineCreateRestore(ctx, clusterId).SearchEngineClusterRestoreRequest(searchEngineClusterRestoreRequest).Execute()

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
	searchEngineClusterRestoreRequest := *openapiclient.NewSearchEngineClusterRestoreRequest("BackupHistoryNumber_example", []openapiclient.SearchEngineRestoreInstanceGroup{*openapiclient.NewSearchEngineRestoreInstanceGroup([]openapiclient.BlockStorageGroupRequest{*openapiclient.NewBlockStorageGroupRequest(openapiclient.BlockStorageGroupRoleType("OS"), int32(123))}, openapiclient.InstanceGroupRoleType("ACTIVE"))}, "InstanceNamePrefix_example", "Name_example") // SearchEngineClusterRestoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchengineV1SearchEngineClustersApiAPI.SearchengineCreateRestore(context.Background(), clusterId).SearchEngineClusterRestoreRequest(searchEngineClusterRestoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineClustersApiAPI.SearchengineCreateRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineCreateRestore`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineClustersApiAPI.SearchengineCreateRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineCreateRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchEngineClusterRestoreRequest** | [**SearchEngineClusterRestoreRequest**](SearchEngineClusterRestoreRequest.md) |  | 

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


## SearchengineDeleteCluster

> AsyncResponse SearchengineDeleteCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineClustersApiAPI.SearchengineDeleteCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineClustersApiAPI.SearchengineDeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineDeleteCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineClustersApiAPI.SearchengineDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineDeleteClusterRequest struct via the builder pattern


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


## SearchengineListClusters

> ClusterPageResponse SearchengineListClusters(ctx).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineClustersApiAPI.SearchengineListClusters(context.Background()).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineClustersApiAPI.SearchengineListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineListClusters`: ClusterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineClustersApiAPI.SearchengineListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineListClustersRequest struct via the builder pattern


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


## SearchengineRestartCluster

> AsyncResponse SearchengineRestartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineClustersApiAPI.SearchengineRestartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineClustersApiAPI.SearchengineRestartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineRestartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineClustersApiAPI.SearchengineRestartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineRestartClusterRequest struct via the builder pattern


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


## SearchengineShowCluster

> SearchEngineClusterDetailResponse SearchengineShowCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineClustersApiAPI.SearchengineShowCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineClustersApiAPI.SearchengineShowCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineShowCluster`: SearchEngineClusterDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineClustersApiAPI.SearchengineShowCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineShowClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchEngineClusterDetailResponse**](SearchEngineClusterDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchengineStartCluster

> AsyncResponse SearchengineStartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineClustersApiAPI.SearchengineStartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineClustersApiAPI.SearchengineStartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineStartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineClustersApiAPI.SearchengineStartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineStartClusterRequest struct via the builder pattern


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


## SearchengineStopCluster

> AsyncResponse SearchengineStopCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineClustersApiAPI.SearchengineStopCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineClustersApiAPI.SearchengineStopCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineStopCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineClustersApiAPI.SearchengineStopCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineStopClusterRequest struct via the builder pattern


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


## SearchengineSyncClusterState

> AsyncResponse SearchengineSyncClusterState(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineClustersApiAPI.SearchengineSyncClusterState(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineClustersApiAPI.SearchengineSyncClusterState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineSyncClusterState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineClustersApiAPI.SearchengineSyncClusterState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineSyncClusterStateRequest struct via the builder pattern


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

