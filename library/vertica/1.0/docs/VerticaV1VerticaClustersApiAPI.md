# \VerticaV1VerticaClustersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerticaCreateCluster**](VerticaV1VerticaClustersApiAPI.md#VerticaCreateCluster) | **Post** /v1/clusters | Create Cluster
[**VerticaCreateRestore**](VerticaV1VerticaClustersApiAPI.md#VerticaCreateRestore) | **Post** /v1/clusters/{cluster_id}/restore | Create Restore Cluster
[**VerticaDeleteCluster**](VerticaV1VerticaClustersApiAPI.md#VerticaDeleteCluster) | **Delete** /v1/clusters/{cluster_id} | Terminate Cluster
[**VerticaListClusters**](VerticaV1VerticaClustersApiAPI.md#VerticaListClusters) | **Get** /v1/clusters | List Clusters
[**VerticaRestartCluster**](VerticaV1VerticaClustersApiAPI.md#VerticaRestartCluster) | **Post** /v1/clusters/{cluster_id}/restart | Restart Cluster
[**VerticaShowCluster**](VerticaV1VerticaClustersApiAPI.md#VerticaShowCluster) | **Get** /v1/clusters/{cluster_id} | Show Cluster
[**VerticaStartCluster**](VerticaV1VerticaClustersApiAPI.md#VerticaStartCluster) | **Post** /v1/clusters/{cluster_id}/start | Start Cluster
[**VerticaStopCluster**](VerticaV1VerticaClustersApiAPI.md#VerticaStopCluster) | **Post** /v1/clusters/{cluster_id}/stop | Stop Cluster
[**VerticaSyncClusterState**](VerticaV1VerticaClustersApiAPI.md#VerticaSyncClusterState) | **Post** /v1/clusters/{cluster_id}/sync-state | Synchronize Cluster State



## VerticaCreateCluster

> AsyncResponse VerticaCreateCluster(ctx).VerticaClusterCreateRequest(verticaClusterCreateRequest).Execute()

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
	verticaClusterCreateRequest := *openapiclient.NewVerticaClusterCreateRequest("DbaasEngineVersionId_example", *openapiclient.NewVerticaInitConfigOptionRequest("DatabaseName_example", "DatabaseUserName_example", "DatabaseUserPassword_example"), []openapiclient.InstanceGroupRequest{*openapiclient.NewInstanceGroupRequest([]openapiclient.BlockStorageGroupRequest{*openapiclient.NewBlockStorageGroupRequest(openapiclient.BlockStorageGroupRoleType("OS"), int32(123))}, openapiclient.InstanceGroupRoleType("ACTIVE"), "ServerTypeName_example")}, "InstanceNamePrefix_example", "Name_example", "SubnetId_example", "Timezone_example") // VerticaClusterCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerticaV1VerticaClustersApiAPI.VerticaCreateCluster(context.Background()).VerticaClusterCreateRequest(verticaClusterCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaClustersApiAPI.VerticaCreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaCreateCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaClustersApiAPI.VerticaCreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerticaCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verticaClusterCreateRequest** | [**VerticaClusterCreateRequest**](VerticaClusterCreateRequest.md) |  | 

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


## VerticaCreateRestore

> AsyncResponse VerticaCreateRestore(ctx, clusterId).VerticaClusterRestoreRequest(verticaClusterRestoreRequest).Execute()

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
	verticaClusterRestoreRequest := *openapiclient.NewVerticaClusterRestoreRequest("BackupHistoryNumber_example", []openapiclient.VerticaRestoreInstanceGroup{*openapiclient.NewVerticaRestoreInstanceGroup([]openapiclient.BlockStorageGroupRequest{*openapiclient.NewBlockStorageGroupRequest(openapiclient.BlockStorageGroupRoleType("OS"), int32(123))}, openapiclient.InstanceGroupRoleType("ACTIVE"), "ServerTypeName_example")}, "InstanceNamePrefix_example", "Name_example") // VerticaClusterRestoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerticaV1VerticaClustersApiAPI.VerticaCreateRestore(context.Background(), clusterId).VerticaClusterRestoreRequest(verticaClusterRestoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaClustersApiAPI.VerticaCreateRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaCreateRestore`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaClustersApiAPI.VerticaCreateRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerticaCreateRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verticaClusterRestoreRequest** | [**VerticaClusterRestoreRequest**](VerticaClusterRestoreRequest.md) |  | 

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


## VerticaDeleteCluster

> AsyncResponse VerticaDeleteCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.VerticaV1VerticaClustersApiAPI.VerticaDeleteCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaClustersApiAPI.VerticaDeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaDeleteCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaClustersApiAPI.VerticaDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerticaDeleteClusterRequest struct via the builder pattern


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


## VerticaListClusters

> VerticaClusterPageResponse VerticaListClusters(ctx).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).Execute()

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
	resp, r, err := apiClient.VerticaV1VerticaClustersApiAPI.VerticaListClusters(context.Background()).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaClustersApiAPI.VerticaListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaListClusters`: VerticaClusterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaClustersApiAPI.VerticaListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerticaListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Cluster name | 
 **serviceState** | **string** | Service state | 

### Return type

[**VerticaClusterPageResponse**](VerticaClusterPageResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerticaRestartCluster

> AsyncResponse VerticaRestartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.VerticaV1VerticaClustersApiAPI.VerticaRestartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaClustersApiAPI.VerticaRestartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaRestartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaClustersApiAPI.VerticaRestartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerticaRestartClusterRequest struct via the builder pattern


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


## VerticaShowCluster

> VerticaClusterDetailResponse VerticaShowCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.VerticaV1VerticaClustersApiAPI.VerticaShowCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaClustersApiAPI.VerticaShowCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaShowCluster`: VerticaClusterDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaClustersApiAPI.VerticaShowCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerticaShowClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VerticaClusterDetailResponse**](VerticaClusterDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerticaStartCluster

> AsyncResponse VerticaStartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.VerticaV1VerticaClustersApiAPI.VerticaStartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaClustersApiAPI.VerticaStartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaStartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaClustersApiAPI.VerticaStartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerticaStartClusterRequest struct via the builder pattern


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


## VerticaStopCluster

> AsyncResponse VerticaStopCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.VerticaV1VerticaClustersApiAPI.VerticaStopCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaClustersApiAPI.VerticaStopCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaStopCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaClustersApiAPI.VerticaStopCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerticaStopClusterRequest struct via the builder pattern


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


## VerticaSyncClusterState

> AsyncResponse VerticaSyncClusterState(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.VerticaV1VerticaClustersApiAPI.VerticaSyncClusterState(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaClustersApiAPI.VerticaSyncClusterState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaSyncClusterState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaClustersApiAPI.VerticaSyncClusterState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerticaSyncClusterStateRequest struct via the builder pattern


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

