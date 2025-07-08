# AutoScalingGroupNotificationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Count | 
**Notifications** | [**[]AutoScalingGroupNotificationShowResponse**](AutoScalingGroupNotificationShowResponse.md) | Auto-Scaling Group notification list | 

## Methods

### NewAutoScalingGroupNotificationListResponse

`func NewAutoScalingGroupNotificationListResponse(count int32, notifications []AutoScalingGroupNotificationShowResponse, ) *AutoScalingGroupNotificationListResponse`

NewAutoScalingGroupNotificationListResponse instantiates a new AutoScalingGroupNotificationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupNotificationListResponseWithDefaults

`func NewAutoScalingGroupNotificationListResponseWithDefaults() *AutoScalingGroupNotificationListResponse`

NewAutoScalingGroupNotificationListResponseWithDefaults instantiates a new AutoScalingGroupNotificationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AutoScalingGroupNotificationListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AutoScalingGroupNotificationListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AutoScalingGroupNotificationListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNotifications

`func (o *AutoScalingGroupNotificationListResponse) GetNotifications() []AutoScalingGroupNotificationShowResponse`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *AutoScalingGroupNotificationListResponse) GetNotificationsOk() (*[]AutoScalingGroupNotificationShowResponse, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *AutoScalingGroupNotificationListResponse) SetNotifications(v []AutoScalingGroupNotificationShowResponse)`

SetNotifications sets Notifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


