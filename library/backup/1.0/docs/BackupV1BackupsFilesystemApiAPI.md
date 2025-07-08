# \BackupV1BackupsFilesystemApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckFilesystemDuplication**](BackupV1BackupsFilesystemApiAPI.md#CheckFilesystemDuplication) | **Get** /v1/backups/check-filesystem-duplication | Check Filesystem Path Duplication
[**SetFilesystemPath**](BackupV1BackupsFilesystemApiAPI.md#SetFilesystemPath) | **Put** /v1/backups/{backup_id}/filesystem-path | Set Filesystem Path



## CheckFilesystemDuplication

> CheckFileSystemDuplicationResponse CheckFilesystemDuplication(ctx).FilesystemPath(filesystemPath).ServerUuid(serverUuid).Execute()

Check Filesystem Path Duplication



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
	filesystemPath := "/aaa" // string | Filesystem backup path
	serverUuid := "89f5ef44-1021-4a5c-8e06-fbb289eac366" // string | Backup server UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupsFilesystemApiAPI.CheckFilesystemDuplication(context.Background()).FilesystemPath(filesystemPath).ServerUuid(serverUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupsFilesystemApiAPI.CheckFilesystemDuplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckFilesystemDuplication`: CheckFileSystemDuplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupsFilesystemApiAPI.CheckFilesystemDuplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckFilesystemDuplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filesystemPath** | **string** | Filesystem backup path | 
 **serverUuid** | **string** | Backup server UUID | 

### Return type

[**CheckFileSystemDuplicationResponse**](CheckFileSystemDuplicationResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFilesystemPath

> SyncResponse SetFilesystemPath(ctx, backupId).FileSystemPathUpdateRequest(fileSystemPathUpdateRequest).Execute()

Set Filesystem Path



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
	fileSystemPathUpdateRequest := *openapiclient.NewFileSystemPathUpdateRequest(*openapiclient.NewBackupFilesystemPaths()) // FileSystemPathUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupsFilesystemApiAPI.SetFilesystemPath(context.Background(), backupId).FileSystemPathUpdateRequest(fileSystemPathUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupsFilesystemApiAPI.SetFilesystemPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFilesystemPath`: SyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupsFilesystemApiAPI.SetFilesystemPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFilesystemPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileSystemPathUpdateRequest** | [**FileSystemPathUpdateRequest**](FileSystemPathUpdateRequest.md) |  | 

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

