# \SqlserverV1SqlserverDatabasesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SqlserverSetDatabases**](SqlserverV1SqlserverDatabasesApiAPI.md#SqlserverSetDatabases) | **Post** /v1/clusters/{cluster_id}/databases | Add/Remove Databases



## SqlserverSetDatabases

> AsyncResponse SqlserverSetDatabases(ctx, clusterId).SqlserverUpdateDatabasesRequest(sqlserverUpdateDatabasesRequest).Execute()

Add/Remove Databases



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
	sqlserverUpdateDatabasesRequest := *openapiclient.NewSqlserverUpdateDatabasesRequest() // SqlserverUpdateDatabasesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SqlserverV1SqlserverDatabasesApiAPI.SqlserverSetDatabases(context.Background(), clusterId).SqlserverUpdateDatabasesRequest(sqlserverUpdateDatabasesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverDatabasesApiAPI.SqlserverSetDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverSetDatabases`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverDatabasesApiAPI.SqlserverSetDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverSetDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sqlserverUpdateDatabasesRequest** | [**SqlserverUpdateDatabasesRequest**](SqlserverUpdateDatabasesRequest.md) |  | 

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

