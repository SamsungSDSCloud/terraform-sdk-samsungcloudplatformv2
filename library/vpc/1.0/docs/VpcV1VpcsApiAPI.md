# \VpcV1VpcsApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckVpcCidrAvailabilities**](VpcV1VpcsApiAPI.md#CheckVpcCidrAvailabilities) | **Get** /v1/vpcs/vpc-cidr-availabilities | CheckVpcCidrAvailabilities
[**CreateVpc**](VpcV1VpcsApiAPI.md#CreateVpc) | **Post** /v1/vpcs | Create VPC
[**DeleteVpc**](VpcV1VpcsApiAPI.md#DeleteVpc) | **Delete** /v1/vpcs/{vpc_id} | Delete VPC
[**ListVpcs**](VpcV1VpcsApiAPI.md#ListVpcs) | **Get** /v1/vpcs | List VPCs
[**SetVpc**](VpcV1VpcsApiAPI.md#SetVpc) | **Put** /v1/vpcs/{vpc_id} | Set VPC
[**ShowMaxCount**](VpcV1VpcsApiAPI.md#ShowMaxCount) | **Get** /v1/vpcs/check-max-count | showMaxCount
[**ShowVpc**](VpcV1VpcsApiAPI.md#ShowVpc) | **Get** /v1/vpcs/{vpc_id} | Show VPC



## CheckVpcCidrAvailabilities

> AvailableVpcCidrResponse CheckVpcCidrAvailabilities(ctx).Cidr(cidr).Execute()

CheckVpcCidrAvailabilities



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
	cidr := "cidr_example" // string | VPC Cidr

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcsApiAPI.CheckVpcCidrAvailabilities(context.Background()).Cidr(cidr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcsApiAPI.CheckVpcCidrAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckVpcCidrAvailabilities`: AvailableVpcCidrResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcsApiAPI.CheckVpcCidrAvailabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckVpcCidrAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cidr** | **string** | VPC Cidr | 

### Return type

[**AvailableVpcCidrResponse**](AvailableVpcCidrResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVpc

> VpcShowResponse CreateVpc(ctx).VpcCreateRequest(vpcCreateRequest).Execute()

Create VPC



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
	vpcCreateRequest := *openapiclient.NewVpcCreateRequest("192.167.0.0/18", "vpcName") // VpcCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcsApiAPI.CreateVpc(context.Background()).VpcCreateRequest(vpcCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcsApiAPI.CreateVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpc`: VpcShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcsApiAPI.CreateVpc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpcCreateRequest** | [**VpcCreateRequest**](VpcCreateRequest.md) |  | 

### Return type

[**VpcShowResponse**](VpcShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpc

> DeleteVpc(ctx, vpcId).Execute()

Delete VPC



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
	vpcId := "vpcId_example" // string | VPC Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1VpcsApiAPI.DeleteVpc(context.Background(), vpcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcsApiAPI.DeleteVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | VPC Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpcRequest struct via the builder pattern


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


## ListVpcs

> VpcListResponse ListVpcs(ctx).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).Name(name).State(state).Cidr(cidr).Execute()

List VPCs



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
	id := "7df8abb4912e4709b1cb237daccca7a8" // string | VPC Id (optional)
	name := "vpcName" // string | VPC Name (optional)
	state := openapiclient.VpcState("CREATING") // VpcState | State (optional)
	cidr := "192.167.0.0/18" // string | VPC Cidr (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcsApiAPI.ListVpcs(context.Background()).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).Name(name).State(state).Cidr(cidr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcsApiAPI.ListVpcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpcs`: VpcListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcsApiAPI.ListVpcs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVpcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 
 **id** | **string** | VPC Id | 
 **name** | **string** | VPC Name | 
 **state** | [**VpcState**](VpcState.md) | State | 
 **cidr** | **string** | VPC Cidr | 

### Return type

[**VpcListResponse**](VpcListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVpc

> VpcShowResponse SetVpc(ctx, vpcId).VpcSetRequest(vpcSetRequest).Execute()

Set VPC



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
	vpcId := "vpcId_example" // string | VPC Id
	vpcSetRequest := *openapiclient.NewVpcSetRequest() // VpcSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcsApiAPI.SetVpc(context.Background(), vpcId).VpcSetRequest(vpcSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcsApiAPI.SetVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVpc`: VpcShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcsApiAPI.SetVpc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | VPC Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpcSetRequest** | [**VpcSetRequest**](VpcSetRequest.md) |  | 

### Return type

[**VpcShowResponse**](VpcShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowMaxCount

> CheckMaxCountShowResponse ShowMaxCount(ctx).ConfigName(configName).Execute()

showMaxCount



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
	configName := openapiclient.NetworkConfigMaxValue("VPC.PROJECT.MAX") // NetworkConfigMaxValue | Config Name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcsApiAPI.ShowMaxCount(context.Background()).ConfigName(configName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcsApiAPI.ShowMaxCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowMaxCount`: CheckMaxCountShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcsApiAPI.ShowMaxCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowMaxCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configName** | [**NetworkConfigMaxValue**](NetworkConfigMaxValue.md) | Config Name | 

### Return type

[**CheckMaxCountShowResponse**](CheckMaxCountShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVpc

> VpcShowResponse ShowVpc(ctx, vpcId).Execute()

Show VPC



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
	vpcId := "vpcId_example" // string | VPC Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcsApiAPI.ShowVpc(context.Background(), vpcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcsApiAPI.ShowVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVpc`: VpcShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcsApiAPI.ShowVpc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | VPC Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VpcShowResponse**](VpcShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

