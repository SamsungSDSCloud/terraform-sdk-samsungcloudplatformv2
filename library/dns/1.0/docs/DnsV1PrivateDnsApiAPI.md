# \DnsV1PrivateDnsApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePrivateDns**](DnsV1PrivateDnsApiAPI.md#ActivatePrivateDns) | **Post** /v1/private-dns/activate | Activate Private DNS
[**CreatePrivateDns**](DnsV1PrivateDnsApiAPI.md#CreatePrivateDns) | **Post** /v1/private-dns | Create Private DNS
[**DeletePrivateDns**](DnsV1PrivateDnsApiAPI.md#DeletePrivateDns) | **Delete** /v1/private-dns/{private_dns_id} | Delete Private DNS
[**ListPrivateDns**](DnsV1PrivateDnsApiAPI.md#ListPrivateDns) | **Get** /v1/private-dns | List Private DNS
[**SetPrivateDns**](DnsV1PrivateDnsApiAPI.md#SetPrivateDns) | **Put** /v1/private-dns/{private_dns_id} | Set Private DNS
[**ShowPrivateDns**](DnsV1PrivateDnsApiAPI.md#ShowPrivateDns) | **Get** /v1/private-dns/{private_dns_id} | Show Private DNS



## ActivatePrivateDns

> PrivateDnsShowResponse ActivatePrivateDns(ctx).PrivateDnsActivateRequest(privateDnsActivateRequest).Execute()

Activate Private DNS



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
	privateDnsActivateRequest := *openapiclient.NewPrivateDnsActivateRequest("private-dns01") // PrivateDnsActivateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1PrivateDnsApiAPI.ActivatePrivateDns(context.Background()).PrivateDnsActivateRequest(privateDnsActivateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PrivateDnsApiAPI.ActivatePrivateDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivatePrivateDns`: PrivateDnsShowResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1PrivateDnsApiAPI.ActivatePrivateDns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivatePrivateDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privateDnsActivateRequest** | [**PrivateDnsActivateRequest**](PrivateDnsActivateRequest.md) |  | 

### Return type

[**PrivateDnsShowResponse**](PrivateDnsShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateDns

> PrivateDnsShowResponse CreatePrivateDns(ctx).PrivateDnsCreateRequest(privateDnsCreateRequest).Execute()

Create Private DNS



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
	privateDnsCreateRequest := *openapiclient.NewPrivateDnsCreateRequest("private-dns01") // PrivateDnsCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1PrivateDnsApiAPI.CreatePrivateDns(context.Background()).PrivateDnsCreateRequest(privateDnsCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PrivateDnsApiAPI.CreatePrivateDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrivateDns`: PrivateDnsShowResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1PrivateDnsApiAPI.CreatePrivateDns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privateDnsCreateRequest** | [**PrivateDnsCreateRequest**](PrivateDnsCreateRequest.md) |  | 

### Return type

[**PrivateDnsShowResponse**](PrivateDnsShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivateDns

> DeletePrivateDns(ctx, privateDnsId).Execute()

Delete Private DNS



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
	privateDnsId := "10fjkewefprivatedns3193rud543" // string | Private DNS ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsV1PrivateDnsApiAPI.DeletePrivateDns(context.Background(), privateDnsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PrivateDnsApiAPI.DeletePrivateDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateDnsId** | **string** | Private DNS ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateDnsRequest struct via the builder pattern


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


## ListPrivateDns

> PrivateDnsListResponse ListPrivateDns(ctx).Size(size).Page(page).Sort(sort).Id(id).Name(name).VpcId(vpcId).State(state).Execute()

List Private DNS



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
	id := "10fjkewefprivatedns3193rud543" // string | Private DNS ID (optional)
	name := "name_example" // string | Private DNS Name (optional)
	vpcId := "7df8abb4912e4709b1cb237daccca7a8" // string | Connected VPC ID (optional)
	state := openapiclient.PrivateDnsState("CREATING") // PrivateDnsState | State (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1PrivateDnsApiAPI.ListPrivateDns(context.Background()).Size(size).Page(page).Sort(sort).Id(id).Name(name).VpcId(vpcId).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PrivateDnsApiAPI.ListPrivateDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrivateDns`: PrivateDnsListResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1PrivateDnsApiAPI.ListPrivateDns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **id** | **string** | Private DNS ID | 
 **name** | **string** | Private DNS Name | 
 **vpcId** | **string** | Connected VPC ID | 
 **state** | [**PrivateDnsState**](PrivateDnsState.md) | State | 

### Return type

[**PrivateDnsListResponse**](PrivateDnsListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPrivateDns

> PrivateDnsShowResponse SetPrivateDns(ctx, privateDnsId).PrivateDnsSetRequest(privateDnsSetRequest).Execute()

Set Private DNS



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
	privateDnsId := "10fjkewefprivatedns3193rud543" // string | Private DNS ID
	privateDnsSetRequest := *openapiclient.NewPrivateDnsSetRequest() // PrivateDnsSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1PrivateDnsApiAPI.SetPrivateDns(context.Background(), privateDnsId).PrivateDnsSetRequest(privateDnsSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PrivateDnsApiAPI.SetPrivateDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPrivateDns`: PrivateDnsShowResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1PrivateDnsApiAPI.SetPrivateDns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateDnsId** | **string** | Private DNS ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPrivateDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateDnsSetRequest** | [**PrivateDnsSetRequest**](PrivateDnsSetRequest.md) |  | 

### Return type

[**PrivateDnsShowResponse**](PrivateDnsShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPrivateDns

> PrivateDnsShowResponse ShowPrivateDns(ctx, privateDnsId).Execute()

Show Private DNS



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
	privateDnsId := "10fjkewefprivatedns3193rud543" // string | Private DNS ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1PrivateDnsApiAPI.ShowPrivateDns(context.Background(), privateDnsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PrivateDnsApiAPI.ShowPrivateDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowPrivateDns`: PrivateDnsShowResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1PrivateDnsApiAPI.ShowPrivateDns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateDnsId** | **string** | Private DNS ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowPrivateDnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateDnsShowResponse**](PrivateDnsShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

