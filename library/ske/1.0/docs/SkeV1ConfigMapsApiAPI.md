# \SkeV1ConfigMapsApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfigMap**](SkeV1ConfigMapsApiAPI.md#DeleteConfigMap) | **Delete** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/configmaps | Delete Config Map
[**ListConfigMaps**](SkeV1ConfigMapsApiAPI.md#ListConfigMaps) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/configmaps | List Config Maps
[**ShowConfigMap**](SkeV1ConfigMapsApiAPI.md#ShowConfigMap) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/configmaps/{config_map_name} | Show Config Map



## DeleteConfigMap

> DeleteConfigMap(ctx, clusterId, namespaceName).Name(name).Execute()

Delete Config Map



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
	name := *openapiclient.NewName() // Name | Names

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SkeV1ConfigMapsApiAPI.DeleteConfigMap(context.Background(), clusterId, namespaceName).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ConfigMapsApiAPI.DeleteConfigMap``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteConfigMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | [**Name**](Name.md) | Names | 

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


## ListConfigMaps

> ConfigMapListResponse ListConfigMaps(ctx, clusterId, namespaceName).Size(size).Page(page).Sort(sort).Name(name).SystemObject(systemObject).Execute()

List Config Maps



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
	name := "name_example" // string | Config Map Name (optional)
	systemObject := true // bool | System Object (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1ConfigMapsApiAPI.ListConfigMaps(context.Background(), clusterId, namespaceName).Size(size).Page(page).Sort(sort).Name(name).SystemObject(systemObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ConfigMapsApiAPI.ListConfigMaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConfigMaps`: ConfigMapListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ConfigMapsApiAPI.ListConfigMaps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConfigMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Config Map Name | 
 **systemObject** | **bool** | System Object | 

### Return type

[**ConfigMapListResponse**](ConfigMapListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowConfigMap

> ConfigMapShowResponse ShowConfigMap(ctx, clusterId, namespaceName, configMapName).Execute()

Show Config Map



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
	configMapName := "configMapName_example" // string | Config Map Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1ConfigMapsApiAPI.ShowConfigMap(context.Background(), clusterId, namespaceName, configMapName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1ConfigMapsApiAPI.ShowConfigMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowConfigMap`: ConfigMapShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1ConfigMapsApiAPI.ShowConfigMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 
**configMapName** | **string** | Config Map Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowConfigMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ConfigMapShowResponse**](ConfigMapShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

