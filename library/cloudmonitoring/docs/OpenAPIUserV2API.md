# \OpenAPIUserV2API

All URIs are relative to *https://cloudmonitoring.kr-west1.dev2.samsungsdscloud.com/monitoring/event*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdressBookList**](OpenAPIUserV2API.md#GetAdressBookList) | **Get** /v1/cloudmonitorings/product/v2/users/addrbooks | ListAddressBooks



## GetAdressBookList

> PageResponseAlarmAddrBookDto GetAdressBookList(ctx).Page(page).Size(size).Sort(sort).Execute()

ListAddressBooks



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
	page := int32(0) // int32 | Page Number (default: 0) (optional)
	size := int32(10) // int32 | Page Contents Size (default: 10) (optional)
	sort := []string{"["email:asc"]"} // []string | Sorting Field List (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIUserV2API.GetAdressBookList(context.Background()).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIUserV2API.GetAdressBookList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdressBookList`: PageResponseAlarmAddrBookDto
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIUserV2API.GetAdressBookList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdressBookListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page Number (default: 0) | 
 **size** | **int32** | Page Contents Size (default: 10) | 
 **sort** | **[]string** | Sorting Field List | 

### Return type

[**PageResponseAlarmAddrBookDto**](PageResponseAlarmAddrBookDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

