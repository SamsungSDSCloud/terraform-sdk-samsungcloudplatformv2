# \EventstreamsV1EventStreamsClustersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventstreamsAddInstances**](EventstreamsV1EventStreamsClustersApiAPI.md#EventstreamsAddInstances) | **Post** /v1/clusters/{cluster_id}/add-instances | Add Instances
[**EventstreamsCreateCluster**](EventstreamsV1EventStreamsClustersApiAPI.md#EventstreamsCreateCluster) | **Post** /v1/clusters | Create Cluster
[**EventstreamsDeleteCluster**](EventstreamsV1EventStreamsClustersApiAPI.md#EventstreamsDeleteCluster) | **Delete** /v1/clusters/{cluster_id} | Terminate Cluster
[**EventstreamsListClusters**](EventstreamsV1EventStreamsClustersApiAPI.md#EventstreamsListClusters) | **Get** /v1/clusters | List Clusters
[**EventstreamsRestartCluster**](EventstreamsV1EventStreamsClustersApiAPI.md#EventstreamsRestartCluster) | **Post** /v1/clusters/{cluster_id}/restart | Restart Cluster
[**EventstreamsShowCluster**](EventstreamsV1EventStreamsClustersApiAPI.md#EventstreamsShowCluster) | **Get** /v1/clusters/{cluster_id} | Show Cluster
[**EventstreamsStartCluster**](EventstreamsV1EventStreamsClustersApiAPI.md#EventstreamsStartCluster) | **Post** /v1/clusters/{cluster_id}/start | Start Cluster
[**EventstreamsStopCluster**](EventstreamsV1EventStreamsClustersApiAPI.md#EventstreamsStopCluster) | **Post** /v1/clusters/{cluster_id}/stop | Stop Cluster
[**EventstreamsSyncClusterState**](EventstreamsV1EventStreamsClustersApiAPI.md#EventstreamsSyncClusterState) | **Post** /v1/clusters/{cluster_id}/sync-state | Synchronize Cluster State



## EventstreamsAddInstances

> AsyncResponse EventstreamsAddInstances(ctx, clusterId).EventStreamsClusterAddInstancesRequest(eventStreamsClusterAddInstancesRequest).Execute()

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
	eventStreamsClusterAddInstancesRequest := *openapiclient.NewEventStreamsClusterAddInstancesRequest(int32(123), []string{"ServiceIpAddresses_example"}) // EventStreamsClusterAddInstancesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventstreamsV1EventStreamsClustersApiAPI.EventstreamsAddInstances(context.Background(), clusterId).EventStreamsClusterAddInstancesRequest(eventStreamsClusterAddInstancesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsAddInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsAddInstances`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsAddInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsAddInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventStreamsClusterAddInstancesRequest** | [**EventStreamsClusterAddInstancesRequest**](EventStreamsClusterAddInstancesRequest.md) |  | 

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


## EventstreamsCreateCluster

> AsyncResponse EventstreamsCreateCluster(ctx).EventStreamsClusterCreateRequest(eventStreamsClusterCreateRequest).Execute()

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
	eventStreamsClusterCreateRequest := *openapiclient.NewEventStreamsClusterCreateRequest("DbaasEngineVersionId_example", *openapiclient.NewEventStreamsInitConfigOptionRequest("BrokerSaslId_example", "BrokerSaslPassword_example", "ZookeeperSaslId_example", "ZookeeperSaslPassword_example"), []openapiclient.InstanceGroupRequest{*openapiclient.NewInstanceGroupRequest([]openapiclient.BlockStorageGroupRequest{*openapiclient.NewBlockStorageGroupRequest(openapiclient.BlockStorageGroupRoleType("OS"), int32(123))}, openapiclient.InstanceGroupRoleType("ACTIVE"), "ServerTypeName_example")}, "InstanceNamePrefix_example", "Name_example", "SubnetId_example", "Timezone_example") // EventStreamsClusterCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventstreamsV1EventStreamsClustersApiAPI.EventstreamsCreateCluster(context.Background()).EventStreamsClusterCreateRequest(eventStreamsClusterCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsCreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsCreateCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsCreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventStreamsClusterCreateRequest** | [**EventStreamsClusterCreateRequest**](EventStreamsClusterCreateRequest.md) |  | 

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


## EventstreamsDeleteCluster

> AsyncResponse EventstreamsDeleteCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EventstreamsV1EventStreamsClustersApiAPI.EventstreamsDeleteCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsDeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsDeleteCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsDeleteClusterRequest struct via the builder pattern


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


## EventstreamsListClusters

> ClusterPageResponse EventstreamsListClusters(ctx).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).Execute()

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
	resp, r, err := apiClient.EventstreamsV1EventStreamsClustersApiAPI.EventstreamsListClusters(context.Background()).Size(size).Page(page).Sort(sort).Name(name).ServiceState(serviceState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsListClusters`: ClusterPageResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsListClustersRequest struct via the builder pattern


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


## EventstreamsRestartCluster

> AsyncResponse EventstreamsRestartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EventstreamsV1EventStreamsClustersApiAPI.EventstreamsRestartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsRestartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsRestartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsRestartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsRestartClusterRequest struct via the builder pattern


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


## EventstreamsShowCluster

> EventStreamsClusterDetailResponse EventstreamsShowCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EventstreamsV1EventStreamsClustersApiAPI.EventstreamsShowCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsShowCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsShowCluster`: EventStreamsClusterDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsShowCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsShowClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventStreamsClusterDetailResponse**](EventStreamsClusterDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventstreamsStartCluster

> AsyncResponse EventstreamsStartCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EventstreamsV1EventStreamsClustersApiAPI.EventstreamsStartCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsStartCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsStartCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsStartCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsStartClusterRequest struct via the builder pattern


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


## EventstreamsStopCluster

> AsyncResponse EventstreamsStopCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EventstreamsV1EventStreamsClustersApiAPI.EventstreamsStopCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsStopCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsStopCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsStopCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsStopClusterRequest struct via the builder pattern


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


## EventstreamsSyncClusterState

> AsyncResponse EventstreamsSyncClusterState(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EventstreamsV1EventStreamsClustersApiAPI.EventstreamsSyncClusterState(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsSyncClusterState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsSyncClusterState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsClustersApiAPI.EventstreamsSyncClusterState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsSyncClusterStateRequest struct via the builder pattern


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

