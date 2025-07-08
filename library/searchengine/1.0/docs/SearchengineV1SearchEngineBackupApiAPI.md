# \SearchengineV1SearchEngineBackupApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchengineListBackupHistories**](SearchengineV1SearchEngineBackupApiAPI.md#SearchengineListBackupHistories) | **Get** /v1/clusters/{cluster_id}/backup-histories | List Backup Histories
[**SearchengineRemoveBackupHistories**](SearchengineV1SearchEngineBackupApiAPI.md#SearchengineRemoveBackupHistories) | **Put** /v1/clusters/{cluster_id}/backup-histories | Remove Backup Histories
[**SearchengineSetBackup**](SearchengineV1SearchEngineBackupApiAPI.md#SearchengineSetBackup) | **Post** /v1/clusters/{cluster_id}/backups | Set Backup
[**SearchengineUnsetBackup**](SearchengineV1SearchEngineBackupApiAPI.md#SearchengineUnsetBackup) | **Delete** /v1/clusters/{cluster_id}/backups | Unset Backup



## SearchengineListBackupHistories

> BackupHistoryListApiResponse SearchengineListBackupHistories(ctx, clusterId).Limit(limit).Page(page).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineBackupApiAPI.SearchengineListBackupHistories(context.Background(), clusterId).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineBackupApiAPI.SearchengineListBackupHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineListBackupHistories`: BackupHistoryListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineBackupApiAPI.SearchengineListBackupHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineListBackupHistoriesRequest struct via the builder pattern


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


## SearchengineRemoveBackupHistories

> AsyncResponse SearchengineRemoveBackupHistories(ctx, clusterId).BackupHistoryNumberRequest(backupHistoryNumberRequest).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineBackupApiAPI.SearchengineRemoveBackupHistories(context.Background(), clusterId).BackupHistoryNumberRequest(backupHistoryNumberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineBackupApiAPI.SearchengineRemoveBackupHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineRemoveBackupHistories`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineBackupApiAPI.SearchengineRemoveBackupHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineRemoveBackupHistoriesRequest struct via the builder pattern


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


## SearchengineSetBackup

> AsyncResponse SearchengineSetBackup(ctx, clusterId).BackupSettingExcludingArchiveRequest(backupSettingExcludingArchiveRequest).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineBackupApiAPI.SearchengineSetBackup(context.Background(), clusterId).BackupSettingExcludingArchiveRequest(backupSettingExcludingArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineBackupApiAPI.SearchengineSetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineSetBackup`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineBackupApiAPI.SearchengineSetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineSetBackupRequest struct via the builder pattern


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


## SearchengineUnsetBackup

> AsyncResponse SearchengineUnsetBackup(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.SearchengineV1SearchEngineBackupApiAPI.SearchengineUnsetBackup(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchengineV1SearchEngineBackupApiAPI.SearchengineUnsetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchengineUnsetBackup`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchengineV1SearchEngineBackupApiAPI.SearchengineUnsetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchengineUnsetBackupRequest struct via the builder pattern


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

