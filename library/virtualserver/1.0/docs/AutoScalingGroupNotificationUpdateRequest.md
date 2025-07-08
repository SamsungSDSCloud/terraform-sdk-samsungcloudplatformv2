# AutoScalingGroupNotificationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationEvents** | Pointer to **[]string** |  | [optional] 
**NotificationState** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAutoScalingGroupNotificationUpdateRequest

`func NewAutoScalingGroupNotificationUpdateRequest() *AutoScalingGroupNotificationUpdateRequest`

NewAutoScalingGroupNotificationUpdateRequest instantiates a new AutoScalingGroupNotificationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupNotificationUpdateRequestWithDefaults

`func NewAutoScalingGroupNotificationUpdateRequestWithDefaults() *AutoScalingGroupNotificationUpdateRequest`

NewAutoScalingGroupNotificationUpdateRequestWithDefaults instantiates a new AutoScalingGroupNotificationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationEvents

`func (o *AutoScalingGroupNotificationUpdateRequest) GetNotificationEvents() []string`

GetNotificationEvents returns the NotificationEvents field if non-nil, zero value otherwise.

### GetNotificationEventsOk

`func (o *AutoScalingGroupNotificationUpdateRequest) GetNotificationEventsOk() (*[]string, bool)`

GetNotificationEventsOk returns a tuple with the NotificationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEvents

`func (o *AutoScalingGroupNotificationUpdateRequest) SetNotificationEvents(v []string)`

SetNotificationEvents sets NotificationEvents field to given value.

### HasNotificationEvents

`func (o *AutoScalingGroupNotificationUpdateRequest) HasNotificationEvents() bool`

HasNotificationEvents returns a boolean if a field has been set.

### SetNotificationEventsNil

`func (o *AutoScalingGroupNotificationUpdateRequest) SetNotificationEventsNil(b bool)`

 SetNotificationEventsNil sets the value for NotificationEvents to be an explicit nil

### UnsetNotificationEvents
`func (o *AutoScalingGroupNotificationUpdateRequest) UnsetNotificationEvents()`

UnsetNotificationEvents ensures that no value is present for NotificationEvents, not even an explicit nil
### GetNotificationState

`func (o *AutoScalingGroupNotificationUpdateRequest) GetNotificationState() string`

GetNotificationState returns the NotificationState field if non-nil, zero value otherwise.

### GetNotificationStateOk

`func (o *AutoScalingGroupNotificationUpdateRequest) GetNotificationStateOk() (*string, bool)`

GetNotificationStateOk returns a tuple with the NotificationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationState

`func (o *AutoScalingGroupNotificationUpdateRequest) SetNotificationState(v string)`

SetNotificationState sets NotificationState field to given value.

### HasNotificationState

`func (o *AutoScalingGroupNotificationUpdateRequest) HasNotificationState() bool`

HasNotificationState returns a boolean if a field has been set.

### SetNotificationStateNil

`func (o *AutoScalingGroupNotificationUpdateRequest) SetNotificationStateNil(b bool)`

 SetNotificationStateNil sets the value for NotificationState to be an explicit nil

### UnsetNotificationState
`func (o *AutoScalingGroupNotificationUpdateRequest) UnsetNotificationState()`

UnsetNotificationState ensures that no value is present for NotificationState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


