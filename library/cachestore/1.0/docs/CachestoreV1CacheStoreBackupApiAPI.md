# \CachestoreV1CacheStoreBackupApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CachestoreListBackupHistories**](CachestoreV1CacheStoreBackupApiAPI.md#CachestoreListBackupHistories) | **Get** /v1/clusters/{cluster_id}/backup-histories | List Backup Histories
[**CachestoreRemoveBackupHistories**](CachestoreV1CacheStoreBackupApiAPI.md#CachestoreRemoveBackupHistories) | **Put** /v1/clusters/{cluster_id}/backup-histories | Remove Backup Histories
[**CachestoreSetBackup**](CachestoreV1CacheStoreBackupApiAPI.md#CachestoreSetBackup) | **Post** /v1/clusters/{cluster_id}/backups | Set Backup
[**CachestoreUnsetBackup**](CachestoreV1CacheStoreBackupApiAPI.md#CachestoreUnsetBackup) | **Delete** /v1/clusters/{cluster_id}/backups | Unset Backup



## CachestoreListBackupHistories

> BackupHistoryListApiResponse CachestoreListBackupHistories(ctx, clusterId).Limit(limit).Page(page).Execute()

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
	resp, r, err := apiClient.CachestoreV1CacheStoreBackupApiAPI.CachestoreListBackupHistories(context.Background(), clusterId).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreBackupApiAPI.CachestoreListBackupHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreListBackupHistories`: BackupHistoryListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreBackupApiAPI.CachestoreListBackupHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreListBackupHistoriesRequest struct via the builder pattern


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


## CachestoreRemoveBackupHistories

> AsyncResponse CachestoreRemoveBackupHistories(ctx, clusterId).BackupHistoryNumberRequest(backupHistoryNumberRequest).Execute()

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
	resp, r, err := apiClient.CachestoreV1CacheStoreBackupApiAPI.CachestoreRemoveBackupHistories(context.Background(), clusterId).BackupHistoryNumberRequest(backupHistoryNumberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreBackupApiAPI.CachestoreRemoveBackupHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreRemoveBackupHistories`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreBackupApiAPI.CachestoreRemoveBackupHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreRemoveBackupHistoriesRequest struct via the builder pattern


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


## CachestoreSetBackup

> AsyncResponse CachestoreSetBackup(ctx, clusterId).BackupSettingExcludingArchiveRequest(backupSettingExcludingArchiveRequest).Execute()

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
	backupSettingExcludingArchiveRequest := *openapiclient.NewBackupSettingExcludingArchiveRequest("RetentionPeriodDay_example", "StartingTimeHour_example") // BackupSettingExcludingArchiveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CachestoreV1CacheStoreBackupApiAPI.CachestoreSetBackup(context.Background(), clusterId).BackupSettingExcludingArchiveRequest(backupSettingExcludingArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreBackupApiAPI.CachestoreSetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreSetBackup`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreBackupApiAPI.CachestoreSetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreSetBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupSettingExcludingArchiveRequest** | [**BackupSettingExcludingArchiveRequest**](BackupSettingExcludingArchiveRequest.md) |  | 

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


## CachestoreUnsetBackup

> AsyncResponse CachestoreUnsetBackup(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.CachestoreV1CacheStoreBackupApiAPI.CachestoreUnsetBackup(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CachestoreV1CacheStoreBackupApiAPI.CachestoreUnsetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CachestoreUnsetBackup`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `CachestoreV1CacheStoreBackupApiAPI.CachestoreUnsetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachestoreUnsetBackupRequest struct via the builder pattern


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

