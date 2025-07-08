# \OpenAPIEventV2API

All URIs are relative to *https://cloudmonitoring.kr-west1.dev2.samsungsdscloud.com/monitoring/event*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountEventList**](OpenAPIEventV2API.md#GetAccountEventList) | **Get** /v1/cloudmonitorings/event/v2/accounts/events | ListAccountEvents
[**GetEventDetail**](OpenAPIEventV2API.md#GetEventDetail) | **Get** /v1/cloudmonitorings/event/v2/events/{eventId} | ShowResourceEvent
[**GetEventNotificationStates**](OpenAPIEventV2API.md#GetEventNotificationStates) | **Get** /v1/cloudmonitorings/event/v2/events/{eventId}/notification-states | ShowEventNotificationStates
[**GetProductEventList**](OpenAPIEventV2API.md#GetProductEventList) | **Get** /v1/cloudmonitorings/event/v2/events | ListResourceEvents



## GetAccountEventList

> PageResponseOpenApiEventResponse GetAccountEventList(ctx).EventState(eventState).QueryStartDt(queryStartDt).QueryEndDt(queryEndDt).Page(page).Size(size).Sort(sort).Execute()

ListAccountEvents



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
	eventState := "ALL" // string | Event State
	queryStartDt := "2021-10-23T18:00:00.000Z" // string | Query Start Date
	queryEndDt := "2021-10-23T23:59:59.000Z" // string | Query End Date
	page := int32(0) // int32 | Page Number (default: 0) (optional)
	size := int32(10) // int32 | Page Contents Size (default: 10) (optional)
	sort := []string{"["email:asc"]"} // []string | Sorting Field List (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIEventV2API.GetAccountEventList(context.Background()).EventState(eventState).QueryStartDt(queryStartDt).QueryEndDt(queryEndDt).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIEventV2API.GetAccountEventList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountEventList`: PageResponseOpenApiEventResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIEventV2API.GetAccountEventList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountEventListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventState** | **string** | Event State | 
 **queryStartDt** | **string** | Query Start Date | 
 **queryEndDt** | **string** | Query End Date | 
 **page** | **int32** | Page Number (default: 0) | 
 **size** | **int32** | Page Contents Size (default: 10) | 
 **sort** | **[]string** | Sorting Field List | 

### Return type

[**PageResponseOpenApiEventResponse**](PageResponseOpenApiEventResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventDetail

> EventDetailResponse GetEventDetail(ctx, eventId).XResourceType(xResourceType).Execute()

ShowResourceEvent



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
	eventId := "eventId_example" // string | Event ID - Event ID can be viewed using @[ListResourceEvents].
	xResourceType := "xResourceType_example" // string | Resource Type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIEventV2API.GetEventDetail(context.Background(), eventId).XResourceType(xResourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIEventV2API.GetEventDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventDetail`: EventDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIEventV2API.GetEventDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Event ID - Event ID can be viewed using @[ListResourceEvents]. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xResourceType** | **string** | Resource Type | 

### Return type

[**EventDetailResponse**](EventDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventNotificationStates

> PageResponseEventNotificationResponse GetEventNotificationStates(ctx, eventId).XResourceType(xResourceType).Page(page).Size(size).Sort(sort).Execute()

ShowEventNotificationStates



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
	eventId := "eventId_example" // string | Event ID - Event ID can be viewed using @[ListResourceEvents].
	xResourceType := "xResourceType_example" // string | Resource Type
	page := int32(0) // int32 | Page Number (default: 0) (optional)
	size := int32(10) // int32 | Page Contents Size (default: 10) (optional)
	sort := []string{"["email:asc"]"} // []string | Sorting Field List (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIEventV2API.GetEventNotificationStates(context.Background(), eventId).XResourceType(xResourceType).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIEventV2API.GetEventNotificationStates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventNotificationStates`: PageResponseEventNotificationResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIEventV2API.GetEventNotificationStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | Event ID - Event ID can be viewed using @[ListResourceEvents]. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventNotificationStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xResourceType** | **string** | Resource Type | 
 **page** | **int32** | Page Number (default: 0) | 
 **size** | **int32** | Page Contents Size (default: 10) | 
 **sort** | **[]string** | Sorting Field List | 

### Return type

[**PageResponseEventNotificationResponse**](PageResponseEventNotificationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductEventList

> PageResponseOpenApiEventResponse GetProductEventList(ctx).ProductResourceId(productResourceId).EventState(eventState).QueryStartDt(queryStartDt).QueryEndDt(queryEndDt).XResourceType(xResourceType).MetricKey(metricKey).Page(page).Size(size).Sort(sort).Execute()

ListResourceEvents



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
	productResourceId := "INSTANCE-c4Hsd27ttDaLw533X4B6Sp" // string | Product Resource ID - Product Resource ID can be viewed using @[ListAccountResources].
	eventState := "ALL" // string | Event State
	queryStartDt := "2021-10-23T18:00:00.000Z" // string | Query Start Date
	queryEndDt := "2021-10-23T23:59:59.000Z" // string | Query End Date
	xResourceType := "xResourceType_example" // string | Resource Type
	metricKey := "system.cpu.system.pct" // string | Metric Key - Metric Key can be viewed using @[ListMetrics]. (optional)
	page := int32(0) // int32 | Page Number (default: 0) (optional)
	size := int32(10) // int32 | Page Contents Size (default: 10) (optional)
	sort := []string{"["email:asc"]"} // []string | Sorting Field List (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIEventV2API.GetProductEventList(context.Background()).ProductResourceId(productResourceId).EventState(eventState).QueryStartDt(queryStartDt).QueryEndDt(queryEndDt).XResourceType(xResourceType).MetricKey(metricKey).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIEventV2API.GetProductEventList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductEventList`: PageResponseOpenApiEventResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIEventV2API.GetProductEventList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductEventListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productResourceId** | **string** | Product Resource ID - Product Resource ID can be viewed using @[ListAccountResources]. | 
 **eventState** | **string** | Event State | 
 **queryStartDt** | **string** | Query Start Date | 
 **queryEndDt** | **string** | Query End Date | 
 **xResourceType** | **string** | Resource Type | 
 **metricKey** | **string** | Metric Key - Metric Key can be viewed using @[ListMetrics]. | 
 **page** | **int32** | Page Number (default: 0) | 
 **size** | **int32** | Page Contents Size (default: 10) | 
 **sort** | **[]string** | Sorting Field List | 

### Return type

[**PageResponseOpenApiEventResponse**](PageResponseOpenApiEventResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

