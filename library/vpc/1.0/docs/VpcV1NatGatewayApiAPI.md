# \VpcV1NatGatewayApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNatGateway**](VpcV1NatGatewayApiAPI.md#CreateNatGateway) | **Post** /v1/nat-gateways | Create NAT Gateway
[**DeleteNatGateway**](VpcV1NatGatewayApiAPI.md#DeleteNatGateway) | **Delete** /v1/nat-gateways/{nat_gateway_id} | Delete NAT Gateway
[**ListNatGateways**](VpcV1NatGatewayApiAPI.md#ListNatGateways) | **Get** /v1/nat-gateways | List NAT Gateways
[**SetNatGateway**](VpcV1NatGatewayApiAPI.md#SetNatGateway) | **Put** /v1/nat-gateways/{nat_gateway_id} | Set NAT Gateway
[**ShowNatGateway**](VpcV1NatGatewayApiAPI.md#ShowNatGateway) | **Get** /v1/nat-gateways/{nat_gateway_id} | Show NAT Gateway



## CreateNatGateway

> NatGatewayShowResponse CreateNatGateway(ctx).NatGatewayCreateRequest(natGatewayCreateRequest).Execute()

Create NAT Gateway



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
	natGatewayCreateRequest := *openapiclient.NewNatGatewayCreateRequest("PublicipId_example", "SubnetId_example") // NatGatewayCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1NatGatewayApiAPI.CreateNatGateway(context.Background()).NatGatewayCreateRequest(natGatewayCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1NatGatewayApiAPI.CreateNatGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNatGateway`: NatGatewayShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1NatGatewayApiAPI.CreateNatGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNatGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **natGatewayCreateRequest** | [**NatGatewayCreateRequest**](NatGatewayCreateRequest.md) |  | 

### Return type

[**NatGatewayShowResponse**](NatGatewayShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNatGateway

> DeleteNatGateway(ctx, natGatewayId).Execute()

Delete NAT Gateway



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
	natGatewayId := "natGatewayId_example" // string | NAT Gateway ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1NatGatewayApiAPI.DeleteNatGateway(context.Background(), natGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1NatGatewayApiAPI.DeleteNatGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**natGatewayId** | **string** | NAT Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNatGatewayRequest struct via the builder pattern


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


## ListNatGateways

> NatGatewayListResponse ListNatGateways(ctx).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Name(name).VpcId(vpcId).VpcName(vpcName).SubnetId(subnetId).SubnetName(subnetName).NatGatewayIpAddress(natGatewayIpAddress).State(state).Execute()

List NAT Gateways



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
	withCount := "true" // string | with count (optional)
	limit := int32(20) // int32 | limit (optional)
	marker := "607e0938521643b5b4b266f343fae693" // string | marker (optional)
	sort := "created_at:desc" // string | sort (optional)
	name := "name_example" // string | NAT Gateway Name (optional)
	vpcId := "vpcId_example" // string | VPC Id (optional)
	vpcName := "vpcName_example" // string | VPC Name (optional)
	subnetId := "subnetId_example" // string | Subnet Id (optional)
	subnetName := "subnetName_example" // string | Subnet Name (optional)
	natGatewayIpAddress := "natGatewayIpAddress_example" // string | NAT Gateway IP Address (optional)
	state := openapiclient.NatGatewayState("CREATING") // NatGatewayState | NAT Gateway State (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1NatGatewayApiAPI.ListNatGateways(context.Background()).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Name(name).VpcId(vpcId).VpcName(vpcName).SubnetId(subnetId).SubnetName(subnetName).NatGatewayIpAddress(natGatewayIpAddress).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1NatGatewayApiAPI.ListNatGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNatGateways`: NatGatewayListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1NatGatewayApiAPI.ListNatGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNatGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 
 **name** | **string** | NAT Gateway Name | 
 **vpcId** | **string** | VPC Id | 
 **vpcName** | **string** | VPC Name | 
 **subnetId** | **string** | Subnet Id | 
 **subnetName** | **string** | Subnet Name | 
 **natGatewayIpAddress** | **string** | NAT Gateway IP Address | 
 **state** | [**NatGatewayState**](NatGatewayState.md) | NAT Gateway State | 

### Return type

[**NatGatewayListResponse**](NatGatewayListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNatGateway

> NatGatewayShowResponse SetNatGateway(ctx, natGatewayId).NatGatewaySetRequest(natGatewaySetRequest).Execute()

Set NAT Gateway



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
	natGatewayId := "natGatewayId_example" // string | NAT Gateway ID
	natGatewaySetRequest := *openapiclient.NewNatGatewaySetRequest() // NatGatewaySetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1NatGatewayApiAPI.SetNatGateway(context.Background(), natGatewayId).NatGatewaySetRequest(natGatewaySetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1NatGatewayApiAPI.SetNatGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNatGateway`: NatGatewayShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1NatGatewayApiAPI.SetNatGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**natGatewayId** | **string** | NAT Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNatGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **natGatewaySetRequest** | [**NatGatewaySetRequest**](NatGatewaySetRequest.md) |  | 

### Return type

[**NatGatewayShowResponse**](NatGatewayShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowNatGateway

> NatGatewayShowResponse ShowNatGateway(ctx, natGatewayId).Execute()

Show NAT Gateway



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
	natGatewayId := "natGatewayId_example" // string | NAT Gateway ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1NatGatewayApiAPI.ShowNatGateway(context.Background(), natGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1NatGatewayApiAPI.ShowNatGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowNatGateway`: NatGatewayShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1NatGatewayApiAPI.ShowNatGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**natGatewayId** | **string** | NAT Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowNatGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NatGatewayShowResponse**](NatGatewayShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

