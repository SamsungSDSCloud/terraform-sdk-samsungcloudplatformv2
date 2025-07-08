# \MariadbV1MariadbArchiveApiAPI

All URIs are relative to *https://scp-dbaas-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariadbDeleteArchiveLog**](MariadbV1MariadbArchiveApiAPI.md#MariadbDeleteArchiveLog) | **Delete** /v1/clusters/{cluster_id}/archive | Delete Archive Log
[**MariadbSetArchiveConfig**](MariadbV1MariadbArchiveApiAPI.md#MariadbSetArchiveConfig) | **Put** /v1/clusters/{cluster_id}/archive | Set Archive Config
[**MariadbShowArchiveConfig**](MariadbV1MariadbArchiveApiAPI.md#MariadbShowArchiveConfig) | **Get** /v1/clusters/{cluster_id}/archive | Show Archive Config
[**MariadbSyncArchiveConfig**](MariadbV1MariadbArchiveApiAPI.md#MariadbSyncArchiveConfig) | **Post** /v1/clusters/{cluster_id}/archive/sync | Synchronize Archive Config



## MariadbDeleteArchiveLog

> AsyncResponse MariadbDeleteArchiveLog(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbArchiveApiAPI.MariadbDeleteArchiveLog(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbArchiveApiAPI.MariadbDeleteArchiveLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbDeleteArchiveLog`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbArchiveApiAPI.MariadbDeleteArchiveLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbDeleteArchiveLogRequest struct via the builder pattern


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


## MariadbSetArchiveConfig

> AsyncResponse MariadbSetArchiveConfig(ctx, clusterId).ArchiveConfigSetRequest(archiveConfigSetRequest).Execute()

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
	archiveConfigSetRequest := *openapiclient.NewArchiveConfigSetRequest() // ArchiveConfigSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MariadbV1MariadbArchiveApiAPI.MariadbSetArchiveConfig(context.Background(), clusterId).ArchiveConfigSetRequest(archiveConfigSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbArchiveApiAPI.MariadbSetArchiveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbSetArchiveConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbArchiveApiAPI.MariadbSetArchiveConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSetArchiveConfigRequest struct via the builder pattern


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


## MariadbShowArchiveConfig

> ArchiveConfigResponse MariadbShowArchiveConfig(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbArchiveApiAPI.MariadbShowArchiveConfig(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbArchiveApiAPI.MariadbShowArchiveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbShowArchiveConfig`: ArchiveConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbArchiveApiAPI.MariadbShowArchiveConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbShowArchiveConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArchiveConfigResponse**](ArchiveConfigResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MariadbSyncArchiveConfig

> AsyncResponse MariadbSyncArchiveConfig(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbArchiveApiAPI.MariadbSyncArchiveConfig(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbArchiveApiAPI.MariadbSyncArchiveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbSyncArchiveConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbArchiveApiAPI.MariadbSyncArchiveConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSyncArchiveConfigRequest struct via the builder pattern


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

