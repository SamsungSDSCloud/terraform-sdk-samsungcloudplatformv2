# \VpcV1SubnetsApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckSubnetCidrAvailabilities**](VpcV1SubnetsApiAPI.md#CheckSubnetCidrAvailabilities) | **Get** /v1/subnets/subnet-cidr-availabilities | CheckSubnetCidrAvailabilities
[**CreateSubnet**](VpcV1SubnetsApiAPI.md#CreateSubnet) | **Post** /v1/subnets | Create Subnet
[**DeleteSubnet**](VpcV1SubnetsApiAPI.md#DeleteSubnet) | **Delete** /v1/subnets/{subnet_id} | Delete Subnet
[**ListSubnets**](VpcV1SubnetsApiAPI.md#ListSubnets) | **Get** /v1/subnets | List Subnets
[**SetSubnet**](VpcV1SubnetsApiAPI.md#SetSubnet) | **Put** /v1/subnets/{subnet_id} | Set Subnet
[**ShowSubnet**](VpcV1SubnetsApiAPI.md#ShowSubnet) | **Get** /v1/subnets/{subnet_id} | Show Subnet



## CheckSubnetCidrAvailabilities

> AvailableSubnetCidrResponse CheckSubnetCidrAvailabilities(ctx).VpcId(vpcId).Cidr(cidr).Execute()

CheckSubnetCidrAvailabilities



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
	vpcId := "7df8abb4912e4709b1cb237daccca7a8" // string | VPC Id
	cidr := "192.167.1.0/24" // string | Subnet Cidr

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1SubnetsApiAPI.CheckSubnetCidrAvailabilities(context.Background()).VpcId(vpcId).Cidr(cidr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1SubnetsApiAPI.CheckSubnetCidrAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckSubnetCidrAvailabilities`: AvailableSubnetCidrResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1SubnetsApiAPI.CheckSubnetCidrAvailabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckSubnetCidrAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpcId** | **string** | VPC Id | 
 **cidr** | **string** | Subnet Cidr | 

### Return type

[**AvailableSubnetCidrResponse**](AvailableSubnetCidrResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubnet

> SubnetShowResponse CreateSubnet(ctx).SubnetCreateRequest(subnetCreateRequest).Execute()

Create Subnet



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
	subnetCreateRequest := *openapiclient.NewSubnetCreateRequest("192.167.1.0/24", "subnetName", openapiclient.SubnetType("GENERAL"), "7df8abb4912e4709b1cb237daccca7a8") // SubnetCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1SubnetsApiAPI.CreateSubnet(context.Background()).SubnetCreateRequest(subnetCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1SubnetsApiAPI.CreateSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubnet`: SubnetShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1SubnetsApiAPI.CreateSubnet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subnetCreateRequest** | [**SubnetCreateRequest**](SubnetCreateRequest.md) |  | 

### Return type

[**SubnetShowResponse**](SubnetShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubnet

> DeleteSubnet(ctx, subnetId).Execute()

Delete Subnet



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
	subnetId := "023c57b14f11483689338d085e061492" // string | Subnet Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1SubnetsApiAPI.DeleteSubnet(context.Background(), subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1SubnetsApiAPI.DeleteSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | Subnet Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubnetRequest struct via the builder pattern


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


## ListSubnets

> SubnetListResponse ListSubnets(ctx).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).Name(name).Type_(type_).State(state).VpcId(vpcId).VpcName(vpcName).Execute()

List Subnets



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
	id := "023c57b14f11483689338d085e061492" // string | Subnet Id (optional) (default to "")
	name := "subnetName" // string | Subnet Name (optional) (default to "")
	type_ := *openapiclient.NewType() // Type | Subnet Type (optional) (default to )
	state := openapiclient.SubnetState("CREATING") // SubnetState | State (optional) (default to "")
	vpcId := "7df8abb4912e4709b1cb237daccca7a8" // string | VPC Id (optional) (default to "")
	vpcName := "vpcName" // string | VPC Name (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1SubnetsApiAPI.ListSubnets(context.Background()).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).Name(name).Type_(type_).State(state).VpcId(vpcId).VpcName(vpcName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1SubnetsApiAPI.ListSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubnets`: SubnetListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1SubnetsApiAPI.ListSubnets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 
 **id** | **string** | Subnet Id | [default to &quot;&quot;]
 **name** | **string** | Subnet Name | [default to &quot;&quot;]
 **type_** | [**Type**](Type.md) | Subnet Type | [default to ]
 **state** | [**SubnetState**](SubnetState.md) | State | [default to &quot;&quot;]
 **vpcId** | **string** | VPC Id | [default to &quot;&quot;]
 **vpcName** | **string** | VPC Name | [default to &quot;&quot;]

### Return type

[**SubnetListResponse**](SubnetListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSubnet

> SubnetShowResponse SetSubnet(ctx, subnetId).SubnetSetRequest(subnetSetRequest).Execute()

Set Subnet



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
	subnetId := "023c57b14f11483689338d085e061492" // string | Subnet Id
	subnetSetRequest := *openapiclient.NewSubnetSetRequest("Subnet Description") // SubnetSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1SubnetsApiAPI.SetSubnet(context.Background(), subnetId).SubnetSetRequest(subnetSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1SubnetsApiAPI.SetSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSubnet`: SubnetShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1SubnetsApiAPI.SetSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | Subnet Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subnetSetRequest** | [**SubnetSetRequest**](SubnetSetRequest.md) |  | 

### Return type

[**SubnetShowResponse**](SubnetShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowSubnet

> SubnetShowResponse ShowSubnet(ctx, subnetId).Execute()

Show Subnet



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
	subnetId := "023c57b14f11483689338d085e061492" // string | Subnet Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1SubnetsApiAPI.ShowSubnet(context.Background(), subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1SubnetsApiAPI.ShowSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowSubnet`: SubnetShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1SubnetsApiAPI.ShowSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | Subnet Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubnetShowResponse**](SubnetShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

