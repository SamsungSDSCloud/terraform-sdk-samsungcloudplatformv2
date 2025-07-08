# NotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddrbookAdditionalInfo** | Pointer to [**AddrbookAdditionalInfo**](AddrbookAdditionalInfo.md) |  | [optional] 
**NotificationMethods** | [**[]NotificationMethod**](NotificationMethod.md) | 알림 메소드 정보 | 
**RecipientKey** | **string** | 수신인 키 - 수신인 키는 @[프로젝트 멤버 조회] 또는 @[사용자의 주소록 목록 조회]를 이용하여 조회합니다. | 
**RecipientName** | **string** | 수신인 이름 | 
**RecipientType** | **string** | 수신인 유형 - USER : 사용자, ADDRBOOK : 주소록 | 
**UserAdditionalInfo** | Pointer to [**Member**](Member.md) |  | [optional] 

## Methods

### NewNotificationResponse

`func NewNotificationResponse(notificationMethods []NotificationMethod, recipientKey string, recipientName string, recipientType string, ) *NotificationResponse`

NewNotificationResponse instantiates a new NotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationResponseWithDefaults

`func NewNotificationResponseWithDefaults() *NotificationResponse`

NewNotificationResponseWithDefaults instantiates a new NotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddrbookAdditionalInfo

`func (o *NotificationResponse) GetAddrbookAdditionalInfo() AddrbookAdditionalInfo`

GetAddrbookAdditionalInfo returns the AddrbookAdditionalInfo field if non-nil, zero value otherwise.

### GetAddrbookAdditionalInfoOk

`func (o *NotificationResponse) GetAddrbookAdditionalInfoOk() (*AddrbookAdditionalInfo, bool)`

GetAddrbookAdditionalInfoOk returns a tuple with the AddrbookAdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrbookAdditionalInfo

`func (o *NotificationResponse) SetAddrbookAdditionalInfo(v AddrbookAdditionalInfo)`

SetAddrbookAdditionalInfo sets AddrbookAdditionalInfo field to given value.

### HasAddrbookAdditionalInfo

`func (o *NotificationResponse) HasAddrbookAdditionalInfo() bool`

HasAddrbookAdditionalInfo returns a boolean if a field has been set.

### GetNotificationMethods

`func (o *NotificationResponse) GetNotificationMethods() []NotificationMethod`

GetNotificationMethods returns the NotificationMethods field if non-nil, zero value otherwise.

### GetNotificationMethodsOk

`func (o *NotificationResponse) GetNotificationMethodsOk() (*[]NotificationMethod, bool)`

GetNotificationMethodsOk returns a tuple with the NotificationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMethods

`func (o *NotificationResponse) SetNotificationMethods(v []NotificationMethod)`

SetNotificationMethods sets NotificationMethods field to given value.


### GetRecipientKey

`func (o *NotificationResponse) GetRecipientKey() string`

GetRecipientKey returns the RecipientKey field if non-nil, zero value otherwise.

### GetRecipientKeyOk

`func (o *NotificationResponse) GetRecipientKeyOk() (*string, bool)`

GetRecipientKeyOk returns a tuple with the RecipientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientKey

`func (o *NotificationResponse) SetRecipientKey(v string)`

SetRecipientKey sets RecipientKey field to given value.


### GetRecipientName

`func (o *NotificationResponse) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *NotificationResponse) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *NotificationResponse) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.


### GetRecipientType

`func (o *NotificationResponse) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *NotificationResponse) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *NotificationResponse) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.


### GetUserAdditionalInfo

`func (o *NotificationResponse) GetUserAdditionalInfo() Member`

GetUserAdditionalInfo returns the UserAdditionalInfo field if non-nil, zero value otherwise.

### GetUserAdditionalInfoOk

`func (o *NotificationResponse) GetUserAdditionalInfoOk() (*Member, bool)`

GetUserAdditionalInfoOk returns a tuple with the UserAdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAdditionalInfo

`func (o *NotificationResponse) SetUserAdditionalInfo(v Member)`

SetUserAdditionalInfo sets UserAdditionalInfo field to given value.

### HasUserAdditionalInfo

`func (o *NotificationResponse) HasUserAdditionalInfo() bool`

HasUserAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


