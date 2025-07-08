# \QuotaV1QuotaRequestsAPIsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListQuotaRequests**](QuotaV1QuotaRequestsAPIsAPI.md#ListQuotaRequests) | **Get** /v1/quota-requests | ListQuotaRequests
[**ShowQuotaRequest**](QuotaV1QuotaRequestsAPIsAPI.md#ShowQuotaRequest) | **Get** /v1/quota-requests/{request_id} | ShowQuotaRequest



## ListQuotaRequests

> QuotaRequestListResponse ListQuotaRequests(ctx).Size(size).Page(page).Sort(sort).Service(service).QuotaItem(quotaItem).RequestClass(requestClass).RequestedStartAt(requestedStartAt).RequestedEndAt(requestedEndAt).AppliedStartAt(appliedStartAt).AppliedEndAt(appliedEndAt).State(state).Execute()

ListQuotaRequests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	service := "Virtual Server" // string | Service Name (optional)
	quotaItem := "VIRTUAL_SERVER_DISK.MAX.SIZE" // string | Quota Item Name (optional)
	requestClass := "Account" // string | Request Class (optional)
	requestedStartAt := time.Now() // time.Time | Start date of the request (optional)
	requestedEndAt := time.Now() // time.Time | End date of the request (optional)
	appliedStartAt := time.Now() // time.Time | Start date of the application (optional)
	appliedEndAt := time.Now() // time.Time | End date of the application (optional)
	state := "BEFORE_APPROVE" // string | Request State (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotaV1QuotaRequestsAPIsAPI.ListQuotaRequests(context.Background()).Size(size).Page(page).Sort(sort).Service(service).QuotaItem(quotaItem).RequestClass(requestClass).RequestedStartAt(requestedStartAt).RequestedEndAt(requestedEndAt).AppliedStartAt(appliedStartAt).AppliedEndAt(appliedEndAt).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotaV1QuotaRequestsAPIsAPI.ListQuotaRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQuotaRequests`: QuotaRequestListResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotaV1QuotaRequestsAPIsAPI.ListQuotaRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQuotaRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **service** | **string** | Service Name | 
 **quotaItem** | **string** | Quota Item Name | 
 **requestClass** | **string** | Request Class | 
 **requestedStartAt** | **time.Time** | Start date of the request | 
 **requestedEndAt** | **time.Time** | End date of the request | 
 **appliedStartAt** | **time.Time** | Start date of the application | 
 **appliedEndAt** | **time.Time** | End date of the application | 
 **state** | **string** | Request State | 

### Return type

[**QuotaRequestListResponse**](QuotaRequestListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowQuotaRequest

> QuotaRequestShowResponse ShowQuotaRequest(ctx, requestId).Execute()

ShowQuotaRequest



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
	requestId := "0c2a06f2f06d42248a6bb88b1a78d4cc" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotaV1QuotaRequestsAPIsAPI.ShowQuotaRequest(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotaV1QuotaRequestsAPIsAPI.ShowQuotaRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowQuotaRequest`: QuotaRequestShowResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotaV1QuotaRequestsAPIsAPI.ShowQuotaRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowQuotaRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuotaRequestShowResponse**](QuotaRequestShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

