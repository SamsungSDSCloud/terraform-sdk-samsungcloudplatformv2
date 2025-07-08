# \MysqlV1MysqlArchiveApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlSetArchiveConfig**](MysqlV1MysqlArchiveApiAPI.md#MysqlSetArchiveConfig) | **Put** /v1/clusters/{cluster_id}/archive | Set Archive Config
[**MysqlShowArchiveConfig**](MysqlV1MysqlArchiveApiAPI.md#MysqlShowArchiveConfig) | **Get** /v1/clusters/{cluster_id}/archive | Show Archive Config
[**MysqlSyncArchiveConfig**](MysqlV1MysqlArchiveApiAPI.md#MysqlSyncArchiveConfig) | **Post** /v1/clusters/{cluster_id}/archive/sync | Synchronize Archive Config



## MysqlSetArchiveConfig

> AsyncResponse MysqlSetArchiveConfig(ctx, clusterId).ArchiveConfigSetRequest(archiveConfigSetRequest).Execute()

Set Archive Config



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
	clusterId := "109a585ae67b4e8482fdafc8a4a5be74" // string | Cluster ID
	archiveConfigSetRequest := *openapiclient.NewArchiveConfigSetRequest(true) // ArchiveConfigSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlArchiveApiAPI.MysqlSetArchiveConfig(context.Background(), clusterId).ArchiveConfigSetRequest(archiveConfigSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlArchiveApiAPI.MysqlSetArchiveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSetArchiveConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlArchiveApiAPI.MysqlSetArchiveConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSetArchiveConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archiveConfigSetRequest** | [**ArchiveConfigSetRequest**](ArchiveConfigSetRequest.md) |  | 

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


## MysqlShowArchiveConfig

> ArchiveConfigDetailResponse MysqlShowArchiveConfig(ctx, clusterId).Execute()

Show Archive Config



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
	clusterId := "109a585ae67b4e8482fdafc8a4a5be74" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlArchiveApiAPI.MysqlShowArchiveConfig(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlArchiveApiAPI.MysqlShowArchiveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlShowArchiveConfig`: ArchiveConfigDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlArchiveApiAPI.MysqlShowArchiveConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlShowArchiveConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArchiveConfigDetailResponse**](ArchiveConfigDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlSyncArchiveConfig

> AsyncResponse MysqlSyncArchiveConfig(ctx, clusterId).Execute()

Synchronize Archive Config



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
	clusterId := "109a585ae67b4e8482fdafc8a4a5be74" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlArchiveApiAPI.MysqlSyncArchiveConfig(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlArchiveApiAPI.MysqlSyncArchiveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSyncArchiveConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlArchiveApiAPI.MysqlSyncArchiveConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSyncArchiveConfigRequest struct via the builder pattern


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

