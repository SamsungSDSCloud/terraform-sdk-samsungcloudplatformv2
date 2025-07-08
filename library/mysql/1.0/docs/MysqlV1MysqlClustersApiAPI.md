# \MysqlV1MysqlClustersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlCreateCluster**](MysqlV1MysqlClustersApiAPI.md#MysqlCreateCluster) | **Post** /v1/clusters | Create Cluster
[**MysqlCreateRestore**](MysqlV1MysqlClustersApiAPI.md#MysqlCreateRestore) | **Post** /v1/clusters/{cluster_id}/restore | Create Restore Cluster
[**MysqlDeleteCluster**](MysqlV1MysqlClustersApiAPI.md#MysqlDeleteCluster) | **Delete** /v1/clusters/{cluster_id} | Terminate Cluster
[**MysqlListClusters**](MysqlV1MysqlClustersApiAPI.md#MysqlListClusters) | **Get** /v1/clusters | List Clusters
[**MysqlRestartCluster**](MysqlV1MysqlClustersApiAPI.md#MysqlRestartCluster) | **Post** /v1/clusters/{cluster_id}/restart | Restart Cluster
[**MysqlShowCluster**](MysqlV1MysqlClustersApiAPI.md#MysqlShowCluster) | **Get** /v1/clusters/{cluster_id} | Show Cluster
[**MysqlStartCluster**](MysqlV1MysqlClustersApiAPI.md#MysqlStartCluster) | **Post** /v1/clusters/{cluster_id}/start | Start Cluster
[**MysqlStopCluster**](MysqlV1MysqlClustersApiAPI.md#MysqlStopCluster) | **Post** /v1/clusters/{cluster_id}/stop | Stop Cluster
[**MysqlSwitchoverCluster**](MysqlV1MysqlClustersApiAPI.md#MysqlSwitchoverCluster) | **Post** /v1/clusters/{cluster_id}/switchover | Switchover Cluster
[**MysqlSyncClusterState**](MysqlV1MysqlClustersApiAPI.md#MysqlSyncClusterState) | **Post** /v1/clusters/{cluster_id}/sync-state | Synchronize Cluster State



## MysqlCreateCluster

> AsyncResponse MysqlCreateCluster(ctx).MysqlClusterCreateRequest(mysqlClusterCreateRequest).Execute()

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
	mysqlClusterCreateRequest := *openapiclient.NewMysqlClusterCreateRequest("DbaasEngineVersionId_example", *openapiclient.NewMysqlInitConfigOptionRequest("DatabaseName_example", "DatabaseUserName_example", "DatabaseUserPassword_example"), []openapiclient.InstanceGroupRequest{*openapiclient.NewInstanceGroupRequest([]openapiclient.BlockStorageGroupRequest{*openapiclient.NewBlockStorageGroupRequest(openapiclient.BlockStorageGroupRoleType("OS"), int32(123))}, openapiclient.InstanceGroupRoleType("ACTIVE"), "ServerTypeName_example")}, "InstanceNamePrefix_example", "Name_example", "SubnetId_example", "Timezone_example") // MysqlClusterCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlClustersApiAPI.MysqlCreateCluster(context.Background()).MysqlClusterCreateRequest(mysqlClusterCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlClustersApiAPI.MysqlCreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlCreateCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlClustersApiAPI.MysqlCreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMysqlCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mysqlClusterCreateRequest** | [**MysqlClusterCreateRequest**](MysqlClusterCreateRequest.md) |  | 

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


## MysqlCreateRestore

> AsyncResponse MysqlCreateRestore(ctx, clusterId).ClusterRestoreRequest(clusterRestoreRequest).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlClustersApiAPI.MysqlCreateRestore(context.Background(), clusterId).ClusterRestoreRequest(clusterRestoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlClustersApiAPI.MysqlCreateRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlCreateRestore`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlClustersApiAPI.MysqlCreateRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlCreateRestoreRequest struct via the builder pattern


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


## MysqlDeleteCluster

> AsyncResponse MysqlDeleteCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlClustersApiAPI.MysqlDeleteCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlClustersApiAPI.MysqlDeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlDeleteCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlClustersApiAPI.MysqlDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlDeleteClusterRequest struct via the builder pattern


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


## MysqlListClusters

> RdbClusterPageResponse MysqlListClusters(ctx).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).DatabaseName(databaseName).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlClustersApiAPI.MysqlListClusters(context.Background()).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).DatabaseName(databaseName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlClustersApiAPI.MysqlListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlListClusters`: RdbClusterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlClustersApiAPI.MysqlListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMysqlListClustersRequest struct via the builder pattern


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


## MysqlRestartCluster

> AsyncResponse MysqlRestartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlClustersApiAPI.MysqlRestartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlClustersApiAPI.MysqlRestartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlRestartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlClustersApiAPI.MysqlRestartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlRestartClusterRequest struct via the builder pattern


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


## MysqlShowCluster

> MysqlClusterDetailResponse MysqlShowCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlClustersApiAPI.MysqlShowCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlClustersApiAPI.MysqlShowCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlShowCluster`: MysqlClusterDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlClustersApiAPI.MysqlShowCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlShowClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MysqlClusterDetailResponse**](MysqlClusterDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlStartCluster

> AsyncResponse MysqlStartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlClustersApiAPI.MysqlStartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlClustersApiAPI.MysqlStartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlStartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlClustersApiAPI.MysqlStartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlStartClusterRequest struct via the builder pattern


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


## MysqlStopCluster

> AsyncResponse MysqlStopCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlClustersApiAPI.MysqlStopCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlClustersApiAPI.MysqlStopCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlStopCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlClustersApiAPI.MysqlStopCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlStopClusterRequest struct via the builder pattern


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


## MysqlSwitchoverCluster

> AsyncResponse MysqlSwitchoverCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlClustersApiAPI.MysqlSwitchoverCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlClustersApiAPI.MysqlSwitchoverCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSwitchoverCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlClustersApiAPI.MysqlSwitchoverCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSwitchoverClusterRequest struct via the builder pattern


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


## MysqlSyncClusterState

> AsyncResponse MysqlSyncClusterState(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlClustersApiAPI.MysqlSyncClusterState(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlClustersApiAPI.MysqlSyncClusterState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSyncClusterState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlClustersApiAPI.MysqlSyncClusterState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSyncClusterStateRequest struct via the builder pattern


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

