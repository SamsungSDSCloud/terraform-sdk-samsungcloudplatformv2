# \MariadbV1MariadbLogExportApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariadbExportLog**](MariadbV1MariadbLogExportApiAPI.md#MariadbExportLog) | **Post** /v1/clusters/{cluster_id}/log-export-configs/{log_type}/export | Export Log
[**MariadbListLogExportConfigs**](MariadbV1MariadbLogExportApiAPI.md#MariadbListLogExportConfigs) | **Get** /v1/clusters/{cluster_id}/log-export-configs | List Log Export Configs
[**MariadbRegisterLogExportConfig**](MariadbV1MariadbLogExportApiAPI.md#MariadbRegisterLogExportConfig) | **Post** /v1/clusters/{cluster_id}/log-export-configs | Register Log Export Config
[**MariadbSetLogExportConfig**](MariadbV1MariadbLogExportApiAPI.md#MariadbSetLogExportConfig) | **Put** /v1/clusters/{cluster_id}/log-export-configs/{log_type} | Set Log Export Config
[**MariadbUnregisterLogExportConfig**](MariadbV1MariadbLogExportApiAPI.md#MariadbUnregisterLogExportConfig) | **Delete** /v1/clusters/{cluster_id}/log-export-configs/{log_type} | Unregister Log Export Config



## MariadbExportLog

> AsyncResponse MariadbExportLog(ctx, clusterId, logType).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbLogExportApiAPI.MariadbExportLog(context.Background(), clusterId, logType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbLogExportApiAPI.MariadbExportLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbExportLog`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbLogExportApiAPI.MariadbExportLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbExportLogRequest struct via the builder pattern


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


## MariadbListLogExportConfigs

> LogExportConfigListResponse MariadbListLogExportConfigs(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbLogExportApiAPI.MariadbListLogExportConfigs(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbLogExportApiAPI.MariadbListLogExportConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbListLogExportConfigs`: LogExportConfigListResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbLogExportApiAPI.MariadbListLogExportConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbListLogExportConfigsRequest struct via the builder pattern


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


## MariadbRegisterLogExportConfig

> AsyncResponse MariadbRegisterLogExportConfig(ctx, clusterId).LogExportConfigCreateRequest(logExportConfigCreateRequest).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbLogExportApiAPI.MariadbRegisterLogExportConfig(context.Background(), clusterId).LogExportConfigCreateRequest(logExportConfigCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbLogExportApiAPI.MariadbRegisterLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbRegisterLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbLogExportApiAPI.MariadbRegisterLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbRegisterLogExportConfigRequest struct via the builder pattern


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


## MariadbSetLogExportConfig

> AsyncResponse MariadbSetLogExportConfig(ctx, clusterId, logType).LogExportConfigModifyRequest(logExportConfigModifyRequest).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbLogExportApiAPI.MariadbSetLogExportConfig(context.Background(), clusterId, logType).LogExportConfigModifyRequest(logExportConfigModifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbLogExportApiAPI.MariadbSetLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbSetLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbLogExportApiAPI.MariadbSetLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSetLogExportConfigRequest struct via the builder pattern


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


## MariadbUnregisterLogExportConfig

> AsyncResponse MariadbUnregisterLogExportConfig(ctx, clusterId, logType).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbLogExportApiAPI.MariadbUnregisterLogExportConfig(context.Background(), clusterId, logType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbLogExportApiAPI.MariadbUnregisterLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbUnregisterLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbLogExportApiAPI.MariadbUnregisterLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logType** | [**LogType**](.md) | Log type | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbUnregisterLogExportConfigRequest struct via the builder pattern


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

