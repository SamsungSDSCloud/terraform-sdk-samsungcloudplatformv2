# \SqlserverV1SqlserverMinorPatchApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SqlserverPatchMinorVersion**](SqlserverV1SqlserverMinorPatchApiAPI.md#SqlserverPatchMinorVersion) | **Put** /v1/clusters/{cluster_id}/patch | Patch Minor Version



## SqlserverPatchMinorVersion

> AsyncResponse SqlserverPatchMinorVersion(ctx, clusterId).MinorPatchRequest(minorPatchRequest).Execute()

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
	clusterId := "109a585ae67b4e8482fdafc8a4a5be74" // string | Cluster ID
	minorPatchRequest := *openapiclient.NewMinorPatchRequest(false, "SoftwareVersion_example") // MinorPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SqlserverV1SqlserverMinorPatchApiAPI.SqlserverPatchMinorVersion(context.Background(), clusterId).MinorPatchRequest(minorPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SqlserverV1SqlserverMinorPatchApiAPI.SqlserverPatchMinorVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SqlserverPatchMinorVersion`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SqlserverV1SqlserverMinorPatchApiAPI.SqlserverPatchMinorVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSqlserverPatchMinorVersionRequest struct via the builder pattern


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

