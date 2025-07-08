# \VirtualserverV1ImageImportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportImage**](VirtualserverV1ImageImportAPI.md#ImportImage) | **Post** /v1/images/{image_id}/import | Import Image



## ImportImage

> ImportImage(ctx, imageId).ImageImportRequest(imageImportRequest).Execute()

Import Image



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
	imageImportRequest := *openapiclient.NewImageImportRequest("https://object-store.kr-west1.s.samsungsdscloud.com/8989447062e04a818baf9e073fd04fa7/bucket/object.qcow2") // ImageImportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1ImageImportAPI.ImportImage(context.Background(), imageId).ImageImportRequest(imageImportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ImageImportAPI.ImportImage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiImportImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imageImportRequest** | [**ImageImportRequest**](ImageImportRequest.md) |  | 

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

