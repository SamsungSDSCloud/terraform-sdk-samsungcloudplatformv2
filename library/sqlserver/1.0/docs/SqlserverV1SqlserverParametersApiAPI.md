# \SqlserverV1SqlserverParametersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SqlserverListParameterValues**](SqlserverV1SqlserverParametersApiAPI.md#SqlserverListParameterValues) | **Get** /v1/clusters/{cluster_id}/parameters | List Parameter Values
[**SqlserverSetParameterValues**](SqlserverV1SqlserverParametersApiAPI.md#SqlserverSetParameterValues) | **Put** /v1/clusters/{cluster_id}/parameters | Set Parameter Values
[**SqlserverSyncParameterValues**](SqlserverV1SqlserverParametersApiAPI.md#SqlserverSyncParameterValues) | **Post** /v1/clusters/{cluster_id}/parameters/sync | Synchronize Parameter Values



## SqlserverListParameterValues

> ParametersResponse SqlserverListParameterValues(ctx, clusterId).Execute()

List Parameter Values



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
	resp, r, err := apiClient.SqlserverV1SqlserverParametersApiAPI.SqlserverListParameterValues(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverParametersApiAPI.SqlserverListParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverListParameterValues`: ParametersResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverParametersApiAPI.SqlserverListParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverListParameterValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParametersResponse**](ParametersResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SqlserverSetParameterValues

> AsyncResponse SqlserverSetParameterValues(ctx, clusterId).ParametersRequest(parametersRequest).Execute()

Set Parameter Values



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
	parametersRequest := *openapiclient.NewParametersRequest([]openapiclient.ParameterRequest{*openapiclient.NewParameterRequest("Id_example", "NewValue_example", "OldValue_example")}) // ParametersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SqlserverV1SqlserverParametersApiAPI.SqlserverSetParameterValues(context.Background(), clusterId).ParametersRequest(parametersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverParametersApiAPI.SqlserverSetParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverSetParameterValues`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverParametersApiAPI.SqlserverSetParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverSetParameterValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parametersRequest** | [**ParametersRequest**](ParametersRequest.md) |  | 

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


## SqlserverSyncParameterValues

> AsyncResponse SqlserverSyncParameterValues(ctx, clusterId).Execute()

Synchronize Parameter Values



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
	resp, r, err := apiClient.SqlserverV1SqlserverParametersApiAPI.SqlserverSyncParameterValues(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverParametersApiAPI.SqlserverSyncParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverSyncParameterValues`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverParametersApiAPI.SqlserverSyncParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverSyncParameterValuesRequest struct via the builder pattern


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

