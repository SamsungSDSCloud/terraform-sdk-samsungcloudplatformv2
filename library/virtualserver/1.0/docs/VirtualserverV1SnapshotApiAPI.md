# \VirtualserverV1SnapshotApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshot**](VirtualserverV1SnapshotApiAPI.md#CreateSnapshot) | **Post** /v1/snapshots | Create a snapshot
[**DeleteSnapshot**](VirtualserverV1SnapshotApiAPI.md#DeleteSnapshot) | **Delete** /v1/snapshots/{snapshot_id} | Delete a snapshot
[**ListSnapshot**](VirtualserverV1SnapshotApiAPI.md#ListSnapshot) | **Get** /v1/snapshots | List snapshots and details
[**ShowSnapshot**](VirtualserverV1SnapshotApiAPI.md#ShowSnapshot) | **Get** /v1/snapshots/{snapshot_id} | Show a snapshot’s details
[**UpdateSnapshot**](VirtualserverV1SnapshotApiAPI.md#UpdateSnapshot) | **Put** /v1/snapshots/{snapshot_id} | Update a snapshot



## CreateSnapshot

> SnapshotDetailResponse CreateSnapshot(ctx).SnapshotCreateRequest(snapshotCreateRequest).Execute()

Create a snapshot



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
	snapshotCreateRequest := *openapiclient.NewSnapshotCreateRequest("snapshot-1", "760fee21-42b7-4717-adb7-511b31f8af74") // SnapshotCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1SnapshotApiAPI.CreateSnapshot(context.Background()).SnapshotCreateRequest(snapshotCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1SnapshotApiAPI.CreateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnapshot`: SnapshotDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1SnapshotApiAPI.CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **snapshotCreateRequest** | [**SnapshotCreateRequest**](SnapshotCreateRequest.md) |  | 

### Return type

[**SnapshotDetailResponse**](SnapshotDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshot

> DeleteSnapshot(ctx, snapshotId).Execute()

Delete a snapshot



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
	snapshotId := "cceed636-1f1e-4bb0-b85c-4e5b9e0bf790" // string | Snapshot ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1SnapshotApiAPI.DeleteSnapshot(context.Background(), snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1SnapshotApiAPI.DeleteSnapshot``: %v\n", err)
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


### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSnapshot

> SnapshotListResponse ListSnapshot(ctx).VolumeId(volumeId).Limit(limit).Offset(offset).Marker(marker).WithCount(withCount).Name(name).Execute()

List snapshots and details



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
	volumeId := "760fee21-42b7-4717-adb7-511b31f8af74" // string | Volume Id (optional)
	limit := int32(20) // int32 | Page size of items (optional)
	offset := int32(0) // int32 | Index of where to start in the list (optional)
	marker := "a7803edaea1b4235956691f150d6c5b1" // string | The last item's ID from previous request (optional)
	withCount := true // bool | Whether to show count (optional)
	name := "snapshot-1" // string | Snapshot name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1SnapshotApiAPI.ListSnapshot(context.Background()).VolumeId(volumeId).Limit(limit).Offset(offset).Marker(marker).WithCount(withCount).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1SnapshotApiAPI.ListSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSnapshot`: SnapshotListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1SnapshotApiAPI.ListSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | **string** | Volume Id | 
 **limit** | **int32** | Page size of items | 
 **offset** | **int32** | Index of where to start in the list | 
 **marker** | **string** | The last item&#39;s ID from previous request | 
 **withCount** | **bool** | Whether to show count | 
 **name** | **string** | Snapshot name | 

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


## ShowSnapshot

> SnapshotDetailResponse ShowSnapshot(ctx, snapshotId).Execute()

Show a snapshot’s details



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
	snapshotId := "cceed636-1f1e-4bb0-b85c-4e5b9e0bf790" // string | Snapshot ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1SnapshotApiAPI.ShowSnapshot(context.Background(), snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1SnapshotApiAPI.ShowSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowSnapshot`: SnapshotDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1SnapshotApiAPI.ShowSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnapshotDetailResponse**](SnapshotDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnapshot

> SnapshotDetailResponse UpdateSnapshot(ctx, snapshotId).SnapshotUpdateRequest(snapshotUpdateRequest).Execute()

Update a snapshot



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
	snapshotId := "cceed636-1f1e-4bb0-b85c-4e5b9e0bf790" // string | Snapshot ID
	snapshotUpdateRequest := *openapiclient.NewSnapshotUpdateRequest("Name_example") // SnapshotUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1SnapshotApiAPI.UpdateSnapshot(context.Background(), snapshotId).SnapshotUpdateRequest(snapshotUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1SnapshotApiAPI.UpdateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSnapshot`: SnapshotDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1SnapshotApiAPI.UpdateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | Snapshot ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotUpdateRequest** | [**SnapshotUpdateRequest**](SnapshotUpdateRequest.md) |  | 

### Return type

[**SnapshotDetailResponse**](SnapshotDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

