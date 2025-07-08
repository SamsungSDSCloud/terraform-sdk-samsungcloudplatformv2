# EventNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationStates** | [**[]NotificationStatus**](NotificationStatus.md) | 이벤트 알림 상태 | 
**User** | Pointer to **map[string]interface{}** |  | [optional] 
**UserEmail** | Pointer to **string** | 사용자 이메일 | [optional] 
**UserId** | **string** | 사용자 아이디 | 
**UserMobileTelNo** | Pointer to **string** | 사용자 휴대 전화번호 | [optional] 
**UserNameNotification** | **string** | 사용자 이름 | 

## Methods

### NewEventNotificationResponse

`func NewEventNotificationResponse(notificationStates []NotificationStatus, userId string, userNameNotification string, ) *EventNotificationResponse`

NewEventNotificationResponse instantiates a new EventNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationResponseWithDefaults

`func NewEventNotificationResponseWithDefaults() *EventNotificationResponse`

NewEventNotificationResponseWithDefaults instantiates a new EventNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationStates

`func (o *EventNotificationResponse) GetNotificationStates() []NotificationStatus`

GetNotificationStates returns the NotificationStates field if non-nil, zero value otherwise.

### GetNotificationStatesOk

`func (o *EventNotificationResponse) GetNotificationStatesOk() (*[]NotificationStatus, bool)`

GetNotificationStatesOk returns a tuple with the NotificationStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationStates

`func (o *EventNotificationResponse) SetNotificationStates(v []NotificationStatus)`

SetNotificationStates sets NotificationStates field to given value.


### GetUser

`func (o *EventNotificationResponse) GetUser() map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventNotificationResponse) GetUserOk() (*map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventNotificationResponse) SetUser(v map[string]interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *EventNotificationResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserEmail

`func (o *EventNotificationResponse) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *EventNotificationResponse) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *EventNotificationResponse) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *EventNotificationResponse) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserId

`func (o *EventNotificationResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventNotificationResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventNotificationResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserMobileTelNo

`func (o *EventNotificationResponse) GetUserMobileTelNo() string`

GetUserMobileTelNo returns the UserMobileTelNo field if non-nil, zero value otherwise.

### GetUserMobileTelNoOk

`func (o *EventNotificationResponse) GetUserMobileTelNoOk() (*string, bool)`

GetUserMobileTelNoOk returns a tuple with the UserMobileTelNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMobileTelNo

`func (o *EventNotificationResponse) SetUserMobileTelNo(v string)`

SetUserMobileTelNo sets UserMobileTelNo field to given value.

### HasUserMobileTelNo

`func (o *EventNotificationResponse) HasUserMobileTelNo() bool`

HasUserMobileTelNo returns a boolean if a field has been set.

### GetUserNameNotification

`func (o *EventNotificationResponse) GetUserNameNotification() string`

GetUserNameNotification returns the UserNameNotification field if non-nil, zero value otherwise.

### GetUserNameNotificationOk

`func (o *EventNotificationResponse) GetUserNameNotificationOk() (*string, bool)`

GetUserNameNotificationOk returns a tuple with the UserNameNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameNotification

`func (o *EventNotificationResponse) SetUserNameNotification(v string)`

SetUserNameNotification sets UserNameNotification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


