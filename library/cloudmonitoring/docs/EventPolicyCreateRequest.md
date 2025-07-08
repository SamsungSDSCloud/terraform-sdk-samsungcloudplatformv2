# EventPolicyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventPolicyRequest** | [**EventPolicyInfo**](EventPolicyInfo.md) |  | 
**NotificationRecipients** | Pointer to [**[]NotificationRecipient**](NotificationRecipient.md) | 이벤트 알림 수신인 정보 - 알림 대상자 정보는 10개까지 지정할 수 있습니다. | [optional] 
**ProductResourceId** | **string** | Product Resource ID - Product Resource ID can be viewed using @[ListAccountResources]. | 

## Methods

### NewEventPolicyCreateRequest

`func NewEventPolicyCreateRequest(eventPolicyRequest EventPolicyInfo, productResourceId string, ) *EventPolicyCreateRequest`

NewEventPolicyCreateRequest instantiates a new EventPolicyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPolicyCreateRequestWithDefaults

`func NewEventPolicyCreateRequestWithDefaults() *EventPolicyCreateRequest`

NewEventPolicyCreateRequestWithDefaults instantiates a new EventPolicyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventPolicyRequest

`func (o *EventPolicyCreateRequest) GetEventPolicyRequest() EventPolicyInfo`

GetEventPolicyRequest returns the EventPolicyRequest field if non-nil, zero value otherwise.

### GetEventPolicyRequestOk

`func (o *EventPolicyCreateRequest) GetEventPolicyRequestOk() (*EventPolicyInfo, bool)`

GetEventPolicyRequestOk returns a tuple with the EventPolicyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyRequest

`func (o *EventPolicyCreateRequest) SetEventPolicyRequest(v EventPolicyInfo)`

SetEventPolicyRequest sets EventPolicyRequest field to given value.


### GetNotificationRecipients

`func (o *EventPolicyCreateRequest) GetNotificationRecipients() []NotificationRecipient`

GetNotificationRecipients returns the NotificationRecipients field if non-nil, zero value otherwise.

### GetNotificationRecipientsOk

`func (o *EventPolicyCreateRequest) GetNotificationRecipientsOk() (*[]NotificationRecipient, bool)`

GetNotificationRecipientsOk returns a tuple with the NotificationRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationRecipients

`func (o *EventPolicyCreateRequest) SetNotificationRecipients(v []NotificationRecipient)`

SetNotificationRecipients sets NotificationRecipients field to given value.

### HasNotificationRecipients

`func (o *EventPolicyCreateRequest) HasNotificationRecipients() bool`

HasNotificationRecipients returns a boolean if a field has been set.

### GetProductResourceId

`func (o *EventPolicyCreateRequest) GetProductResourceId() string`

GetProductResourceId returns the ProductResourceId field if non-nil, zero value otherwise.

### GetProductResourceIdOk

`func (o *EventPolicyCreateRequest) GetProductResourceIdOk() (*string, bool)`

GetProductResourceIdOk returns a tuple with the ProductResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductResourceId

`func (o *EventPolicyCreateRequest) SetProductResourceId(v string)`

SetProductResourceId sets ProductResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


