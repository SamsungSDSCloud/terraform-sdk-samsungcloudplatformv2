# NotificationRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationMethods** | [**[]NotificationMethod**](NotificationMethod.md) | 알림 메소드 | 
**RecipientKey** | **string** | 수신인 키 - 수신인 키는 @[프로젝트 멤버 조회] 또는 @[사용자의 주소록 목록 조회]를 이용하여 조회합니다. | 
**RecipientType** | **string** | 수신인 유형 - USER : 사용자, ADDRBOOK : 주소록 | 

## Methods

### NewNotificationRecipient

`func NewNotificationRecipient(notificationMethods []NotificationMethod, recipientKey string, recipientType string, ) *NotificationRecipient`

NewNotificationRecipient instantiates a new NotificationRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRecipientWithDefaults

`func NewNotificationRecipientWithDefaults() *NotificationRecipient`

NewNotificationRecipientWithDefaults instantiates a new NotificationRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationMethods

`func (o *NotificationRecipient) GetNotificationMethods() []NotificationMethod`

GetNotificationMethods returns the NotificationMethods field if non-nil, zero value otherwise.

### GetNotificationMethodsOk

`func (o *NotificationRecipient) GetNotificationMethodsOk() (*[]NotificationMethod, bool)`

GetNotificationMethodsOk returns a tuple with the NotificationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMethods

`func (o *NotificationRecipient) SetNotificationMethods(v []NotificationMethod)`

SetNotificationMethods sets NotificationMethods field to given value.


### GetRecipientKey

`func (o *NotificationRecipient) GetRecipientKey() string`

GetRecipientKey returns the RecipientKey field if non-nil, zero value otherwise.

### GetRecipientKeyOk

`func (o *NotificationRecipient) GetRecipientKeyOk() (*string, bool)`

GetRecipientKeyOk returns a tuple with the RecipientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientKey

`func (o *NotificationRecipient) SetRecipientKey(v string)`

SetRecipientKey sets RecipientKey field to given value.


### GetRecipientType

`func (o *NotificationRecipient) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *NotificationRecipient) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *NotificationRecipient) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


