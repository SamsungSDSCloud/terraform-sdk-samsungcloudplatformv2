# EventPolicyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventPolicyRequest** | [**EventPolicyInfoEditable**](EventPolicyInfoEditable.md) |  | 
**NotificationRecipients** | Pointer to [**[]NotificationRecipient**](NotificationRecipient.md) | 이벤트 알림 수신인 정보 | [optional] 

## Methods

### NewEventPolicyUpdateRequest

`func NewEventPolicyUpdateRequest(eventPolicyRequest EventPolicyInfoEditable, ) *EventPolicyUpdateRequest`

NewEventPolicyUpdateRequest instantiates a new EventPolicyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPolicyUpdateRequestWithDefaults

`func NewEventPolicyUpdateRequestWithDefaults() *EventPolicyUpdateRequest`

NewEventPolicyUpdateRequestWithDefaults instantiates a new EventPolicyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventPolicyRequest

`func (o *EventPolicyUpdateRequest) GetEventPolicyRequest() EventPolicyInfoEditable`

GetEventPolicyRequest returns the EventPolicyRequest field if non-nil, zero value otherwise.

### GetEventPolicyRequestOk

`func (o *EventPolicyUpdateRequest) GetEventPolicyRequestOk() (*EventPolicyInfoEditable, bool)`

GetEventPolicyRequestOk returns a tuple with the EventPolicyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyRequest

`func (o *EventPolicyUpdateRequest) SetEventPolicyRequest(v EventPolicyInfoEditable)`

SetEventPolicyRequest sets EventPolicyRequest field to given value.


### GetNotificationRecipients

`func (o *EventPolicyUpdateRequest) GetNotificationRecipients() []NotificationRecipient`

GetNotificationRecipients returns the NotificationRecipients field if non-nil, zero value otherwise.

### GetNotificationRecipientsOk

`func (o *EventPolicyUpdateRequest) GetNotificationRecipientsOk() (*[]NotificationRecipient, bool)`

GetNotificationRecipientsOk returns a tuple with the NotificationRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationRecipients

`func (o *EventPolicyUpdateRequest) SetNotificationRecipients(v []NotificationRecipient)`

SetNotificationRecipients sets NotificationRecipients field to given value.

### HasNotificationRecipients

`func (o *EventPolicyUpdateRequest) HasNotificationRecipients() bool`

HasNotificationRecipients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


