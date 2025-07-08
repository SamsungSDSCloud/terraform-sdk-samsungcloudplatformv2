# \VpcV1PrivateNatIpApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrivateNatIp**](VpcV1PrivateNatIpApiAPI.md#CreatePrivateNatIp) | **Post** /v1/private-nats/{private_nat_id}/private-nat-ips | Create Private NAT IP
[**DeletePrivateNatIp**](VpcV1PrivateNatIpApiAPI.md#DeletePrivateNatIp) | **Delete** /v1/private-nats/{private_nat_id}/private-nat-ips/{private_nat_ip_id} | Delete Private NAT IP
[**ListPrivateNatIps**](VpcV1PrivateNatIpApiAPI.md#ListPrivateNatIps) | **Get** /v1/private-nats/{private_nat_id}/private-nat-ips | List Private NAT IP



## CreatePrivateNatIp

> PrivateNatIpShowResponse CreatePrivateNatIp(ctx, privateNatId).PrivateNatIpCreateRequest(privateNatIpCreateRequest).Execute()

Create Private NAT IP



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
	privateNatId := "privateNatId_example" // string | Private NAT ID
	privateNatIpCreateRequest := *openapiclient.NewPrivateNatIpCreateRequest("IpAddress_example") // PrivateNatIpCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PrivateNatIpApiAPI.CreatePrivateNatIp(context.Background(), privateNatId).PrivateNatIpCreateRequest(privateNatIpCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PrivateNatIpApiAPI.CreatePrivateNatIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrivateNatIp`: PrivateNatIpShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PrivateNatIpApiAPI.CreatePrivateNatIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNatId** | **string** | Private NAT ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateNatIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateNatIpCreateRequest** | [**PrivateNatIpCreateRequest**](PrivateNatIpCreateRequest.md) |  | 

### Return type

[**PrivateNatIpShowResponse**](PrivateNatIpShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivateNatIp

> DeletePrivateNatIp(ctx, privateNatId, privateNatIpId).Execute()

Delete Private NAT IP



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
	privateNatId := "privateNatId_example" // string | Private NAT ID
	privateNatIpId := "privateNatIpId_example" // string | Private NAT IP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1PrivateNatIpApiAPI.DeletePrivateNatIp(context.Background(), privateNatId, privateNatIpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PrivateNatIpApiAPI.DeletePrivateNatIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNatId** | **string** | Private NAT ID | 
**privateNatIpId** | **string** | Private NAT IP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateNatIpRequest struct via the builder pattern


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


## ListPrivateNatIps

> PrivateNatIpListResponse ListPrivateNatIps(ctx, privateNatId).Size(size).Page(page).Sort(sort).IpAddress(ipAddress).State(state).AttachedResourceType(attachedResourceType).AttachedResourceId(attachedResourceId).AttachedResourceName(attachedResourceName).Execute()

List Private NAT IP



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
	privateNatId := "privateNatId_example" // string | Private NAT ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	ipAddress := "ipAddress_example" // string | Private NAT IP Address (optional)
	state := openapiclient.PrivateNatIpState("RESERVED") // PrivateNatIpState | Private NAT IP State (optional)
	attachedResourceType := openapiclient.PrivateNatIpAttachedResourceType("VM") // PrivateNatIpAttachedResourceType | Private NAT IP Attached Resource Type (optional)
	attachedResourceId := "attachedResourceId_example" // string | Private NAT IP Attached Resource ID (optional)
	attachedResourceName := "attachedResourceName_example" // string | Private NAT IP Attached Resource Name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PrivateNatIpApiAPI.ListPrivateNatIps(context.Background(), privateNatId).Size(size).Page(page).Sort(sort).IpAddress(ipAddress).State(state).AttachedResourceType(attachedResourceType).AttachedResourceId(attachedResourceId).AttachedResourceName(attachedResourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PrivateNatIpApiAPI.ListPrivateNatIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrivateNatIps`: PrivateNatIpListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PrivateNatIpApiAPI.ListPrivateNatIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNatId** | **string** | Private NAT ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateNatIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **ipAddress** | **string** | Private NAT IP Address | 
 **state** | [**PrivateNatIpState**](PrivateNatIpState.md) | Private NAT IP State | 
 **attachedResourceType** | [**PrivateNatIpAttachedResourceType**](PrivateNatIpAttachedResourceType.md) | Private NAT IP Attached Resource Type | 
 **attachedResourceId** | **string** | Private NAT IP Attached Resource ID | 
 **attachedResourceName** | **string** | Private NAT IP Attached Resource Name | 

### Return type

[**PrivateNatIpListResponse**](PrivateNatIpListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

