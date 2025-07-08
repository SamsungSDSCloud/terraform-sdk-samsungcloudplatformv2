# \BackupV1BackupRestoreApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckBackupRestoreServerNameDuplicate**](BackupV1BackupRestoreApiAPI.md#CheckBackupRestoreServerNameDuplicate) | **Get** /v1/backups/restore/check-name-duplication | Check Backup Restore Server Name Duplicate
[**DeleteBackupRestoreTarget**](BackupV1BackupRestoreApiAPI.md#DeleteBackupRestoreTarget) | **Delete** /v1/backups/{backup_id}/restore-targets | Delete Backup Restore Target
[**ListBackupRestoreHistories**](BackupV1BackupRestoreApiAPI.md#ListBackupRestoreHistories) | **Get** /v1/backups/{backup_id}/restore-histories | List Backup Restore Histories
[**ListBackupRestoreSubnets**](BackupV1BackupRestoreApiAPI.md#ListBackupRestoreSubnets) | **Get** /v1/backups/{backup_id}/restore/restorable-subnets | List Backup Restore Subnets
[**ListBackupRestoreTarget**](BackupV1BackupRestoreApiAPI.md#ListBackupRestoreTarget) | **Get** /v1/backups/{backup_id}/restore-targets | List Backup Restore Targets
[**RestoreBackup**](BackupV1BackupRestoreApiAPI.md#RestoreBackup) | **Post** /v1/backups/{backup_id}/restore | Restore Backup



## CheckBackupRestoreServerNameDuplicate

> BackupRestoreServerNameDuplicateResponse CheckBackupRestoreServerNameDuplicate(ctx).RestoreServerName(restoreServerName).Execute()

Check Backup Restore Server Name Duplicate



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
	restoreServerName := "restoretest1" // string | Restore server name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupRestoreApiAPI.CheckBackupRestoreServerNameDuplicate(context.Background()).RestoreServerName(restoreServerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupRestoreApiAPI.CheckBackupRestoreServerNameDuplicate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckBackupRestoreServerNameDuplicate`: BackupRestoreServerNameDuplicateResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupRestoreApiAPI.CheckBackupRestoreServerNameDuplicate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckBackupRestoreServerNameDuplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restoreServerName** | **string** | Restore server name | 

### Return type

[**BackupRestoreServerNameDuplicateResponse**](BackupRestoreServerNameDuplicateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackupRestoreTarget

> SyncResponse DeleteBackupRestoreTarget(ctx, backupId).BackupRestoreTargetDeleteRequest(backupRestoreTargetDeleteRequest).Execute()

Delete Backup Restore Target



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
	backupRestoreTargetDeleteRequest := *openapiclient.NewBackupRestoreTargetDeleteRequest([]string{"RestoreTargetIds_example"}) // BackupRestoreTargetDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupRestoreApiAPI.DeleteBackupRestoreTarget(context.Background(), backupId).BackupRestoreTargetDeleteRequest(backupRestoreTargetDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupRestoreApiAPI.DeleteBackupRestoreTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBackupRestoreTarget`: SyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupRestoreApiAPI.DeleteBackupRestoreTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupRestoreTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupRestoreTargetDeleteRequest** | [**BackupRestoreTargetDeleteRequest**](BackupRestoreTargetDeleteRequest.md) |  | 

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


## ListBackupRestoreHistories

> BackupRestoreHistoryListResponse ListBackupRestoreHistories(ctx, backupId).Page(page).Size(size).Sort(sort).RestoreState(restoreState).StartTime(startTime).EndTime(endTime).Execute()

List Backup Restore Histories



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
	page := int32(0) // int32 | Page (optional)
	size := int32(20) // int32 | Size (optional)
	sort := "backup_start_time:desc" // string | Sort (optional)
	restoreState := openapiclient.BackupRestoreState("COMPLETED") // BackupRestoreState | Restore state (optional)
	startTime := time.Now() // time.Time | Start time (optional)
	endTime := time.Now() // time.Time | End time (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupRestoreApiAPI.ListBackupRestoreHistories(context.Background(), backupId).Page(page).Size(size).Sort(sort).RestoreState(restoreState).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupRestoreApiAPI.ListBackupRestoreHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupRestoreHistories`: BackupRestoreHistoryListResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupRestoreApiAPI.ListBackupRestoreHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupRestoreHistoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page | 
 **size** | **int32** | Size | 
 **sort** | **string** | Sort | 
 **restoreState** | [**BackupRestoreState**](BackupRestoreState.md) | Restore state | 
 **startTime** | **time.Time** | Start time | 
 **endTime** | **time.Time** | End time | 

### Return type

[**BackupRestoreHistoryListResponse**](BackupRestoreHistoryListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupRestoreSubnets

> BackupRestoreSubnetListResponse ListBackupRestoreSubnets(ctx, backupId).VpcId(vpcId).State(state).Region(region).Execute()

List Backup Restore Subnets



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
	vpcId := "7df8abb4912e4709b1cb237daccca7a8" // string | VPC ID (optional)
	state := "ACTIVE" // string | Subnet state (optional)
	region := "kr-west1" // string | Region to restore server (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupRestoreApiAPI.ListBackupRestoreSubnets(context.Background(), backupId).VpcId(vpcId).State(state).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupRestoreApiAPI.ListBackupRestoreSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupRestoreSubnets`: BackupRestoreSubnetListResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupRestoreApiAPI.ListBackupRestoreSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupRestoreSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpcId** | **string** | VPC ID | 
 **state** | **string** | Subnet state | 
 **region** | **string** | Region to restore server | 

### Return type

[**BackupRestoreSubnetListResponse**](BackupRestoreSubnetListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupRestoreTarget

> BackupRestoreTargetListResponse ListBackupRestoreTarget(ctx, backupId).BackupStartTime(backupStartTime).BackupEndTime(backupEndTime).Page(page).Size(size).Sort(sort).Execute()

List Backup Restore Targets



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
	page := int32(0) // int32 | Page (optional) (default to 0)
	size := int32(20) // int32 | Size (optional) (default to 20)
	sort := "backup_start_time:desc" // string | Sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupRestoreApiAPI.ListBackupRestoreTarget(context.Background(), backupId).BackupStartTime(backupStartTime).BackupEndTime(backupEndTime).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupRestoreApiAPI.ListBackupRestoreTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupRestoreTarget`: BackupRestoreTargetListResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupRestoreApiAPI.ListBackupRestoreTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupRestoreTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupStartTime** | **time.Time** | Start time | 
 **backupEndTime** | **time.Time** | End time | 
 **page** | **int32** | Page | [default to 0]
 **size** | **int32** | Size | [default to 20]
 **sort** | **string** | Sort | 

### Return type

[**BackupRestoreTargetListResponse**](BackupRestoreTargetListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreBackup

> AsyncResponse RestoreBackup(ctx, backupId).BackupRestoreRequest(backupRestoreRequest).Execute()

Restore Backup



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
	backupRestoreRequest := *openapiclient.NewBackupRestoreRequest("restoretest1", "ca1a8e76-af49-40ea-93f8-a5989ffe138b") // BackupRestoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupRestoreApiAPI.RestoreBackup(context.Background(), backupId).BackupRestoreRequest(backupRestoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupRestoreApiAPI.RestoreBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreBackup`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupRestoreApiAPI.RestoreBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupRestoreRequest** | [**BackupRestoreRequest**](BackupRestoreRequest.md) |  | 

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

