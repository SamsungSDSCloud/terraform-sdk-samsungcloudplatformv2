# \MariadbV1MariadbMaintenanceApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariadbSetMaintenance**](MariadbV1MariadbMaintenanceApiAPI.md#MariadbSetMaintenance) | **Post** /v1/clusters/{cluster_id}/maintenance | Set Maintenance
[**MariadbUnsetMaintenance**](MariadbV1MariadbMaintenanceApiAPI.md#MariadbUnsetMaintenance) | **Delete** /v1/clusters/{cluster_id}/maintenance | Unset Maintenance



## MariadbSetMaintenance

> AsyncResponse MariadbSetMaintenance(ctx, clusterId).MaintenanceRequest(maintenanceRequest).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbMaintenanceApiAPI.MariadbSetMaintenance(context.Background(), clusterId).MaintenanceRequest(maintenanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbMaintenanceApiAPI.MariadbSetMaintenance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbSetMaintenance`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbMaintenanceApiAPI.MariadbSetMaintenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | DB cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSetMaintenanceRequest struct via the builder pattern


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


## MariadbUnsetMaintenance

> AsyncResponse MariadbUnsetMaintenance(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbMaintenanceApiAPI.MariadbUnsetMaintenance(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbMaintenanceApiAPI.MariadbUnsetMaintenance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbUnsetMaintenance`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbMaintenanceApiAPI.MariadbUnsetMaintenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | DB cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbUnsetMaintenanceRequest struct via the builder pattern


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

