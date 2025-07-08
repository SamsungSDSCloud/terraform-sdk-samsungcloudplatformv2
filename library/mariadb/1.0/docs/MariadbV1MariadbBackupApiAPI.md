# \MariadbV1MariadbBackupApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariadbListBackupHistories**](MariadbV1MariadbBackupApiAPI.md#MariadbListBackupHistories) | **Get** /v1/clusters/{cluster_id}/backup-histories | List Backup Histories
[**MariadbRemoveBackupHistories**](MariadbV1MariadbBackupApiAPI.md#MariadbRemoveBackupHistories) | **Put** /v1/clusters/{cluster_id}/backup-histories | Remove Backup Histories
[**MariadbSetBackup**](MariadbV1MariadbBackupApiAPI.md#MariadbSetBackup) | **Post** /v1/clusters/{cluster_id}/backups | Set Backup
[**MariadbUnsetBackup**](MariadbV1MariadbBackupApiAPI.md#MariadbUnsetBackup) | **Delete** /v1/clusters/{cluster_id}/backups | Unset Backup



## MariadbListBackupHistories

> BackupHistoryListApiResponse MariadbListBackupHistories(ctx, clusterId).Limit(limit).Page(page).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbBackupApiAPI.MariadbListBackupHistories(context.Background(), clusterId).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbBackupApiAPI.MariadbListBackupHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbListBackupHistories`: BackupHistoryListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbBackupApiAPI.MariadbListBackupHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbListBackupHistoriesRequest struct via the builder pattern


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


## MariadbRemoveBackupHistories

> AsyncResponse MariadbRemoveBackupHistories(ctx, clusterId).BackupHistoryNumberRequest(backupHistoryNumberRequest).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbBackupApiAPI.MariadbRemoveBackupHistories(context.Background(), clusterId).BackupHistoryNumberRequest(backupHistoryNumberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbBackupApiAPI.MariadbRemoveBackupHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbRemoveBackupHistories`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbBackupApiAPI.MariadbRemoveBackupHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbRemoveBackupHistoriesRequest struct via the builder pattern


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


## MariadbSetBackup

> AsyncResponse MariadbSetBackup(ctx, clusterId).BackupSettingRequest(backupSettingRequest).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbBackupApiAPI.MariadbSetBackup(context.Background(), clusterId).BackupSettingRequest(backupSettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbBackupApiAPI.MariadbSetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbSetBackup`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbBackupApiAPI.MariadbSetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbSetBackupRequest struct via the builder pattern


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


## MariadbUnsetBackup

> AsyncResponse MariadbUnsetBackup(ctx, clusterId).Execute()

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
	resp, r, err := apiClient.MariadbV1MariadbBackupApiAPI.MariadbUnsetBackup(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MariadbV1MariadbBackupApiAPI.MariadbUnsetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MariadbUnsetBackup`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MariadbV1MariadbBackupApiAPI.MariadbUnsetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMariadbUnsetBackupRequest struct via the builder pattern


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

