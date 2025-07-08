# \BackupV1AgentBackupsRestoreAPIAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAgentBackupRestoreTargetServers**](BackupV1AgentBackupsRestoreAPIAPI.md#ListAgentBackupRestoreTargetServers) | **Get** /v1/backups/{backup_id}/agent-backup-restore-targets | List Agent Backup Restore Target Servers
[**RestoreAgentBackup**](BackupV1AgentBackupsRestoreAPIAPI.md#RestoreAgentBackup) | **Post** /v1/backups/{backup_id}/restore-agent-backup | Restore Agent Backup



## ListAgentBackupRestoreTargetServers

> AgentBackupRestoreTargetServerListResponse ListAgentBackupRestoreTargetServers(ctx, backupId).ServerName(serverName).Page(page).Size(size).Execute()

List Agent Backup Restore Target Servers



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
	serverName := "server-001" // string | Backup server name (optional)
	page := int32(0) // int32 | Page (optional)
	size := int32(20) // int32 | Size (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1AgentBackupsRestoreAPIAPI.ListAgentBackupRestoreTargetServers(context.Background(), backupId).ServerName(serverName).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1AgentBackupsRestoreAPIAPI.ListAgentBackupRestoreTargetServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentBackupRestoreTargetServers`: AgentBackupRestoreTargetServerListResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1AgentBackupsRestoreAPIAPI.ListAgentBackupRestoreTargetServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentBackupRestoreTargetServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serverName** | **string** | Backup server name | 
 **page** | **int32** | Page | 
 **size** | **int32** | Size | 

### Return type

[**AgentBackupRestoreTargetServerListResponse**](AgentBackupRestoreTargetServerListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreAgentBackup

> SyncResponse RestoreAgentBackup(ctx, backupId).AgentBackupRestoreRequest(agentBackupRestoreRequest).Execute()

Restore Agent Backup



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
	agentBackupRestoreRequest := *openapiclient.NewAgentBackupRestoreRequest([]openapiclient.RestoreFilesystemPath{*openapiclient.NewRestoreFilesystemPath("/data", "/data")}, "89f5ef44-1021-4a5c-8e06-fbb289eac366", "ca1a8e76-af49-40ea-93f8-a5989ffe138b") // AgentBackupRestoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1AgentBackupsRestoreAPIAPI.RestoreAgentBackup(context.Background(), backupId).AgentBackupRestoreRequest(agentBackupRestoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1AgentBackupsRestoreAPIAPI.RestoreAgentBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreAgentBackup`: SyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1AgentBackupsRestoreAPIAPI.RestoreAgentBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | Backup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreAgentBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentBackupRestoreRequest** | [**AgentBackupRestoreRequest**](AgentBackupRestoreRequest.md) |  | 

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

