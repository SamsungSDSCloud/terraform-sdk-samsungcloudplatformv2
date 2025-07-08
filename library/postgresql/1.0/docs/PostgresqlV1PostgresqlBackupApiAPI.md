# \PostgresqlV1PostgresqlBackupApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostgresqlListBackupHistories**](PostgresqlV1PostgresqlBackupApiAPI.md#PostgresqlListBackupHistories) | **Get** /v1/clusters/{cluster_id}/backup-histories | List Backup Histories
[**PostgresqlRemoveBackupHistories**](PostgresqlV1PostgresqlBackupApiAPI.md#PostgresqlRemoveBackupHistories) | **Put** /v1/clusters/{cluster_id}/backup-histories | Remove Backup Histories
[**PostgresqlSetBackup**](PostgresqlV1PostgresqlBackupApiAPI.md#PostgresqlSetBackup) | **Post** /v1/clusters/{cluster_id}/backups | Set Backup
[**PostgresqlUnsetBackup**](PostgresqlV1PostgresqlBackupApiAPI.md#PostgresqlUnsetBackup) | **Delete** /v1/clusters/{cluster_id}/backups | Unset Backup



## PostgresqlListBackupHistories

> BackupHistoryListApiResponse PostgresqlListBackupHistories(ctx, clusterId).Limit(limit).Page(page).Execute()

List Backup Histories



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
	limit := int32(56) // int32 | Number of backup lists (optional)
	page := int32(56) // int32 | Backup list page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostgresqlV1PostgresqlBackupApiAPI.PostgresqlListBackupHistories(context.Background(), clusterId).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlBackupApiAPI.PostgresqlListBackupHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlListBackupHistories`: BackupHistoryListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlBackupApiAPI.PostgresqlListBackupHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlListBackupHistoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of backup lists | 
 **page** | **int32** | Backup list page | 

### Return type

[**BackupHistoryListApiResponse**](BackupHistoryListApiResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlRemoveBackupHistories

> AsyncResponse PostgresqlRemoveBackupHistories(ctx, clusterId).BackupHistoryNumberRequest(backupHistoryNumberRequest).Execute()

Remove Backup Histories



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
	backupHistoryNumberRequest := *openapiclient.NewBackupHistoryNumberRequest([]string{"BackupHistoryNumber_example"}) // BackupHistoryNumberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostgresqlV1PostgresqlBackupApiAPI.PostgresqlRemoveBackupHistories(context.Background(), clusterId).BackupHistoryNumberRequest(backupHistoryNumberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlBackupApiAPI.PostgresqlRemoveBackupHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlRemoveBackupHistories`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlBackupApiAPI.PostgresqlRemoveBackupHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlRemoveBackupHistoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupHistoryNumberRequest** | [**BackupHistoryNumberRequest**](BackupHistoryNumberRequest.md) |  | 

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


## PostgresqlSetBackup

> AsyncResponse PostgresqlSetBackup(ctx, clusterId).BackupSettingRequest(backupSettingRequest).Execute()

Set Backup



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
	backupSettingRequest := *openapiclient.NewBackupSettingRequest("ArchiveFrequencyMinute_example", "RetentionPeriodDay_example", "StartingTimeHour_example") // BackupSettingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostgresqlV1PostgresqlBackupApiAPI.PostgresqlSetBackup(context.Background(), clusterId).BackupSettingRequest(backupSettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlBackupApiAPI.PostgresqlSetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlSetBackup`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlBackupApiAPI.PostgresqlSetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlSetBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupSettingRequest** | [**BackupSettingRequest**](BackupSettingRequest.md) |  | 

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


## PostgresqlUnsetBackup

> AsyncResponse PostgresqlUnsetBackup(ctx, clusterId).Execute()

Unset Backup



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
	resp, r, err := apiClient.PostgresqlV1PostgresqlBackupApiAPI.PostgresqlUnsetBackup(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlV1PostgresqlBackupApiAPI.PostgresqlUnsetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostgresqlUnsetBackup`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `PostgresqlV1PostgresqlBackupApiAPI.PostgresqlUnsetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlUnsetBackupRequest struct via the builder pattern


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

