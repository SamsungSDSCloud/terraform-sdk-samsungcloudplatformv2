# \VpnV1VpnGatewaysApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVpnGateway**](VpnV1VpnGatewaysApiAPI.md#CreateVpnGateway) | **Post** /v1/vpn-gateways | Create VPN Gateway
[**DeleteVpnGateway**](VpnV1VpnGatewaysApiAPI.md#DeleteVpnGateway) | **Delete** /v1/vpn-gateways/{vpn_gateway_id} | Delete VPN Gateway
[**GetVpnGatewayQuotas**](VpnV1VpnGatewaysApiAPI.md#GetVpnGatewayQuotas) | **Get** /v1/vpn-gateways/metrics/quotas | Get VPN Gateway Quotas
[**ListVpnGateways**](VpnV1VpnGatewaysApiAPI.md#ListVpnGateways) | **Get** /v1/vpn-gateways | List VPN Gateways
[**SetVpnGateway**](VpnV1VpnGatewaysApiAPI.md#SetVpnGateway) | **Put** /v1/vpn-gateways/{vpn_gateway_id} | Set VPN Gateway
[**ShowVpnGateway**](VpnV1VpnGatewaysApiAPI.md#ShowVpnGateway) | **Get** /v1/vpn-gateways/{vpn_gateway_id} | Show VPN Gateway



## CreateVpnGateway

> VpnGatewayShowResponse CreateVpnGateway(ctx).VpnGatewayCreateRequest(vpnGatewayCreateRequest).Execute()

Create VPN Gateway



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
	vpnGatewayCreateRequest := *openapiclient.NewVpnGatewayCreateRequest("123.0.0.1", "fcde872f75c145a0893d656cc698f13e", "PUBLIC", "ExampleVpnGW1", "ceb44ea5ecb34a49b16495f9a63b0718") // VpnGatewayCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnV1VpnGatewaysApiAPI.CreateVpnGateway(context.Background()).VpnGatewayCreateRequest(vpnGatewayCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnGatewaysApiAPI.CreateVpnGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpnGateway`: VpnGatewayShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpnV1VpnGatewaysApiAPI.CreateVpnGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpnGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpnGatewayCreateRequest** | [**VpnGatewayCreateRequest**](VpnGatewayCreateRequest.md) |  | 

### Return type

[**VpnGatewayShowResponse**](VpnGatewayShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpnGateway

> DeleteVpnGateway(ctx, vpnGatewayId).Execute()

Delete VPN Gateway



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
	vpnGatewayId := "b156740b6335468d8354eb9ef8eddf5a" // string | VPN Gateway ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpnV1VpnGatewaysApiAPI.DeleteVpnGateway(context.Background(), vpnGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnGatewaysApiAPI.DeleteVpnGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpnGatewayId** | **string** | VPN Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpnGatewayRequest struct via the builder pattern


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


## GetVpnGatewayQuotas

> VpnGatewayQuotaResponse GetVpnGatewayQuotas(ctx).VpcId(vpcId).Execute()

Get VPN Gateway Quotas



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
	vpcId := "ceb44ea5ecb34a49b16495f9a63b0718" // string | VPC Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnV1VpnGatewaysApiAPI.GetVpnGatewayQuotas(context.Background()).VpcId(vpcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnGatewaysApiAPI.GetVpnGatewayQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnGatewayQuotas`: VpnGatewayQuotaResponse
	fmt.Fprintf(os.Stdout, "Response from `VpnV1VpnGatewaysApiAPI.GetVpnGatewayQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVpnGatewayQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpcId** | **string** | VPC Id | 

### Return type

[**VpnGatewayQuotaResponse**](VpnGatewayQuotaResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVpnGateways

> VpnGatewayListResponse ListVpnGateways(ctx).Size(size).Page(page).Sort(sort).Name(name).IpAddress(ipAddress).VpcId(vpcId).VpcName(vpcName).Execute()

List VPN Gateways



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
	name := "ExampleVpnGW1" // string | VPN Gateway Name (optional)
	ipAddress := "123.0.0.1" // string | VPN Gateway IP Address (optional)
	vpcId := "ceb44ea5ecb34a49b16495f9a63b0718" // string | VPC Id (optional)
	vpcName := "ExampleVPC1" // string | VPC Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnV1VpnGatewaysApiAPI.ListVpnGateways(context.Background()).Size(size).Page(page).Sort(sort).Name(name).IpAddress(ipAddress).VpcId(vpcId).VpcName(vpcName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnGatewaysApiAPI.ListVpnGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpnGateways`: VpnGatewayListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpnV1VpnGatewaysApiAPI.ListVpnGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVpnGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | VPN Gateway Name | 
 **ipAddress** | **string** | VPN Gateway IP Address | 
 **vpcId** | **string** | VPC Id | 
 **vpcName** | **string** | VPC Name | 

### Return type

[**VpnGatewayListResponse**](VpnGatewayListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVpnGateway

> VpnGatewayShowResponse SetVpnGateway(ctx, vpnGatewayId).VpnGatewaySetRequest(vpnGatewaySetRequest).Execute()

Set VPN Gateway



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
	vpnGatewayId := "b156740b6335468d8354eb9ef8eddf5a" // string | VPN Gateway ID
	vpnGatewaySetRequest := *openapiclient.NewVpnGatewaySetRequest() // VpnGatewaySetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnV1VpnGatewaysApiAPI.SetVpnGateway(context.Background(), vpnGatewayId).VpnGatewaySetRequest(vpnGatewaySetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnGatewaysApiAPI.SetVpnGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVpnGateway`: VpnGatewayShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpnV1VpnGatewaysApiAPI.SetVpnGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpnGatewayId** | **string** | VPN Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVpnGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpnGatewaySetRequest** | [**VpnGatewaySetRequest**](VpnGatewaySetRequest.md) |  | 

### Return type

[**VpnGatewayShowResponse**](VpnGatewayShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVpnGateway

> VpnGatewayShowResponse ShowVpnGateway(ctx, vpnGatewayId).Execute()

Show VPN Gateway



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
	vpnGatewayId := "b156740b6335468d8354eb9ef8eddf5a" // string | VPN Gateway ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnV1VpnGatewaysApiAPI.ShowVpnGateway(context.Background(), vpnGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnGatewaysApiAPI.ShowVpnGateway``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVpnGateway`: VpnGatewayShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpnV1VpnGatewaysApiAPI.ShowVpnGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpnGatewayId** | **string** | VPN Gateway ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVpnGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VpnGatewayShowResponse**](VpnGatewayShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

