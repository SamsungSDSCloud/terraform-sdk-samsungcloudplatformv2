# \BaremetalBlockstorageV1VolumeV1APIsAPI

All URIs are relative to *http://70.50.166.119:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolume**](BaremetalBlockstorageV1VolumeV1APIsAPI.md#CreateVolume) | **Post** /v1/volumes | Create Volume
[**CreateVolumeAttachments**](BaremetalBlockstorageV1VolumeV1APIsAPI.md#CreateVolumeAttachments) | **Post** /v1/volumes/{volume_id}/attachments | Attach Volume
[**DeleteVolume**](BaremetalBlockstorageV1VolumeV1APIsAPI.md#DeleteVolume) | **Delete** /v1/volumes/{volume_id} | Delete Volume
[**DeleteVolumeAttachments**](BaremetalBlockstorageV1VolumeV1APIsAPI.md#DeleteVolumeAttachments) | **Delete** /v1/volumes/{volume_id}/attachments | Detach Volume
[**ListVolumes**](BaremetalBlockstorageV1VolumeV1APIsAPI.md#ListVolumes) | **Get** /v1/volumes | List Volumes
[**ShowVolume**](BaremetalBlockstorageV1VolumeV1APIsAPI.md#ShowVolume) | **Get** /v1/volumes/{volume_id} | Show Volume



## CreateVolume

> AsyncResponse CreateVolume(ctx).VolumeCreateRequest(volumeCreateRequest).Execute()

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
	volumeCreateRequest := *openapiclient.NewVolumeCreateRequest([]openapiclient.AttachmentListModel{*openapiclient.NewAttachmentListModel()}, openapiclient.DiskType("SSD"), "my-bs-01", int32(10)) // VolumeCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeV1APIsAPI.CreateVolume(context.Background()).VolumeCreateRequest(volumeCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeV1APIsAPI.CreateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolume`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeV1APIsAPI.CreateVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeCreateRequest** | [**VolumeCreateRequest**](VolumeCreateRequest.md) |  | 

### Return type

[**AsyncResponse**](AsyncResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolumeAttachments

> VolumeResponse CreateVolumeAttachments(ctx, volumeId).VolumeAttachmentRequest(volumeAttachmentRequest).Execute()

Attach Volume



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
	volumeAttachmentRequest := *openapiclient.NewVolumeAttachmentRequest([]openapiclient.AttachmentListModel{*openapiclient.NewAttachmentListModel()}) // VolumeAttachmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeV1APIsAPI.CreateVolumeAttachments(context.Background(), volumeId).VolumeAttachmentRequest(volumeAttachmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeV1APIsAPI.CreateVolumeAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolumeAttachments`: VolumeResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeV1APIsAPI.CreateVolumeAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeAttachmentRequest** | [**VolumeAttachmentRequest**](VolumeAttachmentRequest.md) |  | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolume

> AsyncResponse DeleteVolume(ctx, volumeId).Execute()

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
	volumeId := "001e77a81de74b7883ce105760c97cf9" // string | Volume id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeV1APIsAPI.DeleteVolume(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeV1APIsAPI.DeleteVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolume`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeV1APIsAPI.DeleteVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


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


## DeleteVolumeAttachments

> VolumeResponse DeleteVolumeAttachments(ctx, volumeId).VolumeDetachRequest(volumeDetachRequest).Execute()

Detach Volume



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
	volumeDetachRequest := *openapiclient.NewVolumeDetachRequest([]string{"Attachments_example"}) // VolumeDetachRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeV1APIsAPI.DeleteVolumeAttachments(context.Background(), volumeId).VolumeDetachRequest(volumeDetachRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeV1APIsAPI.DeleteVolumeAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolumeAttachments`: VolumeResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeV1APIsAPI.DeleteVolumeAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeDetachRequest** | [**VolumeDetachRequest**](VolumeDetachRequest.md) |  | 

### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumes

> VolumeListResponse ListVolumes(ctx).Limit(limit).Offset(offset).Sort(sort).Name(name).ObjectName(objectName).Execute()

List Volumes



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
	limit := int32(20) // int32 | Number to be displayed on the page (optional) (default to 20)
	offset := int32(0) // int32 | Offset (optional) (default to 0)
	sort := "created_at:desc,name:asc" // string | Sort (optional)
	name := "storage-01" // string | Volume name (optional)
	objectName := "baremetal-server-name" // string | Attached server name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeV1APIsAPI.ListVolumes(context.Background()).Limit(limit).Offset(offset).Sort(sort).Name(name).ObjectName(objectName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeV1APIsAPI.ListVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumes`: VolumeListResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeV1APIsAPI.ListVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number to be displayed on the page | [default to 20]
 **offset** | **int32** | Offset | [default to 0]
 **sort** | **string** | Sort | 
 **name** | **string** | Volume name | 
 **objectName** | **string** | Attached server name | 

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


## ShowVolume

> VolumeResponse ShowVolume(ctx, volumeId).Execute()

Show Volume



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
	resp, r, err := apiClient.BaremetalBlockstorageV1VolumeV1APIsAPI.ShowVolume(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaremetalBlockstorageV1VolumeV1APIsAPI.ShowVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVolume`: VolumeResponse
	fmt.Fprintf(os.Stdout, "Response from `BaremetalBlockstorageV1VolumeV1APIsAPI.ShowVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeResponse**](VolumeResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

