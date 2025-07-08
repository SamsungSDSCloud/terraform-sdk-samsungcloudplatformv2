# \FilestorageV1SnapshotAPIsAPI

All URIs are relative to *https://filestorage-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshot**](FilestorageV1SnapshotAPIsAPI.md#CreateSnapshot) | **Post** /v1/snapshots | CreateSnapshot
[**DeleteSnapshot**](FilestorageV1SnapshotAPIsAPI.md#DeleteSnapshot) | **Delete** /v1/snapshots/{snapshot_id} | DeleteSnapshot
[**ListSnapshots**](FilestorageV1SnapshotAPIsAPI.md#ListSnapshots) | **Get** /v1/snapshots | ListSnapshots
[**RestoreSnapshot**](FilestorageV1SnapshotAPIsAPI.md#RestoreSnapshot) | **Put** /v1/snapshots/{snapshot_id}/restore | RestoreSnapshot



## CreateSnapshot

> SnapshotShowResponse CreateSnapshot(ctx).SnapshotCreateRequest(snapshotCreateRequest).Execute()

CreateSnapshot



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
	snapshotCreateRequest := *openapiclient.NewSnapshotCreateRequest("bfdbabf2-04d9-4e8b-a205-020f8e6da438") // SnapshotCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1SnapshotAPIsAPI.CreateSnapshot(context.Background()).SnapshotCreateRequest(snapshotCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1SnapshotAPIsAPI.CreateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnapshot`: SnapshotShowResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1SnapshotAPIsAPI.CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **snapshotCreateRequest** | [**SnapshotCreateRequest**](SnapshotCreateRequest.md) |  | 

### Return type

[**SnapshotShowResponse**](SnapshotShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshot

> DeleteSnapshot(ctx, snapshotId).VolumeId(volumeId).Execute()

DeleteSnapshot



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
	snapshotId := "40htrre5b-eww25-ff2b-a426-6afefas65a3b" // string | Snapshot ID
	volumeId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilestorageV1SnapshotAPIsAPI.DeleteSnapshot(context.Background(), snapshotId).VolumeId(volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1SnapshotAPIsAPI.DeleteSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRequest struct via the builder pattern


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


## ListSnapshots

> SnapshotListResponse ListSnapshots(ctx).VolumeId(volumeId).Execute()

ListSnapshots



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
	resp, r, err := apiClient.FilestorageV1SnapshotAPIsAPI.ListSnapshots(context.Background()).VolumeId(volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1SnapshotAPIsAPI.ListSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSnapshots`: SnapshotListResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1SnapshotAPIsAPI.ListSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | **string** | Volume ID | 

### Return type

[**SnapshotListResponse**](SnapshotListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreSnapshot

> SnapshotRestoreResponse RestoreSnapshot(ctx, snapshotId).VolumeId(volumeId).Execute()

RestoreSnapshot



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
	snapshotId := "40htrre5b-eww25-ff2b-a426-6afefas65a3b" // string | Snapshot ID
	volumeId := "bfdbabf2-04d9-4e8b-a205-020f8e6da438" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1SnapshotAPIsAPI.RestoreSnapshot(context.Background(), snapshotId).VolumeId(volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1SnapshotAPIsAPI.RestoreSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreSnapshot`: SnapshotRestoreResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1SnapshotAPIsAPI.RestoreSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeId** | **string** | Volume ID | 

### Return type

[**SnapshotRestoreResponse**](SnapshotRestoreResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

