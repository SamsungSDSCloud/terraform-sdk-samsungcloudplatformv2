# \VpcV1PublicIpApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePublicip**](VpcV1PublicIpApiAPI.md#CreatePublicip) | **Post** /v1/publicips | Create PublicIP
[**DeletePublicip**](VpcV1PublicIpApiAPI.md#DeletePublicip) | **Delete** /v1/publicips/{publicip_id} | Delete PublicIP
[**ListPublicip**](VpcV1PublicIpApiAPI.md#ListPublicip) | **Get** /v1/publicips | List PublicIPs
[**SetPublicip**](VpcV1PublicIpApiAPI.md#SetPublicip) | **Put** /v1/publicips/{publicip_id} | Set PublicIP
[**ShowPublicip**](VpcV1PublicIpApiAPI.md#ShowPublicip) | **Get** /v1/publicips/{publicip_id} | Show PublicIP



## CreatePublicip

> PublicipShowResponse CreatePublicip(ctx).PublicipCreateRequest(publicipCreateRequest).Execute()

Create PublicIP



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
	publicipCreateRequest := *openapiclient.NewPublicipCreateRequest(openapiclient.PublicipType("IGW")) // PublicipCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PublicIpApiAPI.CreatePublicip(context.Background()).PublicipCreateRequest(publicipCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PublicIpApiAPI.CreatePublicip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePublicip`: PublicipShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PublicIpApiAPI.CreatePublicip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePublicipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicipCreateRequest** | [**PublicipCreateRequest**](PublicipCreateRequest.md) |  | 

### Return type

[**PublicipShowResponse**](PublicipShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePublicip

> DeletePublicip(ctx, publicipId).Execute()

Delete PublicIP



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
	publicipId := "publicipId_example" // string | PublicIP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VpcV1PublicIpApiAPI.DeletePublicip(context.Background(), publicipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PublicIpApiAPI.DeletePublicip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicipId** | **string** | PublicIP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePublicipRequest struct via the builder pattern


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


## ListPublicip

> PublicipListResponse ListPublicip(ctx).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).IpAddress(ipAddress).State(state).AttachedResourceType(attachedResourceType).AttachedResourceId(attachedResourceId).AttachedResourceName(attachedResourceName).VpcId(vpcId).Type_(type_).Execute()

List PublicIPs



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
	ipAddress := "ipAddress_example" // string | IP Address (optional)
	state := "state_example" // string | PublicIP State (optional)
	attachedResourceType := "attachedResourceType_example" // string | PublicIP Attached Resource Type (optional)
	attachedResourceId := "attachedResourceId_example" // string | PublicIP Attached Resource ID (optional)
	attachedResourceName := "attachedResourceName_example" // string | PublicIP Attached Resource Name (optional)
	vpcId := "vpcId_example" // string | VPC Id (optional)
	type_ := openapiclient.PublicipType("IGW") // PublicipType | PublicIP Type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PublicIpApiAPI.ListPublicip(context.Background()).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).IpAddress(ipAddress).State(state).AttachedResourceType(attachedResourceType).AttachedResourceId(attachedResourceId).AttachedResourceName(attachedResourceName).VpcId(vpcId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PublicIpApiAPI.ListPublicip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPublicip`: PublicipListResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PublicIpApiAPI.ListPublicip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPublicipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 
 **ipAddress** | **string** | IP Address | 
 **state** | **string** | PublicIP State | 
 **attachedResourceType** | **string** | PublicIP Attached Resource Type | 
 **attachedResourceId** | **string** | PublicIP Attached Resource ID | 
 **attachedResourceName** | **string** | PublicIP Attached Resource Name | 
 **vpcId** | **string** | VPC Id | 
 **type_** | [**PublicipType**](PublicipType.md) | PublicIP Type | 

### Return type

[**PublicipListResponse**](PublicipListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPublicip

> PublicipShowResponse SetPublicip(ctx, publicipId).PublicipSetRequest(publicipSetRequest).Execute()

Set PublicIP



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
	publicipId := "publicipId_example" // string | PublicIP ID
	publicipSetRequest := *openapiclient.NewPublicipSetRequest("Description_example") // PublicipSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PublicIpApiAPI.SetPublicip(context.Background(), publicipId).PublicipSetRequest(publicipSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PublicIpApiAPI.SetPublicip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPublicip`: PublicipShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PublicIpApiAPI.SetPublicip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicipId** | **string** | PublicIP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPublicipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicipSetRequest** | [**PublicipSetRequest**](PublicipSetRequest.md) |  | 

### Return type

[**PublicipShowResponse**](PublicipShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPublicip

> PublicipShowResponse ShowPublicip(ctx, publicipId).Execute()

Show PublicIP



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
	publicipId := "publicipId_example" // string | PublicIP ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpcV1PublicIpApiAPI.ShowPublicip(context.Background(), publicipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpcV1PublicIpApiAPI.ShowPublicip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowPublicip`: PublicipShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VpcV1PublicIpApiAPI.ShowPublicip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicipId** | **string** | PublicIP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowPublicipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicipShowResponse**](PublicipShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

