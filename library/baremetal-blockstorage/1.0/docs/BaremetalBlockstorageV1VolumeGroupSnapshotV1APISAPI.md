# \BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI

All URIs are relative to *http://70.50.166.119:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolumeGroupSnapshot**](BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.md#CreateVolumeGroupSnapshot) | **Post** /v1/volume-groups/{volume_group_id}/snapshots | Create Volume Group Snapshot
[**CreateVolumeGroupSnapshotSchedule**](BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.md#CreateVolumeGroupSnapshotSchedule) | **Post** /v1/volume-groups/{volume_group_id}/snapshot-schedules | Create Volume Group Snapshot Schedule
[**DeleteVolumeGroupSnapshot**](BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.md#DeleteVolumeGroupSnapshot) | **Delete** /v1/volume-groups/{volume_group_id}/snapshots/{snapshot_id} | Delete Volume Group Snapshot
[**DeleteVolumeGroupSnapshotSchedule**](BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.md#DeleteVolumeGroupSnapshotSchedule) | **Delete** /v1/volume-groups/{volume_group_id}/snapshot-schedules | Delete Volume Group Snapshot Schedule
[**ListVolumeGroupSnapshots**](BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.md#ListVolumeGroupSnapshots) | **Get** /v1/volume-groups/{volume_group_id}/snapshots | List Volume Group Snapshots
[**RestoreVolumeGroupSnapshot**](BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.md#RestoreVolumeGroupSnapshot) | **Put** /v1/volume-groups/{volume_group_id}/snapshots/{snapshot_id}/restore | Restore Volume Group Snapshot
[**SetVolumeGroupSnapshotSchedule**](BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.md#SetVolumeGroupSnapshotSchedule) | **Put** /v1/volume-groups/{volume_group_id}/snapshot-schedules | Set Volume Group Snapshot Schedule



## CreateVolumeGroupSnapshot

> VolumeGroupSnapshotResponse CreateVolumeGroupSnapshot(ctx, volumeGroupId).Execute()

Create Volume Group Snapshot



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.CreateVolumeGroupSnapshot(context.Background(), volumeGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.CreateVolumeGroupSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolumeGroupSnapshot`: VolumeGroupSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.CreateVolumeGroupSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeGroupSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeGroupSnapshotResponse**](VolumeGroupSnapshotResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolumeGroupSnapshotSchedule

> VolumeGroupSnapshotScheduleResponse CreateVolumeGroupSnapshotSchedule(ctx, volumeGroupId).VolumeGroupSnapshotScheduleRequest(volumeGroupSnapshotScheduleRequest).Execute()

Create Volume Group Snapshot Schedule



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id
	volumeGroupSnapshotScheduleRequest := *openapiclient.NewVolumeGroupSnapshotScheduleRequest(int32(18)) // VolumeGroupSnapshotScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.CreateVolumeGroupSnapshotSchedule(context.Background(), volumeGroupId).VolumeGroupSnapshotScheduleRequest(volumeGroupSnapshotScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.CreateVolumeGroupSnapshotSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolumeGroupSnapshotSchedule`: VolumeGroupSnapshotScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.CreateVolumeGroupSnapshotSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeGroupSnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeGroupSnapshotScheduleRequest** | [**VolumeGroupSnapshotScheduleRequest**](VolumeGroupSnapshotScheduleRequest.md) |  | 

### Return type

[**VolumeGroupSnapshotScheduleResponse**](VolumeGroupSnapshotScheduleResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolumeGroupSnapshot

> VolumeGroupSnapshotDeleteResponse DeleteVolumeGroupSnapshot(ctx, volumeGroupId, snapshotId).Execute()

Delete Volume Group Snapshot



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
	volumeGroupId := "d9b05e9d377d4dbd9c14e4fa20b4cadd" // string | Volume group id
	snapshotId := "f765e80a-8cfb-11ef-9f89-00a0b8c90a2c" // string | Snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.DeleteVolumeGroupSnapshot(context.Background(), volumeGroupId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.DeleteVolumeGroupSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolumeGroupSnapshot`: VolumeGroupSnapshotDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.DeleteVolumeGroupSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 
**snapshotId** | **string** | Snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeGroupSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VolumeGroupSnapshotDeleteResponse**](VolumeGroupSnapshotDeleteResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolumeGroupSnapshotSchedule

> VolumeGroupSnapshotScheduleResponse DeleteVolumeGroupSnapshotSchedule(ctx, volumeGroupId).Execute()

Delete Volume Group Snapshot Schedule



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.DeleteVolumeGroupSnapshotSchedule(context.Background(), volumeGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.DeleteVolumeGroupSnapshotSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolumeGroupSnapshotSchedule`: VolumeGroupSnapshotScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.DeleteVolumeGroupSnapshotSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeGroupSnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeGroupSnapshotScheduleResponse**](VolumeGroupSnapshotScheduleResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumeGroupSnapshots

> VolumeGroupSnapshotListResponse ListVolumeGroupSnapshots(ctx, volumeGroupId).Execute()

List Volume Group Snapshots



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.ListVolumeGroupSnapshots(context.Background(), volumeGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.ListVolumeGroupSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumeGroupSnapshots`: VolumeGroupSnapshotListResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.ListVolumeGroupSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeGroupSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeGroupSnapshotListResponse**](VolumeGroupSnapshotListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreVolumeGroupSnapshot

> VolumeGroupSnapshotRestoreResponse RestoreVolumeGroupSnapshot(ctx, volumeGroupId, snapshotId).Execute()

Restore Volume Group Snapshot



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
	volumeGroupId := "d9b05e9d377d4dbd9c14e4fa20b4cadd" // string | Volume group id
	snapshotId := "f765e80a-8cfb-11ef-9f89-00a0b8c90a2c" // string | Snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.RestoreVolumeGroupSnapshot(context.Background(), volumeGroupId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.RestoreVolumeGroupSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreVolumeGroupSnapshot`: VolumeGroupSnapshotRestoreResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.RestoreVolumeGroupSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 
**snapshotId** | **string** | Snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreVolumeGroupSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VolumeGroupSnapshotRestoreResponse**](VolumeGroupSnapshotRestoreResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVolumeGroupSnapshotSchedule

> VolumeGroupSnapshotScheduleResponse SetVolumeGroupSnapshotSchedule(ctx, volumeGroupId).VolumeGroupSnapshotScheduleRequest(volumeGroupSnapshotScheduleRequest).Execute()

Set Volume Group Snapshot Schedule



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
	volumeGroupId := "c9ecae8b973b425c81c3817893cd8063" // string | Volume group id
	volumeGroupSnapshotScheduleRequest := *openapiclient.NewVolumeGroupSnapshotScheduleRequest(int32(18)) // VolumeGroupSnapshotScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.SetVolumeGroupSnapshotSchedule(context.Background(), volumeGroupId).VolumeGroupSnapshotScheduleRequest(volumeGroupSnapshotScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.SetVolumeGroupSnapshotSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVolumeGroupSnapshotSchedule`: VolumeGroupSnapshotScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeGroupSnapshotV1APISAPI.SetVolumeGroupSnapshotSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeGroupId** | **string** | Volume group id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVolumeGroupSnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeGroupSnapshotScheduleRequest** | [**VolumeGroupSnapshotScheduleRequest**](VolumeGroupSnapshotScheduleRequest.md) |  | 

### Return type

[**VolumeGroupSnapshotScheduleResponse**](VolumeGroupSnapshotScheduleResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

