# \SqlserverV1SqlserverLogExportApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SqlserverExportLog**](SqlserverV1SqlserverLogExportApiAPI.md#SqlserverExportLog) | **Post** /v1/clusters/{cluster_id}/log-export-configs/{log_type}/export | Export Log
[**SqlserverListLogExportConfigs**](SqlserverV1SqlserverLogExportApiAPI.md#SqlserverListLogExportConfigs) | **Get** /v1/clusters/{cluster_id}/log-export-configs | List Log Export Configs
[**SqlserverRegisterLogExportConfig**](SqlserverV1SqlserverLogExportApiAPI.md#SqlserverRegisterLogExportConfig) | **Post** /v1/clusters/{cluster_id}/log-export-configs | Register Log Export Config
[**SqlserverSetLogExportConfig**](SqlserverV1SqlserverLogExportApiAPI.md#SqlserverSetLogExportConfig) | **Put** /v1/clusters/{cluster_id}/log-export-configs/{log_type} | Set Log Export Config
[**SqlserverUnregisterLogExportConfig**](SqlserverV1SqlserverLogExportApiAPI.md#SqlserverUnregisterLogExportConfig) | **Delete** /v1/clusters/{cluster_id}/log-export-configs/{log_type} | Unregister Log Export Config



## SqlserverExportLog

> AsyncResponse SqlserverExportLog(ctx, clusterId, logType).Execute()

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
	resp, r, err := apiClient.SqlserverV1SqlserverLogExportApiAPI.SqlserverExportLog(context.Background(), clusterId, logType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverLogExportApiAPI.SqlserverExportLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverExportLog`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverLogExportApiAPI.SqlserverExportLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverExportLogRequest struct via the builder pattern


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


## SqlserverListLogExportConfigs

> LogExportConfigListResponse SqlserverListLogExportConfigs(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.SqlserverV1SqlserverLogExportApiAPI.SqlserverListLogExportConfigs(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverLogExportApiAPI.SqlserverListLogExportConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverListLogExportConfigs`: LogExportConfigListResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverLogExportApiAPI.SqlserverListLogExportConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverListLogExportConfigsRequest struct via the builder pattern


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


## SqlserverRegisterLogExportConfig

> AsyncResponse SqlserverRegisterLogExportConfig(ctx, clusterId).LogExportConfigCreateRequest(logExportConfigCreateRequest).Execute()

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
	resp, r, err := apiClient.SqlserverV1SqlserverLogExportApiAPI.SqlserverRegisterLogExportConfig(context.Background(), clusterId).LogExportConfigCreateRequest(logExportConfigCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverLogExportApiAPI.SqlserverRegisterLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverRegisterLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverLogExportApiAPI.SqlserverRegisterLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverRegisterLogExportConfigRequest struct via the builder pattern


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


## SqlserverSetLogExportConfig

> AsyncResponse SqlserverSetLogExportConfig(ctx, clusterId, logType).LogExportConfigModifyRequest(logExportConfigModifyRequest).Execute()

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
	resp, r, err := apiClient.SqlserverV1SqlserverLogExportApiAPI.SqlserverSetLogExportConfig(context.Background(), clusterId, logType).LogExportConfigModifyRequest(logExportConfigModifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverLogExportApiAPI.SqlserverSetLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverSetLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverLogExportApiAPI.SqlserverSetLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverSetLogExportConfigRequest struct via the builder pattern


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


## SqlserverUnregisterLogExportConfig

> AsyncResponse SqlserverUnregisterLogExportConfig(ctx, clusterId, logType).Execute()

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
	resp, r, err := apiClient.SqlserverV1SqlserverLogExportApiAPI.SqlserverUnregisterLogExportConfig(context.Background(), clusterId, logType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverLogExportApiAPI.SqlserverUnregisterLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverUnregisterLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverLogExportApiAPI.SqlserverUnregisterLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverUnregisterLogExportConfigRequest struct via the builder pattern


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

