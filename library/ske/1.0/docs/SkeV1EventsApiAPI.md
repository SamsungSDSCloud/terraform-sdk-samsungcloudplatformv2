# \SkeV1EventsApiAPI

All URIs are relative to *https://scp-kubernetes-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEvents**](SkeV1EventsApiAPI.md#ListEvents) | **Get** /v1/clusters/{cluster_id}/events | List Events
[**ListNamespacedEvents**](SkeV1EventsApiAPI.md#ListNamespacedEvents) | **Get** /v1/clusters/{cluster_id}/namespaces/{namespace_name}/events | List Namespaced Events



## ListEvents

> EventListResponse ListEvents(ctx, clusterId).Execute()

List Events



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
	resp, r, err := apiClient.SkeV1EventsApiAPI.ListEvents(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1EventsApiAPI.ListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEvents`: EventListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1EventsApiAPI.ListEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventListResponse**](EventListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespacedEvents

> EventListResponse ListNamespacedEvents(ctx, clusterId, namespaceName).Execute()

List Namespaced Events



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
	resp, r, err := apiClient.SkeV1EventsApiAPI.ListNamespacedEvents(context.Background(), clusterId, namespaceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkeV1EventsApiAPI.ListNamespacedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespacedEvents`: EventListResponse
	fmt.Fprintf(os.Stdout, "Response from `SkeV1EventsApiAPI.ListNamespacedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**namespaceName** | **string** | Namespace Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EventListResponse**](EventListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

