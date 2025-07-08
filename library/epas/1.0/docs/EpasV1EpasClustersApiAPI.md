# \EpasV1EpasClustersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EpasCreateCluster**](EpasV1EpasClustersApiAPI.md#EpasCreateCluster) | **Post** /v1/clusters | Create Cluster
[**EpasCreateRestore**](EpasV1EpasClustersApiAPI.md#EpasCreateRestore) | **Post** /v1/clusters/{cluster_id}/restore | Create Restore Cluster
[**EpasDeleteCluster**](EpasV1EpasClustersApiAPI.md#EpasDeleteCluster) | **Delete** /v1/clusters/{cluster_id} | Terminate Cluster
[**EpasListClusters**](EpasV1EpasClustersApiAPI.md#EpasListClusters) | **Get** /v1/clusters | List Clusters
[**EpasRestartCluster**](EpasV1EpasClustersApiAPI.md#EpasRestartCluster) | **Post** /v1/clusters/{cluster_id}/restart | Restart Cluster
[**EpasShowCluster**](EpasV1EpasClustersApiAPI.md#EpasShowCluster) | **Get** /v1/clusters/{cluster_id} | Show Cluster
[**EpasStartCluster**](EpasV1EpasClustersApiAPI.md#EpasStartCluster) | **Post** /v1/clusters/{cluster_id}/start | Start Cluster
[**EpasStopCluster**](EpasV1EpasClustersApiAPI.md#EpasStopCluster) | **Post** /v1/clusters/{cluster_id}/stop | Stop Cluster
[**EpasSwitchoverCluster**](EpasV1EpasClustersApiAPI.md#EpasSwitchoverCluster) | **Post** /v1/clusters/{cluster_id}/switchover | Switchover Cluster
[**EpasSyncClusterState**](EpasV1EpasClustersApiAPI.md#EpasSyncClusterState) | **Post** /v1/clusters/{cluster_id}/sync-state | Synchronize Cluster State



## EpasCreateCluster

> AsyncResponse EpasCreateCluster(ctx).EpasClusterCreateRequest(epasClusterCreateRequest).Execute()

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
	epasClusterCreateRequest := *openapiclient.NewEpasClusterCreateRequest("DbaasEngineVersionId_example", *openapiclient.NewEpasInitConfigOptionRequest("DatabaseName_example", "DatabaseUserName_example", "DatabaseUserPassword_example"), []openapiclient.InstanceGroupRequest{*openapiclient.NewInstanceGroupRequest([]openapiclient.BlockStorageGroupRequest{*openapiclient.NewBlockStorageGroupRequest(openapiclient.BlockStorageGroupRoleType("OS"), int32(123))}, openapiclient.InstanceGroupRoleType("ACTIVE"), "ServerTypeName_example")}, "InstanceNamePrefix_example", "Name_example", "SubnetId_example", "Timezone_example") // EpasClusterCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasClustersApiAPI.EpasCreateCluster(context.Background()).EpasClusterCreateRequest(epasClusterCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasClustersApiAPI.EpasCreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasCreateCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasClustersApiAPI.EpasCreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEpasCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **epasClusterCreateRequest** | [**EpasClusterCreateRequest**](EpasClusterCreateRequest.md) |  | 

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


## EpasCreateRestore

> AsyncResponse EpasCreateRestore(ctx, clusterId).ClusterRestoreRequest(clusterRestoreRequest).Execute()

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
	resp, r, err := apiClient.EpasV1EpasClustersApiAPI.EpasCreateRestore(context.Background(), clusterId).ClusterRestoreRequest(clusterRestoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasClustersApiAPI.EpasCreateRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasCreateRestore`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasClustersApiAPI.EpasCreateRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasCreateRestoreRequest struct via the builder pattern


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


## EpasDeleteCluster

> AsyncResponse EpasDeleteCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EpasV1EpasClustersApiAPI.EpasDeleteCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasClustersApiAPI.EpasDeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasDeleteCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasClustersApiAPI.EpasDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasDeleteClusterRequest struct via the builder pattern


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


## EpasListClusters

> RdbClusterPageResponse EpasListClusters(ctx).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).DatabaseName(databaseName).Execute()

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
	databaseName := "databaseName_example" // string | Database Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasClustersApiAPI.EpasListClusters(context.Background()).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).DatabaseName(databaseName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasClustersApiAPI.EpasListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasListClusters`: RdbClusterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasClustersApiAPI.EpasListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEpasListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Cluster name | 
 **serviceState** | **string** | Service state | 
 **databaseName** | **string** | Database Name | 

### Return type

[**RdbClusterPageResponse**](RdbClusterPageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EpasRestartCluster

> AsyncResponse EpasRestartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EpasV1EpasClustersApiAPI.EpasRestartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasClustersApiAPI.EpasRestartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasRestartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasClustersApiAPI.EpasRestartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasRestartClusterRequest struct via the builder pattern


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


## EpasShowCluster

> EpasClusterDetailResponse EpasShowCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EpasV1EpasClustersApiAPI.EpasShowCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasClustersApiAPI.EpasShowCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasShowCluster`: EpasClusterDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasClustersApiAPI.EpasShowCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasShowClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EpasClusterDetailResponse**](EpasClusterDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EpasStartCluster

> AsyncResponse EpasStartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EpasV1EpasClustersApiAPI.EpasStartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasClustersApiAPI.EpasStartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasStartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasClustersApiAPI.EpasStartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasStartClusterRequest struct via the builder pattern


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


## EpasStopCluster

> AsyncResponse EpasStopCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EpasV1EpasClustersApiAPI.EpasStopCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasClustersApiAPI.EpasStopCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasStopCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasClustersApiAPI.EpasStopCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasStopClusterRequest struct via the builder pattern


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


## EpasSwitchoverCluster

> AsyncResponse EpasSwitchoverCluster(ctx, clusterId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasClustersApiAPI.EpasSwitchoverCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasClustersApiAPI.EpasSwitchoverCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSwitchoverCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasClustersApiAPI.EpasSwitchoverCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSwitchoverClusterRequest struct via the builder pattern


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


## EpasSyncClusterState

> AsyncResponse EpasSyncClusterState(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EpasV1EpasClustersApiAPI.EpasSyncClusterState(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasClustersApiAPI.EpasSyncClusterState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSyncClusterState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasClustersApiAPI.EpasSyncClusterState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSyncClusterStateRequest struct via the builder pattern


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

