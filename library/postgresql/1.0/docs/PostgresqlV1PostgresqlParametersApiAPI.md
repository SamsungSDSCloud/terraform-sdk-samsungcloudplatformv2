# \PostgresqlV1PostgresqlParametersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostgresqlListParameterValues**](PostgresqlV1PostgresqlParametersApiAPI.md#PostgresqlListParameterValues) | **Get** /v1/clusters/{cluster_id}/parameters | List Parameter Values
[**PostgresqlSetParameterValues**](PostgresqlV1PostgresqlParametersApiAPI.md#PostgresqlSetParameterValues) | **Put** /v1/clusters/{cluster_id}/parameters | Set Parameter Values
[**PostgresqlSyncParameterValues**](PostgresqlV1PostgresqlParametersApiAPI.md#PostgresqlSyncParameterValues) | **Post** /v1/clusters/{cluster_id}/parameters/sync | Synchronize Parameter Values



## PostgresqlListParameterValues

> ParametersResponse PostgresqlListParameterValues(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlParametersApiAPI.PostgresqlListParameterValues(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlParametersApiAPI.PostgresqlListParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlListParameterValues`: ParametersResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlParametersApiAPI.PostgresqlListParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlListParameterValuesRequest struct via the builder pattern


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


## PostgresqlSetParameterValues

> AsyncResponse PostgresqlSetParameterValues(ctx, clusterId).ParametersRequest(parametersRequest).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlParametersApiAPI.PostgresqlSetParameterValues(context.Background(), clusterId).ParametersRequest(parametersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlParametersApiAPI.PostgresqlSetParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlSetParameterValues`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlParametersApiAPI.PostgresqlSetParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlSetParameterValuesRequest struct via the builder pattern


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


## PostgresqlSyncParameterValues

> AsyncResponse PostgresqlSyncParameterValues(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.PostgresqlV1PostgresqlParametersApiAPI.PostgresqlSyncParameterValues(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlParametersApiAPI.PostgresqlSyncParameterValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlSyncParameterValues`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlParametersApiAPI.PostgresqlSyncParameterValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlSyncParameterValuesRequest struct via the builder pattern


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

