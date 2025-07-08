# \EpasV1EpasConsoleApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EpasCheckDbConnection**](EpasV1EpasConsoleApiAPI.md#EpasCheckDbConnection) | **Post** /v1/console/clusters/{cluster_id}/migration/check-db-connection | Check Db Connection
[**EpasListAccessControl**](EpasV1EpasConsoleApiAPI.md#EpasListAccessControl) | **Get** /v1/console/clusters/{cluster_id}/access-control | List Access Control
[**EpasListDatabaseUsers**](EpasV1EpasConsoleApiAPI.md#EpasListDatabaseUsers) | **Get** /v1/console/clusters/{cluster_id}/database-users | List DataBase Users
[**EpasMigrateCluster**](EpasV1EpasConsoleApiAPI.md#EpasMigrateCluster) | **Post** /v1/console/clusters/{cluster_id}/migration | Migrate Cluster
[**EpasPromoteMigrationCluster**](EpasV1EpasConsoleApiAPI.md#EpasPromoteMigrationCluster) | **Post** /v1/console/clusters/{cluster_id}/migration/promote | Promote Migration Cluster
[**EpasSetAccessControl**](EpasV1EpasConsoleApiAPI.md#EpasSetAccessControl) | **Put** /v1/console/clusters/{cluster_id}/access-control | Set Access Control
[**EpasSetDatabaseUsers**](EpasV1EpasConsoleApiAPI.md#EpasSetDatabaseUsers) | **Put** /v1/console/clusters/{cluster_id}/database-users | Set DataBase Users
[**EpasSyncAccessControl**](EpasV1EpasConsoleApiAPI.md#EpasSyncAccessControl) | **Post** /v1/console/clusters/{cluster_id}/access-control/sync | Synchronize Access Control
[**EpasSyncDatabaseUsers**](EpasV1EpasConsoleApiAPI.md#EpasSyncDatabaseUsers) | **Post** /v1/console/clusters/{cluster_id}/database-users/sync | Synchronize DataBase Users



## EpasCheckDbConnection

> AsyncResponse EpasCheckDbConnection(ctx, clusterId).ClusterMigrationRequest(clusterMigrationRequest).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID
	clusterMigrationRequest := *openapiclient.NewClusterMigrationRequest("SourceDatabaseName_example", int32(123), "SourceDatabaseUserName_example", "SourceDatabaseUserPassword_example", "SourceIpAddress_example") // ClusterMigrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasConsoleApiAPI.EpasCheckDbConnection(context.Background(), clusterId).ClusterMigrationRequest(clusterMigrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasConsoleApiAPI.EpasCheckDbConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasCheckDbConnection`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasConsoleApiAPI.EpasCheckDbConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasCheckDbConnectionRequest struct via the builder pattern


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


## EpasListAccessControl

> AsyncResponse EpasListAccessControl(ctx, clusterId).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasConsoleApiAPI.EpasListAccessControl(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasConsoleApiAPI.EpasListAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasListAccessControl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasConsoleApiAPI.EpasListAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasListAccessControlRequest struct via the builder pattern


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


## EpasListDatabaseUsers

> AsyncResponse EpasListDatabaseUsers(ctx, clusterId).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasConsoleApiAPI.EpasListDatabaseUsers(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasConsoleApiAPI.EpasListDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasListDatabaseUsers`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasConsoleApiAPI.EpasListDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasListDatabaseUsersRequest struct via the builder pattern


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


## EpasMigrateCluster

> AsyncResponse EpasMigrateCluster(ctx, clusterId).ClusterMigrationRequest(clusterMigrationRequest).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID
	clusterMigrationRequest := *openapiclient.NewClusterMigrationRequest("SourceDatabaseName_example", int32(123), "SourceDatabaseUserName_example", "SourceDatabaseUserPassword_example", "SourceIpAddress_example") // ClusterMigrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasConsoleApiAPI.EpasMigrateCluster(context.Background(), clusterId).ClusterMigrationRequest(clusterMigrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasConsoleApiAPI.EpasMigrateCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasMigrateCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasConsoleApiAPI.EpasMigrateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasMigrateClusterRequest struct via the builder pattern


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


## EpasPromoteMigrationCluster

> AsyncResponse EpasPromoteMigrationCluster(ctx, clusterId).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasConsoleApiAPI.EpasPromoteMigrationCluster(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasConsoleApiAPI.EpasPromoteMigrationCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasPromoteMigrationCluster`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasConsoleApiAPI.EpasPromoteMigrationCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasPromoteMigrationClusterRequest struct via the builder pattern


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


## EpasSetAccessControl

> AsyncResponse EpasSetAccessControl(ctx, clusterId).AccessControlSetRequest(accessControlSetRequest).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID
	accessControlSetRequest := *openapiclient.NewAccessControlSetRequest() // AccessControlSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasConsoleApiAPI.EpasSetAccessControl(context.Background(), clusterId).AccessControlSetRequest(accessControlSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasConsoleApiAPI.EpasSetAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSetAccessControl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasConsoleApiAPI.EpasSetAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSetAccessControlRequest struct via the builder pattern


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


## EpasSetDatabaseUsers

> AsyncResponse EpasSetDatabaseUsers(ctx, clusterId).DatabaseUsersSetRequest(databaseUsersSetRequest).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID
	databaseUsersSetRequest := *openapiclient.NewDatabaseUsersSetRequest([]openapiclient.DatabaseUserRequest{*openapiclient.NewDatabaseUserRequest("Name_example")}) // DatabaseUsersSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasConsoleApiAPI.EpasSetDatabaseUsers(context.Background(), clusterId).DatabaseUsersSetRequest(databaseUsersSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasConsoleApiAPI.EpasSetDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSetDatabaseUsers`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasConsoleApiAPI.EpasSetDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSetDatabaseUsersRequest struct via the builder pattern


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


## EpasSyncAccessControl

> AsyncResponse EpasSyncAccessControl(ctx, clusterId).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasConsoleApiAPI.EpasSyncAccessControl(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasConsoleApiAPI.EpasSyncAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSyncAccessControl`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasConsoleApiAPI.EpasSyncAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSyncAccessControlRequest struct via the builder pattern


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


## EpasSyncDatabaseUsers

> AsyncResponse EpasSyncDatabaseUsers(ctx, clusterId).Execute()

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
	clusterId := "clusterId_example" // string | Cluster ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasConsoleApiAPI.EpasSyncDatabaseUsers(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasConsoleApiAPI.EpasSyncDatabaseUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasSyncDatabaseUsers`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasConsoleApiAPI.EpasSyncDatabaseUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasSyncDatabaseUsersRequest struct via the builder pattern


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

