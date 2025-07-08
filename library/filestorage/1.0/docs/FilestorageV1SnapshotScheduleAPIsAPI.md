# \FilestorageV1SnapshotScheduleAPIsAPI

All URIs are relative to *https://filestorage-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshotSchedule**](FilestorageV1SnapshotScheduleAPIsAPI.md#CreateSnapshotSchedule) | **Post** /v1/snapshot-schedules | CreateSnapshotSchedule
[**DeleteSnapshotSchedule**](FilestorageV1SnapshotScheduleAPIsAPI.md#DeleteSnapshotSchedule) | **Delete** /v1/snapshot-schedules/{snapshot_schedule_id} | DeleteSnapshotSchedule
[**ListSnapshotSchedule**](FilestorageV1SnapshotScheduleAPIsAPI.md#ListSnapshotSchedule) | **Get** /v1/snapshot-schedules | ListSnapshotSchedule
[**SetSnapshotSchedule**](FilestorageV1SnapshotScheduleAPIsAPI.md#SetSnapshotSchedule) | **Put** /v1/snapshot-schedules/{snapshot_schedule_id} | SetSnapshotSchedule



## CreateSnapshotSchedule

> SnapshotScheduleCreateResponse CreateSnapshotSchedule(ctx).SnapshotScheduleCreateRequest(snapshotScheduleCreateRequest).Execute()

CreateSnapshotSchedule



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
	snapshotScheduleCreateRequest := *openapiclient.NewSnapshotScheduleCreateRequest(*openapiclient.NewSnapshotSchedule("WEEKLY", "23"), "bfdbabf2-04d9-4e8b-a205-020f8e6da438") // SnapshotScheduleCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1SnapshotScheduleAPIsAPI.CreateSnapshotSchedule(context.Background()).SnapshotScheduleCreateRequest(snapshotScheduleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1SnapshotScheduleAPIsAPI.CreateSnapshotSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnapshotSchedule`: SnapshotScheduleCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1SnapshotScheduleAPIsAPI.CreateSnapshotSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **snapshotScheduleCreateRequest** | [**SnapshotScheduleCreateRequest**](SnapshotScheduleCreateRequest.md) |  | 

### Return type

[**SnapshotScheduleCreateResponse**](SnapshotScheduleCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshotSchedule

> DeleteSnapshotSchedule(ctx, snapshotScheduleId).VolumeId(volumeId).Execute()

DeleteSnapshotSchedule



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
	snapshotScheduleId := "d02d1a74-1871-4a3d-bdfc-6e60e9ea0f31" // string | Snapshot Schedule ID
	volumeId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilestorageV1SnapshotScheduleAPIsAPI.DeleteSnapshotSchedule(context.Background(), snapshotScheduleId).VolumeId(volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1SnapshotScheduleAPIsAPI.DeleteSnapshotSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotScheduleId** | **string** | Snapshot Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeId** | **string** | Volume ID | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSnapshotSchedule

> SnapshotScheduleListResponse ListSnapshotSchedule(ctx).VolumeId(volumeId).Execute()

ListSnapshotSchedule



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
	volumeId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1SnapshotScheduleAPIsAPI.ListSnapshotSchedule(context.Background()).VolumeId(volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1SnapshotScheduleAPIsAPI.ListSnapshotSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSnapshotSchedule`: SnapshotScheduleListResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1SnapshotScheduleAPIsAPI.ListSnapshotSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | **string** | Volume ID | 

### Return type

[**SnapshotScheduleListResponse**](SnapshotScheduleListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSnapshotSchedule

> SnapshotScheduleSetResponse SetSnapshotSchedule(ctx, snapshotScheduleId).VolumeId(volumeId).SnapshotScheduleSetRequest(snapshotScheduleSetRequest).Execute()

SetSnapshotSchedule



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
	snapshotScheduleId := "d02d1a74-1871-4a3d-bdfc-6e60e9ea0f31" // string | Snapshot Schedule ID
	volumeId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Volume ID
	snapshotScheduleSetRequest := *openapiclient.NewSnapshotScheduleSetRequest(*openapiclient.NewSnapshotSchedule("WEEKLY", "23")) // SnapshotScheduleSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1SnapshotScheduleAPIsAPI.SetSnapshotSchedule(context.Background(), snapshotScheduleId).VolumeId(volumeId).SnapshotScheduleSetRequest(snapshotScheduleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1SnapshotScheduleAPIsAPI.SetSnapshotSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSnapshotSchedule`: SnapshotScheduleSetResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1SnapshotScheduleAPIsAPI.SetSnapshotSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotScheduleId** | **string** | Snapshot Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeId** | **string** | Volume ID | 
 **snapshotScheduleSetRequest** | [**SnapshotScheduleSetRequest**](SnapshotScheduleSetRequest.md) |  | 

### Return type

[**SnapshotScheduleSetResponse**](SnapshotScheduleSetResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

