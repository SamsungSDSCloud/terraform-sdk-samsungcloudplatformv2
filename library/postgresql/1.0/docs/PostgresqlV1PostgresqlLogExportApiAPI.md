# \PostgresqlV1PostgresqlLogExportApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostgresqlExportLog**](PostgresqlV1PostgresqlLogExportApiAPI.md#PostgresqlExportLog) | **Post** /v1/clusters/{cluster_id}/log-export-configs/{log_type}/export | Export Log
[**PostgresqlListLogExportConfigs**](PostgresqlV1PostgresqlLogExportApiAPI.md#PostgresqlListLogExportConfigs) | **Get** /v1/clusters/{cluster_id}/log-export-configs | List Log Export Configs
[**PostgresqlRegisterLogExportConfig**](PostgresqlV1PostgresqlLogExportApiAPI.md#PostgresqlRegisterLogExportConfig) | **Post** /v1/clusters/{cluster_id}/log-export-configs | Register Log Export Config
[**PostgresqlSetLogExportConfig**](PostgresqlV1PostgresqlLogExportApiAPI.md#PostgresqlSetLogExportConfig) | **Put** /v1/clusters/{cluster_id}/log-export-configs/{log_type} | Set Log Export Config
[**PostgresqlUnregisterLogExportConfig**](PostgresqlV1PostgresqlLogExportApiAPI.md#PostgresqlUnregisterLogExportConfig) | **Delete** /v1/clusters/{cluster_id}/log-export-configs/{log_type} | Unregister Log Export Config



## PostgresqlExportLog

> AsyncResponse PostgresqlExportLog(ctx, clusterId, logType).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlExportLog(context.Background(), clusterId, logType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlExportLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlExportLog`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlExportLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlExportLogRequest struct via the builder pattern


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


## PostgresqlListLogExportConfigs

> LogExportConfigListResponse PostgresqlListLogExportConfigs(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlListLogExportConfigs(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlListLogExportConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlListLogExportConfigs`: LogExportConfigListResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlListLogExportConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlListLogExportConfigsRequest struct via the builder pattern


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


## PostgresqlRegisterLogExportConfig

> AsyncResponse PostgresqlRegisterLogExportConfig(ctx, clusterId).LogExportConfigCreateRequestDTO(logExportConfigCreateRequestDTO).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlRegisterLogExportConfig(context.Background(), clusterId).LogExportConfigCreateRequestDTO(logExportConfigCreateRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlRegisterLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlRegisterLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlRegisterLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlRegisterLogExportConfigRequest struct via the builder pattern


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


## PostgresqlSetLogExportConfig

> AsyncResponse PostgresqlSetLogExportConfig(ctx, clusterId, logType).LogExportConfigModifyRequestDTO(logExportConfigModifyRequestDTO).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlSetLogExportConfig(context.Background(), clusterId, logType).LogExportConfigModifyRequestDTO(logExportConfigModifyRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlSetLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlSetLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlSetLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlSetLogExportConfigRequest struct via the builder pattern


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


## PostgresqlUnregisterLogExportConfig

> AsyncResponse PostgresqlUnregisterLogExportConfig(ctx, clusterId, logType).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlUnregisterLogExportConfig(context.Background(), clusterId, logType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlUnregisterLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlUnregisterLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlLogExportApiAPI.PostgresqlUnregisterLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlUnregisterLogExportConfigRequest struct via the builder pattern


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

