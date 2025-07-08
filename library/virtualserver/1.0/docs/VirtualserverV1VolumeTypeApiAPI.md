# \VirtualserverV1VolumeTypeApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListVolumeTypes**](VirtualserverV1VolumeTypeApiAPI.md#ListVolumeTypes) | **Get** /v1/volume-types | List volume types and details
[**ShowDefaultVolumeType**](VirtualserverV1VolumeTypeApiAPI.md#ShowDefaultVolumeType) | **Get** /v1/volume-types/default | Show default Volume type’s details
[**ShowVolumeType**](VirtualserverV1VolumeTypeApiAPI.md#ShowVolumeType) | **Get** /v1/volume-types/{volume_type_id} | Show Volume type’s details



## ListVolumeTypes

> VolumeTypeListResponse ListVolumeTypes(ctx).Limit(limit).Offset(offset).Marker(marker).Sort(sort).Execute()

List volume types and details



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
	limit := int32(20) // int32 | Page size of items (optional)
	offset := int32(0) // int32 | Index of where to start in the list (optional)
	marker := "c643ac35-3aec-43e6-b97a-3d6f666c5207" // string | The last item's ID from previous request (optional)
	sort := "name:asc" // string | Comma-separated list of sort keys and optional sort directions in the form of <key>[:<direction>] (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumeTypeApiAPI.ListVolumeTypes(context.Background()).Limit(limit).Offset(offset).Marker(marker).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumeTypeApiAPI.ListVolumeTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumeTypes`: VolumeTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumeTypeApiAPI.ListVolumeTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Page size of items | 
 **offset** | **int32** | Index of where to start in the list | 
 **marker** | **string** | The last item&#39;s ID from previous request | 
 **sort** | **string** | Comma-separated list of sort keys and optional sort directions in the form of &lt;key&gt;[:&lt;direction&gt;] | 

### Return type

[**VolumeTypeListResponse**](VolumeTypeListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowDefaultVolumeType

> VolumeTypeDetailResponse ShowDefaultVolumeType(ctx).Execute()

Show default Volume type’s details



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumeTypeApiAPI.ShowDefaultVolumeType(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumeTypeApiAPI.ShowDefaultVolumeType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowDefaultVolumeType`: VolumeTypeDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumeTypeApiAPI.ShowDefaultVolumeType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowDefaultVolumeTypeRequest struct via the builder pattern


### Return type

[**VolumeTypeDetailResponse**](VolumeTypeDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVolumeType

> VolumeTypeDetailResponse ShowVolumeType(ctx, volumeTypeId).Execute()

Show Volume type’s details



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
	volumeTypeId := "c643ac35-3aec-43e6-b97a-3d6f666c5207" // string | Volume type ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1VolumeTypeApiAPI.ShowVolumeType(context.Background(), volumeTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1VolumeTypeApiAPI.ShowVolumeType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVolumeType`: VolumeTypeDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1VolumeTypeApiAPI.ShowVolumeType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeTypeId** | **string** | Volume type ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVolumeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeTypeDetailResponse**](VolumeTypeDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

