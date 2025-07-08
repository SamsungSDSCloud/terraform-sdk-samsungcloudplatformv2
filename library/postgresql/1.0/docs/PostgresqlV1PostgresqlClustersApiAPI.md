# \PostgresqlV1PostgresqlClustersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostgresqlCreateCluster**](PostgresqlV1PostgresqlClustersApiAPI.md#PostgresqlCreateCluster) | **Post** /v1/clusters | Create Cluster
[**PostgresqlCreateRestore**](PostgresqlV1PostgresqlClustersApiAPI.md#PostgresqlCreateRestore) | **Post** /v1/clusters/{cluster_id}/restore | Create Restore Cluster
[**PostgresqlDeleteCluster**](PostgresqlV1PostgresqlClustersApiAPI.md#PostgresqlDeleteCluster) | **Delete** /v1/clusters/{cluster_id} | Terminate Cluster
[**PostgresqlListClusters**](PostgresqlV1PostgresqlClustersApiAPI.md#PostgresqlListClusters) | **Get** /v1/clusters | List Clusters
[**PostgresqlRestartCluster**](PostgresqlV1PostgresqlClustersApiAPI.md#PostgresqlRestartCluster) | **Post** /v1/clusters/{cluster_id}/restart | Restart Cluster
[**PostgresqlShowCluster**](PostgresqlV1PostgresqlClustersApiAPI.md#PostgresqlShowCluster) | **Get** /v1/clusters/{cluster_id} | Show Cluster
[**PostgresqlStartCluster**](PostgresqlV1PostgresqlClustersApiAPI.md#PostgresqlStartCluster) | **Post** /v1/clusters/{cluster_id}/start | Start Cluster
[**PostgresqlStopCluster**](PostgresqlV1PostgresqlClustersApiAPI.md#PostgresqlStopCluster) | **Post** /v1/clusters/{cluster_id}/stop | Stop Cluster
[**PostgresqlSwitchoverCluster**](PostgresqlV1PostgresqlClustersApiAPI.md#PostgresqlSwitchoverCluster) | **Post** /v1/clusters/{cluster_id}/switchover | Switchover Cluster
[**PostgresqlSyncClusterState**](PostgresqlV1PostgresqlClustersApiAPI.md#PostgresqlSyncClusterState) | **Post** /v1/clusters/{cluster_id}/sync-state | Synchronize Cluster State



## PostgresqlCreateCluster

> AsyncResponse PostgresqlCreateCluster(ctx).PostgresqlClusterCreateRequest(postgresqlClusterCreateRequest).Execute()

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
	postgresqlClusterCreateRequest := *openapiclient.NewPostgresqlClusterCreateRequest("DbaasEngineVersionId_example", *openapiclient.NewPostgresqlInitConfigOptionRequest("DatabaseName_example", "DatabaseUserName_example", "DatabaseUserPassword_example"), []openapiclient.InstanceGroupRequest{*openapiclient.NewInstanceGroupRequest([]openapiclient.BlockStorageGroupRequest{*openapiclient.NewBlockStorageGroupRequest(openapiclient.BlockStorageGroupRoleType("OS"), int32(123))}, openapiclient.InstanceGroupRoleType("ACTIVE"), "ServerTypeName_example")}, "InstanceNamePrefix_example", "Name_example", "SubnetId_example", "Timezone_example") // PostgresqlClusterCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostgresqlV1PostgresqlClustersApiAPI.PostgresqlCreateCluster(context.Background()).PostgresqlClusterCreateRequest(postgresqlClusterCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlCreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlCreateCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlCreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postgresqlClusterCreateRequest** | [**PostgresqlClusterCreateRequest**](PostgresqlClusterCreateRequest.md) |  | 

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


## PostgresqlCreateRestore

> AsyncResponse PostgresqlCreateRestore(ctx, clusterId).ClusterRestoreRequest(clusterRestoreRequest).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlClustersApiAPI.PostgresqlCreateRestore(context.Background(), clusterId).ClusterRestoreRequest(clusterRestoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlCreateRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlCreateRestore`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlCreateRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlCreateRestoreRequest struct via the builder pattern


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


## PostgresqlDeleteCluster

> AsyncResponse PostgresqlDeleteCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlClustersApiAPI.PostgresqlDeleteCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlDeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlDeleteCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlDeleteClusterRequest struct via the builder pattern


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


## PostgresqlListClusters

> RdbClusterPageResponse PostgresqlListClusters(ctx).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).DatabaseName(databaseName).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlClustersApiAPI.PostgresqlListClusters(context.Background()).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).DatabaseName(databaseName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlListClusters`: RdbClusterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlListClustersRequest struct via the builder pattern


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


## PostgresqlRestartCluster

> AsyncResponse PostgresqlRestartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlClustersApiAPI.PostgresqlRestartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlRestartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlRestartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlRestartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlRestartClusterRequest struct via the builder pattern


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


## PostgresqlShowCluster

> PostgresqlClusterDetailResponse PostgresqlShowCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlClustersApiAPI.PostgresqlShowCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlShowCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlShowCluster`: PostgresqlClusterDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlShowCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlShowClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostgresqlClusterDetailResponse**](PostgresqlClusterDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlStartCluster

> AsyncResponse PostgresqlStartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlClustersApiAPI.PostgresqlStartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlStartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlStartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlStartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlStartClusterRequest struct via the builder pattern


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


## PostgresqlStopCluster

> AsyncResponse PostgresqlStopCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlClustersApiAPI.PostgresqlStopCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlStopCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlStopCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlStopCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlStopClusterRequest struct via the builder pattern


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


## PostgresqlSwitchoverCluster

> AsyncResponse PostgresqlSwitchoverCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlClustersApiAPI.PostgresqlSwitchoverCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlSwitchoverCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlSwitchoverCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlSwitchoverCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlSwitchoverClusterRequest struct via the builder pattern


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


## PostgresqlSyncClusterState

> AsyncResponse PostgresqlSyncClusterState(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlClustersApiAPI.PostgresqlSyncClusterState(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlSyncClusterState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlSyncClusterState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlClustersApiAPI.PostgresqlSyncClusterState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlSyncClusterStateRequest struct via the builder pattern


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

