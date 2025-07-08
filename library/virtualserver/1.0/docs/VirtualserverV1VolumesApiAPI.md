# \VirtualserverV1VolumesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolume**](VirtualserverV1VolumesApiAPI.md#CreateVolume) | **Post** /v1/volumes | Create Volume
[**DeleteVolume**](VirtualserverV1VolumesApiAPI.md#DeleteVolume) | **Delete** /v1/volumes/{volume_id} | Delete Volume
[**ExtendVolume**](VirtualserverV1VolumesApiAPI.md#ExtendVolume) | **Put** /v1/volumes/{volume_id}/size | Extend Volume
[**ListVolumes**](VirtualserverV1VolumesApiAPI.md#ListVolumes) | **Get** /v1/volumes | List volumes and details
[**RevertVolumeToSnapshot**](VirtualserverV1VolumesApiAPI.md#RevertVolumeToSnapshot) | **Post** /v1/volumes/{volume_id}/revert | Revert Volume to snapshot
[**ShowVolume**](VirtualserverV1VolumesApiAPI.md#ShowVolume) | **Get** /v1/volumes/{volume_id} | Show Volume&#39;s details
[**UpdateVolume**](VirtualserverV1VolumesApiAPI.md#UpdateVolume) | **Put** /v1/volumes/{volume_id} | Update Volume



## CreateVolume

> VolumeShowResponse CreateVolume(ctx).VolumeCreateRequest(volumeCreateRequest).Execute()

Create Volume



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
	volumeCreateRequest := *openapiclient.NewVolumeCreateRequest("volume01", int32(8)) // VolumeCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumesApiAPI.CreateVolume(context.Background()).VolumeCreateRequest(volumeCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumesApiAPI.CreateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolume`: VolumeShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumesApiAPI.CreateVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeCreateRequest** | [**VolumeCreateRequest**](VolumeCreateRequest.md) |  | 

### Return type

[**VolumeShowResponse**](VolumeShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolume

> DeleteVolume(ctx, volumeId).Cascade(cascade).Execute()

Delete Volume



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
	volumeId := "760fee21-42b7-4717-adb7-511b31f8af74" // string | Volume ID
	cascade := true // bool | Whether to delete snapshots together (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1VolumesApiAPI.DeleteVolume(context.Background(), volumeId).Cascade(cascade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumesApiAPI.DeleteVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **bool** | Whether to delete snapshots together | 

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


## ExtendVolume

> VolumeShowResponse ExtendVolume(ctx, volumeId).VolumeExtendRequest(volumeExtendRequest).Execute()

Extend Volume



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
	volumeId := "760fee21-42b7-4717-adb7-511b31f8af74" // string | Volume ID
	volumeExtendRequest := *openapiclient.NewVolumeExtendRequest(int32(8)) // VolumeExtendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumesApiAPI.ExtendVolume(context.Background(), volumeId).VolumeExtendRequest(volumeExtendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumesApiAPI.ExtendVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExtendVolume`: VolumeShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumesApiAPI.ExtendVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeExtendRequest** | [**VolumeExtendRequest**](VolumeExtendRequest.md) |  | 

### Return type

[**VolumeShowResponse**](VolumeShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumes

> VolumeListResponse ListVolumes(ctx).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Offset(offset).Name(name).State(state).Bootable(bootable).Execute()

List volumes and details



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
	withCount := true // bool | Whether to show count (optional)
	limit := int32(20) // int32 | Page size of items (optional)
	marker := "760fee21-42b7-4717-adb7-511b31f8af74" // string | The last item's ID from previous request (optional)
	sort := "name:asc" // string | Comma-separated list of sort keys and optional sort directions in the form of <key>[:<direction>] (optional)
	offset := int32(0) // int32 | Index of where to start in the list (optional)
	name := "volume01" // string | Volume name (optional)
	state := "available" // string | Volume state (optional)
	bootable := false // bool | Bootable (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumesApiAPI.ListVolumes(context.Background()).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Offset(offset).Name(name).State(state).Bootable(bootable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumesApiAPI.ListVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumes`: VolumeListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumesApiAPI.ListVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withCount** | **bool** | Whether to show count | 
 **limit** | **int32** | Page size of items | 
 **marker** | **string** | The last item&#39;s ID from previous request | 
 **sort** | **string** | Comma-separated list of sort keys and optional sort directions in the form of &lt;key&gt;[:&lt;direction&gt;] | 
 **offset** | **int32** | Index of where to start in the list | 
 **name** | **string** | Volume name | 
 **state** | **string** | Volume state | 
 **bootable** | **bool** | Bootable | 

### Return type

[**VolumeListResponse**](VolumeListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertVolumeToSnapshot

> RevertVolumeToSnapshot(ctx, volumeId).VolumeRevertRequest(volumeRevertRequest).Execute()

Revert Volume to snapshot



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
	volumeId := "760fee21-42b7-4717-adb7-511b31f8af74" // string | Volume ID
	volumeRevertRequest := *openapiclient.NewVolumeRevertRequest("cceed636-1f1e-4bb0-b85c-4e5b9e0bf790") // VolumeRevertRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1VolumesApiAPI.RevertVolumeToSnapshot(context.Background(), volumeId).VolumeRevertRequest(volumeRevertRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumesApiAPI.RevertVolumeToSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertVolumeToSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeRevertRequest** | [**VolumeRevertRequest**](VolumeRevertRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVolume

> VolumeShowResponse ShowVolume(ctx, volumeId).Execute()

Show Volume's details



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
	volumeId := "760fee21-42b7-4717-adb7-511b31f8af74" // string | Volume ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumesApiAPI.ShowVolume(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumesApiAPI.ShowVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVolume`: VolumeShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumesApiAPI.ShowVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeShowResponse**](VolumeShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolume

> VolumeShowResponse UpdateVolume(ctx, volumeId).VolumeUpdateRequest(volumeUpdateRequest).Execute()

Update Volume



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
	volumeId := "760fee21-42b7-4717-adb7-511b31f8af74" // string | Volume ID
	volumeUpdateRequest := *openapiclient.NewVolumeUpdateRequest() // VolumeUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumesApiAPI.UpdateVolume(context.Background(), volumeId).VolumeUpdateRequest(volumeUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumesApiAPI.UpdateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVolume`: VolumeShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumesApiAPI.UpdateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeUpdateRequest** | [**VolumeUpdateRequest**](VolumeUpdateRequest.md) |  | 

### Return type

[**VolumeShowResponse**](VolumeShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

