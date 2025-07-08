# \SkeV1IngressesApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIngress**](SkeV1IngressesApiAPI.md#DeleteIngress) | **Delete** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/ingresses | Delete Ingress
[**ListIngresses**](SkeV1IngressesApiAPI.md#ListIngresses) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/ingresses | List Ingresses
[**ShowIngress**](SkeV1IngressesApiAPI.md#ShowIngress) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/ingresses/{ingress_name} | Show Ingress



## DeleteIngress

> DeleteIngress(ctx, clusterId, namespaceName).Execute()

Delete Ingress



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
	r, err := apiClient.SkeV1IngressesApiAPI.DeleteIngress(context.Background(), clusterId, namespaceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1IngressesApiAPI.DeleteIngress``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIngressRequest struct via the builder pattern


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


## ListIngresses

> IngressListResponse ListIngresses(ctx, clusterId, namespaceName).Execute()

List Ingresses



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
	resp, r, err := apiClient.SkeV1IngressesApiAPI.ListIngresses(context.Background(), clusterId, namespaceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1IngressesApiAPI.ListIngresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIngresses`: IngressListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1IngressesApiAPI.ListIngresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIngressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IngressListResponse**](IngressListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowIngress

> IngressShowResponse ShowIngress(ctx, clusterId, namespaceName, ingressName).Execute()

Show Ingress



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
	ingressName := "ingressName_example" // string | Ingress Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkeV1IngressesApiAPI.ShowIngress(context.Background(), clusterId, namespaceName, ingressName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1IngressesApiAPI.ShowIngress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowIngress`: IngressShowResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1IngressesApiAPI.ShowIngress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 
**ingressName** | **string** | Ingress Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowIngressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IngressShowResponse**](IngressShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

