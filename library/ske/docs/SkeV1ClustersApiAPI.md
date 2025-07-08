# \SkeV1ClustersApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](SkeV1ClustersApiAPI.md#CreateCluster) | **Post** /v1/clusters | Create Cluster
[**DeleteCluster**](SkeV1ClustersApiAPI.md#DeleteCluster) | **Delete** /v1/clusters/{cluster_id} | Delete Cluster
[**ListClusters**](SkeV1ClustersApiAPI.md#ListClusters) | **Get** /v1/clusters | List Clusters
[**SetClusterLogging**](SkeV1ClustersApiAPI.md#SetClusterLogging) | **Put** /v1/clusters/{cluster_id}/logging | Set Cluster Logging
[**SetClusterPrivateAccessControl**](SkeV1ClustersApiAPI.md#SetClusterPrivateAccessControl) | **Put** /v1/clusters/{cluster_id}/private-access-control | Set Cluster Private Access Control
[**SetClusterPublicAccessControl**](SkeV1ClustersApiAPI.md#SetClusterPublicAccessControl) | **Put** /v1/clusters/{cluster_id}/public-access-control | Set Cluster Public Access Control
[**SetClusterSecurityGroups**](SkeV1ClustersApiAPI.md#SetClusterSecurityGroups) | **Put** /v1/clusters/{cluster_id}/security-groups | Set Cluster Security Groups
[**SetClusterUpgrade**](SkeV1ClustersApiAPI.md#SetClusterUpgrade) | **Put** /v1/clusters/{cluster_id}/upgrade | Set Cluster Upgrade
[**ShowCluster**](SkeV1ClustersApiAPI.md#ShowCluster) | **Get** /v1/clusters/{cluster_id} | Show Cluster
[**ShowClusterKubeconfig**](SkeV1ClustersApiAPI.md#ShowClusterKubeconfig) | **Get** /v1/clusters/{cluster_id}/kubeconfig | Show Cluster Kubeconfig



## CreateCluster

> AsyncResponse CreateCluster(ctx).ClusterCreateRequest(clusterCreateRequest).Execute()

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
	clusterCreateRequest := *openapiclient.NewClusterCreateRequest(false, "KubernetesVersion_example", "Name_example", []string{"SecurityGroupIdList_example"}, "SubnetId_example", "VolumeId_example", "VpcId_example") // ClusterCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1ClustersApiAPI.CreateCluster(context.Background()).ClusterCreateRequest(clusterCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClustersApiAPI.CreateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ClustersApiAPI.CreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterCreateRequest** | [**ClusterCreateRequest**](ClusterCreateRequest.md) |  | 

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


## DeleteCluster

> AsyncResponse DeleteCluster(ctx, clusterId).Execute()

Delete Cluster



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
	clusterId := "10a599e031e749b7b260868f441e862b" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1ClustersApiAPI.DeleteCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClustersApiAPI.DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ClustersApiAPI.DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterRequest struct via the builder pattern


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


## ListClusters

> ClusterListResponse ListClusters(ctx).Size(size).Page(page).Sort(sort).Name(name).Status(status).KubernetesVersion(kubernetesVersion).SubnetId(subnetId).Execute()

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
	name := "name_example" // string | Cluster Name (optional)
	status := *openapiclient.NewStatus() // Status | Cluster Status (optional)
	kubernetesVersion := *openapiclient.NewKubernetesVersion() // KubernetesVersion | Cluster Version (optional)
	subnetId := "subnetId_example" // string | Cluster Subnet ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1ClustersApiAPI.ListClusters(context.Background()).Size(size).Page(page).Sort(sort).Name(name).Status(status).KubernetesVersion(kubernetesVersion).SubnetId(subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClustersApiAPI.ListClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusters`: ClusterListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ClustersApiAPI.ListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Cluster Name | 
 **status** | [**Status**](Status.md) | Cluster Status | 
 **kubernetesVersion** | [**KubernetesVersion**](KubernetesVersion.md) | Cluster Version | 
 **subnetId** | **string** | Cluster Subnet ID | 

### Return type

[**ClusterListResponse**](ClusterListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetClusterLogging

> ClusterSetResponse SetClusterLogging(ctx, clusterId).ClusterLoggingSetRequest(clusterLoggingSetRequest).Execute()

Set Cluster Logging



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
	clusterId := "10a599e031e749b7b260868f441e862b" // string | Cluster ID
	clusterLoggingSetRequest := *openapiclient.NewClusterLoggingSetRequest(false) // ClusterLoggingSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1ClustersApiAPI.SetClusterLogging(context.Background(), clusterId).ClusterLoggingSetRequest(clusterLoggingSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClustersApiAPI.SetClusterLogging``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetClusterLogging`: ClusterSetResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ClustersApiAPI.SetClusterLogging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetClusterLoggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterLoggingSetRequest** | [**ClusterLoggingSetRequest**](ClusterLoggingSetRequest.md) |  | 

### Return type

[**ClusterSetResponse**](ClusterSetResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetClusterPrivateAccessControl

> ClusterSetResponse SetClusterPrivateAccessControl(ctx, clusterId).ClusterPrivateAccessControlSetRequest(clusterPrivateAccessControlSetRequest).Execute()

Set Cluster Private Access Control



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
	clusterId := "10a599e031e749b7b260868f441e862b" // string | Cluster ID
	clusterPrivateAccessControlSetRequest := *openapiclient.NewClusterPrivateAccessControlSetRequest([]openapiclient.PrivateEndpointAccessControlResource{*openapiclient.NewPrivateEndpointAccessControlResource()}) // ClusterPrivateAccessControlSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1ClustersApiAPI.SetClusterPrivateAccessControl(context.Background(), clusterId).ClusterPrivateAccessControlSetRequest(clusterPrivateAccessControlSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClustersApiAPI.SetClusterPrivateAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetClusterPrivateAccessControl`: ClusterSetResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ClustersApiAPI.SetClusterPrivateAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetClusterPrivateAccessControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterPrivateAccessControlSetRequest** | [**ClusterPrivateAccessControlSetRequest**](ClusterPrivateAccessControlSetRequest.md) |  | 

### Return type

[**ClusterSetResponse**](ClusterSetResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetClusterPublicAccessControl

> ClusterSetResponse SetClusterPublicAccessControl(ctx, clusterId).ClusterPublicAccessControlSetRequest(clusterPublicAccessControlSetRequest).Execute()

Set Cluster Public Access Control



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
	clusterId := "10a599e031e749b7b260868f441e862b" // string | Cluster ID
	clusterPublicAccessControlSetRequest := *openapiclient.NewClusterPublicAccessControlSetRequest("PublicEndpointAccessControlIp_example") // ClusterPublicAccessControlSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1ClustersApiAPI.SetClusterPublicAccessControl(context.Background(), clusterId).ClusterPublicAccessControlSetRequest(clusterPublicAccessControlSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClustersApiAPI.SetClusterPublicAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetClusterPublicAccessControl`: ClusterSetResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ClustersApiAPI.SetClusterPublicAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetClusterPublicAccessControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterPublicAccessControlSetRequest** | [**ClusterPublicAccessControlSetRequest**](ClusterPublicAccessControlSetRequest.md) |  | 

### Return type

[**ClusterSetResponse**](ClusterSetResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetClusterSecurityGroups

> ClusterShowResponse SetClusterSecurityGroups(ctx, clusterId).ClusterSecurityGroupsSetRequest(clusterSecurityGroupsSetRequest).Execute()

Set Cluster Security Groups



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
	clusterId := "10a599e031e749b7b260868f441e862b" // string | Cluster ID
	clusterSecurityGroupsSetRequest := *openapiclient.NewClusterSecurityGroupsSetRequest([]string{"SecurityGroupIdList_example"}) // ClusterSecurityGroupsSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1ClustersApiAPI.SetClusterSecurityGroups(context.Background(), clusterId).ClusterSecurityGroupsSetRequest(clusterSecurityGroupsSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClustersApiAPI.SetClusterSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetClusterSecurityGroups`: ClusterShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ClustersApiAPI.SetClusterSecurityGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetClusterSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterSecurityGroupsSetRequest** | [**ClusterSecurityGroupsSetRequest**](ClusterSecurityGroupsSetRequest.md) |  | 

### Return type

[**ClusterShowResponse**](ClusterShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetClusterUpgrade

> ClusterSetResponse SetClusterUpgrade(ctx, clusterId).ClusterUpgradeSetRequest(clusterUpgradeSetRequest).Execute()

Set Cluster Upgrade



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
	clusterId := "10a599e031e749b7b260868f441e862b" // string | Cluster ID
	clusterUpgradeSetRequest := *openapiclient.NewClusterUpgradeSetRequest("KubernetesVersion_example") // ClusterUpgradeSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1ClustersApiAPI.SetClusterUpgrade(context.Background(), clusterId).ClusterUpgradeSetRequest(clusterUpgradeSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClustersApiAPI.SetClusterUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetClusterUpgrade`: ClusterSetResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ClustersApiAPI.SetClusterUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetClusterUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterUpgradeSetRequest** | [**ClusterUpgradeSetRequest**](ClusterUpgradeSetRequest.md) |  | 

### Return type

[**ClusterSetResponse**](ClusterSetResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowCluster

> ClusterShowResponse ShowCluster(ctx, clusterId).Execute()

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
	clusterId := "10a599e031e749b7b260868f441e862b" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1ClustersApiAPI.ShowCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClustersApiAPI.ShowCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowCluster`: ClusterShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ClustersApiAPI.ShowCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterShowResponse**](ClusterShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowClusterKubeconfig

> ShowClusterKubeconfig(ctx, clusterId).KubeconfigType(kubeconfigType).Execute()

Show Cluster Kubeconfig



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
	clusterId := "10a599e031e749b7b260868f441e862b" // string | Cluster ID
	kubeconfigType := openapiclient.ClusterKubeconfigType("private") // ClusterKubeconfigType | Kubeconfig Type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SkeV1ClustersApiAPI.ShowClusterKubeconfig(context.Background(), clusterId).KubeconfigType(kubeconfigType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClustersApiAPI.ShowClusterKubeconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowClusterKubeconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubeconfigType** | [**ClusterKubeconfigType**](ClusterKubeconfigType.md) | Kubeconfig Type | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

