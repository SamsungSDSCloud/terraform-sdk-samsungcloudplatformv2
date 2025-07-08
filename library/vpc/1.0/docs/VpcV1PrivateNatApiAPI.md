# \VpcV1PrivateNatApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckPrivateNatCidrAvailabilities**](VpcV1PrivateNatApiAPI.md#CheckPrivateNatCidrAvailabilities) | **Get** /v1/private-nats/private-nat-cidr-availabilities | Check Private NAT CIDR availabilities
[**CreatePrivateNat**](VpcV1PrivateNatApiAPI.md#CreatePrivateNat) | **Post** /v1/private-nats | Create Private NAT
[**DeletePrivateNat**](VpcV1PrivateNatApiAPI.md#DeletePrivateNat) | **Delete** /v1/private-nats/{private_nat_id} | Delete Private NAT
[**ListPrivateNats**](VpcV1PrivateNatApiAPI.md#ListPrivateNats) | **Get** /v1/private-nats | List Private NAT
[**SetPrivateNat**](VpcV1PrivateNatApiAPI.md#SetPrivateNat) | **Put** /v1/private-nats/{private_nat_id} | Set Private NAT
[**ShowPrivateNat**](VpcV1PrivateNatApiAPI.md#ShowPrivateNat) | **Get** /v1/private-nats/{private_nat_id} | Show Private NAT



## CheckPrivateNatCidrAvailabilities

> PrivateNatCidrAvailabilityShowResponse CheckPrivateNatCidrAvailabilities(ctx).DirectConnectId(directConnectId).Cidr(cidr).Execute()

Check Private NAT CIDR availabilities



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
	directConnectId := "directConnectId_example" // string | Direct Connect ID
	cidr := "cidr_example" // string | Private NAT IP range

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PrivateNatApiAPI.CheckPrivateNatCidrAvailabilities(context.Background()).DirectConnectId(directConnectId).Cidr(cidr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PrivateNatApiAPI.CheckPrivateNatCidrAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckPrivateNatCidrAvailabilities`: PrivateNatCidrAvailabilityShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PrivateNatApiAPI.CheckPrivateNatCidrAvailabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckPrivateNatCidrAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directConnectId** | **string** | Direct Connect ID | 
 **cidr** | **string** | Private NAT IP range | 

### Return type

[**PrivateNatCidrAvailabilityShowResponse**](PrivateNatCidrAvailabilityShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateNat

> PrivateNatShowResponse CreatePrivateNat(ctx).PrivateNatCreateRequest(privateNatCreateRequest).Execute()

Create Private NAT



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
	privateNatCreateRequest := *openapiclient.NewPrivateNatCreateRequest("Cidr_example", "DirectConnectId_example", "Name_example") // PrivateNatCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PrivateNatApiAPI.CreatePrivateNat(context.Background()).PrivateNatCreateRequest(privateNatCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PrivateNatApiAPI.CreatePrivateNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrivateNat`: PrivateNatShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PrivateNatApiAPI.CreatePrivateNat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privateNatCreateRequest** | [**PrivateNatCreateRequest**](PrivateNatCreateRequest.md) |  | 

### Return type

[**PrivateNatShowResponse**](PrivateNatShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivateNat

> DeletePrivateNat(ctx, privateNatId).Execute()

Delete Private NAT



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1PrivateNatApiAPI.DeletePrivateNat(context.Background(), privateNatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PrivateNatApiAPI.DeletePrivateNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNatId** | **string** | Private NAT ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateNatRequest struct via the builder pattern


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


## ListPrivateNats

> PrivateNatListResponse ListPrivateNats(ctx).Size(size).Page(page).Sort(sort).Name(name).Cidr(cidr).VpcId(vpcId).VpcName(vpcName).DirectConnectId(directConnectId).DirectConnectName(directConnectName).State(state).Execute()

List Private NAT



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
	name := "name_example" // string | Private NAT Name (optional)
	cidr := "cidr_example" // string | Private NAT IP range (optional)
	vpcId := "vpcId_example" // string | VPC Id (optional)
	vpcName := "vpcName_example" // string | VPC Name (optional)
	directConnectId := "directConnectId_example" // string | Direct Connect ID (optional)
	directConnectName := "directConnectName_example" // string | Direct Connect Name (optional)
	state := openapiclient.PrivateNatState("CREATING") // PrivateNatState | Private NAT State (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PrivateNatApiAPI.ListPrivateNats(context.Background()).Size(size).Page(page).Sort(sort).Name(name).Cidr(cidr).VpcId(vpcId).VpcName(vpcName).DirectConnectId(directConnectId).DirectConnectName(directConnectName).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PrivateNatApiAPI.ListPrivateNats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrivateNats`: PrivateNatListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PrivateNatApiAPI.ListPrivateNats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateNatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Private NAT Name | 
 **cidr** | **string** | Private NAT IP range | 
 **vpcId** | **string** | VPC Id | 
 **vpcName** | **string** | VPC Name | 
 **directConnectId** | **string** | Direct Connect ID | 
 **directConnectName** | **string** | Direct Connect Name | 
 **state** | [**PrivateNatState**](PrivateNatState.md) | Private NAT State | 

### Return type

[**PrivateNatListResponse**](PrivateNatListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPrivateNat

> PrivateNatShowResponse SetPrivateNat(ctx, privateNatId).PrivateNatSetRequest(privateNatSetRequest).Execute()

Set Private NAT



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
	privateNatSetRequest := *openapiclient.NewPrivateNatSetRequest() // PrivateNatSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PrivateNatApiAPI.SetPrivateNat(context.Background(), privateNatId).PrivateNatSetRequest(privateNatSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PrivateNatApiAPI.SetPrivateNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPrivateNat`: PrivateNatShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PrivateNatApiAPI.SetPrivateNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNatId** | **string** | Private NAT ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPrivateNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateNatSetRequest** | [**PrivateNatSetRequest**](PrivateNatSetRequest.md) |  | 

### Return type

[**PrivateNatShowResponse**](PrivateNatShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPrivateNat

> PrivateNatShowResponse ShowPrivateNat(ctx, privateNatId).Execute()

Show Private NAT



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PrivateNatApiAPI.ShowPrivateNat(context.Background(), privateNatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PrivateNatApiAPI.ShowPrivateNat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowPrivateNat`: PrivateNatShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PrivateNatApiAPI.ShowPrivateNat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privateNatId** | **string** | Private NAT ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowPrivateNatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateNatShowResponse**](PrivateNatShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

