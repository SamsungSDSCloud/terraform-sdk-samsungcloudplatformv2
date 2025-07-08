# \EpasV1EpasMinorPatchApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EpasPatchMinorVersion**](EpasV1EpasMinorPatchApiAPI.md#EpasPatchMinorVersion) | **Put** /v1/clusters/{cluster_id}/patch | Patch Minor Version



## EpasPatchMinorVersion

> AsyncResponse EpasPatchMinorVersion(ctx, clusterId).MinorPatchRequest(minorPatchRequest).Execute()

Patch Minor Version



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
	clusterId := "clusterId_example" // string | DB cluster id
	minorPatchRequest := *openapiclient.NewMinorPatchRequest(false, "SoftwareVersion_example") // MinorPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EpasV1EpasMinorPatchApiAPI.EpasPatchMinorVersion(context.Background(), clusterId).MinorPatchRequest(minorPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpasV1EpasMinorPatchApiAPI.EpasPatchMinorVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EpasPatchMinorVersion`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `EpasV1EpasMinorPatchApiAPI.EpasPatchMinorVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | DB cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEpasPatchMinorVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **minorPatchRequest** | [**MinorPatchRequest**](MinorPatchRequest.md) |  | 

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

