# \LoadbalancerV1LbListenersApiAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLbListener**](LoadbalancerV1LbListenersApiAPI.md#CreateLbListener) | **Post** /v1/lb-listeners | CreateLbListener
[**DeleteLbListener**](LoadbalancerV1LbListenersApiAPI.md#DeleteLbListener) | **Delete** /v1/lb-listeners/{listener_id} | DeleteLbListener
[**ListLbListeners**](LoadbalancerV1LbListenersApiAPI.md#ListLbListeners) | **Get** /v1/lb-listeners | ListLbListeners
[**SetLbListener**](LoadbalancerV1LbListenersApiAPI.md#SetLbListener) | **Put** /v1/lb-listeners/{listener_id} | SetLbListener
[**ShowLbListener**](LoadbalancerV1LbListenersApiAPI.md#ShowLbListener) | **Get** /v1/lb-listeners/{listener_id} | ShowLbListener



## CreateLbListener

> LbListenerShowResponse CreateLbListener(ctx).LbListenerCreateRequest(lbListenerCreateRequest).Execute()

CreateLbListener



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
	lbListenerCreateRequest := *openapiclient.NewLbListenerCreateRequest(*openapiclient.NewListenerForCreate("LoadbalancerId_example", "Name_example", "Protocol_example", int32(123), int32(123))) // LbListenerCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LbListenersApiAPI.CreateLbListener(context.Background()).LbListenerCreateRequest(lbListenerCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LbListenersApiAPI.CreateLbListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLbListener`: LbListenerShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LbListenersApiAPI.CreateLbListener`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLbListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lbListenerCreateRequest** | [**LbListenerCreateRequest**](LbListenerCreateRequest.md) |  | 

### Return type

[**LbListenerShowResponse**](LbListenerShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLbListener

> DeleteLbListener(ctx, listenerId).Execute()

DeleteLbListener



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
	listenerId := "listenerId_example" // string | The ID of the listener.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadbalancerV1LbListenersApiAPI.DeleteLbListener(context.Background(), listenerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LbListenersApiAPI.DeleteLbListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listenerId** | **string** | The ID of the listener. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLbListenerRequest struct via the builder pattern


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


## ListLbListeners

> LbListenerListResponse ListLbListeners(ctx).Size(size).Page(page).Sort(sort).LoadbalancerId(loadbalancerId).State(state).Name(name).ServicePort(servicePort).Protocol(protocol).Execute()

ListLbListeners



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
	loadbalancerId := "loadbalancerId_example" // string | The ID of the load balancer. (optional)
	state := "state_example" // string | The State of the listener. (optional)
	name := "name_example" // string | The Name of the listener. (optional)
	servicePort := int32(56) // int32 | The Port of the listener. (optional)
	protocol := *openapiclient.NewProtocol1() // Protocol1 | Protocol (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LbListenersApiAPI.ListLbListeners(context.Background()).Size(size).Page(page).Sort(sort).LoadbalancerId(loadbalancerId).State(state).Name(name).ServicePort(servicePort).Protocol(protocol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LbListenersApiAPI.ListLbListeners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLbListeners`: LbListenerListResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LbListenersApiAPI.ListLbListeners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLbListenersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **loadbalancerId** | **string** | The ID of the load balancer. | 
 **state** | **string** | The State of the listener. | 
 **name** | **string** | The Name of the listener. | 
 **servicePort** | **int32** | The Port of the listener. | 
 **protocol** | [**Protocol1**](Protocol1.md) | Protocol | 

### Return type

[**LbListenerListResponse**](LbListenerListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLbListener

> LbListenerShowResponse SetLbListener(ctx, listenerId).LbListenerSetRequest(lbListenerSetRequest).Execute()

SetLbListener



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
	listenerId := "listenerId_example" // string | The ID of the listener.
	lbListenerSetRequest := *openapiclient.NewLbListenerSetRequest(*openapiclient.NewListenerForSet()) // LbListenerSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LbListenersApiAPI.SetLbListener(context.Background(), listenerId).LbListenerSetRequest(lbListenerSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LbListenersApiAPI.SetLbListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetLbListener`: LbListenerShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LbListenersApiAPI.SetLbListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listenerId** | **string** | The ID of the listener. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLbListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lbListenerSetRequest** | [**LbListenerSetRequest**](LbListenerSetRequest.md) |  | 

### Return type

[**LbListenerShowResponse**](LbListenerShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowLbListener

> LbListenerShowResponse ShowLbListener(ctx, listenerId).Execute()

ShowLbListener



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
	listenerId := "listenerId_example" // string | The ID of the listener.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LbListenersApiAPI.ShowLbListener(context.Background(), listenerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LbListenersApiAPI.ShowLbListener``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowLbListener`: LbListenerShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LbListenersApiAPI.ShowLbListener`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listenerId** | **string** | The ID of the listener. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowLbListenerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LbListenerShowResponse**](LbListenerShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

