# \VirtualserverV1VirtualserverV1ImagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImage**](VirtualserverV1VirtualserverV1ImagesAPI.md#CreateImage) | **Post** /v1/images | Create Image
[**DeleteImage**](VirtualserverV1VirtualserverV1ImagesAPI.md#DeleteImage) | **Delete** /v1/images/{image_id} | Delete Image
[**ListImages**](VirtualserverV1VirtualserverV1ImagesAPI.md#ListImages) | **Get** /v1/images | List Images
[**ListPendingImages**](VirtualserverV1VirtualserverV1ImagesAPI.md#ListPendingImages) | **Get** /v1/pending-images | List pending Images
[**ShowImage**](VirtualserverV1VirtualserverV1ImagesAPI.md#ShowImage) | **Get** /v1/images/{image_id} | Show Image
[**UpdateImage**](VirtualserverV1VirtualserverV1ImagesAPI.md#UpdateImage) | **Put** /v1/images/{image_id} | Update Image



## CreateImage

> ImageShowResponse CreateImage(ctx).ImageCreateRequest(imageCreateRequest).Execute()

Create Image



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
	imageCreateRequest := *openapiclient.NewImageCreateRequest("image_name", openapiclient.ImageOsDistroEnum("alma"), "https://object-store.kr-west1.s.samsungsdscloud.com/8989447062e04a818baf9e073fd04fa7/bucket/object.qcow2") // ImageCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VirtualserverV1ImagesAPI.CreateImage(context.Background()).ImageCreateRequest(imageCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VirtualserverV1ImagesAPI.CreateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateImage`: ImageShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VirtualserverV1ImagesAPI.CreateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageCreateRequest** | [**ImageCreateRequest**](ImageCreateRequest.md) |  | 

### Return type

[**ImageShowResponse**](ImageShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImage

> DeleteImage(ctx, imageId).Execute()

Delete Image



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
	imageId := "1bea47ed-f6a9-463b-b423-14b9cca9ad27" // string | Image ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1VirtualserverV1ImagesAPI.DeleteImage(context.Background(), imageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VirtualserverV1ImagesAPI.DeleteImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


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


## ListImages

> ImageListResponse ListImages(ctx).Limit(limit).Marker(marker).ScpImageType(scpImageType).ScpOriginalImageType(scpOriginalImageType).Name(name).OsDistro(osDistro).Status(status).Visibility(visibility).Sort(sort).Execute()

List Images



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
	limit := int32(20) // int32 | Limit (optional)
	marker := "1bea47ed-f6a9-463b-b423-14b9cca9ad27" // string | Marker (optional)
	scpImageType := "8.8" // string | Image type (optional)
	scpOriginalImageType := "standard" // string | Original Image type (optional)
	name := "image_name" // string | Name (optional)
	osDistro := "alma" // string | OS distribution (optional)
	status := "active" // string | Status (optional)
	visibility := openapiclient.ImageVisibilityEnum("public") // ImageVisibilityEnum | Visibility (optional)
	sort := "created_at:desc" // string | Sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VirtualserverV1ImagesAPI.ListImages(context.Background()).Limit(limit).Marker(marker).ScpImageType(scpImageType).ScpOriginalImageType(scpOriginalImageType).Name(name).OsDistro(osDistro).Status(status).Visibility(visibility).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VirtualserverV1ImagesAPI.ListImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImages`: ImageListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VirtualserverV1ImagesAPI.ListImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit | 
 **marker** | **string** | Marker | 
 **scpImageType** | **string** | Image type | 
 **scpOriginalImageType** | **string** | Original Image type | 
 **name** | **string** | Name | 
 **osDistro** | **string** | OS distribution | 
 **status** | **string** | Status | 
 **visibility** | [**ImageVisibilityEnum**](ImageVisibilityEnum.md) | Visibility | 
 **sort** | **string** | Sort | 

### Return type

[**ImageListResponse**](ImageListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPendingImages

> ImageListResponse ListPendingImages(ctx).ScpImageType(scpImageType).ScpOriginalImageType(scpOriginalImageType).Execute()

List pending Images



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
	scpImageType := "custom" // string | Image type (optional)
	scpOriginalImageType := "standard" // string | Original Image type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VirtualserverV1ImagesAPI.ListPendingImages(context.Background()).ScpImageType(scpImageType).ScpOriginalImageType(scpOriginalImageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VirtualserverV1ImagesAPI.ListPendingImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPendingImages`: ImageListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VirtualserverV1ImagesAPI.ListPendingImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPendingImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scpImageType** | **string** | Image type | 
 **scpOriginalImageType** | **string** | Original Image type | 

### Return type

[**ImageListResponse**](ImageListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowImage

> ImageShowResponse ShowImage(ctx, imageId).Execute()

Show Image



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
	imageId := "1bea47ed-f6a9-463b-b423-14b9cca9ad27" // string | Image ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VirtualserverV1ImagesAPI.ShowImage(context.Background(), imageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VirtualserverV1ImagesAPI.ShowImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowImage`: ImageShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VirtualserverV1ImagesAPI.ShowImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImageShowResponse**](ImageShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImage

> ImageShowResponse UpdateImage(ctx, imageId).ImageSetRequest(imageSetRequest).Execute()

Update Image



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
	imageId := "1bea47ed-f6a9-463b-b423-14b9cca9ad27" // string | Image ID
	imageSetRequest := *openapiclient.NewImageSetRequest() // ImageSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VirtualserverV1ImagesAPI.UpdateImage(context.Background(), imageId).ImageSetRequest(imageSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VirtualserverV1ImagesAPI.UpdateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateImage`: ImageShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VirtualserverV1ImagesAPI.UpdateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imageSetRequest** | [**ImageSetRequest**](ImageSetRequest.md) |  | 

### Return type

[**ImageShowResponse**](ImageShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

