# NotificationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventLevel** | **string** | 이벤트 레벨 | 
**SendMethod** | **[]string** | 전송 메소드 - 허용 값: MESSENGER, SMS, MAIL | 

## Methods

### NewNotificationMethod

`func NewNotificationMethod(eventLevel string, sendMethod []string, ) *NotificationMethod`

NewNotificationMethod instantiates a new NotificationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationMethodWithDefaults

`func NewNotificationMethodWithDefaults() *NotificationMethod`

NewNotificationMethodWithDefaults instantiates a new NotificationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventLevel

`func (o *NotificationMethod) GetEventLevel() string`

GetEventLevel returns the EventLevel field if non-nil, zero value otherwise.

### GetEventLevelOk

`func (o *NotificationMethod) GetEventLevelOk() (*string, bool)`

GetEventLevelOk returns a tuple with the EventLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLevel

`func (o *NotificationMethod) SetEventLevel(v string)`

SetEventLevel sets EventLevel field to given value.


### GetSendMethod

`func (o *NotificationMethod) GetSendMethod() []string`

GetSendMethod returns the SendMethod field if non-nil, zero value otherwise.

### GetSendMethodOk

`func (o *NotificationMethod) GetSendMethodOk() (*[]string, bool)`

GetSendMethodOk returns a tuple with the SendMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendMethod

`func (o *NotificationMethod) SetSendMethod(v []string)`

SetSendMethod sets SendMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


