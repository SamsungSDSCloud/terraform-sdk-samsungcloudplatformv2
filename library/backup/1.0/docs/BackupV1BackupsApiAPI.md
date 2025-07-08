# \BackupV1BackupsApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckBackupNameDuplicate**](BackupV1BackupsApiAPI.md#CheckBackupNameDuplicate) | **Get** /v1/backups/check-name-duplication | Check Backup Name Duplicate
[**CreateBackup**](BackupV1BackupsApiAPI.md#CreateBackup) | **Post** /v1/backups | Create Backup
[**DeleteBackup**](BackupV1BackupsApiAPI.md#DeleteBackup) | **Delete** /v1/backups/{backup_id} | Delete Backup
[**GetBackupTargetList**](BackupV1BackupsApiAPI.md#GetBackupTargetList) | **Get** /v1/backups/backup-targets | List Backup Targets
[**ListBackupRegionRelationship**](BackupV1BackupsApiAPI.md#ListBackupRegionRelationship) | **Get** /v1/backups/region-relationship | List Backup Region Relationship
[**ListBackups**](BackupV1BackupsApiAPI.md#ListBackups) | **Get** /v1/backups | List Backups
[**ManualBackup**](BackupV1BackupsApiAPI.md#ManualBackup) | **Post** /v1/backups/{backup_id}/manual-backup | Manual Backup
[**ShowBackup**](BackupV1BackupsApiAPI.md#ShowBackup) | **Get** /v1/backups/{backup_id} | Show Backup
[**UpdateRetentionPeriod**](BackupV1BackupsApiAPI.md#UpdateRetentionPeriod) | **Put** /v1/backups/{backup_id}/retention-period | Update Retention Period



## CheckBackupNameDuplicate

> BackupNameDuplicateResponse CheckBackupNameDuplicate(ctx).BackupName(backupName).Execute()

Check Backup Name Duplicate



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
	backupName := "backup-001" // string | Backup name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupsApiAPI.CheckBackupNameDuplicate(context.Background()).BackupName(backupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupsApiAPI.CheckBackupNameDuplicate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckBackupNameDuplicate`: BackupNameDuplicateResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupsApiAPI.CheckBackupNameDuplicate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckBackupNameDuplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupName** | **string** | Backup name | 

### Return type

[**BackupNameDuplicateResponse**](BackupNameDuplicateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBackup

> SyncResponse CreateBackup(ctx).BackupCreateRequest(backupCreateRequest).Execute()

Create Backup



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
	backupCreateRequest := *openapiclient.NewBackupCreateRequest("backup-001", openapiclient.BackupPolicyCategory("AGENTLESS"), openapiclient.BackupPolicyType("VM_IMAGE"), []openapiclient.BackupScheduleCreateRequest{*openapiclient.NewBackupScheduleCreateRequest(openapiclient.BackupScheduleFrequency("MONTHLY"), "09:00:00", openapiclient.BackupScheduleType("FULL"))}, openapiclient.ServerCategory("VIRTUAL_SERVER"), "89f5ef44-1021-4a5c-8e06-fbb289eac366") // BackupCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupsApiAPI.CreateBackup(context.Background()).BackupCreateRequest(backupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupsApiAPI.CreateBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBackup`: SyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupsApiAPI.CreateBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupCreateRequest** | [**BackupCreateRequest**](BackupCreateRequest.md) |  | 

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


## DeleteBackup

> SyncResponse DeleteBackup(ctx, backupId).Execute()

Delete Backup



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
	resp, r, err := apiClient.BackupV1BackupsApiAPI.DeleteBackup(context.Background(), backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupsApiAPI.DeleteBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBackup`: SyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupsApiAPI.DeleteBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncResponse**](SyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupTargetList

> BackupTargetListResponse GetBackupTargetList(ctx).ServerCategory(serverCategory).ServerName(serverName).Region(region).Page(page).Size(size).Execute()

List Backup Targets



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
	serverCategory := openapiclient.ServerCategory("VIRTUAL_SERVER") // ServerCategory | Backup server category
	serverName := "server-001" // string | Backup server name (optional)
	region := "kr-west1" // string | Region (optional)
	page := int32(0) // int32 | Page (optional)
	size := int32(20) // int32 | Size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupsApiAPI.GetBackupTargetList(context.Background()).ServerCategory(serverCategory).ServerName(serverName).Region(region).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupsApiAPI.GetBackupTargetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupTargetList`: BackupTargetListResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupsApiAPI.GetBackupTargetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupTargetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverCategory** | [**ServerCategory**](ServerCategory.md) | Backup server category | 
 **serverName** | **string** | Backup server name | 
 **region** | **string** | Region | 
 **page** | **int32** | Page | 
 **size** | **int32** | Size | 

### Return type

[**BackupTargetListResponse**](BackupTargetListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupRegionRelationship

> BackupListRegionRelationshipResponse ListBackupRegionRelationship(ctx).FromRegion(fromRegion).FromAvailabilityZone(fromAvailabilityZone).ToRegion(toRegion).ToAvailabilityZone(toAvailabilityZone).Execute()

List Backup Region Relationship



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
	fromRegion := "kr-west1" // string | Region (optional)
	fromAvailabilityZone := "kr-west1" // string | Availability zone name (optional)
	toRegion := "kr-west1" // string | Region (optional)
	toAvailabilityZone := "kr-west1" // string | Availability zone name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupsApiAPI.ListBackupRegionRelationship(context.Background()).FromRegion(fromRegion).FromAvailabilityZone(fromAvailabilityZone).ToRegion(toRegion).ToAvailabilityZone(toAvailabilityZone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupsApiAPI.ListBackupRegionRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupRegionRelationship`: BackupListRegionRelationshipResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupsApiAPI.ListBackupRegionRelationship`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBackupRegionRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromRegion** | **string** | Region | 
 **fromAvailabilityZone** | **string** | Availability zone name | 
 **toRegion** | **string** | Region | 
 **toAvailabilityZone** | **string** | Availability zone name | 

### Return type

[**BackupListRegionRelationshipResponse**](BackupListRegionRelationshipResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackups

> BackupListResponse ListBackups(ctx).Size(size).Page(page).Sort(sort).Name(name).ServerName(serverName).Execute()

List Backups



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
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	name := "backup-001" // string | Backup name (optional)
	serverName := "server-001" // string | Backup server name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupsApiAPI.ListBackups(context.Background()).Size(size).Page(page).Sort(sort).Name(name).ServerName(serverName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupsApiAPI.ListBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackups`: BackupListResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupsApiAPI.ListBackups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Backup name | 
 **serverName** | **string** | Backup server name | 

### Return type

[**BackupListResponse**](BackupListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManualBackup

> SyncResponse ManualBackup(ctx, backupId).Execute()

Manual Backup



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
	resp, r, err := apiClient.BackupV1BackupsApiAPI.ManualBackup(context.Background(), backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupsApiAPI.ManualBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManualBackup`: SyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupsApiAPI.ManualBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiManualBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyncResponse**](SyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowBackup

> BackupDetailResponse ShowBackup(ctx, backupId).Execute()

Show Backup



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
	resp, r, err := apiClient.BackupV1BackupsApiAPI.ShowBackup(context.Background(), backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupsApiAPI.ShowBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowBackup`: BackupDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupsApiAPI.ShowBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupDetailResponse**](BackupDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRetentionPeriod

> SyncResponse UpdateRetentionPeriod(ctx, backupId).RetentionPeriodUpdateBody(retentionPeriodUpdateBody).Execute()

Update Retention Period



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
	retentionPeriodUpdateBody := *openapiclient.NewRetentionPeriodUpdateBody(openapiclient.BackupRetentionPeriod("WEEK_2")) // RetentionPeriodUpdateBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupsApiAPI.UpdateRetentionPeriod(context.Background(), backupId).RetentionPeriodUpdateBody(retentionPeriodUpdateBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupsApiAPI.UpdateRetentionPeriod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRetentionPeriod`: SyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupsApiAPI.UpdateRetentionPeriod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRetentionPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **retentionPeriodUpdateBody** | [**RetentionPeriodUpdateBody**](RetentionPeriodUpdateBody.md) |  | 

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

