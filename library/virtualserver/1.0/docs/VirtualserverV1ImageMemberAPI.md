# \VirtualserverV1ImageMemberAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImageMember**](VirtualserverV1ImageMemberAPI.md#CreateImageMember) | **Post** /v1/images/{image_id}/members | Create Image member
[**DeleteImageMember**](VirtualserverV1ImageMemberAPI.md#DeleteImageMember) | **Delete** /v1/images/{image_id}/members/{member_id} | Delete Image member
[**ListImageMembers**](VirtualserverV1ImageMemberAPI.md#ListImageMembers) | **Get** /v1/images/{image_id}/members | List Image members
[**ShowImageMember**](VirtualserverV1ImageMemberAPI.md#ShowImageMember) | **Get** /v1/images/{image_id}/members/{member_id} | Show Image member
[**UpdateImageMember**](VirtualserverV1ImageMemberAPI.md#UpdateImageMember) | **Put** /v1/images/{image_id}/members/{member_id} | Update Image member



## CreateImageMember

> ImageMemberShowResponse CreateImageMember(ctx, imageId).ImageMemberCreateRequest(imageMemberCreateRequest).Execute()

Create Image member



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
	imageMemberCreateRequest := *openapiclient.NewImageMemberCreateRequest("8989447062e04a818baf9e073fd04fa7") // ImageMemberCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ImageMemberAPI.CreateImageMember(context.Background(), imageId).ImageMemberCreateRequest(imageMemberCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ImageMemberAPI.CreateImageMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateImageMember`: ImageMemberShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ImageMemberAPI.CreateImageMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imageMemberCreateRequest** | [**ImageMemberCreateRequest**](ImageMemberCreateRequest.md) |  | 

### Return type

[**ImageMemberShowResponse**](ImageMemberShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImageMember

> DeleteImageMember(ctx, imageId, memberId).Execute()

Delete Image member



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
	memberId := "8989447062e04a818baf9e073fd04fa7" // string | Member ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ImageMemberAPI.DeleteImageMember(context.Background(), imageId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ImageMemberAPI.DeleteImageMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID | 
**memberId** | **string** | Member ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageMemberRequest struct via the builder pattern


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


## ListImageMembers

> ImageMemberListResponse ListImageMembers(ctx, imageId).Execute()

List Image members



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
	resp, r, err := apiClient.VirtualserverV1ImageMemberAPI.ListImageMembers(context.Background(), imageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ImageMemberAPI.ListImageMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImageMembers`: ImageMemberListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ImageMemberAPI.ListImageMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImageMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImageMemberListResponse**](ImageMemberListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowImageMember

> ImageMemberShowResponse ShowImageMember(ctx, imageId, memberId).Execute()

Show Image member



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
	memberId := "8989447062e04a818baf9e073fd04fa7" // string | Member ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ImageMemberAPI.ShowImageMember(context.Background(), imageId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ImageMemberAPI.ShowImageMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowImageMember`: ImageMemberShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ImageMemberAPI.ShowImageMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID | 
**memberId** | **string** | Member ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowImageMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ImageMemberShowResponse**](ImageMemberShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateImageMember

> ImageMemberShowResponse UpdateImageMember(ctx, imageId, memberId).ImageMemberSetRequest(imageMemberSetRequest).Execute()

Update Image member



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
	memberId := "8989447062e04a818baf9e073fd04fa7" // string | Member ID
	imageMemberSetRequest := *openapiclient.NewImageMemberSetRequest(openapiclient.ImageMemberStatusEnum("pending")) // ImageMemberSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ImageMemberAPI.UpdateImageMember(context.Background(), imageId, memberId).ImageMemberSetRequest(imageMemberSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ImageMemberAPI.UpdateImageMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateImageMember`: ImageMemberShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ImageMemberAPI.UpdateImageMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | Image ID | 
**memberId** | **string** | Member ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **imageMemberSetRequest** | [**ImageMemberSetRequest**](ImageMemberSetRequest.md) |  | 

### Return type

[**ImageMemberShowResponse**](ImageMemberShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

