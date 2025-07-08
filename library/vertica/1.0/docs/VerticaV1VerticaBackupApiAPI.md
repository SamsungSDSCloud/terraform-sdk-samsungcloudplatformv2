# \VerticaV1VerticaBackupApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerticaListBackupHistories**](VerticaV1VerticaBackupApiAPI.md#VerticaListBackupHistories) | **Get** /v1/clusters/{cluster_id}/backup-histories | List Backup Histories
[**VerticaRemoveBackupHistories**](VerticaV1VerticaBackupApiAPI.md#VerticaRemoveBackupHistories) | **Put** /v1/clusters/{cluster_id}/backup-histories | Remove Backup Histories
[**VerticaSetBackup**](VerticaV1VerticaBackupApiAPI.md#VerticaSetBackup) | **Post** /v1/clusters/{cluster_id}/backups | Set Backup
[**VerticaUnsetBackup**](VerticaV1VerticaBackupApiAPI.md#VerticaUnsetBackup) | **Delete** /v1/clusters/{cluster_id}/backups | Unset Backup



## VerticaListBackupHistories

> BackupHistoryListApiResponse VerticaListBackupHistories(ctx, clusterId).Limit(limit).Page(page).Execute()

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
	resp, r, err := apiClient.VerticaV1VerticaBackupApiAPI.VerticaListBackupHistories(context.Background(), clusterId).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaBackupApiAPI.VerticaListBackupHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaListBackupHistories`: BackupHistoryListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaBackupApiAPI.VerticaListBackupHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerticaListBackupHistoriesRequest struct via the builder pattern


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


## VerticaRemoveBackupHistories

> AsyncResponse VerticaRemoveBackupHistories(ctx, clusterId).BackupHistoryNumberRequest(backupHistoryNumberRequest).Execute()

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
	resp, r, err := apiClient.VerticaV1VerticaBackupApiAPI.VerticaRemoveBackupHistories(context.Background(), clusterId).BackupHistoryNumberRequest(backupHistoryNumberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaBackupApiAPI.VerticaRemoveBackupHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaRemoveBackupHistories`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaBackupApiAPI.VerticaRemoveBackupHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerticaRemoveBackupHistoriesRequest struct via the builder pattern


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


## VerticaSetBackup

> AsyncResponse VerticaSetBackup(ctx, clusterId).BackupSettingExcludingArchiveRequest(backupSettingExcludingArchiveRequest).Execute()

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
	resp, r, err := apiClient.VerticaV1VerticaBackupApiAPI.VerticaSetBackup(context.Background(), clusterId).BackupSettingExcludingArchiveRequest(backupSettingExcludingArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaBackupApiAPI.VerticaSetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaSetBackup`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaBackupApiAPI.VerticaSetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerticaSetBackupRequest struct via the builder pattern


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


## VerticaUnsetBackup

> AsyncResponse VerticaUnsetBackup(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.VerticaV1VerticaBackupApiAPI.VerticaUnsetBackup(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerticaV1VerticaBackupApiAPI.VerticaUnsetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerticaUnsetBackup`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `VerticaV1VerticaBackupApiAPI.VerticaUnsetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerticaUnsetBackupRequest struct via the builder pattern


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

