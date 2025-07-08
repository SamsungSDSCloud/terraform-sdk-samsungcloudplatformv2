# \SkeV1ClusterRoleBindingsApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClusterRoleBinding**](SkeV1ClusterRoleBindingsApiAPI.md#DeleteClusterRoleBinding) | **Delete** /v1/clusters/{cluster_id}/clusterrolebindings | Delete Cluster RoleBinding
[**ListClusterRoleBindings**](SkeV1ClusterRoleBindingsApiAPI.md#ListClusterRoleBindings) | **Get** /v1/clusters/{cluster_id}/clusterrolebindings | List Cluster Role Bindings
[**ShowClusterRoleBinding**](SkeV1ClusterRoleBindingsApiAPI.md#ShowClusterRoleBinding) | **Get** /v1/clusters/{cluster_id}/clusterrolebindings/{cluster_role_binding_name} | Show Cluster RoleBinding



## DeleteClusterRoleBinding

> DeleteClusterRoleBinding(ctx, clusterId).Execute()

Delete Cluster RoleBinding



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
	r, err := apiClient.SkeV1ClusterRoleBindingsApiAPI.DeleteClusterRoleBinding(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClusterRoleBindingsApiAPI.DeleteClusterRoleBinding``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteClusterRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListClusterRoleBindings

> ClusterRoleBindingListResponse ListClusterRoleBindings(ctx, clusterId).Execute()

List Cluster Role Bindings



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
	resp, r, err := apiClient.SkeV1ClusterRoleBindingsApiAPI.ListClusterRoleBindings(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClusterRoleBindingsApiAPI.ListClusterRoleBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClusterRoleBindings`: ClusterRoleBindingListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ClusterRoleBindingsApiAPI.ListClusterRoleBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterRoleBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterRoleBindingListResponse**](ClusterRoleBindingListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowClusterRoleBinding

> ClusterRoleBindingShowResponse ShowClusterRoleBinding(ctx, clusterId, clusterRoleBindingName).Execute()

Show Cluster RoleBinding



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
	clusterRoleBindingName := "clusterRoleBindingName_example" // string | Cluster Role Binding Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1ClusterRoleBindingsApiAPI.ShowClusterRoleBinding(context.Background(), clusterId, clusterRoleBindingName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ClusterRoleBindingsApiAPI.ShowClusterRoleBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowClusterRoleBinding`: ClusterRoleBindingShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ClusterRoleBindingsApiAPI.ShowClusterRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**clusterRoleBindingName** | **string** | Cluster Role Binding Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowClusterRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClusterRoleBindingShowResponse**](ClusterRoleBindingShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

