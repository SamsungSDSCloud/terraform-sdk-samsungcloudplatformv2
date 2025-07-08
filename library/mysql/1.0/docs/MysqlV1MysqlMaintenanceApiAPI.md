# \MysqlV1MysqlMaintenanceApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlSetMaintenance**](MysqlV1MysqlMaintenanceApiAPI.md#MysqlSetMaintenance) | **Post** /v1/clusters/{cluster_id}/maintenance | Set Maintenance
[**MysqlUnsetMaintenance**](MysqlV1MysqlMaintenanceApiAPI.md#MysqlUnsetMaintenance) | **Delete** /v1/clusters/{cluster_id}/maintenance | Unset Maintenance



## MysqlSetMaintenance

> AsyncResponse MysqlSetMaintenance(ctx, clusterId).MaintenanceRequest(maintenanceRequest).Execute()

Set Maintenance



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
	clusterId := "clusterId_example" // string | DB cluster id
	maintenanceRequest := *openapiclient.NewMaintenanceRequest(openapiclient.DayOfWeek("MON"), "StartMinute_example", "StartTime_example", "TermHour_example") // MaintenanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlMaintenanceApiAPI.MysqlSetMaintenance(context.Background(), clusterId).MaintenanceRequest(maintenanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlMaintenanceApiAPI.MysqlSetMaintenance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSetMaintenance`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlMaintenanceApiAPI.MysqlSetMaintenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | DB cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSetMaintenanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenanceRequest** | [**MaintenanceRequest**](MaintenanceRequest.md) |  | 

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


## MysqlUnsetMaintenance

> AsyncResponse MysqlUnsetMaintenance(ctx, clusterId).Execute()

Unset Maintenance



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
	clusterId := "clusterId_example" // string | DB cluster id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlMaintenanceApiAPI.MysqlUnsetMaintenance(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlMaintenanceApiAPI.MysqlUnsetMaintenance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlUnsetMaintenance`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlMaintenanceApiAPI.MysqlUnsetMaintenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | DB cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlUnsetMaintenanceRequest struct via the builder pattern


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

