# \EpasV1EpasLogExportApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EpasExportLog**](EpasV1EpasLogExportApiAPI.md#EpasExportLog) | **Post** /v1/clusters/{cluster_id}/log-export-configs/{log_type}/export | Export Log
[**EpasListLogExportConfigs**](EpasV1EpasLogExportApiAPI.md#EpasListLogExportConfigs) | **Get** /v1/clusters/{cluster_id}/log-export-configs | List Log Export Configs
[**EpasRegisterLogExportConfig**](EpasV1EpasLogExportApiAPI.md#EpasRegisterLogExportConfig) | **Post** /v1/clusters/{cluster_id}/log-export-configs | Register Log Export Config
[**EpasSetLogExportConfig**](EpasV1EpasLogExportApiAPI.md#EpasSetLogExportConfig) | **Put** /v1/clusters/{cluster_id}/log-export-configs/{log_type} | Set Log Export Config
[**EpasUnregisterLogExportConfig**](EpasV1EpasLogExportApiAPI.md#EpasUnregisterLogExportConfig) | **Delete** /v1/clusters/{cluster_id}/log-export-configs/{log_type} | Unregister Log Export Config



## EpasExportLog

> AsyncResponse EpasExportLog(ctx, clusterId, logType).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID
	logType := openapiclient.LogType("alert") // LogType | Log type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasLogExportApiAPI.EpasExportLog(context.Background(), clusterId, logType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasLogExportApiAPI.EpasExportLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasExportLog`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasLogExportApiAPI.EpasExportLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasExportLogRequest struct via the builder pattern


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


## EpasListLogExportConfigs

> LogExportConfigListResponse EpasListLogExportConfigs(ctx, clusterId).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasLogExportApiAPI.EpasListLogExportConfigs(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasLogExportApiAPI.EpasListLogExportConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasListLogExportConfigs`: LogExportConfigListResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasLogExportApiAPI.EpasListLogExportConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasListLogExportConfigsRequest struct via the builder pattern


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


## EpasRegisterLogExportConfig

> AsyncResponse EpasRegisterLogExportConfig(ctx, clusterId).LogExportConfigCreateRequestDTO(logExportConfigCreateRequestDTO).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID
	logExportConfigCreateRequestDTO := *openapiclient.NewLogExportConfigCreateRequestDTO("AccessKey_example", "BucketName_example", false, "LogType_example", "ScheduleDayOfMonth_example", "TODO", "ScheduleFrequencyType_example", "ScheduleHour_example", "SecretKey_example") // LogExportConfigCreateRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasLogExportApiAPI.EpasRegisterLogExportConfig(context.Background(), clusterId).LogExportConfigCreateRequestDTO(logExportConfigCreateRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasLogExportApiAPI.EpasRegisterLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasRegisterLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasLogExportApiAPI.EpasRegisterLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasRegisterLogExportConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logExportConfigCreateRequestDTO** | [**LogExportConfigCreateRequestDTO**](LogExportConfigCreateRequestDTO.md) |  | 

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


## EpasSetLogExportConfig

> AsyncResponse EpasSetLogExportConfig(ctx, clusterId, logType).LogExportConfigModifyRequestDTO(logExportConfigModifyRequestDTO).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID
	logType := openapiclient.LogType("alert") // LogType | Log type
	logExportConfigModifyRequestDTO := *openapiclient.NewLogExportConfigModifyRequestDTO("AccessKey_example", false, "ScheduleDayOfMonth_example", "TODO", "ScheduleFrequencyType_example", "ScheduleHour_example", "SecretKey_example") // LogExportConfigModifyRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasLogExportApiAPI.EpasSetLogExportConfig(context.Background(), clusterId, logType).LogExportConfigModifyRequestDTO(logExportConfigModifyRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasLogExportApiAPI.EpasSetLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSetLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasLogExportApiAPI.EpasSetLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSetLogExportConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **logExportConfigModifyRequestDTO** | [**LogExportConfigModifyRequestDTO**](LogExportConfigModifyRequestDTO.md) |  | 

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


## EpasUnregisterLogExportConfig

> AsyncResponse EpasUnregisterLogExportConfig(ctx, clusterId, logType).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID
	logType := openapiclient.LogType("alert") // LogType | Log type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasLogExportApiAPI.EpasUnregisterLogExportConfig(context.Background(), clusterId, logType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasLogExportApiAPI.EpasUnregisterLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasUnregisterLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasLogExportApiAPI.EpasUnregisterLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasUnregisterLogExportConfigRequest struct via the builder pattern


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

