# \VirtualserverV1AutoScalingGroupNotificationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoScalingGroupNotification**](VirtualserverV1AutoScalingGroupNotificationsAPI.md#CreateAutoScalingGroupNotification) | **Post** /v1/auto-scaling-groups/{auto_scaling_group_id}/notifications | Create Auto-Scaling Group Notification
[**DeleteAutoScalingGroupNotification**](VirtualserverV1AutoScalingGroupNotificationsAPI.md#DeleteAutoScalingGroupNotification) | **Delete** /v1/auto-scaling-groups/{auto_scaling_group_id}/notifications/{notification_id} | Delete Auto-Scaling Group Notification
[**ListAutoScalingGroupNotifications**](VirtualserverV1AutoScalingGroupNotificationsAPI.md#ListAutoScalingGroupNotifications) | **Get** /v1/auto-scaling-groups/{auto_scaling_group_id}/notifications | List Auto-Scaling Group Notifications
[**ShowAutoScalingGroupNotification**](VirtualserverV1AutoScalingGroupNotificationsAPI.md#ShowAutoScalingGroupNotification) | **Get** /v1/auto-scaling-groups/{auto_scaling_group_id}/notifications/{notification_id} | Show Auto-Scaling Group Notification
[**UpdateAutoScalingGroupNotification**](VirtualserverV1AutoScalingGroupNotificationsAPI.md#UpdateAutoScalingGroupNotification) | **Put** /v1/auto-scaling-groups/{auto_scaling_group_id}/notifications/{notification_id} | Update Auto-Scaling Group Notification



## CreateAutoScalingGroupNotification

> AutoScalingGroupNotificationListResponse CreateAutoScalingGroupNotification(ctx, autoScalingGroupId).AutoScalingGroupNotificationCreateRequest(autoScalingGroupNotificationCreateRequest).Execute()

Create Auto-Scaling Group Notification



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
	autoScalingGroupNotificationCreateRequest := *openapiclient.NewAutoScalingGroupNotificationCreateRequest([]string{"NotificationEvents_example"}, []string{"UserIds_example"}) // AutoScalingGroupNotificationCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupNotificationsAPI.CreateAutoScalingGroupNotification(context.Background(), autoScalingGroupId).AutoScalingGroupNotificationCreateRequest(autoScalingGroupNotificationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupNotificationsAPI.CreateAutoScalingGroupNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutoScalingGroupNotification`: AutoScalingGroupNotificationListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupNotificationsAPI.CreateAutoScalingGroupNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoScalingGroupNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoScalingGroupNotificationCreateRequest** | [**AutoScalingGroupNotificationCreateRequest**](AutoScalingGroupNotificationCreateRequest.md) |  | 

### Return type

[**AutoScalingGroupNotificationListResponse**](AutoScalingGroupNotificationListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoScalingGroupNotification

> DeleteAutoScalingGroupNotification(ctx, autoScalingGroupId, notificationId).Execute()

Delete Auto-Scaling Group Notification



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
	notificationId := "fa7fd191410744cd810a290c0b4d22b9" // string | Notification ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualserverV1AutoScalingGroupNotificationsAPI.DeleteAutoScalingGroupNotification(context.Background(), autoScalingGroupId, notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupNotificationsAPI.DeleteAutoScalingGroupNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 
**notificationId** | **string** | Notification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoScalingGroupNotificationRequest struct via the builder pattern


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


## ListAutoScalingGroupNotifications

> AutoScalingGroupNotificationListResponse ListAutoScalingGroupNotifications(ctx, autoScalingGroupId).UserIds(userIds).NotificationState(notificationState).NotificationEvent(notificationEvent).Offset(offset).Limit(limit).Sort(sort).Execute()

List Auto-Scaling Group Notifications



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
	userIds := []string{"Inner_example"} // []string | User ID list (optional)
	notificationState := "ACTIVE" // string | Auto-Scaling Group notification state (optional)
	notificationEvent := "SCALE_OUT" // string | Auto-Scaling Group notification event (optional)
	offset := int32(0) // int32 | Offset (optional)
	limit := int32(20) // int32 | Limit (optional)
	sort := "created_at:desc" // string | Sort (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupNotificationsAPI.ListAutoScalingGroupNotifications(context.Background(), autoScalingGroupId).UserIds(userIds).NotificationState(notificationState).NotificationEvent(notificationEvent).Offset(offset).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupNotificationsAPI.ListAutoScalingGroupNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAutoScalingGroupNotifications`: AutoScalingGroupNotificationListResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupNotificationsAPI.ListAutoScalingGroupNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAutoScalingGroupNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userIds** | **[]string** | User ID list | 
 **notificationState** | **string** | Auto-Scaling Group notification state | 
 **notificationEvent** | **string** | Auto-Scaling Group notification event | 
 **offset** | **int32** | Offset | 
 **limit** | **int32** | Limit | 
 **sort** | **string** | Sort | 

### Return type

[**AutoScalingGroupNotificationListResponse**](AutoScalingGroupNotificationListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAutoScalingGroupNotification

> AutoScalingGroupNotificationShowResponse ShowAutoScalingGroupNotification(ctx, autoScalingGroupId, notificationId).Execute()

Show Auto-Scaling Group Notification



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
	notificationId := "fa7fd191410744cd810a290c0b4d22b9" // string | Notification ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupNotificationsAPI.ShowAutoScalingGroupNotification(context.Background(), autoScalingGroupId, notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupNotificationsAPI.ShowAutoScalingGroupNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowAutoScalingGroupNotification`: AutoScalingGroupNotificationShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupNotificationsAPI.ShowAutoScalingGroupNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 
**notificationId** | **string** | Notification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowAutoScalingGroupNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AutoScalingGroupNotificationShowResponse**](AutoScalingGroupNotificationShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoScalingGroupNotification

> AutoScalingGroupNotificationShowResponse UpdateAutoScalingGroupNotification(ctx, autoScalingGroupId, notificationId).AutoScalingGroupNotificationUpdateRequest(autoScalingGroupNotificationUpdateRequest).Execute()

Update Auto-Scaling Group Notification



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
	notificationId := "fa7fd191410744cd810a290c0b4d22b9" // string | Notification ID
	autoScalingGroupNotificationUpdateRequest := *openapiclient.NewAutoScalingGroupNotificationUpdateRequest() // AutoScalingGroupNotificationUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualserverV1AutoScalingGroupNotificationsAPI.UpdateAutoScalingGroupNotification(context.Background(), autoScalingGroupId, notificationId).AutoScalingGroupNotificationUpdateRequest(autoScalingGroupNotificationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualserverV1AutoScalingGroupNotificationsAPI.UpdateAutoScalingGroupNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoScalingGroupNotification`: AutoScalingGroupNotificationShowResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualserverV1AutoScalingGroupNotificationsAPI.UpdateAutoScalingGroupNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autoScalingGroupId** | **string** | Auto-Scaling Group ID | 
**notificationId** | **string** | Notification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoScalingGroupNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoScalingGroupNotificationUpdateRequest** | [**AutoScalingGroupNotificationUpdateRequest**](AutoScalingGroupNotificationUpdateRequest.md) |  | 

### Return type

[**AutoScalingGroupNotificationShowResponse**](AutoScalingGroupNotificationShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

