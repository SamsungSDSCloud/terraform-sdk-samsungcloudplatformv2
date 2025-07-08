# \BackupV1BackupAgentsApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckConnectionState**](BackupV1BackupAgentsApiAPI.md#CheckConnectionState) | **Get** /v1/backup-agents/{backup_agent_id}/check-connection-state | Check Backup Master Connection State
[**CreateBackupAgent**](BackupV1BackupAgentsApiAPI.md#CreateBackupAgent) | **Post** /v1/backup-agents | Create Backup Agent
[**DeleteBackupAgent**](BackupV1BackupAgentsApiAPI.md#DeleteBackupAgent) | **Delete** /v1/backup-agents/{backup_agent_id} | Delete Backup Agent
[**ListBackupAgentTargets**](BackupV1BackupAgentsApiAPI.md#ListBackupAgentTargets) | **Get** /v1/backup-agents/targets | List Targets of Backup Agent
[**ListBackupAgents**](BackupV1BackupAgentsApiAPI.md#ListBackupAgents) | **Get** /v1/backup-agents | List Backup Agents
[**ShowBackupAgent**](BackupV1BackupAgentsApiAPI.md#ShowBackupAgent) | **Get** /v1/backup-agents/{backup_agent_id} | Show Backup Agent
[**ShowInstallFilePath**](BackupV1BackupAgentsApiAPI.md#ShowInstallFilePath) | **Get** /v1/backup-agents/agent-install-file-path | Show Backup Agent Install File Path



## CheckConnectionState

> AsyncResponse CheckConnectionState(ctx, backupAgentId).Execute()

Check Backup Master Connection State



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
	backupAgentId := "BACKUP_AGENT-abcd" // string | Backup Agent ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupAgentsApiAPI.CheckConnectionState(context.Background(), backupAgentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupAgentsApiAPI.CheckConnectionState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckConnectionState`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupAgentsApiAPI.CheckConnectionState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupAgentId** | **string** | Backup Agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckConnectionStateRequest struct via the builder pattern


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


## CreateBackupAgent

> SyncResponse CreateBackupAgent(ctx).BackupAgentCreateRequest(backupAgentCreateRequest).Execute()

Create Backup Agent



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
	backupAgentCreateRequest := *openapiclient.NewBackupAgentCreateRequest("89f5ef44-1021-4a5c-8e06-fbb289eac366") // BackupAgentCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupAgentsApiAPI.CreateBackupAgent(context.Background()).BackupAgentCreateRequest(backupAgentCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupAgentsApiAPI.CreateBackupAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBackupAgent`: SyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupAgentsApiAPI.CreateBackupAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupAgentCreateRequest** | [**BackupAgentCreateRequest**](BackupAgentCreateRequest.md) |  | 

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


## DeleteBackupAgent

> SyncResponse DeleteBackupAgent(ctx, backupAgentId).Execute()

Delete Backup Agent



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
	backupAgentId := "BACKUP_AGENT-abcd" // string | Backup Agent ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupAgentsApiAPI.DeleteBackupAgent(context.Background(), backupAgentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupAgentsApiAPI.DeleteBackupAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBackupAgent`: SyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupAgentsApiAPI.DeleteBackupAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupAgentId** | **string** | Backup Agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupAgentRequest struct via the builder pattern


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


## ListBackupAgentTargets

> BackupAgentTargetListResponse ListBackupAgentTargets(ctx).ServerCategory(serverCategory).Size(size).Page(page).Sort(sort).ServerName(serverName).Execute()

List Targets of Backup Agent



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
	serverCategory := "BAREMETAL_SERVER" // string | Backup server category
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "sort_example" // string | Sort (optional)
	serverName := "server-001" // string | Backup server name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupAgentsApiAPI.ListBackupAgentTargets(context.Background()).ServerCategory(serverCategory).Size(size).Page(page).Sort(sort).ServerName(serverName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupAgentsApiAPI.ListBackupAgentTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupAgentTargets`: BackupAgentTargetListResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupAgentsApiAPI.ListBackupAgentTargets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBackupAgentTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverCategory** | **string** | Backup server category | 
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | Sort | 
 **serverName** | **string** | Backup server name | 

### Return type

[**BackupAgentTargetListResponse**](BackupAgentTargetListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupAgents

> BackupAgentListResponse ListBackupAgents(ctx).Size(size).Page(page).Sort(sort).Name(name).ServerBackupAgentIp(serverBackupAgentIp).ServerName(serverName).Execute()

List Backup Agents



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
	name := "agent_abcd12" // string | Backup Agent Name (optional)
	serverBackupAgentIp := "11.22.33.44" // string | Server's Backup Agent IP (optional)
	serverName := "server-001" // string | Backup Agent target server name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupAgentsApiAPI.ListBackupAgents(context.Background()).Size(size).Page(page).Sort(sort).Name(name).ServerBackupAgentIp(serverBackupAgentIp).ServerName(serverName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupAgentsApiAPI.ListBackupAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackupAgents`: BackupAgentListResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupAgentsApiAPI.ListBackupAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBackupAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | Backup Agent Name | 
 **serverBackupAgentIp** | **string** | Server&#39;s Backup Agent IP | 
 **serverName** | **string** | Backup Agent target server name | 

### Return type

[**BackupAgentListResponse**](BackupAgentListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowBackupAgent

> BackupAgentDetailResponse ShowBackupAgent(ctx, backupAgentId).Execute()

Show Backup Agent



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
	backupAgentId := "BACKUP_AGENT-abcd" // string | Backup Agent ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupAgentsApiAPI.ShowBackupAgent(context.Background(), backupAgentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupAgentsApiAPI.ShowBackupAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowBackupAgent`: BackupAgentDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupAgentsApiAPI.ShowBackupAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupAgentId** | **string** | Backup Agent ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowBackupAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupAgentDetailResponse**](BackupAgentDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowInstallFilePath

> InstallFilePathResponse ShowInstallFilePath(ctx).OsType(osType).Execute()

Show Backup Agent Install File Path



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
	osType := "osType_example" // string | OS type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupV1BackupAgentsApiAPI.ShowInstallFilePath(context.Background()).OsType(osType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupV1BackupAgentsApiAPI.ShowInstallFilePath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowInstallFilePath`: InstallFilePathResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupV1BackupAgentsApiAPI.ShowInstallFilePath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowInstallFilePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **osType** | **string** | OS type | 

### Return type

[**InstallFilePathResponse**](InstallFilePathResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

