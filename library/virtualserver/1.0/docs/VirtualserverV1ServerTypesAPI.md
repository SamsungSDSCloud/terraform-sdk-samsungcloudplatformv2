# \VirtualserverV1ServerTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListServerTypes**](VirtualserverV1ServerTypesAPI.md#ListServerTypes) | **Get** /v1/server-types | List Server Types
[**ShowServerType**](VirtualserverV1ServerTypesAPI.md#ShowServerType) | **Get** /v1/server-types/{server_type_id} | Show Server Type



## ListServerTypes

> ServerTypeListResponse ListServerTypes(ctx).Limit(limit).Marker(marker).MinDisk(minDisk).MinRam(minRam).Sort(sort).Execute()

List Server Types



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
	marker := "s1v10m20" // string | Marker (optional)
	minDisk := int32(100) // int32 | Minimum disk size (optional)
	minRam := int32(10) // int32 | Minimum ram size (optional)
	sort := "vcpus:asc" // string | Sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServerTypesAPI.ListServerTypes(context.Background()).Limit(limit).Marker(marker).MinDisk(minDisk).MinRam(minRam).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerTypesAPI.ListServerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServerTypes`: ServerTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServerTypesAPI.ListServerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit | 
 **marker** | **string** | Marker | 
 **minDisk** | **int32** | Minimum disk size | 
 **minRam** | **int32** | Minimum ram size | 
 **sort** | **string** | Sort | 

### Return type

[**ServerTypeListResponse**](ServerTypeListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowServerType

> ServerType ShowServerType(ctx, serverTypeId).Execute()

Show Server Type



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
	serverTypeId := "s1v10m20" // string | Server type ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1ServerTypesAPI.ShowServerType(context.Background(), serverTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1ServerTypesAPI.ShowServerType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowServerType`: ServerType
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1ServerTypesAPI.ShowServerType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverTypeId** | **string** | Server type ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowServerTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerType**](ServerType.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

