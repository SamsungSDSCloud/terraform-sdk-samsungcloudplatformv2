# \SearchengineV1SearchEngineMaintenanceApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchengineSetMaintenance**](SearchengineV1SearchEngineMaintenanceApiAPI.md#SearchengineSetMaintenance) | **Post** /v1/clusters/{cluster_id}/maintenance | Set Maintenance
[**SearchengineUnsetMaintenance**](SearchengineV1SearchEngineMaintenanceApiAPI.md#SearchengineUnsetMaintenance) | **Delete** /v1/clusters/{cluster_id}/maintenance | Unset Maintenance



## SearchengineSetMaintenance

> AsyncResponse SearchengineSetMaintenance(ctx, clusterId).MaintenanceRequest(maintenanceRequest).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID
	maintenanceRequest := *openapiclient.NewMaintenanceRequest(openapiclient.DayOfWeek("MON"), "StartMinute_example", "StartTime_example", "TermHour_example") // MaintenanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchengineV1SearchEngineMaintenanceApiAPI.SearchengineSetMaintenance(context.Background(), clusterId).MaintenanceRequest(maintenanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineMaintenanceApiAPI.SearchengineSetMaintenance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineSetMaintenance`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineMaintenanceApiAPI.SearchengineSetMaintenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineSetMaintenanceRequest struct via the builder pattern


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


## SearchengineUnsetMaintenance

> AsyncResponse SearchengineUnsetMaintenance(ctx, clusterId).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchengineV1SearchEngineMaintenanceApiAPI.SearchengineUnsetMaintenance(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineMaintenanceApiAPI.SearchengineUnsetMaintenance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineUnsetMaintenance`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineMaintenanceApiAPI.SearchengineUnsetMaintenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineUnsetMaintenanceRequest struct via the builder pattern


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

