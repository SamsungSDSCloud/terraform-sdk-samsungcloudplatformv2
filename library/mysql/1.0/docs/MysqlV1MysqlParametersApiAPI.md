# \MysqlV1MysqlParametersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlListParameterValues**](MysqlV1MysqlParametersApiAPI.md#MysqlListParameterValues) | **Get** /v1/clusters/{cluster_id}/parameters | List Parameter Values
[**MysqlSetParameterValues**](MysqlV1MysqlParametersApiAPI.md#MysqlSetParameterValues) | **Put** /v1/clusters/{cluster_id}/parameters | Set Parameter Values
[**MysqlSyncParameterValues**](MysqlV1MysqlParametersApiAPI.md#MysqlSyncParameterValues) | **Post** /v1/clusters/{cluster_id}/parameters/sync | Synchronize Parameter Values



## MysqlListParameterValues

> ParametersResponse MysqlListParameterValues(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlParametersApiAPI.MysqlListParameterValues(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlParametersApiAPI.MysqlListParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlListParameterValues`: ParametersResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlParametersApiAPI.MysqlListParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlListParameterValuesRequest struct via the builder pattern


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


## MysqlSetParameterValues

> AsyncResponse MysqlSetParameterValues(ctx, clusterId).ParametersRequest(parametersRequest).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlParametersApiAPI.MysqlSetParameterValues(context.Background(), clusterId).ParametersRequest(parametersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlParametersApiAPI.MysqlSetParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSetParameterValues`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlParametersApiAPI.MysqlSetParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSetParameterValuesRequest struct via the builder pattern


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


## MysqlSyncParameterValues

> AsyncResponse MysqlSyncParameterValues(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlParametersApiAPI.MysqlSyncParameterValues(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlParametersApiAPI.MysqlSyncParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSyncParameterValues`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlParametersApiAPI.MysqlSyncParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSyncParameterValuesRequest struct via the builder pattern


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

