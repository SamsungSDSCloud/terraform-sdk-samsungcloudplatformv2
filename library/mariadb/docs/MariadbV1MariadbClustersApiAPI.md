# \MariadbV1MariadbClustersApiAPI

All URIs are relative to *https://scp-dbaas-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariadbCreateCluster**](MariadbV1MariadbClustersApiAPI.md#MariadbCreateCluster) | **Post** /v1/clusters | Create Cluster
[**MariadbCreateRestore**](MariadbV1MariadbClustersApiAPI.md#MariadbCreateRestore) | **Post** /v1/clusters/{cluster_id}/restore | Create Restore Cluster
[**MariadbDeleteCluster**](MariadbV1MariadbClustersApiAPI.md#MariadbDeleteCluster) | **Delete** /v1/clusters/{cluster_id} | Terminate Cluster
[**MariadbListClusters**](MariadbV1MariadbClustersApiAPI.md#MariadbListClusters) | **Get** /v1/clusters | List Clusters
[**MariadbRestartCluster**](MariadbV1MariadbClustersApiAPI.md#MariadbRestartCluster) | **Post** /v1/clusters/{cluster_id}/restart | Restart Cluster
[**MariadbShowCluster**](MariadbV1MariadbClustersApiAPI.md#MariadbShowCluster) | **Get** /v1/clusters/{cluster_id} | Show Cluster
[**MariadbStartCluster**](MariadbV1MariadbClustersApiAPI.md#MariadbStartCluster) | **Post** /v1/clusters/{cluster_id}/start | Start Cluster
[**MariadbStopCluster**](MariadbV1MariadbClustersApiAPI.md#MariadbStopCluster) | **Post** /v1/clusters/{cluster_id}/stop | Stop Cluster
[**MariadbSwitchoverCluster**](MariadbV1MariadbClustersApiAPI.md#MariadbSwitchoverCluster) | **Post** /v1/clusters/{cluster_id}/switchover | Switchover Cluster
[**MariadbSyncClusterState**](MariadbV1MariadbClustersApiAPI.md#MariadbSyncClusterState) | **Post** /v1/clusters/{cluster_id}/sync-state | Synchronize Cluster State



## MariadbCreateCluster

> AsyncResponse MariadbCreateCluster(ctx).MariadbClusterCreateRequest(mariadbClusterCreateRequest).Execute()

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
	mariadbClusterCreateRequest := *openapiclient.NewMariadbClusterCreateRequest("DbaasEngineVersionId_example", *openapiclient.NewMariadbInitConfigOptionRequest("DatabaseName_example", "DatabaseUserName_example", "DatabaseUserPassword_example"), []openapiclient.InstanceGroupRequest{*openapiclient.NewInstanceGroupRequest([]openapiclient.BlockStorageGroupRequest{*openapiclient.NewBlockStorageGroupRequest(openapiclient.BlockStorageGroupRoleType("OS"), int32(123))}, openapiclient.InstanceGroupRoleType("ACTIVE"), "ServerTypeName_example")}, "InstanceNamePrefix_example", "Name_example", "SubnetId_example", "Timezone_example") // MariadbClusterCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MariadbV1MariadbClustersApiAPI.MariadbCreateCluster(context.Background()).MariadbClusterCreateRequest(mariadbClusterCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbClustersApiAPI.MariadbCreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbCreateCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbClustersApiAPI.MariadbCreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMariadbCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mariadbClusterCreateRequest** | [**MariadbClusterCreateRequest**](MariadbClusterCreateRequest.md) |  | 

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


## MariadbCreateRestore

> AsyncResponse MariadbCreateRestore(ctx, clusterId).ClusterRestoreRequest(clusterRestoreRequest).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbClustersApiAPI.MariadbCreateRestore(context.Background(), clusterId).ClusterRestoreRequest(clusterRestoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbClustersApiAPI.MariadbCreateRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbCreateRestore`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbClustersApiAPI.MariadbCreateRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbCreateRestoreRequest struct via the builder pattern


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


## MariadbDeleteCluster

> AsyncResponse MariadbDeleteCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbClustersApiAPI.MariadbDeleteCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbClustersApiAPI.MariadbDeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbDeleteCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbClustersApiAPI.MariadbDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbDeleteClusterRequest struct via the builder pattern


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


## MariadbListClusters

> RdbClusterPageResponse MariadbListClusters(ctx).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).DatabaseName(databaseName).Execute()

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
	databaseName := "databaseName_example" // string | Database name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MariadbV1MariadbClustersApiAPI.MariadbListClusters(context.Background()).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).DatabaseName(databaseName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbClustersApiAPI.MariadbListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbListClusters`: RdbClusterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbClustersApiAPI.MariadbListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMariadbListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Cluster name | 
 **serviceState** | **string** | Service state | 
 **databaseName** | **string** | Database name | 

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


## MariadbRestartCluster

> AsyncResponse MariadbRestartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbClustersApiAPI.MariadbRestartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbClustersApiAPI.MariadbRestartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbRestartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbClustersApiAPI.MariadbRestartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbRestartClusterRequest struct via the builder pattern


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


## MariadbShowCluster

> MariadbClusterDetailResponse MariadbShowCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbClustersApiAPI.MariadbShowCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbClustersApiAPI.MariadbShowCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbShowCluster`: MariadbClusterDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbClustersApiAPI.MariadbShowCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbShowClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MariadbClusterDetailResponse**](MariadbClusterDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MariadbStartCluster

> AsyncResponse MariadbStartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbClustersApiAPI.MariadbStartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbClustersApiAPI.MariadbStartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbStartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbClustersApiAPI.MariadbStartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbStartClusterRequest struct via the builder pattern


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


## MariadbStopCluster

> AsyncResponse MariadbStopCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbClustersApiAPI.MariadbStopCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbClustersApiAPI.MariadbStopCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbStopCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbClustersApiAPI.MariadbStopCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbStopClusterRequest struct via the builder pattern


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


## MariadbSwitchoverCluster

> AsyncResponse MariadbSwitchoverCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbClustersApiAPI.MariadbSwitchoverCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbClustersApiAPI.MariadbSwitchoverCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbSwitchoverCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbClustersApiAPI.MariadbSwitchoverCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSwitchoverClusterRequest struct via the builder pattern


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


## MariadbSyncClusterState

> AsyncResponse MariadbSyncClusterState(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbClustersApiAPI.MariadbSyncClusterState(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbClustersApiAPI.MariadbSyncClusterState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbSyncClusterState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbClustersApiAPI.MariadbSyncClusterState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSyncClusterStateRequest struct via the builder pattern


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

