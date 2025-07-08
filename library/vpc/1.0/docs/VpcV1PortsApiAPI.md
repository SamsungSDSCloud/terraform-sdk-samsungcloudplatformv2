# \VpcV1PortsApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePort**](VpcV1PortsApiAPI.md#CreatePort) | **Post** /v1/ports | Create Port
[**DeletePort**](VpcV1PortsApiAPI.md#DeletePort) | **Delete** /v1/ports/{port_id} | Delete Port
[**ListPortIpAvailabilities**](VpcV1PortsApiAPI.md#ListPortIpAvailabilities) | **Get** /v1/ports/port-ip-availabilities | ListPortIpAvailabilities
[**ListPorts**](VpcV1PortsApiAPI.md#ListPorts) | **Get** /v1/ports | List Ports
[**SetPort**](VpcV1PortsApiAPI.md#SetPort) | **Put** /v1/ports/{port_id} | Set Port
[**ShowPort**](VpcV1PortsApiAPI.md#ShowPort) | **Get** /v1/ports/{port_id} | Show Port



## CreatePort

> PortShowResponse CreatePort(ctx).PortCreateRequest(portCreateRequest).Execute()

Create Port



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
	portCreateRequest := *openapiclient.NewPortCreateRequest("Name_example", "SubnetId_example") // PortCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PortsApiAPI.CreatePort(context.Background()).PortCreateRequest(portCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PortsApiAPI.CreatePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePort`: PortShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PortsApiAPI.CreatePort`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portCreateRequest** | [**PortCreateRequest**](PortCreateRequest.md) |  | 

### Return type

[**PortShowResponse**](PortShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePort

> DeletePort(ctx, portId).Execute()

Delete Port



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
	portId := "portId_example" // string | Port ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1PortsApiAPI.DeletePort(context.Background(), portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PortsApiAPI.DeletePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortRequest struct via the builder pattern


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


## ListPortIpAvailabilities

> PortIpAvailabilityListResponse ListPortIpAvailabilities(ctx).SubnetId(subnetId).FixedIpAddress(fixedIpAddress).Execute()

ListPortIpAvailabilities



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
	fixedIpAddress := "fixedIpAddress_example" // string | Fixed IP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PortsApiAPI.ListPortIpAvailabilities(context.Background()).SubnetId(subnetId).FixedIpAddress(fixedIpAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PortsApiAPI.ListPortIpAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPortIpAvailabilities`: PortIpAvailabilityListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PortsApiAPI.ListPortIpAvailabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPortIpAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subnetId** | **string** | Subnet Id | 
 **fixedIpAddress** | **string** | Fixed IP | 

### Return type

[**PortIpAvailabilityListResponse**](PortIpAvailabilityListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPorts

> PortListResponse ListPorts(ctx).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Name(name).SubnetId(subnetId).SubnetName(subnetName).Id(id).AttachedResourceId(attachedResourceId).FixedIpAddress(fixedIpAddress).MacAddress(macAddress).State(state).SecurityGroups(securityGroups).Execute()

List Ports



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
	limit := int32(56) // int32 | Port Limit (optional)
	marker := "marker_example" // string | Port Marker (optional)
	sort := "sort_example" // string | Port Sort (optional)
	name := "name_example" // string | Port Name (optional)
	subnetId := "subnetId_example" // string | Subnet Id (optional)
	subnetName := "subnetName_example" // string | Subnet Name (optional)
	id := "id_example" // string | Port ID (optional)
	attachedResourceId := "attachedResourceId_example" // string | Connected resource ID (optional)
	fixedIpAddress := "fixedIpAddress_example" // string | Fixed IP (optional)
	macAddress := "macAddress_example" // string | MAC Address (optional)
	state := "state_example" // string | State (optional)
	securityGroups := "securityGroups_example" // string | Security Group List (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PortsApiAPI.ListPorts(context.Background()).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Name(name).SubnetId(subnetId).SubnetName(subnetName).Id(id).AttachedResourceId(attachedResourceId).FixedIpAddress(fixedIpAddress).MacAddress(macAddress).State(state).SecurityGroups(securityGroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PortsApiAPI.ListPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPorts`: PortListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PortsApiAPI.ListPorts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withCount** | **string** | with count | 
 **limit** | **int32** | Port Limit | 
 **marker** | **string** | Port Marker | 
 **sort** | **string** | Port Sort | 
 **name** | **string** | Port Name | 
 **subnetId** | **string** | Subnet Id | 
 **subnetName** | **string** | Subnet Name | 
 **id** | **string** | Port ID | 
 **attachedResourceId** | **string** | Connected resource ID | 
 **fixedIpAddress** | **string** | Fixed IP | 
 **macAddress** | **string** | MAC Address | 
 **state** | **string** | State | 
 **securityGroups** | **string** | Security Group List | 

### Return type

[**PortListResponse**](PortListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPort

> PortShowResponse SetPort(ctx, portId).PortSetRequest(portSetRequest).Execute()

Set Port



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
	portId := "portId_example" // string | Port ID
	portSetRequest := *openapiclient.NewPortSetRequest() // PortSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PortsApiAPI.SetPort(context.Background(), portId).PortSetRequest(portSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PortsApiAPI.SetPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPort`: PortShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PortsApiAPI.SetPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **portSetRequest** | [**PortSetRequest**](PortSetRequest.md) |  | 

### Return type

[**PortShowResponse**](PortShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPort

> PortShowResponse ShowPort(ctx, portId).Execute()

Show Port



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
	portId := "portId_example" // string | Port ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PortsApiAPI.ShowPort(context.Background(), portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PortsApiAPI.ShowPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowPort`: PortShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PortsApiAPI.ShowPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortShowResponse**](PortShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

