# \EpasV1EpasParametersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EpasListParameterValues**](EpasV1EpasParametersApiAPI.md#EpasListParameterValues) | **Get** /v1/clusters/{cluster_id}/parameters | List Parameter Values
[**EpasSetParameterValues**](EpasV1EpasParametersApiAPI.md#EpasSetParameterValues) | **Put** /v1/clusters/{cluster_id}/parameters | Set Parameter Values
[**EpasSyncParameterValues**](EpasV1EpasParametersApiAPI.md#EpasSyncParameterValues) | **Post** /v1/clusters/{cluster_id}/parameters/sync | Synchronize Parameter Values



## EpasListParameterValues

> ParametersResponse EpasListParameterValues(ctx, clusterId).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasParametersApiAPI.EpasListParameterValues(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasParametersApiAPI.EpasListParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasListParameterValues`: ParametersResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasParametersApiAPI.EpasListParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasListParameterValuesRequest struct via the builder pattern


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


## EpasSetParameterValues

> AsyncResponse EpasSetParameterValues(ctx, clusterId).ParametersRequest(parametersRequest).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID
	parametersRequest := *openapiclient.NewParametersRequest([]openapiclient.ParameterRequest{*openapiclient.NewParameterRequest("Id_example", "NewValue_example", "OldValue_example")}) // ParametersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasParametersApiAPI.EpasSetParameterValues(context.Background(), clusterId).ParametersRequest(parametersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasParametersApiAPI.EpasSetParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSetParameterValues`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasParametersApiAPI.EpasSetParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSetParameterValuesRequest struct via the builder pattern


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


## EpasSyncParameterValues

> AsyncResponse EpasSyncParameterValues(ctx, clusterId).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasParametersApiAPI.EpasSyncParameterValues(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasParametersApiAPI.EpasSyncParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSyncParameterValues`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasParametersApiAPI.EpasSyncParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSyncParameterValuesRequest struct via the builder pattern


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

