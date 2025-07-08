# \EpasV1EpasArchiveApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EpasDeleteArchiveLog**](EpasV1EpasArchiveApiAPI.md#EpasDeleteArchiveLog) | **Delete** /v1/clusters/{cluster_id}/archive | Delete Archive Log
[**EpasSetArchiveConfig**](EpasV1EpasArchiveApiAPI.md#EpasSetArchiveConfig) | **Put** /v1/clusters/{cluster_id}/archive | Set Archive Config
[**EpasShowArchiveConfig**](EpasV1EpasArchiveApiAPI.md#EpasShowArchiveConfig) | **Get** /v1/clusters/{cluster_id}/archive | Show Archive Config
[**EpasSyncArchiveConfig**](EpasV1EpasArchiveApiAPI.md#EpasSyncArchiveConfig) | **Post** /v1/clusters/{cluster_id}/archive/sync | Synchronize Archive Config



## EpasDeleteArchiveLog

> AsyncResponse EpasDeleteArchiveLog(ctx, clusterId).Execute()

Delete Archive Log



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
	resp, r, err := apiClient.EpasV1EpasArchiveApiAPI.EpasDeleteArchiveLog(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasArchiveApiAPI.EpasDeleteArchiveLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasDeleteArchiveLog`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasArchiveApiAPI.EpasDeleteArchiveLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasDeleteArchiveLogRequest struct via the builder pattern


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


## EpasSetArchiveConfig

> AsyncResponse EpasSetArchiveConfig(ctx, clusterId).ArchiveConfigSetRequest(archiveConfigSetRequest).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID
	archiveConfigSetRequest := *openapiclient.NewArchiveConfigSetRequest(false) // ArchiveConfigSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasArchiveApiAPI.EpasSetArchiveConfig(context.Background(), clusterId).ArchiveConfigSetRequest(archiveConfigSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasArchiveApiAPI.EpasSetArchiveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSetArchiveConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasArchiveApiAPI.EpasSetArchiveConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSetArchiveConfigRequest struct via the builder pattern


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


## EpasShowArchiveConfig

> ArchiveConfigDetailResponse EpasShowArchiveConfig(ctx, clusterId).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasArchiveApiAPI.EpasShowArchiveConfig(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasArchiveApiAPI.EpasShowArchiveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasShowArchiveConfig`: ArchiveConfigDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasArchiveApiAPI.EpasShowArchiveConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasShowArchiveConfigRequest struct via the builder pattern


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


## EpasSyncArchiveConfig

> AsyncResponse EpasSyncArchiveConfig(ctx, clusterId).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasArchiveApiAPI.EpasSyncArchiveConfig(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasArchiveApiAPI.EpasSyncArchiveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSyncArchiveConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasArchiveApiAPI.EpasSyncArchiveConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSyncArchiveConfigRequest struct via the builder pattern


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

