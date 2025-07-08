# \SkeV1DaemonSetsApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDaemonSet**](SkeV1DaemonSetsApiAPI.md#DeleteDaemonSet) | **Delete** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/daemonsets | Delete Daemon Set
[**ListDaemonSets**](SkeV1DaemonSetsApiAPI.md#ListDaemonSets) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/daemonsets | List Daemon Sets
[**ShowDaemonSet**](SkeV1DaemonSetsApiAPI.md#ShowDaemonSet) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/daemonsets/{daemon_set_name} | Show Daemon Set



## DeleteDaemonSet

> DeleteDaemonSet(ctx, clusterId, namespaceName).Name(name).Execute()

Delete Daemon Set



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
	namespaceName := "namespaceName_example" // string | Namespace Name
	name := *openapiclient.NewName1() // Name1 | Names

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SkeV1DaemonSetsApiAPI.DeleteDaemonSet(context.Background(), clusterId, namespaceName).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1DaemonSetsApiAPI.DeleteDaemonSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDaemonSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | [**Name1**](Name1.md) | Names | 

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


## ListDaemonSets

> DaemonSetListResponse ListDaemonSets(ctx, clusterId, namespaceName).Size(size).Page(page).Sort(sort).Name(name).SystemObject(systemObject).Execute()

List Daemon Sets



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
	namespaceName := "namespaceName_example" // string | Namespace Name
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	name := "name_example" // string | Daemon Set Name (optional)
	systemObject := true // bool | System Object (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1DaemonSetsApiAPI.ListDaemonSets(context.Background(), clusterId, namespaceName).Size(size).Page(page).Sort(sort).Name(name).SystemObject(systemObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1DaemonSetsApiAPI.ListDaemonSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDaemonSets`: DaemonSetListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1DaemonSetsApiAPI.ListDaemonSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDaemonSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Daemon Set Name | 
 **systemObject** | **bool** | System Object | 

### Return type

[**DaemonSetListResponse**](DaemonSetListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowDaemonSet

> DaemonSetShowResponse ShowDaemonSet(ctx, clusterId, namespaceName, daemonSetName).Execute()

Show Daemon Set



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
	namespaceName := "namespaceName_example" // string | Namespace Name
	daemonSetName := "daemonSetName_example" // string | Daemon Set Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1DaemonSetsApiAPI.ShowDaemonSet(context.Background(), clusterId, namespaceName, daemonSetName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1DaemonSetsApiAPI.ShowDaemonSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowDaemonSet`: DaemonSetShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1DaemonSetsApiAPI.ShowDaemonSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 
**daemonSetName** | **string** | Daemon Set Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowDaemonSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DaemonSetShowResponse**](DaemonSetShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

