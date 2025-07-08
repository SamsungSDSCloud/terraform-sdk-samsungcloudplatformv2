# \FirewallV1FirewallApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFirewalls**](FirewallV1FirewallApiAPI.md#ListFirewalls) | **Get** /v1/firewalls | List Firewalls
[**SetFirewall**](FirewallV1FirewallApiAPI.md#SetFirewall) | **Put** /v1/firewalls/{firewall_id} | Set Firewall
[**SetFirewallStatus**](FirewallV1FirewallApiAPI.md#SetFirewallStatus) | **Put** /v1/firewalls/{firewall_id}/status | Set Firewall Status.
[**ShowFirewall**](FirewallV1FirewallApiAPI.md#ShowFirewall) | **Get** /v1/firewalls/{firewall_id} | Show Firewall



## ListFirewalls

> FirewallListResponse ListFirewalls(ctx).Size(size).Page(page).Sort(sort).Name(name).VpcName(vpcName).ProductType(productType).State(state).Execute()

List Firewalls



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
	name := "FW_IGW_secuVPC" // string | Firewall Name (optional)
	vpcName := "secuVPC" // string | VPC Name (optional)
	productType := []openapiclient.FirewallProductType{openapiclient.FirewallProductType("IGW")} // []FirewallProductType | Firewall Product Type (optional) (default to [])
	state := []openapiclient.FirewallState{openapiclient.FirewallState("CREATING")} // []FirewallState | Firewall State (optional) (default to [])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallV1FirewallApiAPI.ListFirewalls(context.Background()).Size(size).Page(page).Sort(sort).Name(name).VpcName(vpcName).ProductType(productType).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallApiAPI.ListFirewalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFirewalls`: FirewallListResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallV1FirewallApiAPI.ListFirewalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFirewallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Firewall Name | 
 **vpcName** | **string** | VPC Name | 
 **productType** | [**[]FirewallProductType**](FirewallProductType.md) | Firewall Product Type | [default to []]
 **state** | [**[]FirewallState**](FirewallState.md) | Firewall State | [default to []]

### Return type

[**FirewallListResponse**](FirewallListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFirewall

> FirewallShowResponse SetFirewall(ctx, firewallId).FirewallSetRequest(firewallSetRequest).Execute()

Set Firewall



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
	firewallId := "b156740b6335468d8354eb9ef8eddf5a" // string | Firewall ID
	firewallSetRequest := *openapiclient.NewFirewallSetRequest() // FirewallSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallV1FirewallApiAPI.SetFirewall(context.Background(), firewallId).FirewallSetRequest(firewallSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallApiAPI.SetFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFirewall`: FirewallShowResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallV1FirewallApiAPI.SetFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** | Firewall ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallSetRequest** | [**FirewallSetRequest**](FirewallSetRequest.md) |  | 

### Return type

[**FirewallShowResponse**](FirewallShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFirewallStatus

> FirewallShowResponse SetFirewallStatus(ctx, firewallId).FirewallSetStatusRequest(firewallSetStatusRequest).Execute()

Set Firewall Status.



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
	firewallId := "b156740b6335468d8354eb9ef8eddf5a" // string | Firewall ID
	firewallSetStatusRequest := *openapiclient.NewFirewallSetStatusRequest(openapiclient.FirewallStatusType("ENABLE")) // FirewallSetStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallV1FirewallApiAPI.SetFirewallStatus(context.Background(), firewallId).FirewallSetStatusRequest(firewallSetStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallApiAPI.SetFirewallStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFirewallStatus`: FirewallShowResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallV1FirewallApiAPI.SetFirewallStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** | Firewall ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFirewallStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallSetStatusRequest** | [**FirewallSetStatusRequest**](FirewallSetStatusRequest.md) |  | 

### Return type

[**FirewallShowResponse**](FirewallShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFirewall

> FirewallShowResponse ShowFirewall(ctx, firewallId).Execute()

Show Firewall



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
	firewallId := "b156740b6335468d8354eb9ef8eddf5a" // string | Firewall ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallV1FirewallApiAPI.ShowFirewall(context.Background(), firewallId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallV1FirewallApiAPI.ShowFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowFirewall`: FirewallShowResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallV1FirewallApiAPI.ShowFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **string** | Firewall ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirewallShowResponse**](FirewallShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

