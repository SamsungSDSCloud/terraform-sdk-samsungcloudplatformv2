# \VpnV1VpnTunnelsApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckVpnTunnelSubnetAvailabilities**](VpnV1VpnTunnelsApiAPI.md#CheckVpnTunnelSubnetAvailabilities) | **Get** /v1/vpn-tunnels/vpn-tunnel-subnet-availabilities | Check VPN Tunnel Subnet Availabilities
[**CreateVpnTunnel**](VpnV1VpnTunnelsApiAPI.md#CreateVpnTunnel) | **Post** /v1/vpn-tunnels | Create VPN Tunnel
[**DeleteVpnTunnel**](VpnV1VpnTunnelsApiAPI.md#DeleteVpnTunnel) | **Delete** /v1/vpn-tunnels/{vpn_tunnel_id} | Delete VPN Tunnel
[**GetVpnTunnelQuotas**](VpnV1VpnTunnelsApiAPI.md#GetVpnTunnelQuotas) | **Get** /v1/vpn-tunnels/metrics/quotas | Get VPN Tunnel Quotas
[**ListVpnTunnels**](VpnV1VpnTunnelsApiAPI.md#ListVpnTunnels) | **Get** /v1/vpn-tunnels | List VPN Tunnels
[**SetVpnTunnel**](VpnV1VpnTunnelsApiAPI.md#SetVpnTunnel) | **Put** /v1/vpn-tunnels/{vpn_tunnel_id} | Set VPN Tunnel
[**ShowVpnTunnel**](VpnV1VpnTunnelsApiAPI.md#ShowVpnTunnel) | **Get** /v1/vpn-tunnels/{vpn_tunnel_id} | Show VPN Tunnel



## CheckVpnTunnelSubnetAvailabilities

> VpnTunnelSubnetAvailabilityResponse CheckVpnTunnelSubnetAvailabilities(ctx).VpnGatewayId(vpnGatewayId).RemoteSubnet(remoteSubnet).Execute()

Check VPN Tunnel Subnet Availabilities



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
	remoteSubnet := "10.1.0.0/16" // string | VPN Tunnel IPSec Remote Subnet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnV1VpnTunnelsApiAPI.CheckVpnTunnelSubnetAvailabilities(context.Background()).VpnGatewayId(vpnGatewayId).RemoteSubnet(remoteSubnet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnTunnelsApiAPI.CheckVpnTunnelSubnetAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckVpnTunnelSubnetAvailabilities`: VpnTunnelSubnetAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `VpnV1VpnTunnelsApiAPI.CheckVpnTunnelSubnetAvailabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckVpnTunnelSubnetAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpnGatewayId** | **string** | VPN Gateway ID | 
 **remoteSubnet** | **string** | VPN Tunnel IPSec Remote Subnet | 

### Return type

[**VpnTunnelSubnetAvailabilityResponse**](VpnTunnelSubnetAvailabilityResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVpnTunnel

> VpnTunnelShowResponse CreateVpnTunnel(ctx).VpnTunnelCreateRequest(vpnTunnelCreateRequest).Execute()

Create VPN Tunnel



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
	vpnTunnelCreateRequest := *openapiclient.NewVpnTunnelCreateRequest("ExampleVpnTunnel1", *openapiclient.NewVpnPhase1CreateRequest(int32(60), int32(2), "123.0.0.2", []int32{int32(123)}, []string{"Phase1Encryptions_example"}, int32(86400), "PreSharedKey1"), *openapiclient.NewVpnPhase2CreateRequest(openapiclient.VpnPerfectForwardSecrecyType("ENABLE"), []int32{int32(123)}, []string{"Phase2Encryptions_example"}, int32(43200), "10.1.0.0/16"), "b156740b6335468d8354eb9ef8eddf5a") // VpnTunnelCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnV1VpnTunnelsApiAPI.CreateVpnTunnel(context.Background()).VpnTunnelCreateRequest(vpnTunnelCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnTunnelsApiAPI.CreateVpnTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpnTunnel`: VpnTunnelShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpnV1VpnTunnelsApiAPI.CreateVpnTunnel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpnTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpnTunnelCreateRequest** | [**VpnTunnelCreateRequest**](VpnTunnelCreateRequest.md) |  | 

### Return type

[**VpnTunnelShowResponse**](VpnTunnelShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpnTunnel

> DeleteVpnTunnel(ctx, vpnTunnelId).Execute()

Delete VPN Tunnel



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
	vpnTunnelId := "2ade2919287040bc92e557fbbe6709c1" // string | VPN Tunnel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpnV1VpnTunnelsApiAPI.DeleteVpnTunnel(context.Background(), vpnTunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnTunnelsApiAPI.DeleteVpnTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpnTunnelId** | **string** | VPN Tunnel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpnTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpnTunnelQuotas

> VpnTunnelQuotaResponse GetVpnTunnelQuotas(ctx).VpnGatewayId(vpnGatewayId).Execute()

Get VPN Tunnel Quotas



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
	vpnGatewayId := "b156740b6335468d8354eb9ef8eddf5a" // string | VPN Gateway ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnV1VpnTunnelsApiAPI.GetVpnTunnelQuotas(context.Background()).VpnGatewayId(vpnGatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnTunnelsApiAPI.GetVpnTunnelQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnTunnelQuotas`: VpnTunnelQuotaResponse
	fmt.Fprintf(os.Stdout, "Response from `VpnV1VpnTunnelsApiAPI.GetVpnTunnelQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVpnTunnelQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpnGatewayId** | **string** | VPN Gateway ID | 

### Return type

[**VpnTunnelQuotaResponse**](VpnTunnelQuotaResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVpnTunnels

> VpnTunnelListResponse ListVpnTunnels(ctx).Size(size).Page(page).Sort(sort).Name(name).VpnGatewayId(vpnGatewayId).VpnGatewayName(vpnGatewayName).PeerGatewayIp(peerGatewayIp).RemoteSubnet(remoteSubnet).Execute()

List VPN Tunnels



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
	name := "ExampleVpnTunnel1" // string | VPN Tunnel Name (optional)
	vpnGatewayId := "b156740b6335468d8354eb9ef8eddf5a" // string | VPN Gateway ID (optional)
	vpnGatewayName := "ExampleVpnGW1" // string | VPN Gateway Name (optional)
	peerGatewayIp := "123.0.0.2" // string | VPN Tunnel Peer Gateway IP Address (optional)
	remoteSubnet := "10.1.0.0/16" // string | VPN Tunnel IPSec Remote Subnet (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnV1VpnTunnelsApiAPI.ListVpnTunnels(context.Background()).Size(size).Page(page).Sort(sort).Name(name).VpnGatewayId(vpnGatewayId).VpnGatewayName(vpnGatewayName).PeerGatewayIp(peerGatewayIp).RemoteSubnet(remoteSubnet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnTunnelsApiAPI.ListVpnTunnels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpnTunnels`: VpnTunnelListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpnV1VpnTunnelsApiAPI.ListVpnTunnels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVpnTunnelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | VPN Tunnel Name | 
 **vpnGatewayId** | **string** | VPN Gateway ID | 
 **vpnGatewayName** | **string** | VPN Gateway Name | 
 **peerGatewayIp** | **string** | VPN Tunnel Peer Gateway IP Address | 
 **remoteSubnet** | **string** | VPN Tunnel IPSec Remote Subnet | 

### Return type

[**VpnTunnelListResponse**](VpnTunnelListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVpnTunnel

> VpnTunnelShowResponse SetVpnTunnel(ctx, vpnTunnelId).VpnTunnelSetRequest(vpnTunnelSetRequest).Execute()

Set VPN Tunnel



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
	vpnTunnelId := "2ade2919287040bc92e557fbbe6709c1" // string | VPN Tunnel ID
	vpnTunnelSetRequest := *openapiclient.NewVpnTunnelSetRequest() // VpnTunnelSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnV1VpnTunnelsApiAPI.SetVpnTunnel(context.Background(), vpnTunnelId).VpnTunnelSetRequest(vpnTunnelSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnTunnelsApiAPI.SetVpnTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVpnTunnel`: VpnTunnelShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpnV1VpnTunnelsApiAPI.SetVpnTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpnTunnelId** | **string** | VPN Tunnel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVpnTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpnTunnelSetRequest** | [**VpnTunnelSetRequest**](VpnTunnelSetRequest.md) |  | 

### Return type

[**VpnTunnelShowResponse**](VpnTunnelShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVpnTunnel

> VpnTunnelShowWithStatusResponse ShowVpnTunnel(ctx, vpnTunnelId).Execute()

Show VPN Tunnel



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
	vpnTunnelId := "2ade2919287040bc92e557fbbe6709c1" // string | VPN Tunnel ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnV1VpnTunnelsApiAPI.ShowVpnTunnel(context.Background(), vpnTunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnV1VpnTunnelsApiAPI.ShowVpnTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVpnTunnel`: VpnTunnelShowWithStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `VpnV1VpnTunnelsApiAPI.ShowVpnTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpnTunnelId** | **string** | VPN Tunnel ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVpnTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VpnTunnelShowWithStatusResponse**](VpnTunnelShowWithStatusResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

