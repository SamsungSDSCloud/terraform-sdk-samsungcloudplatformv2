# \OpenAPIProductTypeV1API

All URIs are relative to *https://cloudmonitoring.kr-west1.dev2.samsungsdscloud.com/monitoring/event*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductTypeList**](OpenAPIProductTypeV1API.md#GetProductTypeList) | **Get** /v1/cloudmonitorings/product/v1/product-types | ListService



## GetProductTypeList

> PageResponseProductTypeInfoDto GetProductTypeList(ctx).ProductCategoryCode(productCategoryCode).Page(page).Size(size).Sort(sort).Execute()

ListService



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
	productCategoryCode := "Compute" // string | Product type category - If not specified, the entire product type code will be retrieved. (optional)
	page := int32(0) // int32 | Page Number (default: 0) (optional)
	size := int32(10) // int32 | Page Contents Size (default: 10) (optional)
	sort := []string{"["email:asc"]"} // []string | Sorting Field List (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIProductTypeV1API.GetProductTypeList(context.Background()).ProductCategoryCode(productCategoryCode).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIProductTypeV1API.GetProductTypeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductTypeList`: PageResponseProductTypeInfoDto
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIProductTypeV1API.GetProductTypeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductTypeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCategoryCode** | **string** | Product type category - If not specified, the entire product type code will be retrieved. | 
 **page** | **int32** | Page Number (default: 0) | 
 **size** | **int32** | Page Contents Size (default: 10) | 
 **sort** | **[]string** | Sorting Field List | 

### Return type

[**PageResponseProductTypeInfoDto**](PageResponseProductTypeInfoDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

