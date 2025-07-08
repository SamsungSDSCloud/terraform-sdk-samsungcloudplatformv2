# \VpcV1TransitGatewayApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransitGateway**](VpcV1TransitGatewayApiAPI.md#CreateTransitGateway) | **Post** /v1/transit-gateways | Create Transit Gateway
[**DeleteTransitGateway**](VpcV1TransitGatewayApiAPI.md#DeleteTransitGateway) | **Delete** /v1/transit-gateways/{transit_gateway_id} | Delete Transit Gateway
[**ListTransitGateways**](VpcV1TransitGatewayApiAPI.md#ListTransitGateways) | **Get** /v1/transit-gateways | List Transit Gateways
[**SetTransitGateway**](VpcV1TransitGatewayApiAPI.md#SetTransitGateway) | **Put** /v1/transit-gateways/{transit_gateway_id} | Set Transit Gateway
[**ShowTransitGateway**](VpcV1TransitGatewayApiAPI.md#ShowTransitGateway) | **Get** /v1/transit-gateways/{transit_gateway_id} | Show Transit Gateway



## CreateTransitGateway

> TransitGatewayShowResponse CreateTransitGateway(ctx).TransitGatewayCreateRequest(transitGatewayCreateRequest).Execute()

Create Transit Gateway



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
	transitGatewayCreateRequest := *openapiclient.NewTransitGatewayCreateRequest("Name_example") // TransitGatewayCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1TransitGatewayApiAPI.CreateTransitGateway(context.Background()).TransitGatewayCreateRequest(transitGatewayCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1TransitGatewayApiAPI.CreateTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTransitGateway`: TransitGatewayShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1TransitGatewayApiAPI.CreateTransitGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transitGatewayCreateRequest** | [**TransitGatewayCreateRequest**](TransitGatewayCreateRequest.md) |  | 

### Return type

[**TransitGatewayShowResponse**](TransitGatewayShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransitGateway

> DeleteTransitGateway(ctx, transitGatewayId).Execute()

Delete Transit Gateway



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
	transitGatewayId := "transitGatewayId_example" // string | Transit Gateway ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1TransitGatewayApiAPI.DeleteTransitGateway(context.Background(), transitGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1TransitGatewayApiAPI.DeleteTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitGatewayId** | **string** | Transit Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransitGatewayRequest struct via the builder pattern


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


## ListTransitGateways

> TransitGatewayListResponse ListTransitGateways(ctx).Size(size).Page(page).Sort(sort).Id(id).Name(name).State(state).Execute()

List Transit Gateways



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
	id := "id_example" // string | Transit Gateway ID (optional)
	name := "name_example" // string | Transit Gateway Name (optional)
	state := openapiclient.TransitGatewayState("CREATING") // TransitGatewayState | State (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1TransitGatewayApiAPI.ListTransitGateways(context.Background()).Size(size).Page(page).Sort(sort).Id(id).Name(name).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1TransitGatewayApiAPI.ListTransitGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransitGateways`: TransitGatewayListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1TransitGatewayApiAPI.ListTransitGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransitGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **id** | **string** | Transit Gateway ID | 
 **name** | **string** | Transit Gateway Name | 
 **state** | [**TransitGatewayState**](TransitGatewayState.md) | State | 

### Return type

[**TransitGatewayListResponse**](TransitGatewayListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTransitGateway

> TransitGatewayShowResponse SetTransitGateway(ctx, transitGatewayId).TransitGatewaySetRequest(transitGatewaySetRequest).Execute()

Set Transit Gateway



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
	transitGatewayId := "transitGatewayId_example" // string | Transit Gateway ID
	transitGatewaySetRequest := *openapiclient.NewTransitGatewaySetRequest() // TransitGatewaySetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1TransitGatewayApiAPI.SetTransitGateway(context.Background(), transitGatewayId).TransitGatewaySetRequest(transitGatewaySetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1TransitGatewayApiAPI.SetTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTransitGateway`: TransitGatewayShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1TransitGatewayApiAPI.SetTransitGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitGatewayId** | **string** | Transit Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitGatewaySetRequest** | [**TransitGatewaySetRequest**](TransitGatewaySetRequest.md) |  | 

### Return type

[**TransitGatewayShowResponse**](TransitGatewayShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowTransitGateway

> TransitGatewayShowResponse ShowTransitGateway(ctx, transitGatewayId).Execute()

Show Transit Gateway



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
	transitGatewayId := "transitGatewayId_example" // string | Transit Gateway ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1TransitGatewayApiAPI.ShowTransitGateway(context.Background(), transitGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1TransitGatewayApiAPI.ShowTransitGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowTransitGateway`: TransitGatewayShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1TransitGatewayApiAPI.ShowTransitGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitGatewayId** | **string** | Transit Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowTransitGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransitGatewayShowResponse**](TransitGatewayShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

