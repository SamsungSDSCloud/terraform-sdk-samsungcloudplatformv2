# \MysqlV1MysqlLogExportApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlExportLog**](MysqlV1MysqlLogExportApiAPI.md#MysqlExportLog) | **Post** /v1/clusters/{cluster_id}/log-export-configs/{log_type}/export | Export Log
[**MysqlListLogExportConfigs**](MysqlV1MysqlLogExportApiAPI.md#MysqlListLogExportConfigs) | **Get** /v1/clusters/{cluster_id}/log-export-configs | List Log Export Configs
[**MysqlRegisterLogExportConfig**](MysqlV1MysqlLogExportApiAPI.md#MysqlRegisterLogExportConfig) | **Post** /v1/clusters/{cluster_id}/log-export-configs | Register Log Export Config
[**MysqlSetLogExportConfig**](MysqlV1MysqlLogExportApiAPI.md#MysqlSetLogExportConfig) | **Put** /v1/clusters/{cluster_id}/log-export-configs/{log_type} | Set Log Export Config
[**MysqlUnregisterLogExportConfig**](MysqlV1MysqlLogExportApiAPI.md#MysqlUnregisterLogExportConfig) | **Delete** /v1/clusters/{cluster_id}/log-export-configs/{log_type} | Unregister Log Export Config



## MysqlExportLog

> AsyncResponse MysqlExportLog(ctx, clusterId, logType).Execute()

Export Log



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
	logType := openapiclient.LogType("alert") // LogType | Log type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlLogExportApiAPI.MysqlExportLog(context.Background(), clusterId, logType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlLogExportApiAPI.MysqlExportLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlExportLog`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlLogExportApiAPI.MysqlExportLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlExportLogRequest struct via the builder pattern


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


## MysqlListLogExportConfigs

> LogExportConfigListResponse MysqlListLogExportConfigs(ctx, clusterId).Execute()

List Log Export Configs



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
	resp, r, err := apiClient.MysqlV1MysqlLogExportApiAPI.MysqlListLogExportConfigs(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlLogExportApiAPI.MysqlListLogExportConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlListLogExportConfigs`: LogExportConfigListResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlLogExportApiAPI.MysqlListLogExportConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlListLogExportConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogExportConfigListResponse**](LogExportConfigListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlRegisterLogExportConfig

> AsyncResponse MysqlRegisterLogExportConfig(ctx, clusterId).LogExportConfigCreateRequest(logExportConfigCreateRequest).Execute()

Register Log Export Config



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
	logExportConfigCreateRequest := *openapiclient.NewLogExportConfigCreateRequest("AccessKey_example", "examplebucket", false, "alert", "ScheduleDayOfMonth_example", "TODO", "ScheduleFrequencyType_example", "ScheduleHour_example", "SecretKey_example") // LogExportConfigCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlLogExportApiAPI.MysqlRegisterLogExportConfig(context.Background(), clusterId).LogExportConfigCreateRequest(logExportConfigCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlLogExportApiAPI.MysqlRegisterLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlRegisterLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlLogExportApiAPI.MysqlRegisterLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlRegisterLogExportConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logExportConfigCreateRequest** | [**LogExportConfigCreateRequest**](LogExportConfigCreateRequest.md) |  | 

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


## MysqlSetLogExportConfig

> AsyncResponse MysqlSetLogExportConfig(ctx, clusterId, logType).LogExportConfigModifyRequest(logExportConfigModifyRequest).Execute()

Set Log Export Config



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
	logType := openapiclient.LogType("alert") // LogType | Log type
	logExportConfigModifyRequest := *openapiclient.NewLogExportConfigModifyRequest("AccessKey_example", false, "ScheduleDayOfMonth_example", "TODO", "ScheduleFrequencyType_example", "ScheduleHour_example", "SecretKey_example") // LogExportConfigModifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlLogExportApiAPI.MysqlSetLogExportConfig(context.Background(), clusterId, logType).LogExportConfigModifyRequest(logExportConfigModifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlLogExportApiAPI.MysqlSetLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSetLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlLogExportApiAPI.MysqlSetLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSetLogExportConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **logExportConfigModifyRequest** | [**LogExportConfigModifyRequest**](LogExportConfigModifyRequest.md) |  | 

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


## MysqlUnregisterLogExportConfig

> AsyncResponse MysqlUnregisterLogExportConfig(ctx, clusterId, logType).Execute()

Unregister Log Export Config



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
	logType := openapiclient.LogType("alert") // LogType | Log type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlLogExportApiAPI.MysqlUnregisterLogExportConfig(context.Background(), clusterId, logType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlLogExportApiAPI.MysqlUnregisterLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlUnregisterLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlLogExportApiAPI.MysqlUnregisterLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlUnregisterLogExportConfigRequest struct via the builder pattern


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

