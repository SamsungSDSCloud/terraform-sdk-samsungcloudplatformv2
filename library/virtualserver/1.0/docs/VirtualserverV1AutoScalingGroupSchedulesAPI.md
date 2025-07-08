# \VirtualserverV1AutoScalingGroupSchedulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoScalingGroupSchedule**](VirtualserverV1AutoScalingGroupSchedulesAPI.md#CreateAutoScalingGroupSchedule) | **Post** /v1/auto-scaling-groups/{auto_scaling_group_id}/schedules | Create Auto-Scaling Group Schedule
[**DeleteAutoScalingGroupSchedule**](VirtualserverV1AutoScalingGroupSchedulesAPI.md#DeleteAutoScalingGroupSchedule) | **Delete** /v1/auto-scaling-groups/{auto_scaling_group_id}/schedules/{schedule_id} | Delete Auto-Scaling Group Schedule
[**ListAutoScalingGroupSchedules**](VirtualserverV1AutoScalingGroupSchedulesAPI.md#ListAutoScalingGroupSchedules) | **Get** /v1/auto-scaling-groups/{auto_scaling_group_id}/schedules | List Auto-Scaling Group Schedules
[**ShowAutoScalingGroupSchedule**](VirtualserverV1AutoScalingGroupSchedulesAPI.md#ShowAutoScalingGroupSchedule) | **Get** /v1/auto-scaling-groups/{auto_scaling_group_id}/schedules/{schedule_id} | Show Auto-Scaling Group Schedule
[**UpdateAutoScalingGroupSchedule**](VirtualserverV1AutoScalingGroupSchedulesAPI.md#UpdateAutoScalingGroupSchedule) | **Put** /v1/auto-scaling-groups/{auto_scaling_group_id}/schedules/{schedule_id} | Update Auto-Scaling Group Schedule



## CreateAutoScalingGroupSchedule

> AutoScalingGroupScheduleShowResponse CreateAutoScalingGroupSchedule(ctx, autoScalingGroupId).AutoScalingGroupScheduleCreateRequest(autoScalingGroupScheduleCreateRequest).Execute()

Create Auto-Scaling Group Schedule



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
	autoScalingGroupId := "52613bd852b04b39adcb15a8364d856d" // string | Auto-Scaling Group ID
	autoScalingGroupScheduleCreateRequest := *openapiclient.NewAutoScalingGroupScheduleCreateRequest(openapiclient.JobSchedulingFrequencyEnum("ONCE"), int32(9), int32(0), "schedulename", "2024-01-01", "Asia/Seoul") // AutoScalingGroupScheduleCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupSchedulesAPI.CreateAutoScalingGroupSchedule(context.Background(), autoScalingGroupId).AutoScalingGroupScheduleCreateRequest(autoScalingGroupScheduleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupSchedulesAPI.CreateAutoScalingGroupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoScalingGroupSchedule`: AutoScalingGroupScheduleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupSchedulesAPI.CreateAutoScalingGroupSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoScalingGroupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoScalingGroupScheduleCreateRequest** | [**AutoScalingGroupScheduleCreateRequest**](AutoScalingGroupScheduleCreateRequest.md) |  | 

### Return type

[**AutoScalingGroupScheduleShowResponse**](AutoScalingGroupScheduleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoScalingGroupSchedule

> DeleteAutoScalingGroupSchedule(ctx, autoScalingGroupId, scheduleId).Execute()

Delete Auto-Scaling Group Schedule



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
	autoScalingGroupId := "52613bd852b04b39adcb15a8364d856d" // string | Auto-Scaling Group ID
	scheduleId := "43d4841abd404220af5f0904037aba71" // string | Schedule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1AutoScalingGroupSchedulesAPI.DeleteAutoScalingGroupSchedule(context.Background(), autoScalingGroupId, scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupSchedulesAPI.DeleteAutoScalingGroupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 
**scheduleId** | **string** | Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoScalingGroupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAutoScalingGroupSchedules

> AutoScalingGroupScheduleListResponse ListAutoScalingGroupSchedules(ctx, autoScalingGroupId).Size(size).Page(page).Sort(sort).Frequency(frequency).Execute()

List Auto-Scaling Group Schedules



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
	autoScalingGroupId := "52613bd852b04b39adcb15a8364d856d" // string | Auto-Scaling Group ID
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	frequency := openapiclient.JobSchedulingFrequencyEnum("ONCE") // JobSchedulingFrequencyEnum | Frequency (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupSchedulesAPI.ListAutoScalingGroupSchedules(context.Background(), autoScalingGroupId).Size(size).Page(page).Sort(sort).Frequency(frequency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupSchedulesAPI.ListAutoScalingGroupSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAutoScalingGroupSchedules`: AutoScalingGroupScheduleListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupSchedulesAPI.ListAutoScalingGroupSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAutoScalingGroupSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **frequency** | [**JobSchedulingFrequencyEnum**](JobSchedulingFrequencyEnum.md) | Frequency | 

### Return type

[**AutoScalingGroupScheduleListResponse**](AutoScalingGroupScheduleListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAutoScalingGroupSchedule

> AutoScalingGroupScheduleShowResponse ShowAutoScalingGroupSchedule(ctx, autoScalingGroupId, scheduleId).Execute()

Show Auto-Scaling Group Schedule



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
	autoScalingGroupId := "52613bd852b04b39adcb15a8364d856d" // string | Auto-Scaling Group ID
	scheduleId := "43d4841abd404220af5f0904037aba71" // string | Schedule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupSchedulesAPI.ShowAutoScalingGroupSchedule(context.Background(), autoScalingGroupId, scheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupSchedulesAPI.ShowAutoScalingGroupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowAutoScalingGroupSchedule`: AutoScalingGroupScheduleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupSchedulesAPI.ShowAutoScalingGroupSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 
**scheduleId** | **string** | Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowAutoScalingGroupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AutoScalingGroupScheduleShowResponse**](AutoScalingGroupScheduleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoScalingGroupSchedule

> AutoScalingGroupScheduleShowResponse UpdateAutoScalingGroupSchedule(ctx, autoScalingGroupId, scheduleId).AutoScalingGroupScheduleUpdateRequest(autoScalingGroupScheduleUpdateRequest).Execute()

Update Auto-Scaling Group Schedule



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
	autoScalingGroupId := "52613bd852b04b39adcb15a8364d856d" // string | Auto-Scaling Group ID
	scheduleId := "43d4841abd404220af5f0904037aba71" // string | Schedule ID
	autoScalingGroupScheduleUpdateRequest := *openapiclient.NewAutoScalingGroupScheduleUpdateRequest() // AutoScalingGroupScheduleUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupSchedulesAPI.UpdateAutoScalingGroupSchedule(context.Background(), autoScalingGroupId, scheduleId).AutoScalingGroupScheduleUpdateRequest(autoScalingGroupScheduleUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupSchedulesAPI.UpdateAutoScalingGroupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoScalingGroupSchedule`: AutoScalingGroupScheduleShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupSchedulesAPI.UpdateAutoScalingGroupSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 
**scheduleId** | **string** | Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoScalingGroupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoScalingGroupScheduleUpdateRequest** | [**AutoScalingGroupScheduleUpdateRequest**](AutoScalingGroupScheduleUpdateRequest.md) |  | 

### Return type

[**AutoScalingGroupScheduleShowResponse**](AutoScalingGroupScheduleShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

