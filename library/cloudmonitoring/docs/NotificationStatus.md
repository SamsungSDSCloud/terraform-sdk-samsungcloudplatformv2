# NotificationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMsg** | Pointer to **string** | swagger.event.notificationStatus.errorMsg.value | [optional] 
**SendDt** | Pointer to **time.Time** | 전송일시 | [optional] 
**SendMethod** | Pointer to **string** | 전송 메소드 | [optional] 
**SendResult** | Pointer to **string** | 전송 결과 - SUCCESS : 전송 성공 - FAIL : 전송을 시도했으나 전송 실패 - NOT SEND : 전송 메소드 미설정으로 전송 안함 - FILTERED : 알림 필터링 설정으로 전송 안함 - NO PERMISSION : 권한이 없어 전송 안함 | [optional] 

## Methods

### NewNotificationStatus

`func NewNotificationStatus() *NotificationStatus`

NewNotificationStatus instantiates a new NotificationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationStatusWithDefaults

`func NewNotificationStatusWithDefaults() *NotificationStatus`

NewNotificationStatusWithDefaults instantiates a new NotificationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMsg

`func (o *NotificationStatus) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *NotificationStatus) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *NotificationStatus) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *NotificationStatus) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### GetSendDt

`func (o *NotificationStatus) GetSendDt() time.Time`

GetSendDt returns the SendDt field if non-nil, zero value otherwise.

### GetSendDtOk

`func (o *NotificationStatus) GetSendDtOk() (*time.Time, bool)`

GetSendDtOk returns a tuple with the SendDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDt

`func (o *NotificationStatus) SetSendDt(v time.Time)`

SetSendDt sets SendDt field to given value.

### HasSendDt

`func (o *NotificationStatus) HasSendDt() bool`

HasSendDt returns a boolean if a field has been set.

### GetSendMethod

`func (o *NotificationStatus) GetSendMethod() string`

GetSendMethod returns the SendMethod field if non-nil, zero value otherwise.

### GetSendMethodOk

`func (o *NotificationStatus) GetSendMethodOk() (*string, bool)`

GetSendMethodOk returns a tuple with the SendMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendMethod

`func (o *NotificationStatus) SetSendMethod(v string)`

SetSendMethod sets SendMethod field to given value.

### HasSendMethod

`func (o *NotificationStatus) HasSendMethod() bool`

HasSendMethod returns a boolean if a field has been set.

### GetSendResult

`func (o *NotificationStatus) GetSendResult() string`

GetSendResult returns the SendResult field if non-nil, zero value otherwise.

### GetSendResultOk

`func (o *NotificationStatus) GetSendResultOk() (*string, bool)`

GetSendResultOk returns a tuple with the SendResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResult

`func (o *NotificationStatus) SetSendResult(v string)`

SetSendResult sets SendResult field to given value.

### HasSendResult

`func (o *NotificationStatus) HasSendResult() bool`

HasSendResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


