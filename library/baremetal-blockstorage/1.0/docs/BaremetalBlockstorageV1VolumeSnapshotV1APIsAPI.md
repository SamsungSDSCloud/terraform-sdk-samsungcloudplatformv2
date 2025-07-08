# \BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI

All URIs are relative to *http://70.50.166.119:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolumeSnapshot**](BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.md#CreateVolumeSnapshot) | **Post** /v1/volumes/{volume_id}/snapshots | Create Volume Snapshot
[**CreateVolumeSnapshotRate**](BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.md#CreateVolumeSnapshotRate) | **Post** /v1/volumes/{volume_id}/snapshot-rates | Create Volume Snapshot Rate
[**CreateVolumeSnapshotSchedule**](BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.md#CreateVolumeSnapshotSchedule) | **Post** /v1/volumes/{volume_id}/snapshot-schedules | Create Volume Snapshot Schedule
[**DeleteVolumeSnapshot**](BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.md#DeleteVolumeSnapshot) | **Delete** /v1/volumes/{volume_id}/snapshots/{snapshot_id} | Delete Volume Snapshot
[**DeleteVolumeSnapshotRate**](BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.md#DeleteVolumeSnapshotRate) | **Delete** /v1/volumes/{volume_id}/snapshot-rates | Delete Volume Snapshot Rate
[**DeleteVolumeSnapshotSchedule**](BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.md#DeleteVolumeSnapshotSchedule) | **Delete** /v1/volumes/{volume_id}/snapshot-schedules | Delete Volume Snapshot Schedule
[**ListVolumeSnapshots**](BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.md#ListVolumeSnapshots) | **Get** /v1/volumes/{volume_id}/snapshots | List Volume Snapshots
[**RestoreVolumeSnapshot**](BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.md#RestoreVolumeSnapshot) | **Put** /v1/volumes/{volume_id}/snapshots/{snapshot_id}/restore | Restore Volume Snapshot
[**SetVolumeSnapshotRate**](BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.md#SetVolumeSnapshotRate) | **Put** /v1/volumes/{volume_id}/snapshot-rates | Set Volume Snapshot Rate
[**SetVolumeSnapshotSchedule**](BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.md#SetVolumeSnapshotSchedule) | **Put** /v1/volumes/{volume_id}/snapshot-schedules | Set Volume Snapshot Schedule



## CreateVolumeSnapshot

> VolumeSnapshotResponse CreateVolumeSnapshot(ctx, volumeId).Execute()

Create Volume Snapshot



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.CreateVolumeSnapshot(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.CreateVolumeSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolumeSnapshot`: VolumeSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.CreateVolumeSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeSnapshotResponse**](VolumeSnapshotResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolumeSnapshotRate

> VolumeSnapshotRateResponse CreateVolumeSnapshotRate(ctx, volumeId).VolumeSnapshotRateRequest(volumeSnapshotRateRequest).Execute()

Create Volume Snapshot Rate



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id
	volumeSnapshotRateRequest := *openapiclient.NewVolumeSnapshotRateRequest(int32(100)) // VolumeSnapshotRateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.CreateVolumeSnapshotRate(context.Background(), volumeId).VolumeSnapshotRateRequest(volumeSnapshotRateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.CreateVolumeSnapshotRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolumeSnapshotRate`: VolumeSnapshotRateResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.CreateVolumeSnapshotRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeSnapshotRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeSnapshotRateRequest** | [**VolumeSnapshotRateRequest**](VolumeSnapshotRateRequest.md) |  | 

### Return type

[**VolumeSnapshotRateResponse**](VolumeSnapshotRateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolumeSnapshotSchedule

> VolumeSnapshotScheduleResponse CreateVolumeSnapshotSchedule(ctx, volumeId).VolumeSnapshotScheduleRequest(volumeSnapshotScheduleRequest).Execute()

Create Volume Snapshot Schedule



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id
	volumeSnapshotScheduleRequest := *openapiclient.NewVolumeSnapshotScheduleRequest(int32(18)) // VolumeSnapshotScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.CreateVolumeSnapshotSchedule(context.Background(), volumeId).VolumeSnapshotScheduleRequest(volumeSnapshotScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.CreateVolumeSnapshotSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolumeSnapshotSchedule`: VolumeSnapshotScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.CreateVolumeSnapshotSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeSnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeSnapshotScheduleRequest** | [**VolumeSnapshotScheduleRequest**](VolumeSnapshotScheduleRequest.md) |  | 

### Return type

[**VolumeSnapshotScheduleResponse**](VolumeSnapshotScheduleResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolumeSnapshot

> VolumeSnapshotDeleteResponse DeleteVolumeSnapshot(ctx, volumeId, snapshotId).Execute()

Delete Volume Snapshot



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
	volumeId := "8bf55e738d4e44b5a21dbe133a42ecbe" // string | Volume id
	snapshotId := "gjireo3921kdkwskjwowo32901mkj" // string | Snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.DeleteVolumeSnapshot(context.Background(), volumeId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.DeleteVolumeSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolumeSnapshot`: VolumeSnapshotDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.DeleteVolumeSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 
**snapshotId** | **string** | Snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VolumeSnapshotDeleteResponse**](VolumeSnapshotDeleteResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolumeSnapshotRate

> VolumeSnapshotRateResponse DeleteVolumeSnapshotRate(ctx, volumeId).Execute()

Delete Volume Snapshot Rate



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.DeleteVolumeSnapshotRate(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.DeleteVolumeSnapshotRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolumeSnapshotRate`: VolumeSnapshotRateResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.DeleteVolumeSnapshotRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeSnapshotRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeSnapshotRateResponse**](VolumeSnapshotRateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolumeSnapshotSchedule

> VolumeSnapshotScheduleResponse DeleteVolumeSnapshotSchedule(ctx, volumeId).Execute()

Delete Volume Snapshot Schedule



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.DeleteVolumeSnapshotSchedule(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.DeleteVolumeSnapshotSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolumeSnapshotSchedule`: VolumeSnapshotScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.DeleteVolumeSnapshotSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeSnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeSnapshotScheduleResponse**](VolumeSnapshotScheduleResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumeSnapshots

> VolumeSnapshotListResponse ListVolumeSnapshots(ctx, volumeId).Execute()

List Volume Snapshots



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.ListVolumeSnapshots(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.ListVolumeSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumeSnapshots`: VolumeSnapshotListResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.ListVolumeSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeSnapshotListResponse**](VolumeSnapshotListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreVolumeSnapshot

> VolumeSnapshotRestoreResponse RestoreVolumeSnapshot(ctx, volumeId, snapshotId).Execute()

Restore Volume Snapshot



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
	volumeId := "8bf55e738d4e44b5a21dbe133a42ecbe" // string | Volume id
	snapshotId := "gjireo3921kdkwskjwowo32901mkj" // string | Snapshot id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.RestoreVolumeSnapshot(context.Background(), volumeId, snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.RestoreVolumeSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreVolumeSnapshot`: VolumeSnapshotRestoreResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.RestoreVolumeSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 
**snapshotId** | **string** | Snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreVolumeSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VolumeSnapshotRestoreResponse**](VolumeSnapshotRestoreResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVolumeSnapshotRate

> VolumeSnapshotRateResponse SetVolumeSnapshotRate(ctx, volumeId).VolumeSnapshotRateRequest(volumeSnapshotRateRequest).Execute()

Set Volume Snapshot Rate



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id
	volumeSnapshotRateRequest := *openapiclient.NewVolumeSnapshotRateRequest(int32(100)) // VolumeSnapshotRateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.SetVolumeSnapshotRate(context.Background(), volumeId).VolumeSnapshotRateRequest(volumeSnapshotRateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.SetVolumeSnapshotRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVolumeSnapshotRate`: VolumeSnapshotRateResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.SetVolumeSnapshotRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVolumeSnapshotRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeSnapshotRateRequest** | [**VolumeSnapshotRateRequest**](VolumeSnapshotRateRequest.md) |  | 

### Return type

[**VolumeSnapshotRateResponse**](VolumeSnapshotRateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVolumeSnapshotSchedule

> VolumeSnapshotScheduleResponse SetVolumeSnapshotSchedule(ctx, volumeId).VolumeSnapshotScheduleRequest(volumeSnapshotScheduleRequest).Execute()

Set Volume Snapshot Schedule



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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id
	volumeSnapshotScheduleRequest := *openapiclient.NewVolumeSnapshotScheduleRequest(int32(18)) // VolumeSnapshotScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.SetVolumeSnapshotSchedule(context.Background(), volumeId).VolumeSnapshotScheduleRequest(volumeSnapshotScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.SetVolumeSnapshotSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVolumeSnapshotSchedule`: VolumeSnapshotScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeSnapshotV1APIsAPI.SetVolumeSnapshotSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVolumeSnapshotScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeSnapshotScheduleRequest** | [**VolumeSnapshotScheduleRequest**](VolumeSnapshotScheduleRequest.md) |  | 

### Return type

[**VolumeSnapshotScheduleResponse**](VolumeSnapshotScheduleResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

