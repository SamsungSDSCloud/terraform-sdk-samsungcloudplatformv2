# \LoadbalancerV1LoadbalancersApiAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckLoadbalancerCreation**](LoadbalancerV1LoadbalancersApiAPI.md#CheckLoadbalancerCreation) | **Get** /v1/loadbalancers/{subnet_id}/check-lb-create-available | CheckLoadbalancerCreation
[**CreateLoadbalancer**](LoadbalancerV1LoadbalancersApiAPI.md#CreateLoadbalancer) | **Post** /v1/loadbalancers | CreateLoadbalancer
[**CreateLoadbalancerPublicNatIp**](LoadbalancerV1LoadbalancersApiAPI.md#CreateLoadbalancerPublicNatIp) | **Post** /v1/loadbalancers/{loadbalancer_id}/static-nats | CreateLoadbalancerPublicNatIp
[**DeleteLoadbalancer**](LoadbalancerV1LoadbalancersApiAPI.md#DeleteLoadbalancer) | **Delete** /v1/loadbalancers/{loadbalancer_id} | DeleteLoadbalancer
[**DeleteLoadbalancerPublicNatIp**](LoadbalancerV1LoadbalancersApiAPI.md#DeleteLoadbalancerPublicNatIp) | **Delete** /v1/loadbalancers/{loadbalancer_id}/static-nats | DeleteLoadbalancerPublicNatIp
[**ListLoadbalancers**](LoadbalancerV1LoadbalancersApiAPI.md#ListLoadbalancers) | **Get** /v1/loadbalancers | ListLoadbalancers
[**SetLoadbalancer**](LoadbalancerV1LoadbalancersApiAPI.md#SetLoadbalancer) | **Put** /v1/loadbalancers/{loadbalancer_id} | SetLoadbalancer
[**ShowLoadbalancer**](LoadbalancerV1LoadbalancersApiAPI.md#ShowLoadbalancer) | **Get** /v1/loadbalancers/{loadbalancer_id} | ShowLoadbalancer



## CheckLoadbalancerCreation

> LoadbalancerCreateValidateResponse CheckLoadbalancerCreation(ctx, subnetId).Execute()

CheckLoadbalancerCreation



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
	subnetId := "subnetId_example" // string | Subnet ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LoadbalancersApiAPI.CheckLoadbalancerCreation(context.Background(), subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LoadbalancersApiAPI.CheckLoadbalancerCreation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckLoadbalancerCreation`: LoadbalancerCreateValidateResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LoadbalancersApiAPI.CheckLoadbalancerCreation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | Subnet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckLoadbalancerCreationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoadbalancerCreateValidateResponse**](LoadbalancerCreateValidateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoadbalancer

> LoadbalancerShowResponse CreateLoadbalancer(ctx).LoadbalancerCreateRequest(loadbalancerCreateRequest).Execute()

CreateLoadbalancer



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
	loadbalancerCreateRequest := *openapiclient.NewLoadbalancerCreateRequest(*openapiclient.NewLoadbalancerCreateRequestDetail("LayerType_example", "Name_example", "SubnetId_example", "VpcId_example")) // LoadbalancerCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LoadbalancersApiAPI.CreateLoadbalancer(context.Background()).LoadbalancerCreateRequest(loadbalancerCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LoadbalancersApiAPI.CreateLoadbalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoadbalancer`: LoadbalancerShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LoadbalancersApiAPI.CreateLoadbalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadbalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loadbalancerCreateRequest** | [**LoadbalancerCreateRequest**](LoadbalancerCreateRequest.md) |  | 

### Return type

[**LoadbalancerShowResponse**](LoadbalancerShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoadbalancerPublicNatIp

> StaticNatCreateResponse CreateLoadbalancerPublicNatIp(ctx, loadbalancerId).StaticNatCreateRequest(staticNatCreateRequest).Execute()

CreateLoadbalancerPublicNatIp



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
	loadbalancerId := "loadbalancerId_example" // string | ID
	staticNatCreateRequest := *openapiclient.NewStaticNatCreateRequest(*openapiclient.NewStaticNatCreateRequestDetail("PublicipId_example")) // StaticNatCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LoadbalancersApiAPI.CreateLoadbalancerPublicNatIp(context.Background(), loadbalancerId).StaticNatCreateRequest(staticNatCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LoadbalancersApiAPI.CreateLoadbalancerPublicNatIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoadbalancerPublicNatIp`: StaticNatCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LoadbalancersApiAPI.CreateLoadbalancerPublicNatIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadbalancerId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadbalancerPublicNatIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **staticNatCreateRequest** | [**StaticNatCreateRequest**](StaticNatCreateRequest.md) |  | 

### Return type

[**StaticNatCreateResponse**](StaticNatCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadbalancer

> DeleteLoadbalancer(ctx, loadbalancerId).Execute()

DeleteLoadbalancer



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
	loadbalancerId := "loadbalancerId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadbalancerV1LoadbalancersApiAPI.DeleteLoadbalancer(context.Background(), loadbalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LoadbalancersApiAPI.DeleteLoadbalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadbalancerId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadbalancerRequest struct via the builder pattern


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


## DeleteLoadbalancerPublicNatIp

> DeleteLoadbalancerPublicNatIp(ctx, loadbalancerId).Execute()

DeleteLoadbalancerPublicNatIp



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
	loadbalancerId := "loadbalancerId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadbalancerV1LoadbalancersApiAPI.DeleteLoadbalancerPublicNatIp(context.Background(), loadbalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LoadbalancersApiAPI.DeleteLoadbalancerPublicNatIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadbalancerId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadbalancerPublicNatIpRequest struct via the builder pattern


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


## ListLoadbalancers

> LoadbalancerListResponse ListLoadbalancers(ctx).Size(size).Page(page).Sort(sort).Name(name).State(state).ServiceIp(serviceIp).SubnetId(subnetId).Execute()

ListLoadbalancers



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
	name := "name_example" // string | The name of the load balancer. (optional)
	state := *openapiclient.NewState() // State | The state of the load balancer. (optional) (default to )
	serviceIp := "serviceIp_example" // string | The service IP of the load balancer. (optional)
	subnetId := "subnetId_example" // string | Subnet ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LoadbalancersApiAPI.ListLoadbalancers(context.Background()).Size(size).Page(page).Sort(sort).Name(name).State(state).ServiceIp(serviceIp).SubnetId(subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LoadbalancersApiAPI.ListLoadbalancers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLoadbalancers`: LoadbalancerListResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LoadbalancersApiAPI.ListLoadbalancers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLoadbalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | The name of the load balancer. | 
 **state** | [**State**](State.md) | The state of the load balancer. | [default to ]
 **serviceIp** | **string** | The service IP of the load balancer. | 
 **subnetId** | **string** | Subnet ID | 

### Return type

[**LoadbalancerListResponse**](LoadbalancerListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLoadbalancer

> LoadbalancerShowResponse SetLoadbalancer(ctx, loadbalancerId).LoadbalancerUpdateRequest(loadbalancerUpdateRequest).Execute()

SetLoadbalancer



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
	loadbalancerId := "loadbalancerId_example" // string | ID
	loadbalancerUpdateRequest := *openapiclient.NewLoadbalancerUpdateRequest(*openapiclient.NewLoadbalancerUpdateRequestDetail("Description_example")) // LoadbalancerUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LoadbalancersApiAPI.SetLoadbalancer(context.Background(), loadbalancerId).LoadbalancerUpdateRequest(loadbalancerUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LoadbalancersApiAPI.SetLoadbalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetLoadbalancer`: LoadbalancerShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LoadbalancersApiAPI.SetLoadbalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadbalancerId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLoadbalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadbalancerUpdateRequest** | [**LoadbalancerUpdateRequest**](LoadbalancerUpdateRequest.md) |  | 

### Return type

[**LoadbalancerShowResponse**](LoadbalancerShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowLoadbalancer

> LoadbalancerShowResponse ShowLoadbalancer(ctx, loadbalancerId).Execute()

ShowLoadbalancer



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
	loadbalancerId := "loadbalancerId_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadbalancerV1LoadbalancersApiAPI.ShowLoadbalancer(context.Background(), loadbalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadbalancerV1LoadbalancersApiAPI.ShowLoadbalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowLoadbalancer`: LoadbalancerShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadbalancerV1LoadbalancersApiAPI.ShowLoadbalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadbalancerId** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowLoadbalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoadbalancerShowResponse**](LoadbalancerShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

