# \FilestorageV1VolumeAPIsAPI

All URIs are relative to *https://filestorage-a.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolume**](FilestorageV1VolumeAPIsAPI.md#CreateVolume) | **Post** /v1/volumes | CreateVolume
[**DeleteVolume**](FilestorageV1VolumeAPIsAPI.md#DeleteVolume) | **Delete** /v1/volumes/{volume_id} | DeleteVolume
[**ListVolumes**](FilestorageV1VolumeAPIsAPI.md#ListVolumes) | **Get** /v1/volumes | ListVolumes
[**SetVolume**](FilestorageV1VolumeAPIsAPI.md#SetVolume) | **Put** /v1/volumes/{volume_id} | SetVolume
[**ShowVolume**](FilestorageV1VolumeAPIsAPI.md#ShowVolume) | **Get** /v1/volumes/{volume_id} | ShowVolume



## CreateVolume

> VolumeCreateResponse CreateVolume(ctx).VolumeCreateRequest(volumeCreateRequest).Execute()

CreateVolume



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
	volumeCreateRequest := *openapiclient.NewVolumeCreateRequest("my_volume", "NFS", "HDD") // VolumeCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1VolumeAPIsAPI.CreateVolume(context.Background()).VolumeCreateRequest(volumeCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeAPIsAPI.CreateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolume`: VolumeCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1VolumeAPIsAPI.CreateVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeCreateRequest** | [**VolumeCreateRequest**](VolumeCreateRequest.md) |  | 

### Return type

[**VolumeCreateResponse**](VolumeCreateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolume

> DeleteVolume(ctx, volumeId).Execute()

DeleteVolume



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
	r, err := apiClient.FilestorageV1VolumeAPIsAPI.DeleteVolume(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeAPIsAPI.DeleteVolume``: %v\n", err)
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


## ListVolumes

> VolumeListResponse ListVolumes(ctx).Offset(offset).Limit(limit).Name(name).TypeName(typeName).Sort(sort).Execute()

ListVolumes



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
	offset := int32(56) // int32 | offset (optional) (default to 0)
	limit := int32(56) // int32 | limit (optional) (default to 20)
	name := "my_volume" // string | Volume Name (optional)
	typeName := "HDD" // string | Volume Type Name (optional)
	sort := "sort_example" // string | sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilestorageV1VolumeAPIsAPI.ListVolumes(context.Background()).Offset(offset).Limit(limit).Name(name).TypeName(typeName).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeAPIsAPI.ListVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumes`: VolumeListResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1VolumeAPIsAPI.ListVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | offset | [default to 0]
 **limit** | **int32** | limit | [default to 20]
 **name** | **string** | Volume Name | 
 **typeName** | **string** | Volume Type Name | 
 **sort** | **string** | sort | 

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


## SetVolume

> SetVolume(ctx, volumeId).VolumeSetRequest(volumeSetRequest).Execute()

SetVolume



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
	volumeSetRequest := *openapiclient.NewVolumeSetRequest(true) // VolumeSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilestorageV1VolumeAPIsAPI.SetVolume(context.Background(), volumeId).VolumeSetRequest(volumeSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeAPIsAPI.SetVolume``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeSetRequest** | [**VolumeSetRequest**](VolumeSetRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVolume

> VolumeShowResponse ShowVolume(ctx, volumeId).Execute()

ShowVolume



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
	resp, r, err := apiClient.FilestorageV1VolumeAPIsAPI.ShowVolume(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilestorageV1VolumeAPIsAPI.ShowVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVolume`: VolumeShowResponse
	fmt.Fprintf(os.Stdout, "Response from `FilestorageV1VolumeAPIsAPI.ShowVolume`: %v\n", resp)
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

