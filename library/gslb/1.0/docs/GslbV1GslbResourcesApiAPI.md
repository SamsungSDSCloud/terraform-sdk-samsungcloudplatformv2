# \GslbV1GslbResourcesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGslbResources**](GslbV1GslbResourcesApiAPI.md#ListGslbResources) | **Get** /v1/gslbs/{gslb_id}/resources | ListGslbResources
[**SetGslbResources**](GslbV1GslbResourcesApiAPI.md#SetGslbResources) | **Put** /v1/gslbs/{gslb_id}/resources | SetGslbResources



## ListGslbResources

> GslbResourceListResponse ListGslbResources(ctx, gslbId).Size(size).Page(page).Sort(sort).Execute()

ListGslbResources



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
	gslbId := "gslbId_example" // string | The GSLB ID.
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GslbV1GslbResourcesApiAPI.ListGslbResources(context.Background(), gslbId).Size(size).Page(page).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GslbV1GslbResourcesApiAPI.ListGslbResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGslbResources`: GslbResourceListResponse
	fmt.Fprintf(os.Stdout, "Response from `GslbV1GslbResourcesApiAPI.ListGslbResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gslbId** | **string** | The GSLB ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGslbResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 

### Return type

[**GslbResourceListResponse**](GslbResourceListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGslbResources

> GslbShowResponse SetGslbResources(ctx, gslbId).GslbResourceSetRequest(gslbResourceSetRequest).Execute()

SetGslbResources



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
	gslbId := "gslbId_example" // string | The GSLB ID.
	gslbResourceSetRequest := *openapiclient.NewGslbResourceSetRequest([]openapiclient.GslbResource{*openapiclient.NewGslbResource("Destination_example", "Region_example")}) // GslbResourceSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GslbV1GslbResourcesApiAPI.SetGslbResources(context.Background(), gslbId).GslbResourceSetRequest(gslbResourceSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GslbV1GslbResourcesApiAPI.SetGslbResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGslbResources`: GslbShowResponse
	fmt.Fprintf(os.Stdout, "Response from `GslbV1GslbResourcesApiAPI.SetGslbResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gslbId** | **string** | The GSLB ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGslbResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gslbResourceSetRequest** | [**GslbResourceSetRequest**](GslbResourceSetRequest.md) |  | 

### Return type

[**GslbShowResponse**](GslbShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

