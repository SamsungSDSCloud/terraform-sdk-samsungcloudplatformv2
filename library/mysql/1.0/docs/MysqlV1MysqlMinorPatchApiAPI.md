# \MysqlV1MysqlMinorPatchApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlPatchMinorVersion**](MysqlV1MysqlMinorPatchApiAPI.md#MysqlPatchMinorVersion) | **Put** /v1/clusters/{cluster_id}/patch | Patch Minor Version



## MysqlPatchMinorVersion

> AsyncResponse MysqlPatchMinorVersion(ctx, clusterId).MinorPatchRequest(minorPatchRequest).Execute()

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
	resp, r, err := apiClient.MysqlV1MysqlMinorPatchApiAPI.MysqlPatchMinorVersion(context.Background(), clusterId).MinorPatchRequest(minorPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MysqlV1MysqlMinorPatchApiAPI.MysqlPatchMinorVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MysqlPatchMinorVersion`: AsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MysqlV1MysqlMinorPatchApiAPI.MysqlPatchMinorVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | DB cluster id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlPatchMinorVersionRequest struct via the builder pattern


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

