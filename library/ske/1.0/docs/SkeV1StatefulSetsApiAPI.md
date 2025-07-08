# \SkeV1StatefulSetsApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteStatefulSet**](SkeV1StatefulSetsApiAPI.md#DeleteStatefulSet) | **Delete** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/statefulsets | Delete Stateful Set
[**ListStatefulSets**](SkeV1StatefulSetsApiAPI.md#ListStatefulSets) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/statefulsets | List Stateful Sets
[**ShowStatefulSet**](SkeV1StatefulSetsApiAPI.md#ShowStatefulSet) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/statefulsets/{stateful_set_name} | Show Stateful Set



## DeleteStatefulSet

> DeleteStatefulSet(ctx, clusterId, namespaceName).Name(name).Execute()

Delete Stateful Set



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
	r, err := apiClient.SkeV1StatefulSetsApiAPI.DeleteStatefulSet(context.Background(), clusterId, namespaceName).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1StatefulSetsApiAPI.DeleteStatefulSet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteStatefulSetRequest struct via the builder pattern


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


## ListStatefulSets

> StatefulSetListResponse ListStatefulSets(ctx, clusterId, namespaceName).Size(size).Page(page).Sort(sort).Name(name).SystemObject(systemObject).Execute()

List Stateful Sets



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
	name := "name_example" // string | Stateful Set Name (optional)
	systemObject := true // bool | System Object (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1StatefulSetsApiAPI.ListStatefulSets(context.Background(), clusterId, namespaceName).Size(size).Page(page).Sort(sort).Name(name).SystemObject(systemObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1StatefulSetsApiAPI.ListStatefulSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStatefulSets`: StatefulSetListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1StatefulSetsApiAPI.ListStatefulSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStatefulSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Stateful Set Name | 
 **systemObject** | **bool** | System Object | 

### Return type

[**StatefulSetListResponse**](StatefulSetListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowStatefulSet

> StatefulSetShowResponse ShowStatefulSet(ctx, clusterId, namespaceName, statefulSetName).Execute()

Show Stateful Set



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
	statefulSetName := "statefulSetName_example" // string | Stateful Set Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1StatefulSetsApiAPI.ShowStatefulSet(context.Background(), clusterId, namespaceName, statefulSetName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1StatefulSetsApiAPI.ShowStatefulSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowStatefulSet`: StatefulSetShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1StatefulSetsApiAPI.ShowStatefulSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 
**statefulSetName** | **string** | Stateful Set Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowStatefulSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StatefulSetShowResponse**](StatefulSetShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

