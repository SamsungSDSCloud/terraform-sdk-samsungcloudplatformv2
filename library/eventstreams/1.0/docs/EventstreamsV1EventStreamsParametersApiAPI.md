# \EventstreamsV1EventStreamsParametersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventstreamsListParameterValues**](EventstreamsV1EventStreamsParametersApiAPI.md#EventstreamsListParameterValues) | **Get** /v1/clusters/{cluster_id}/parameters | List Parameter Values
[**EventstreamsSetParameterValues**](EventstreamsV1EventStreamsParametersApiAPI.md#EventstreamsSetParameterValues) | **Put** /v1/clusters/{cluster_id}/parameters | Set Parameter Values
[**EventstreamsSyncParameterValues**](EventstreamsV1EventStreamsParametersApiAPI.md#EventstreamsSyncParameterValues) | **Post** /v1/clusters/{cluster_id}/parameters/sync | Synchronize Parameter Values



## EventstreamsListParameterValues

> ParametersResponse EventstreamsListParameterValues(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EventstreamsV1EventStreamsParametersApiAPI.EventstreamsListParameterValues(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsParametersApiAPI.EventstreamsListParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsListParameterValues`: ParametersResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsParametersApiAPI.EventstreamsListParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsListParameterValuesRequest struct via the builder pattern


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


## EventstreamsSetParameterValues

> AsyncResponse EventstreamsSetParameterValues(ctx, clusterId).ParametersRequest(parametersRequest).Execute()

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
	resp, r, err := apiClient.EventstreamsV1EventStreamsParametersApiAPI.EventstreamsSetParameterValues(context.Background(), clusterId).ParametersRequest(parametersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsParametersApiAPI.EventstreamsSetParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsSetParameterValues`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsParametersApiAPI.EventstreamsSetParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsSetParameterValuesRequest struct via the builder pattern


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


## EventstreamsSyncParameterValues

> AsyncResponse EventstreamsSyncParameterValues(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.EventstreamsV1EventStreamsParametersApiAPI.EventstreamsSyncParameterValues(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventstreamsV1EventStreamsParametersApiAPI.EventstreamsSyncParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventstreamsSyncParameterValues`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EventstreamsV1EventStreamsParametersApiAPI.EventstreamsSyncParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventstreamsSyncParameterValuesRequest struct via the builder pattern


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

