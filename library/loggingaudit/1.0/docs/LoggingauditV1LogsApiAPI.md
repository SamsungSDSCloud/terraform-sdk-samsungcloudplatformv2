# \LoggingauditV1LogsApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadLogs**](LoggingauditV1LogsApiAPI.md#DownloadLogs) | **Post** /v1/logs/download | DownloadLogs
[**ListLogs**](LoggingauditV1LogsApiAPI.md#ListLogs) | **Get** /v1/logs | ListLogs
[**ShowLog**](LoggingauditV1LogsApiAPI.md#ShowLog) | **Get** /v1/logs/{logging_id} | ShowLog



## DownloadLogs

> LogDownloadResponse DownloadLogs(ctx).LogDownloadRequest(logDownloadRequest).Execute()

DownloadLogs



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
	logDownloadRequest := *openapiclient.NewLogDownloadRequest("2025-05-21T06:00:00Z", "JSON", "2025-02-20T06:00:00Z", "Asia/Seoul") // LogDownloadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoggingauditV1LogsApiAPI.DownloadLogs(context.Background()).LogDownloadRequest(logDownloadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoggingauditV1LogsApiAPI.DownloadLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadLogs`: LogDownloadResponse
	fmt.Fprintf(os.Stdout, "Response from `LoggingauditV1LogsApiAPI.DownloadLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logDownloadRequest** | [**LogDownloadRequest**](LogDownloadRequest.md) |  | 

### Return type

[**LogDownloadResponse**](LogDownloadResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogs

> LogListResponse ListLogs(ctx).StartAt(startAt).EndAt(endAt).Size(size).Page(page).Sort(sort).ResourceId(resourceId).ResourceName(resourceName).ResourceType(resourceType).ProductName(productName).ProductType(productType).Status(status).EventType(eventType).UserId(userId).UserName(userName).Region(region).RootResourceId(rootResourceId).ServiceType(serviceType).Execute()

ListLogs



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
	startAt := "2025-02-20T06:00:00Z" // string | Query start date
	endAt := "2025-05-21T06:00:00Z" // string | Query finish date
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	resourceId := "resourceId_example" // string | Resource ID (optional)
	resourceName := "resourceName_example" // string | Resource name (optional)
	resourceType := "resourceType_example" // string | Resource type (optional)
	productName := "productName_example" // string | Product name (optional)
	productType := "productType_example" // string | Product type (optional)
	status := "status_example" // string | Task result (optional)
	eventType := "eventType_example" // string | Event type (optional)
	userId := "userId_example" // string | User ID (optional)
	userName := "userName_example" // string | Username (optional)
	region := "region_example" // string | Region name (optional)
	rootResourceId := "rootResourceId_example" // string | Root resource ID (optional)
	serviceType := "serviceType_example" // string | Service type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoggingauditV1LogsApiAPI.ListLogs(context.Background()).StartAt(startAt).EndAt(endAt).Size(size).Page(page).Sort(sort).ResourceId(resourceId).ResourceName(resourceName).ResourceType(resourceType).ProductName(productName).ProductType(productType).Status(status).EventType(eventType).UserId(userId).UserName(userName).Region(region).RootResourceId(rootResourceId).ServiceType(serviceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoggingauditV1LogsApiAPI.ListLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLogs`: LogListResponse
	fmt.Fprintf(os.Stdout, "Response from `LoggingauditV1LogsApiAPI.ListLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** | Query start date | 
 **endAt** | **string** | Query finish date | 
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **resourceId** | **string** | Resource ID | 
 **resourceName** | **string** | Resource name | 
 **resourceType** | **string** | Resource type | 
 **productName** | **string** | Product name | 
 **productType** | **string** | Product type | 
 **status** | **string** | Task result | 
 **eventType** | **string** | Event type | 
 **userId** | **string** | User ID | 
 **userName** | **string** | Username | 
 **region** | **string** | Region name | 
 **rootResourceId** | **string** | Root resource ID | 
 **serviceType** | **string** | Service type | 

### Return type

[**LogListResponse**](LogListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowLog

> LogShowResponse ShowLog(ctx, loggingId).Execute()

ShowLog



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
	loggingId := "loggingId_example" // string | Log ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoggingauditV1LogsApiAPI.ShowLog(context.Background(), loggingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoggingauditV1LogsApiAPI.ShowLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowLog`: LogShowResponse
	fmt.Fprintf(os.Stdout, "Response from `LoggingauditV1LogsApiAPI.ShowLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loggingId** | **string** | Log ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogShowResponse**](LogShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

