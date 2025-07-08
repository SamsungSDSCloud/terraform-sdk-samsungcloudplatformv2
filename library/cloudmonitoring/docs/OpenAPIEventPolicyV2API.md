# \OpenAPIEventPolicyV2API

All URIs are relative to *https://cloudmonitoring.kr-west1.dev2.samsungsdscloud.com/monitoring/event*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEventPolicy**](OpenAPIEventPolicyV2API.md#DeleteEventPolicy) | **Delete** /v1/cloudmonitorings/event/v2/event-policies/{eventPolicyId} | DeleteEventPolicy
[**GetEventPolicyDetail**](OpenAPIEventPolicyV2API.md#GetEventPolicyDetail) | **Get** /v1/cloudmonitorings/event/v2/event-policies/{eventPolicyId} | ShowEventPolicy
[**GetEventPolicyHistories**](OpenAPIEventPolicyV2API.md#GetEventPolicyHistories) | **Get** /v1/cloudmonitorings/event/v2/event-policies/{eventPolicyId}/histories | ListEventPolicyHistories
[**GetEventPolicyNotification**](OpenAPIEventPolicyV2API.md#GetEventPolicyNotification) | **Get** /v1/cloudmonitorings/event/v2/event-policies/{eventPolicyId}/notifications | ListEventPolicyNotification
[**GetProductEventPolicyList**](OpenAPIEventPolicyV2API.md#GetProductEventPolicyList) | **Get** /v1/cloudmonitorings/event/v2/event-policies | ListEventPolicies
[**ModifyEventPolicy**](OpenAPIEventPolicyV2API.md#ModifyEventPolicy) | **Put** /v1/cloudmonitorings/event/v2/event-policies/{eventPolicyId} | SetEventPolicy
[**PutEventPolicy**](OpenAPIEventPolicyV2API.md#PutEventPolicy) | **Post** /v1/cloudmonitorings/event/v2/event-policies | CreateEventPolicy



## DeleteEventPolicy

> DeleteEventPolicy(ctx, eventPolicyId).XResourceType(xResourceType).Execute()

DeleteEventPolicy



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
	eventPolicyId := int64(7226) // int64 | Event policy ID to delete - Event Policy ID can be viewed using @[ListEventPolicies].
	xResourceType := "xResourceType_example" // string | Resource Type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OpenAPIEventPolicyV2API.DeleteEventPolicy(context.Background(), eventPolicyId).XResourceType(xResourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIEventPolicyV2API.DeleteEventPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventPolicyId** | **int64** | Event policy ID to delete - Event Policy ID can be viewed using @[ListEventPolicies]. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xResourceType** | **string** | Resource Type | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventPolicyDetail

> EventPolicyDetailResponse GetEventPolicyDetail(ctx, eventPolicyId).XResourceType(xResourceType).Execute()

ShowEventPolicy



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
	eventPolicyId := int64(7226) // int64 | Event Policy ID to inquire - Event Policy ID can be viewed using @[ListEventPolicies].
	xResourceType := "xResourceType_example" // string | Resource Type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIEventPolicyV2API.GetEventPolicyDetail(context.Background(), eventPolicyId).XResourceType(xResourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIEventPolicyV2API.GetEventPolicyDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventPolicyDetail`: EventPolicyDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIEventPolicyV2API.GetEventPolicyDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventPolicyId** | **int64** | Event Policy ID to inquire - Event Policy ID can be viewed using @[ListEventPolicies]. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventPolicyDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xResourceType** | **string** | Resource Type | 

### Return type

[**EventPolicyDetailResponse**](EventPolicyDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventPolicyHistories

> PageResponseEventPolicyHistoryResponse GetEventPolicyHistories(ctx, eventPolicyId).QueryStartDt(queryStartDt).QueryEndDt(queryEndDt).XResourceType(xResourceType).Page(page).Size(size).Sort(sort).Execute()

ListEventPolicyHistories



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
	eventPolicyId := int64(7226) // int64 | Event Policy ID to inquire - Event Policy ID can be viewed using @[ListEventPolicies].
	queryStartDt := "2021-10-23T18:00:00.000Z" // string | Query Start Date
	queryEndDt := "2021-10-23T23:59:59.000Z" // string | Query End Date
	xResourceType := "xResourceType_example" // string | Resource Type
	page := int32(0) // int32 | Page Number (default: 0) (optional)
	size := int32(10) // int32 | Page Contents Size (default: 10) (optional)
	sort := []string{"["email:asc"]"} // []string | Sorting Field List (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIEventPolicyV2API.GetEventPolicyHistories(context.Background(), eventPolicyId).QueryStartDt(queryStartDt).QueryEndDt(queryEndDt).XResourceType(xResourceType).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIEventPolicyV2API.GetEventPolicyHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventPolicyHistories`: PageResponseEventPolicyHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIEventPolicyV2API.GetEventPolicyHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventPolicyId** | **int64** | Event Policy ID to inquire - Event Policy ID can be viewed using @[ListEventPolicies]. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventPolicyHistoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryStartDt** | **string** | Query Start Date | 
 **queryEndDt** | **string** | Query End Date | 
 **xResourceType** | **string** | Resource Type | 
 **page** | **int32** | Page Number (default: 0) | 
 **size** | **int32** | Page Contents Size (default: 10) | 
 **sort** | **[]string** | Sorting Field List | 

### Return type

[**PageResponseEventPolicyHistoryResponse**](PageResponseEventPolicyHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventPolicyNotification

> PageResponseNotificationResponse GetEventPolicyNotification(ctx, eventPolicyId).XResourceType(xResourceType).Page(page).Size(size).Sort(sort).Execute()

ListEventPolicyNotification



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
	eventPolicyId := int64(7226) // int64 | Event Policy ID to inquire - Event Policy ID can be viewed using @[ListEventPolicies].
	xResourceType := "xResourceType_example" // string | Resource Type
	page := int32(0) // int32 | Page Number (default: 0) (optional)
	size := int32(10) // int32 | Page Contents Size (default: 10) (optional)
	sort := []string{"["email:asc"]"} // []string | Sorting Field List (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIEventPolicyV2API.GetEventPolicyNotification(context.Background(), eventPolicyId).XResourceType(xResourceType).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIEventPolicyV2API.GetEventPolicyNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventPolicyNotification`: PageResponseNotificationResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIEventPolicyV2API.GetEventPolicyNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventPolicyId** | **int64** | Event Policy ID to inquire - Event Policy ID can be viewed using @[ListEventPolicies]. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventPolicyNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xResourceType** | **string** | Resource Type | 
 **page** | **int32** | Page Number (default: 0) | 
 **size** | **int32** | Page Contents Size (default: 10) | 
 **sort** | **[]string** | Sorting Field List | 

### Return type

[**PageResponseNotificationResponse**](PageResponseNotificationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductEventPolicyList

> PageResponseEventPolicyResponse GetProductEventPolicyList(ctx).ProductResourceId(productResourceId).XResourceType(xResourceType).MetricKey(metricKey).Page(page).Size(size).Sort(sort).Execute()

ListEventPolicies



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
	productResourceId := "c94607c8-6a15-4e13-84f3-99cb77d478e3" // string | Product Resource ID - Product Resource ID can be viewed using @[ListAccountResources].
	xResourceType := "xResourceType_example" // string | Resource Type
	metricKey := "system.cpu.system.pct" // string | Metric Key - Metric Key can be viewed using @[ListMetrics]. (optional)
	page := int32(0) // int32 | Page Number (default: 0) (optional)
	size := int32(10) // int32 | Page Contents Size (default: 10) (optional)
	sort := []string{"["email:asc"]"} // []string | Sorting Field List (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIEventPolicyV2API.GetProductEventPolicyList(context.Background()).ProductResourceId(productResourceId).XResourceType(xResourceType).MetricKey(metricKey).Page(page).Size(size).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIEventPolicyV2API.GetProductEventPolicyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductEventPolicyList`: PageResponseEventPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIEventPolicyV2API.GetProductEventPolicyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductEventPolicyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productResourceId** | **string** | Product Resource ID - Product Resource ID can be viewed using @[ListAccountResources]. | 
 **xResourceType** | **string** | Resource Type | 
 **metricKey** | **string** | Metric Key - Metric Key can be viewed using @[ListMetrics]. | 
 **page** | **int32** | Page Number (default: 0) | 
 **size** | **int32** | Page Contents Size (default: 10) | 
 **sort** | **[]string** | Sorting Field List | 

### Return type

[**PageResponseEventPolicyResponse**](PageResponseEventPolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyEventPolicy

> EventPolicyDetailResponse ModifyEventPolicy(ctx, eventPolicyId).XResourceType(xResourceType).EventPolicyUpdateRequest(eventPolicyUpdateRequest).Execute()

SetEventPolicy



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
	eventPolicyId := int64(7226) // int64 | Event policy ID to modify - Event Policy ID can be viewed using @[ListEventPolicies].
	xResourceType := "xResourceType_example" // string | Resource Type
	eventPolicyUpdateRequest := *openapiclient.NewEventPolicyUpdateRequest(*openapiclient.NewEventPolicyInfoEditable("N", "WARNING", *openapiclient.NewEventThreshold("RANGE_VALUE"), int64(3))) // EventPolicyUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIEventPolicyV2API.ModifyEventPolicy(context.Background(), eventPolicyId).XResourceType(xResourceType).EventPolicyUpdateRequest(eventPolicyUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIEventPolicyV2API.ModifyEventPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyEventPolicy`: EventPolicyDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIEventPolicyV2API.ModifyEventPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventPolicyId** | **int64** | Event policy ID to modify - Event Policy ID can be viewed using @[ListEventPolicies]. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyEventPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xResourceType** | **string** | Resource Type | 
 **eventPolicyUpdateRequest** | [**EventPolicyUpdateRequest**](EventPolicyUpdateRequest.md) |  | 

### Return type

[**EventPolicyDetailResponse**](EventPolicyDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutEventPolicy

> EventPolicyDetailResponse PutEventPolicy(ctx).XResourceType(xResourceType).EventPolicyCreateRequest(eventPolicyCreateRequest).Execute()

CreateEventPolicy



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
	eventPolicyCreateRequest := *openapiclient.NewEventPolicyCreateRequest(*openapiclient.NewEventPolicyInfo("N", "WARNING", *openapiclient.NewEventThreshold("RANGE_VALUE"), int64(3), "N", "system.diskio.read.bytes"), "INSTANCE-c4Hsd27ttDaLw533X4B6Sp") // EventPolicyCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenAPIEventPolicyV2API.PutEventPolicy(context.Background()).XResourceType(xResourceType).EventPolicyCreateRequest(eventPolicyCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIEventPolicyV2API.PutEventPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutEventPolicy`: EventPolicyDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIEventPolicyV2API.PutEventPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutEventPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xResourceType** | **string** | Resource Type | 
 **eventPolicyCreateRequest** | [**EventPolicyCreateRequest**](EventPolicyCreateRequest.md) |  | 

### Return type

[**EventPolicyDetailResponse**](EventPolicyDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

