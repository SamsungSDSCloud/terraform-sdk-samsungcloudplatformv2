# \MariadbV1MariadbConsoleApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariadbCheckDbConnection**](MariadbV1MariadbConsoleApiAPI.md#MariadbCheckDbConnection) | **Post** /v1/console/clusters/{cluster_id}/migration/check-db-connection | Check Db Connection
[**MariadbListAccessControl**](MariadbV1MariadbConsoleApiAPI.md#MariadbListAccessControl) | **Get** /v1/console/clusters/{cluster_id}/access-control | List Access Control
[**MariadbListDatabaseUsers**](MariadbV1MariadbConsoleApiAPI.md#MariadbListDatabaseUsers) | **Get** /v1/console/clusters/{cluster_id}/database-users | List DataBase Users
[**MariadbMigrateCluster**](MariadbV1MariadbConsoleApiAPI.md#MariadbMigrateCluster) | **Post** /v1/console/clusters/{cluster_id}/migration | Migrate Cluster
[**MariadbPromoteMigrationCluster**](MariadbV1MariadbConsoleApiAPI.md#MariadbPromoteMigrationCluster) | **Post** /v1/console/clusters/{cluster_id}/migration/promote | Promote Migration Cluster
[**MariadbSetAccessControl**](MariadbV1MariadbConsoleApiAPI.md#MariadbSetAccessControl) | **Put** /v1/console/clusters/{cluster_id}/access-control | Set Access Control
[**MariadbSetDatabaseUsers**](MariadbV1MariadbConsoleApiAPI.md#MariadbSetDatabaseUsers) | **Put** /v1/console/clusters/{cluster_id}/database-users | Set DataBase Users
[**MariadbSyncAccessControl**](MariadbV1MariadbConsoleApiAPI.md#MariadbSyncAccessControl) | **Post** /v1/console/clusters/{cluster_id}/access-control/sync | Synchronize Access Control
[**MariadbSyncDatabaseUsers**](MariadbV1MariadbConsoleApiAPI.md#MariadbSyncDatabaseUsers) | **Post** /v1/console/clusters/{cluster_id}/database-users/sync | Synchronize DataBase Users



## MariadbCheckDbConnection

> AsyncResponse MariadbCheckDbConnection(ctx, clusterId).ClusterMigrationRequest(clusterMigrationRequest).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbConsoleApiAPI.MariadbCheckDbConnection(context.Background(), clusterId).ClusterMigrationRequest(clusterMigrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbConsoleApiAPI.MariadbCheckDbConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbCheckDbConnection`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbConsoleApiAPI.MariadbCheckDbConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbCheckDbConnectionRequest struct via the builder pattern


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


## MariadbListAccessControl

> AsyncResponse MariadbListAccessControl(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbConsoleApiAPI.MariadbListAccessControl(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbConsoleApiAPI.MariadbListAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbListAccessControl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbConsoleApiAPI.MariadbListAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbListAccessControlRequest struct via the builder pattern


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


## MariadbListDatabaseUsers

> AsyncResponse MariadbListDatabaseUsers(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbConsoleApiAPI.MariadbListDatabaseUsers(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbConsoleApiAPI.MariadbListDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbListDatabaseUsers`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbConsoleApiAPI.MariadbListDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbListDatabaseUsersRequest struct via the builder pattern


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


## MariadbMigrateCluster

> AsyncResponse MariadbMigrateCluster(ctx, clusterId).ClusterMigrationRequest(clusterMigrationRequest).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbConsoleApiAPI.MariadbMigrateCluster(context.Background(), clusterId).ClusterMigrationRequest(clusterMigrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbConsoleApiAPI.MariadbMigrateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbMigrateCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbConsoleApiAPI.MariadbMigrateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbMigrateClusterRequest struct via the builder pattern


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


## MariadbPromoteMigrationCluster

> AsyncResponse MariadbPromoteMigrationCluster(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbConsoleApiAPI.MariadbPromoteMigrationCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbConsoleApiAPI.MariadbPromoteMigrationCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbPromoteMigrationCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbConsoleApiAPI.MariadbPromoteMigrationCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbPromoteMigrationClusterRequest struct via the builder pattern


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


## MariadbSetAccessControl

> AsyncResponse MariadbSetAccessControl(ctx, clusterId).AccessControlSetRequest(accessControlSetRequest).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbConsoleApiAPI.MariadbSetAccessControl(context.Background(), clusterId).AccessControlSetRequest(accessControlSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbConsoleApiAPI.MariadbSetAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbSetAccessControl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbConsoleApiAPI.MariadbSetAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSetAccessControlRequest struct via the builder pattern


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


## MariadbSetDatabaseUsers

> AsyncResponse MariadbSetDatabaseUsers(ctx, clusterId).DatabaseUsersSetRequest(databaseUsersSetRequest).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbConsoleApiAPI.MariadbSetDatabaseUsers(context.Background(), clusterId).DatabaseUsersSetRequest(databaseUsersSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbConsoleApiAPI.MariadbSetDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbSetDatabaseUsers`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbConsoleApiAPI.MariadbSetDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSetDatabaseUsersRequest struct via the builder pattern


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


## MariadbSyncAccessControl

> AsyncResponse MariadbSyncAccessControl(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbConsoleApiAPI.MariadbSyncAccessControl(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbConsoleApiAPI.MariadbSyncAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbSyncAccessControl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbConsoleApiAPI.MariadbSyncAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSyncAccessControlRequest struct via the builder pattern


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


## MariadbSyncDatabaseUsers

> AsyncResponse MariadbSyncDatabaseUsers(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbConsoleApiAPI.MariadbSyncDatabaseUsers(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbConsoleApiAPI.MariadbSyncDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbSyncDatabaseUsers`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbConsoleApiAPI.MariadbSyncDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSyncDatabaseUsersRequest struct via the builder pattern


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

