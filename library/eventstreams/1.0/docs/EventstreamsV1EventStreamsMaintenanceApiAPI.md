# \EventstreamsV1EventStreamsMaintenanceApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventstreamsSetMaintenance**](EventstreamsV1EventStreamsMaintenanceApiAPI.md#EventstreamsSetMaintenance) | **Post** /v1/clusters/{cluster_id}/maintenance | Set Maintenance
[**EventstreamsUnsetMaintenance**](EventstreamsV1EventStreamsMaintenanceApiAPI.md#EventstreamsUnsetMaintenance) | **Delete** /v1/clusters/{cluster_id}/maintenance | Unset Maintenance



## EventstreamsSetMaintenance

> AsyncResponse EventstreamsSetMaintenance(ctx, clusterId).MaintenanceRequest(maintenanceRequest).Execute()

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
	resp, r, err := apiClient.EventstreamsV1EventStreamsMaintenanceApiAPI.EventstreamsSetMaintenance(context.Background(), clusterId).MaintenanceRequest(maintenanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsMaintenanceApiAPI.EventstreamsSetMaintenance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsSetMaintenance`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsMaintenanceApiAPI.EventstreamsSetMaintenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsSetMaintenanceRequest struct via the builder pattern


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


## EventstreamsUnsetMaintenance

> AsyncResponse EventstreamsUnsetMaintenance(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EventstreamsV1EventStreamsMaintenanceApiAPI.EventstreamsUnsetMaintenance(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsMaintenanceApiAPI.EventstreamsUnsetMaintenance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsUnsetMaintenance`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsMaintenanceApiAPI.EventstreamsUnsetMaintenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsUnsetMaintenanceRequest struct via the builder pattern


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

