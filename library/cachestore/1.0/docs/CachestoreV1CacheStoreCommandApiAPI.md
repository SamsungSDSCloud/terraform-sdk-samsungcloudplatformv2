# \CachestoreV1CacheStoreCommandApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CachestoreListCommands**](CachestoreV1CacheStoreCommandApiAPI.md#CachestoreListCommands) | **Get** /v1/clusters/{cluster_id}/commands | List Commands
[**CachestoreSetCommands**](CachestoreV1CacheStoreCommandApiAPI.md#CachestoreSetCommands) | **Put** /v1/clusters/{cluster_id}/commands | Update Commands
[**CachestoreSyncCommands**](CachestoreV1CacheStoreCommandApiAPI.md#CachestoreSyncCommands) | **Post** /v1/clusters/{cluster_id}/commands/sync | Synchronize Commands



## CachestoreListCommands

> CommandListResponse CachestoreListCommands(ctx, clusterId).Execute()

List Commands



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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreCommandApiAPI.CachestoreListCommands(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreCommandApiAPI.CachestoreListCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreListCommands`: CommandListResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreCommandApiAPI.CachestoreListCommands`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreListCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommandListResponse**](CommandListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CachestoreSetCommands

> AsyncResponse CachestoreSetCommands(ctx, clusterId).ModifyCommandRequest(modifyCommandRequest).Execute()

Update Commands



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
	clusterId := "clusterId_example" // string | Cluster ID
	modifyCommandRequest := *openapiclient.NewModifyCommandRequest([]openapiclient.ModifyCommandItem{*openapiclient.NewModifyCommandItem("Id_example", "Name_example", "NewValue_example")}) // ModifyCommandRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreCommandApiAPI.CachestoreSetCommands(context.Background(), clusterId).ModifyCommandRequest(modifyCommandRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreCommandApiAPI.CachestoreSetCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreSetCommands`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreCommandApiAPI.CachestoreSetCommands`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreSetCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyCommandRequest** | [**ModifyCommandRequest**](ModifyCommandRequest.md) |  | 

### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CachestoreSyncCommands

> AsyncResponse CachestoreSyncCommands(ctx, clusterId).Execute()

Synchronize Commands



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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreCommandApiAPI.CachestoreSyncCommands(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreCommandApiAPI.CachestoreSyncCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreSyncCommands`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreCommandApiAPI.CachestoreSyncCommands`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreSyncCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

