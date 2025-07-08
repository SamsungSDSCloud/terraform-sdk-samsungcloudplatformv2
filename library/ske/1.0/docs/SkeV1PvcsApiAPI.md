# \SkeV1PvcsApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePvc**](SkeV1PvcsApiAPI.md#DeletePvc) | **Delete** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/persistentvolumeclaims | Delete PVC
[**ListPvcs**](SkeV1PvcsApiAPI.md#ListPvcs) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/persistentvolumeclaims | List PVCs
[**ShowPvc**](SkeV1PvcsApiAPI.md#ShowPvc) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/persistentvolumeclaims/{pvc_name} | Show PVC



## DeletePvc

> DeletePvc(ctx, clusterId, namespaceName).Execute()

Delete PVC



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SkeV1PvcsApiAPI.DeletePvc(context.Background(), clusterId, namespaceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1PvcsApiAPI.DeletePvc``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePvcRequest struct via the builder pattern


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


## ListPvcs

> PvcListResponse ListPvcs(ctx, clusterId, namespaceName).Execute()

List PVCs



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1PvcsApiAPI.ListPvcs(context.Background(), clusterId, namespaceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1PvcsApiAPI.ListPvcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPvcs`: PvcListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1PvcsApiAPI.ListPvcs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPvcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PvcListResponse**](PvcListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPvc

> PvcShowResponse ShowPvc(ctx, clusterId, namespaceName, pvcName).Execute()

Show PVC



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
	pvcName := "pvcName_example" // string | PVC Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1PvcsApiAPI.ShowPvc(context.Background(), clusterId, namespaceName, pvcName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1PvcsApiAPI.ShowPvc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowPvc`: PvcShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1PvcsApiAPI.ShowPvc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 
**pvcName** | **string** | PVC Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowPvcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PvcShowResponse**](PvcShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

