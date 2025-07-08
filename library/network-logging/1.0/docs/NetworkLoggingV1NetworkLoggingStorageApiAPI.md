# \NetworkLoggingV1NetworkLoggingStorageApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkLoggingStorage**](NetworkLoggingV1NetworkLoggingStorageApiAPI.md#CreateNetworkLoggingStorage) | **Post** /v1/network-logging/storages | Create Network Logging Storage
[**DeleteNetworkLoggingStorage**](NetworkLoggingV1NetworkLoggingStorageApiAPI.md#DeleteNetworkLoggingStorage) | **Delete** /v1/network-logging/storages/{network_logging_storage_id} | Delete Network Logging Storage
[**ListNetworkLoggingStorages**](NetworkLoggingV1NetworkLoggingStorageApiAPI.md#ListNetworkLoggingStorages) | **Get** /v1/network-logging/storages | List Network Logging Storages



## CreateNetworkLoggingStorage

> NetworkLoggingStorageShowResponse CreateNetworkLoggingStorage(ctx).NetworkLoggingStorageCreateRequest(networkLoggingStorageCreateRequest).Execute()

Create Network Logging Storage



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
	networkLoggingStorageCreateRequest := *openapiclient.NewNetworkLoggingStorageCreateRequest("BucketName_example", openapiclient.NetworkLoggingResourceType("FIREWALL")) // NetworkLoggingStorageCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkLoggingV1NetworkLoggingStorageApiAPI.CreateNetworkLoggingStorage(context.Background()).NetworkLoggingStorageCreateRequest(networkLoggingStorageCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoggingV1NetworkLoggingStorageApiAPI.CreateNetworkLoggingStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkLoggingStorage`: NetworkLoggingStorageShowResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkLoggingV1NetworkLoggingStorageApiAPI.CreateNetworkLoggingStorage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkLoggingStorageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkLoggingStorageCreateRequest** | [**NetworkLoggingStorageCreateRequest**](NetworkLoggingStorageCreateRequest.md) |  | 

### Return type

[**NetworkLoggingStorageShowResponse**](NetworkLoggingStorageShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkLoggingStorage

> DeleteNetworkLoggingStorage(ctx, networkLoggingStorageId).Execute()

Delete Network Logging Storage



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
	networkLoggingStorageId := "networkLoggingStorageId_example" // string | Network Logging Storage ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkLoggingV1NetworkLoggingStorageApiAPI.DeleteNetworkLoggingStorage(context.Background(), networkLoggingStorageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoggingV1NetworkLoggingStorageApiAPI.DeleteNetworkLoggingStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkLoggingStorageId** | **string** | Network Logging Storage ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkLoggingStorageRequest struct via the builder pattern


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


## ListNetworkLoggingStorages

> NetworkLoggingStorageListResponse ListNetworkLoggingStorages(ctx).ResourceType(resourceType).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Execute()

List Network Logging Storages



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
	resourceType := openapiclient.NetworkLoggingResourceType("FIREWALL") // NetworkLoggingResourceType | Resource Type
	withCount := "true" // string | with count (optional)
	limit := int32(20) // int32 | limit (optional)
	marker := "607e0938521643b5b4b266f343fae693" // string | marker (optional)
	sort := "created_at:desc" // string | sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkLoggingV1NetworkLoggingStorageApiAPI.ListNetworkLoggingStorages(context.Background()).ResourceType(resourceType).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkLoggingV1NetworkLoggingStorageApiAPI.ListNetworkLoggingStorages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkLoggingStorages`: NetworkLoggingStorageListResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkLoggingV1NetworkLoggingStorageApiAPI.ListNetworkLoggingStorages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkLoggingStoragesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | [**NetworkLoggingResourceType**](NetworkLoggingResourceType.md) | Resource Type | 
 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 

### Return type

[**NetworkLoggingStorageListResponse**](NetworkLoggingStorageListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

