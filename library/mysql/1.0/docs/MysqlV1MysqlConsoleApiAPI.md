# \MysqlV1MysqlConsoleApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlCheckDbConnection**](MysqlV1MysqlConsoleApiAPI.md#MysqlCheckDbConnection) | **Post** /v1/console/clusters/{cluster_id}/migration/check-db-connection | Check Db Connection
[**MysqlListAccessControl**](MysqlV1MysqlConsoleApiAPI.md#MysqlListAccessControl) | **Get** /v1/console/clusters/{cluster_id}/access-control | List Access Control
[**MysqlListDatabaseUsers**](MysqlV1MysqlConsoleApiAPI.md#MysqlListDatabaseUsers) | **Get** /v1/console/clusters/{cluster_id}/database-users | List DataBase Users
[**MysqlMigrateCluster**](MysqlV1MysqlConsoleApiAPI.md#MysqlMigrateCluster) | **Post** /v1/console/clusters/{cluster_id}/migration | Migrate Cluster
[**MysqlPromoteMigrationCluster**](MysqlV1MysqlConsoleApiAPI.md#MysqlPromoteMigrationCluster) | **Post** /v1/console/clusters/{cluster_id}/migration/promote | Promote Migration Cluster
[**MysqlSetAccessControl**](MysqlV1MysqlConsoleApiAPI.md#MysqlSetAccessControl) | **Put** /v1/console/clusters/{cluster_id}/access-control | Set Access Control
[**MysqlSetDatabaseUsers**](MysqlV1MysqlConsoleApiAPI.md#MysqlSetDatabaseUsers) | **Put** /v1/console/clusters/{cluster_id}/database-users | Set DataBase Users
[**MysqlSyncAccessControl**](MysqlV1MysqlConsoleApiAPI.md#MysqlSyncAccessControl) | **Post** /v1/console/clusters/{cluster_id}/access-control/sync | Synchronize Access Control
[**MysqlSyncDatabaseUsers**](MysqlV1MysqlConsoleApiAPI.md#MysqlSyncDatabaseUsers) | **Post** /v1/console/clusters/{cluster_id}/database-users/sync | Synchronize DataBase Users



## MysqlCheckDbConnection

> AsyncResponse MysqlCheckDbConnection(ctx, clusterId).ClusterMigrationRequest(clusterMigrationRequest).Execute()

Check Db Connection



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
	clusterMigrationRequest := *openapiclient.NewClusterMigrationRequest("SourceDatabaseName_example", int32(123), "SourceDatabaseUserName_example", "SourceDatabaseUserPassword_example", "SourceIpAddress_example") // ClusterMigrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlConsoleApiAPI.MysqlCheckDbConnection(context.Background(), clusterId).ClusterMigrationRequest(clusterMigrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlConsoleApiAPI.MysqlCheckDbConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlCheckDbConnection`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlConsoleApiAPI.MysqlCheckDbConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlCheckDbConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterMigrationRequest** | [**ClusterMigrationRequest**](ClusterMigrationRequest.md) |  | 

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


## MysqlListAccessControl

> AsyncResponse MysqlListAccessControl(ctx, clusterId).Execute()

List Access Control



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
	resp, r, err := apiClient.MysqlV1MysqlConsoleApiAPI.MysqlListAccessControl(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlConsoleApiAPI.MysqlListAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlListAccessControl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlConsoleApiAPI.MysqlListAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlListAccessControlRequest struct via the builder pattern


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


## MysqlListDatabaseUsers

> AsyncResponse MysqlListDatabaseUsers(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlConsoleApiAPI.MysqlListDatabaseUsers(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlConsoleApiAPI.MysqlListDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlListDatabaseUsers`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlConsoleApiAPI.MysqlListDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlListDatabaseUsersRequest struct via the builder pattern


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


## MysqlMigrateCluster

> AsyncResponse MysqlMigrateCluster(ctx, clusterId).ClusterMigrationRequest(clusterMigrationRequest).Execute()

Migrate Cluster



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
	clusterMigrationRequest := *openapiclient.NewClusterMigrationRequest("SourceDatabaseName_example", int32(123), "SourceDatabaseUserName_example", "SourceDatabaseUserPassword_example", "SourceIpAddress_example") // ClusterMigrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlConsoleApiAPI.MysqlMigrateCluster(context.Background(), clusterId).ClusterMigrationRequest(clusterMigrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlConsoleApiAPI.MysqlMigrateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlMigrateCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlConsoleApiAPI.MysqlMigrateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlMigrateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterMigrationRequest** | [**ClusterMigrationRequest**](ClusterMigrationRequest.md) |  | 

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


## MysqlPromoteMigrationCluster

> AsyncResponse MysqlPromoteMigrationCluster(ctx, clusterId).Execute()

Promote Migration Cluster



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
	resp, r, err := apiClient.MysqlV1MysqlConsoleApiAPI.MysqlPromoteMigrationCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlConsoleApiAPI.MysqlPromoteMigrationCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlPromoteMigrationCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlConsoleApiAPI.MysqlPromoteMigrationCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlPromoteMigrationClusterRequest struct via the builder pattern


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


## MysqlSetAccessControl

> AsyncResponse MysqlSetAccessControl(ctx, clusterId).AccessControlSetRequest(accessControlSetRequest).Execute()

Set Access Control



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
	accessControlSetRequest := *openapiclient.NewAccessControlSetRequest() // AccessControlSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MysqlV1MysqlConsoleApiAPI.MysqlSetAccessControl(context.Background(), clusterId).AccessControlSetRequest(accessControlSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlConsoleApiAPI.MysqlSetAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSetAccessControl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlConsoleApiAPI.MysqlSetAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSetAccessControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessControlSetRequest** | [**AccessControlSetRequest**](AccessControlSetRequest.md) |  | 

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


## MysqlSetDatabaseUsers

> AsyncResponse MysqlSetDatabaseUsers(ctx, clusterId).DatabaseUsersSetRequest(databaseUsersSetRequest).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlConsoleApiAPI.MysqlSetDatabaseUsers(context.Background(), clusterId).DatabaseUsersSetRequest(databaseUsersSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlConsoleApiAPI.MysqlSetDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSetDatabaseUsers`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlConsoleApiAPI.MysqlSetDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSetDatabaseUsersRequest struct via the builder pattern


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


## MysqlSyncAccessControl

> AsyncResponse MysqlSyncAccessControl(ctx, clusterId).Execute()

Synchronize Access Control



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
	resp, r, err := apiClient.MysqlV1MysqlConsoleApiAPI.MysqlSyncAccessControl(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlConsoleApiAPI.MysqlSyncAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSyncAccessControl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlConsoleApiAPI.MysqlSyncAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSyncAccessControlRequest struct via the builder pattern


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


## MysqlSyncDatabaseUsers

> AsyncResponse MysqlSyncDatabaseUsers(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlConsoleApiAPI.MysqlSyncDatabaseUsers(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlConsoleApiAPI.MysqlSyncDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlSyncDatabaseUsers`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlConsoleApiAPI.MysqlSyncDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlSyncDatabaseUsersRequest struct via the builder pattern


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

