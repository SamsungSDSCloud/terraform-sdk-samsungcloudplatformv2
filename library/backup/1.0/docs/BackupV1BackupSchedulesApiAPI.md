# \BackupV1BackupSchedulesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListBackupHistories**](BackupV1BackupSchedulesApiAPI.md#ListBackupHistories) | **Get** /v1/backups/{backup_id}/backup-histories | List Backup Histories
[**ListBackupSchedules**](BackupV1BackupSchedulesApiAPI.md#ListBackupSchedules) | **Get** /v1/backups/{backup_id}/schedules | List Backup Schedules
[**SetBackupSchedules**](BackupV1BackupSchedulesApiAPI.md#SetBackupSchedules) | **Put** /v1/backups/{backup_id}/schedules | Set Backup Schedule



## ListBackupHistories

> BackupHistoryListResponse ListBackupHistories(ctx, backupId).BackupStartTime(backupStartTime).BackupEndTime(backupEndTime).BackupJobState(backupJobState).Page(page).Size(size).Sort(sort).Execute()

List Backup Histories



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	backupId := "00cd2538fbf94d12b36aabbdd607e974" // string | Backup ID
	backupStartTime := time.Now() // time.Time | Start time (optional)
	backupEndTime := time.Now() // time.Time | End time (optional)
	backupJobState := openapiclient.BackupJobState("SUCCESS") // BackupJobState | Backup history state (optional)
	page := int32(0) // int32 | Page (optional) (default to 0)
	size := int32(20) // int32 | Size (optional) (default to 20)
	sort := "backup_start_time:desc" // string | Sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupSchedulesApiAPI.ListBackupHistories(context.Background(), backupId).BackupStartTime(backupStartTime).BackupEndTime(backupEndTime).BackupJobState(backupJobState).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupSchedulesApiAPI.ListBackupHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupHistories`: BackupHistoryListResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupSchedulesApiAPI.ListBackupHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupHistoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupStartTime** | **time.Time** | Start time | 
 **backupEndTime** | **time.Time** | End time | 
 **backupJobState** | [**BackupJobState**](BackupJobState.md) | Backup history state | 
 **page** | **int32** | Page | [default to 0]
 **size** | **int32** | Size | [default to 20]
 **sort** | **string** | Sort | 

### Return type

[**BackupHistoryListResponse**](BackupHistoryListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupSchedules

> BackupScheduleListResponse ListBackupSchedules(ctx, backupId).Execute()

List Backup Schedules



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
	backupId := "00cd2538fbf94d12b36aabbdd607e974" // string | Backup ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupSchedulesApiAPI.ListBackupSchedules(context.Background(), backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupSchedulesApiAPI.ListBackupSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupSchedules`: BackupScheduleListResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupSchedulesApiAPI.ListBackupSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupScheduleListResponse**](BackupScheduleListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBackupSchedules

> SyncResponse SetBackupSchedules(ctx, backupId).ModifyBackupSchedulesRequest(modifyBackupSchedulesRequest).Execute()

Set Backup Schedule



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
	backupId := "00cd2538fbf94d12b36aabbdd607e974" // string | Backup ID
	modifyBackupSchedulesRequest := *openapiclient.NewModifyBackupSchedulesRequest([]openapiclient.BackupScheduleCreateRequest{*openapiclient.NewBackupScheduleCreateRequest(openapiclient.BackupScheduleFrequency("MONTHLY"), "09:00:00", openapiclient.BackupScheduleType("FULL"))}) // ModifyBackupSchedulesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupSchedulesApiAPI.SetBackupSchedules(context.Background(), backupId).ModifyBackupSchedulesRequest(modifyBackupSchedulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupSchedulesApiAPI.SetBackupSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetBackupSchedules`: SyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupSchedulesApiAPI.SetBackupSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetBackupSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyBackupSchedulesRequest** | [**ModifyBackupSchedulesRequest**](ModifyBackupSchedulesRequest.md) |  | 

### Return type

[**SyncResponse**](SyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

