# \VpcV1TransitGatewayVpcConnectionApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransitGatewayVpcConnection**](VpcV1TransitGatewayVpcConnectionApiAPI.md#CreateTransitGatewayVpcConnection) | **Post** /v1/transit-gateways/{transit_gateway_id}/vpc-connections | Create Transit Gateway VPC Connection
[**DeleteTransitGatewayVpcConnection**](VpcV1TransitGatewayVpcConnectionApiAPI.md#DeleteTransitGatewayVpcConnection) | **Delete** /v1/transit-gateways/{transit_gateway_id}/vpc-connections/{vpc_connection_id} | Delete Transit Gateway VPC Connection
[**ListTransitGatewayVpcConnections**](VpcV1TransitGatewayVpcConnectionApiAPI.md#ListTransitGatewayVpcConnections) | **Get** /v1/transit-gateways/{transit_gateway_id}/vpc-connections | List Transit Gateway VPC Connections



## CreateTransitGatewayVpcConnection

> TransitGatewayVpcConnectionShowResponse CreateTransitGatewayVpcConnection(ctx, transitGatewayId).TransitGatewayVpcConnectionCreateRequest(transitGatewayVpcConnectionCreateRequest).Execute()

Create Transit Gateway VPC Connection



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
	transitGatewayVpcConnectionCreateRequest := *openapiclient.NewTransitGatewayVpcConnectionCreateRequest("VpcId_example") // TransitGatewayVpcConnectionCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1TransitGatewayVpcConnectionApiAPI.CreateTransitGatewayVpcConnection(context.Background(), transitGatewayId).TransitGatewayVpcConnectionCreateRequest(transitGatewayVpcConnectionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1TransitGatewayVpcConnectionApiAPI.CreateTransitGatewayVpcConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTransitGatewayVpcConnection`: TransitGatewayVpcConnectionShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1TransitGatewayVpcConnectionApiAPI.CreateTransitGatewayVpcConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitGatewayId** | **string** | Transit Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransitGatewayVpcConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transitGatewayVpcConnectionCreateRequest** | [**TransitGatewayVpcConnectionCreateRequest**](TransitGatewayVpcConnectionCreateRequest.md) |  | 

### Return type

[**TransitGatewayVpcConnectionShowResponse**](TransitGatewayVpcConnectionShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransitGatewayVpcConnection

> DeleteTransitGatewayVpcConnection(ctx, transitGatewayId, vpcConnectionId).Execute()

Delete Transit Gateway VPC Connection



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
	vpcConnectionId := "vpcConnectionId_example" // string | Transit Gateway VPC Connection ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1TransitGatewayVpcConnectionApiAPI.DeleteTransitGatewayVpcConnection(context.Background(), transitGatewayId, vpcConnectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1TransitGatewayVpcConnectionApiAPI.DeleteTransitGatewayVpcConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitGatewayId** | **string** | Transit Gateway ID | 
**vpcConnectionId** | **string** | Transit Gateway VPC Connection ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransitGatewayVpcConnectionRequest struct via the builder pattern


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


## ListTransitGatewayVpcConnections

> TransitGatewayVpcConnectionListResponse ListTransitGatewayVpcConnections(ctx, transitGatewayId).Size(size).Page(page).Sort(sort).Id(id).VpcId(vpcId).VpcName(vpcName).State(state).Execute()

List Transit Gateway VPC Connections



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
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	id := "id_example" // string | Transit Gateway VPC Connection ID (optional)
	vpcId := "vpcId_example" // string | VPC Id (optional)
	vpcName := "vpcName_example" // string | VPC Name (optional)
	state := openapiclient.TransitGatewayVpcConnectionState("CREATING") // TransitGatewayVpcConnectionState | State (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1TransitGatewayVpcConnectionApiAPI.ListTransitGatewayVpcConnections(context.Background(), transitGatewayId).Size(size).Page(page).Sort(sort).Id(id).VpcId(vpcId).VpcName(vpcName).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1TransitGatewayVpcConnectionApiAPI.ListTransitGatewayVpcConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransitGatewayVpcConnections`: TransitGatewayVpcConnectionListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1TransitGatewayVpcConnectionApiAPI.ListTransitGatewayVpcConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitGatewayId** | **string** | Transit Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransitGatewayVpcConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **id** | **string** | Transit Gateway VPC Connection ID | 
 **vpcId** | **string** | VPC Id | 
 **vpcName** | **string** | VPC Name | 
 **state** | [**TransitGatewayVpcConnectionState**](TransitGatewayVpcConnectionState.md) | State | 

### Return type

[**TransitGatewayVpcConnectionListResponse**](TransitGatewayVpcConnectionListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

