# \MariadbV1MariadbLogExportApiAPI

All URIs are relative to *https://scp-dbaas-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariadbCreateLogExportConfig**](MariadbV1MariadbLogExportApiAPI.md#MariadbCreateLogExportConfig) | **Post** /v1/clusters/{cluster_id}/log-export-configs | Create Log Export Config
[**MariadbDeleteLogExportConfig**](MariadbV1MariadbLogExportApiAPI.md#MariadbDeleteLogExportConfig) | **Delete** /v1/clusters/{cluster_id}/log-export-configs/{log_export_config_id} | Delete Log Export Config
[**MariadbExportLog**](MariadbV1MariadbLogExportApiAPI.md#MariadbExportLog) | **Post** /v1/clusters/{cluster_id}/log-export-configs/{log_export_config_id}/export | Export Log
[**MariadbListLogExportConfigs**](MariadbV1MariadbLogExportApiAPI.md#MariadbListLogExportConfigs) | **Get** /v1/clusters/{cluster_id}/log-export-configs | List Log Export Configs
[**MariadbSetLogExportConfig**](MariadbV1MariadbLogExportApiAPI.md#MariadbSetLogExportConfig) | **Put** /v1/clusters/{cluster_id}/log-export-configs/{log_export_config_id} | Set Log Export Config
[**MariadbShowLogExportConfig**](MariadbV1MariadbLogExportApiAPI.md#MariadbShowLogExportConfig) | **Get** /v1/clusters/{cluster_id}/log-export-configs/{log_export_config_id} | Show Log Export Config



## MariadbCreateLogExportConfig

> AsyncResponse MariadbCreateLogExportConfig(ctx, clusterId).LogExportConfigCreateRequestDTO(logExportConfigCreateRequestDTO).Execute()

Create Log Export Config



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
	resp, r, err := apiClient.MariadbV1MariadbLogExportApiAPI.MariadbCreateLogExportConfig(context.Background(), clusterId).LogExportConfigCreateRequestDTO(logExportConfigCreateRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbLogExportApiAPI.MariadbCreateLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbCreateLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbLogExportApiAPI.MariadbCreateLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbCreateLogExportConfigRequest struct via the builder pattern


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


## MariadbDeleteLogExportConfig

> AsyncResponse MariadbDeleteLogExportConfig(ctx, clusterId, logExportConfigId).Execute()

Delete Log Export Config



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
	logExportConfigId := "logExportConfigId_example" // string | Log export config id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MariadbV1MariadbLogExportApiAPI.MariadbDeleteLogExportConfig(context.Background(), clusterId, logExportConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbLogExportApiAPI.MariadbDeleteLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbDeleteLogExportConfig`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbLogExportApiAPI.MariadbDeleteLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logExportConfigId** | **string** | Log export config id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbDeleteLogExportConfigRequest struct via the builder pattern


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


## MariadbExportLog

> LogExportConfigListResponse MariadbExportLog(ctx, clusterId, logExportConfigId).Execute()

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
	logExportConfigId := "logExportConfigId_example" // string | Log export config id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MariadbV1MariadbLogExportApiAPI.MariadbExportLog(context.Background(), clusterId, logExportConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbLogExportApiAPI.MariadbExportLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbExportLog`: LogExportConfigListResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbLogExportApiAPI.MariadbExportLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logExportConfigId** | **string** | Log export config id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbExportLogRequest struct via the builder pattern


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
	clusterId := "clusterId_example" // string | Cluster ID

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


## MariadbSetLogExportConfig

> AsyncResponse MariadbSetLogExportConfig(ctx, clusterId, logExportConfigId).LogExportConfigModifyRequestDTO(logExportConfigModifyRequestDTO).Execute()

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
	logExportConfigId := "logExportConfigId_example" // string | Log export config id
	logExportConfigModifyRequestDTO := *openapiclient.NewLogExportConfigModifyRequestDTO("AccessKey_example", false, "ScheduleDayOfMonth_example", "TODO", "ScheduleFrequencyType_example", "ScheduleHour_example", "SecretKey_example") // LogExportConfigModifyRequestDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MariadbV1MariadbLogExportApiAPI.MariadbSetLogExportConfig(context.Background(), clusterId, logExportConfigId).LogExportConfigModifyRequestDTO(logExportConfigModifyRequestDTO).Execute()
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
**logExportConfigId** | **string** | Log export config id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSetLogExportConfigRequest struct via the builder pattern


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


## MariadbShowLogExportConfig

> LogExportConfigListResponse MariadbShowLogExportConfig(ctx, clusterId, logExportConfigId).Execute()

Show Log Export Config



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
	logExportConfigId := "logExportConfigId_example" // string | Log export config id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MariadbV1MariadbLogExportApiAPI.MariadbShowLogExportConfig(context.Background(), clusterId, logExportConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbLogExportApiAPI.MariadbShowLogExportConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbShowLogExportConfig`: LogExportConfigListResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbLogExportApiAPI.MariadbShowLogExportConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 
**logExportConfigId** | **string** | Log export config id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbShowLogExportConfigRequest struct via the builder pattern


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

