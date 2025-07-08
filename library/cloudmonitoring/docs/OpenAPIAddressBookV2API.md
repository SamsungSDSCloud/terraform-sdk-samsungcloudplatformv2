# \OpenAPIAddressBookV2API

All URIs are relative to *https://cloudmonitoring.kr-west1.dev2.samsungsdscloud.com/monitoring/event*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdressBookMemberList**](OpenAPIAddressBookV2API.md#GetAdressBookMemberList) | **Get** /v1/cloudmonitorings/product/v2/addrbooks/{addrbookId}/members | ListAddressBookMembers



## GetAdressBookMemberList

> PageResponseAlarmAddrBookMemberDto GetAdressBookMemberList(ctx, addrbookId).Page(page).Size(size).Sort(sort).Execute()

ListAddressBookMembers



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
	addrbookId := int32(56) // int32 | Address book ID - Address book IDs can be obtained from @[ListAddressBooks].
	page := int32(0) // int32 | Page Number (default: 0) (optional)
	size := int32(10) // int32 | Page Contents Size (default: 10) (optional)
	sort := []string{"["email:asc"]"} // []string | Sorting Field List (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIAddressBookV2API.GetAdressBookMemberList(context.Background(), addrbookId).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIAddressBookV2API.GetAdressBookMemberList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdressBookMemberList`: PageResponseAlarmAddrBookMemberDto
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIAddressBookV2API.GetAdressBookMemberList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addrbookId** | **int32** | Address book ID - Address book IDs can be obtained from @[ListAddressBooks]. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdressBookMemberListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page Number (default: 0) | 
 **size** | **int32** | Page Contents Size (default: 10) | 
 **sort** | **[]string** | Sorting Field List | 

### Return type

[**PageResponseAlarmAddrBookMemberDto**](PageResponseAlarmAddrBookMemberDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

