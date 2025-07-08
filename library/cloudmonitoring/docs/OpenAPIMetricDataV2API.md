# \OpenAPIMetricDataV2API

All URIs are relative to *https://cloudmonitoring.kr-west1.dev2.samsungsdscloud.com/monitoring/event*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetricPerfDataList**](OpenAPIMetricDataV2API.md#GetMetricPerfDataList) | **Post** /v1/cloudmonitorings/product/v2/metric-data | ListMetricPefData



## GetMetricPerfDataList

> ListResponseMetricStatisticsDataDtoOpenAPIV2 GetMetricPerfDataList(ctx).XResourceType(xResourceType).MetricDataSearchCriteriaOpenAPIV2(metricDataSearchCriteriaOpenAPIV2).Execute()

ListMetricPefData



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
	xResourceType := "xResourceType_example" // string | Resource Type
	metricDataSearchCriteriaOpenAPIV2 := *openapiclient.NewMetricDataSearchCriteriaOpenAPIV2([]openapiclient.MetricDataConditionOpenAPIV2{*openapiclient.NewMetricDataConditionOpenAPIV2("system.diskio.read.bytes", []openapiclient.ProductResourceInfo{*openapiclient.NewProductResourceInfo("INSTANCE-c4Hsd27ttDaLw533X4B6Sp")})}, "2022-08-07T23:59:00.000Z", "2022-08-07T23:50:00.000Z") // MetricDataSearchCriteriaOpenAPIV2 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIMetricDataV2API.GetMetricPerfDataList(context.Background()).XResourceType(xResourceType).MetricDataSearchCriteriaOpenAPIV2(metricDataSearchCriteriaOpenAPIV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIMetricDataV2API.GetMetricPerfDataList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricPerfDataList`: ListResponseMetricStatisticsDataDtoOpenAPIV2
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIMetricDataV2API.GetMetricPerfDataList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricPerfDataListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xResourceType** | **string** | Resource Type | 
 **metricDataSearchCriteriaOpenAPIV2** | [**MetricDataSearchCriteriaOpenAPIV2**](MetricDataSearchCriteriaOpenAPIV2.md) |  | 

### Return type

[**ListResponseMetricStatisticsDataDtoOpenAPIV2**](ListResponseMetricStatisticsDataDtoOpenAPIV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

