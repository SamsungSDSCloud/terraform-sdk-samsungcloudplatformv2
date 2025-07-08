# \OpenAPIMetricV2API

All URIs are relative to *https://cloudmonitoring.kr-west1.dev2.samsungsdscloud.com/monitoring/event*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetricList**](OpenAPIMetricV2API.md#GetMetricList) | **Get** /v1/cloudmonitorings/product/v2/metrics | ListMetrics



## GetMetricList

> PageResponseMetricInfoDto GetMetricList(ctx).ProductTypeCode(productTypeCode).ObjectType(objectType).Page(page).Size(size).Sort(sort).Execute()

ListMetrics



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
	productTypeCode := "VM" // string | Product type code - Product type codes can be obtained using @[ListServiceMonitoring]. If not specified, the entire metric list will be retrieved. (optional)
	objectType := "node" // string | Object Type - Only available for services that have subtypes, such as when the 'productTypeCode' is 'Kubernetes' or 'Load Balancer'. (optional)
	page := int32(0) // int32 | Page Number (default: 0) (optional)
	size := int32(10) // int32 | Page Contents Size (default: 10) (optional)
	sort := []string{"["email:asc"]"} // []string | Sorting Field List (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIMetricV2API.GetMetricList(context.Background()).ProductTypeCode(productTypeCode).ObjectType(objectType).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIMetricV2API.GetMetricList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricList`: PageResponseMetricInfoDto
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIMetricV2API.GetMetricList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productTypeCode** | **string** | Product type code - Product type codes can be obtained using @[ListServiceMonitoring]. If not specified, the entire metric list will be retrieved. | 
 **objectType** | **string** | Object Type - Only available for services that have subtypes, such as when the &#39;productTypeCode&#39; is &#39;Kubernetes&#39; or &#39;Load Balancer&#39;. | 
 **page** | **int32** | Page Number (default: 0) | 
 **size** | **int32** | Page Contents Size (default: 10) | 
 **sort** | **[]string** | Sorting Field List | 

### Return type

[**PageResponseMetricInfoDto**](PageResponseMetricInfoDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

