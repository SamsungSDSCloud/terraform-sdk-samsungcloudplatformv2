# \VpcV1VpcEndpointApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckVpcEndpointIpAddressAvailabilities**](VpcV1VpcEndpointApiAPI.md#CheckVpcEndpointIpAddressAvailabilities) | **Get** /v1/vpc-endpoints/vpc-endpoint-ip-availabilities | Check VPC Endpoint IP Address Availabilities
[**CheckVpcEndpointResourceAvailabilities**](VpcV1VpcEndpointApiAPI.md#CheckVpcEndpointResourceAvailabilities) | **Get** /v1/vpc-endpoints/vpc-endpoint-resource-availabilities | Check VPC Endpoint resource availabilities
[**CreateVpcEndpoint**](VpcV1VpcEndpointApiAPI.md#CreateVpcEndpoint) | **Post** /v1/vpc-endpoints | Create VPC Endpoint
[**DeleteVpcEndpoint**](VpcV1VpcEndpointApiAPI.md#DeleteVpcEndpoint) | **Delete** /v1/vpc-endpoints/{vpc_endpoint_id} | Delete VPC Endpoint
[**ListVpcEndpointConnectableResources**](VpcV1VpcEndpointApiAPI.md#ListVpcEndpointConnectableResources) | **Get** /v1/vpc-endpoints/connectable-resources | List VPC Endpoint Connectable Resources
[**ListVpcEndpoints**](VpcV1VpcEndpointApiAPI.md#ListVpcEndpoints) | **Get** /v1/vpc-endpoints | List VPC Endpoint
[**SetVpcEndpoint**](VpcV1VpcEndpointApiAPI.md#SetVpcEndpoint) | **Put** /v1/vpc-endpoints/{vpc_endpoint_id} | Set VPC Endpoint
[**ShowVpcEndpoint**](VpcV1VpcEndpointApiAPI.md#ShowVpcEndpoint) | **Get** /v1/vpc-endpoints/{vpc_endpoint_id} | Show VPC Endpoint



## CheckVpcEndpointIpAddressAvailabilities

> VpcEndpointShowIpAvailabilityResponse CheckVpcEndpointIpAddressAvailabilities(ctx).SubnetId(subnetId).EndpointIpAddress(endpointIpAddress).Execute()

Check VPC Endpoint IP Address Availabilities



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
	subnetId := "subnetId_example" // string | Subnet Id
	endpointIpAddress := "endpointIpAddress_example" // string | VPC Endpoint IP Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcEndpointApiAPI.CheckVpcEndpointIpAddressAvailabilities(context.Background()).SubnetId(subnetId).EndpointIpAddress(endpointIpAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcEndpointApiAPI.CheckVpcEndpointIpAddressAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckVpcEndpointIpAddressAvailabilities`: VpcEndpointShowIpAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcEndpointApiAPI.CheckVpcEndpointIpAddressAvailabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckVpcEndpointIpAddressAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subnetId** | **string** | Subnet Id | 
 **endpointIpAddress** | **string** | VPC Endpoint IP Address | 

### Return type

[**VpcEndpointShowIpAvailabilityResponse**](VpcEndpointShowIpAvailabilityResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckVpcEndpointResourceAvailabilities

> VpcEndpointResourceAvailabilityShowResponse CheckVpcEndpointResourceAvailabilities(ctx).VpcId(vpcId).ResourceType(resourceType).ResourceKey(resourceKey).Execute()

Check VPC Endpoint resource availabilities



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
	resourceType := openapiclient.VpcEndpointResourceType("FS") // VpcEndpointResourceType | VPC Endpoint Resource Type
	resourceKey := "resourceKey_example" // string | VPC Endpoint Resource Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcEndpointApiAPI.CheckVpcEndpointResourceAvailabilities(context.Background()).VpcId(vpcId).ResourceType(resourceType).ResourceKey(resourceKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcEndpointApiAPI.CheckVpcEndpointResourceAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckVpcEndpointResourceAvailabilities`: VpcEndpointResourceAvailabilityShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcEndpointApiAPI.CheckVpcEndpointResourceAvailabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckVpcEndpointResourceAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpcId** | **string** | VPC Id | 
 **resourceType** | [**VpcEndpointResourceType**](VpcEndpointResourceType.md) | VPC Endpoint Resource Type | 
 **resourceKey** | **string** | VPC Endpoint Resource Key | 

### Return type

[**VpcEndpointResourceAvailabilityShowResponse**](VpcEndpointResourceAvailabilityShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVpcEndpoint

> VpcEndpointShowResponse CreateVpcEndpoint(ctx).VpcEndpointCreateRequest(vpcEndpointCreateRequest).Execute()

Create VPC Endpoint



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
	vpcEndpointCreateRequest := *openapiclient.NewVpcEndpointCreateRequest("EndpointIpAddress_example", "Name_example", "ResourceInfo_example", "ResourceKey_example", openapiclient.VpcEndpointResourceType("FS"), "SubnetId_example", "VpcId_example") // VpcEndpointCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcEndpointApiAPI.CreateVpcEndpoint(context.Background()).VpcEndpointCreateRequest(vpcEndpointCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcEndpointApiAPI.CreateVpcEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpcEndpoint`: VpcEndpointShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcEndpointApiAPI.CreateVpcEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpcEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vpcEndpointCreateRequest** | [**VpcEndpointCreateRequest**](VpcEndpointCreateRequest.md) |  | 

### Return type

[**VpcEndpointShowResponse**](VpcEndpointShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpcEndpoint

> DeleteVpcEndpoint(ctx, vpcEndpointId).Execute()

Delete VPC Endpoint



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
	vpcEndpointId := "vpcEndpointId_example" // string | VPC Endpoint ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1VpcEndpointApiAPI.DeleteVpcEndpoint(context.Background(), vpcEndpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcEndpointApiAPI.DeleteVpcEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcEndpointId** | **string** | VPC Endpoint ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpcEndpointRequest struct via the builder pattern


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


## ListVpcEndpointConnectableResources

> VpcEndpointConnectableResourceListResponse ListVpcEndpointConnectableResources(ctx).ResourceType(resourceType).Execute()

List VPC Endpoint Connectable Resources



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
	resourceType := openapiclient.VpcEndpointResourceType("FS") // VpcEndpointResourceType | VPC Endpoint Resource Type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcEndpointApiAPI.ListVpcEndpointConnectableResources(context.Background()).ResourceType(resourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcEndpointApiAPI.ListVpcEndpointConnectableResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpcEndpointConnectableResources`: VpcEndpointConnectableResourceListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcEndpointApiAPI.ListVpcEndpointConnectableResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVpcEndpointConnectableResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | [**VpcEndpointResourceType**](VpcEndpointResourceType.md) | VPC Endpoint Resource Type | 

### Return type

[**VpcEndpointConnectableResourceListResponse**](VpcEndpointConnectableResourceListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVpcEndpoints

> VpcEndpointListResponse ListVpcEndpoints(ctx).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).Name(name).VpcId(vpcId).VpcName(vpcName).SubnetId(subnetId).ResourceKey(resourceKey).ResourceType(resourceType).EndpointIpAddress(endpointIpAddress).State(state).Execute()

List VPC Endpoint



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
	id := "id_example" // string | VPC Endpoint ID (optional)
	name := "name_example" // string | VPC Endpoint Name (optional)
	vpcId := "vpcId_example" // string | VPC Id (optional)
	vpcName := "vpcName_example" // string | VPC Name (optional)
	subnetId := "subnetId_example" // string | Subnet Id (optional)
	resourceKey := "resourceKey_example" // string | VPC Endpoint Resource Key (optional)
	resourceType := openapiclient.VpcEndpointResourceType("FS") // VpcEndpointResourceType | VPC Endpoint Resource Type (optional)
	endpointIpAddress := "endpointIpAddress_example" // string | VPC Endpoint IP Address (optional)
	state := openapiclient.VpcEndpointState("CREATING") // VpcEndpointState | VPC Endpoint State (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcEndpointApiAPI.ListVpcEndpoints(context.Background()).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).Name(name).VpcId(vpcId).VpcName(vpcName).SubnetId(subnetId).ResourceKey(resourceKey).ResourceType(resourceType).EndpointIpAddress(endpointIpAddress).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcEndpointApiAPI.ListVpcEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpcEndpoints`: VpcEndpointListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcEndpointApiAPI.ListVpcEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVpcEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 
 **id** | **string** | VPC Endpoint ID | 
 **name** | **string** | VPC Endpoint Name | 
 **vpcId** | **string** | VPC Id | 
 **vpcName** | **string** | VPC Name | 
 **subnetId** | **string** | Subnet Id | 
 **resourceKey** | **string** | VPC Endpoint Resource Key | 
 **resourceType** | [**VpcEndpointResourceType**](VpcEndpointResourceType.md) | VPC Endpoint Resource Type | 
 **endpointIpAddress** | **string** | VPC Endpoint IP Address | 
 **state** | [**VpcEndpointState**](VpcEndpointState.md) | VPC Endpoint State | 

### Return type

[**VpcEndpointListResponse**](VpcEndpointListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVpcEndpoint

> VpcEndpointShowResponse SetVpcEndpoint(ctx, vpcEndpointId).VpcEndpointSetRequest(vpcEndpointSetRequest).Execute()

Set VPC Endpoint



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
	vpcEndpointId := "vpcEndpointId_example" // string | VPC Endpoint ID
	vpcEndpointSetRequest := *openapiclient.NewVpcEndpointSetRequest() // VpcEndpointSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcEndpointApiAPI.SetVpcEndpoint(context.Background(), vpcEndpointId).VpcEndpointSetRequest(vpcEndpointSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcEndpointApiAPI.SetVpcEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVpcEndpoint`: VpcEndpointShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcEndpointApiAPI.SetVpcEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcEndpointId** | **string** | VPC Endpoint ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVpcEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpcEndpointSetRequest** | [**VpcEndpointSetRequest**](VpcEndpointSetRequest.md) |  | 

### Return type

[**VpcEndpointShowResponse**](VpcEndpointShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVpcEndpoint

> VpcEndpointShowResponse ShowVpcEndpoint(ctx, vpcEndpointId).Execute()

Show VPC Endpoint



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
	vpcEndpointId := "vpcEndpointId_example" // string | VPC Endpoint ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1VpcEndpointApiAPI.ShowVpcEndpoint(context.Background(), vpcEndpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1VpcEndpointApiAPI.ShowVpcEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVpcEndpoint`: VpcEndpointShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1VpcEndpointApiAPI.ShowVpcEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcEndpointId** | **string** | VPC Endpoint ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVpcEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VpcEndpointShowResponse**](VpcEndpointShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

