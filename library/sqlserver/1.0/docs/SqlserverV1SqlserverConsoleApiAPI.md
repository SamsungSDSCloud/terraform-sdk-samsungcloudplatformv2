# \SqlserverV1SqlserverConsoleApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SqlserverListDatabaseUsers**](SqlserverV1SqlserverConsoleApiAPI.md#SqlserverListDatabaseUsers) | **Get** /v1/console/clusters/{cluster_id}/database-users | List DataBase Users
[**SqlserverSetDatabaseUsers**](SqlserverV1SqlserverConsoleApiAPI.md#SqlserverSetDatabaseUsers) | **Put** /v1/console/clusters/{cluster_id}/database-users | Set DataBase Users
[**SqlserverSyncDatabaseUsers**](SqlserverV1SqlserverConsoleApiAPI.md#SqlserverSyncDatabaseUsers) | **Post** /v1/console/clusters/{cluster_id}/database-users/sync | Synchronize DataBase Users



## SqlserverListDatabaseUsers

> AsyncResponse SqlserverListDatabaseUsers(ctx, clusterId).Execute()

List DataBase Users



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
	resp, r, err := apiClient.SqlserverV1SqlserverConsoleApiAPI.SqlserverListDatabaseUsers(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverConsoleApiAPI.SqlserverListDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverListDatabaseUsers`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverConsoleApiAPI.SqlserverListDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverListDatabaseUsersRequest struct via the builder pattern


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


## SqlserverSetDatabaseUsers

> AsyncResponse SqlserverSetDatabaseUsers(ctx, clusterId).DatabaseUsersSetRequest(databaseUsersSetRequest).Execute()

Set DataBase Users



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
	databaseUsersSetRequest := *openapiclient.NewDatabaseUsersSetRequest([]openapiclient.DatabaseUserRequest{*openapiclient.NewDatabaseUserRequest("Name_example")}) // DatabaseUsersSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SqlserverV1SqlserverConsoleApiAPI.SqlserverSetDatabaseUsers(context.Background(), clusterId).DatabaseUsersSetRequest(databaseUsersSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverConsoleApiAPI.SqlserverSetDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverSetDatabaseUsers`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverConsoleApiAPI.SqlserverSetDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverSetDatabaseUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **databaseUsersSetRequest** | [**DatabaseUsersSetRequest**](DatabaseUsersSetRequest.md) |  | 

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


## SqlserverSyncDatabaseUsers

> AsyncResponse SqlserverSyncDatabaseUsers(ctx, clusterId).Execute()

Synchronize DataBase Users



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
	resp, r, err := apiClient.SqlserverV1SqlserverConsoleApiAPI.SqlserverSyncDatabaseUsers(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverConsoleApiAPI.SqlserverSyncDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverSyncDatabaseUsers`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverConsoleApiAPI.SqlserverSyncDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverSyncDatabaseUsersRequest struct via the builder pattern


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

