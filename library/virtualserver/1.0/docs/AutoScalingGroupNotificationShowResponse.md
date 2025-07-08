# AutoScalingGroupNotificationShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AutoScalingGroupId** | **string** | Auto-Scaling Group ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | **string** | ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**NotificationEvents** | **[]string** | Auto-Scaling Group notification events | 
**NotificationState** | **string** | Auto-Scaling Group notification state | 
**UserEmail** | **string** | User email | 
**UserId** | **string** | User ID | 
**UserName** | **string** | User name | 

## Methods

### NewAutoScalingGroupNotificationShowResponse

`func NewAutoScalingGroupNotificationShowResponse(accountId string, autoScalingGroupId string, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, notificationEvents []string, notificationState string, userEmail string, userId string, userName string, ) *AutoScalingGroupNotificationShowResponse`

NewAutoScalingGroupNotificationShowResponse instantiates a new AutoScalingGroupNotificationShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupNotificationShowResponseWithDefaults

`func NewAutoScalingGroupNotificationShowResponseWithDefaults() *AutoScalingGroupNotificationShowResponse`

NewAutoScalingGroupNotificationShowResponseWithDefaults instantiates a new AutoScalingGroupNotificationShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AutoScalingGroupNotificationShowResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AutoScalingGroupNotificationShowResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AutoScalingGroupNotificationShowResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAutoScalingGroupId

`func (o *AutoScalingGroupNotificationShowResponse) GetAutoScalingGroupId() string`

GetAutoScalingGroupId returns the AutoScalingGroupId field if non-nil, zero value otherwise.

### GetAutoScalingGroupIdOk

`func (o *AutoScalingGroupNotificationShowResponse) GetAutoScalingGroupIdOk() (*string, bool)`

GetAutoScalingGroupIdOk returns a tuple with the AutoScalingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroupId

`func (o *AutoScalingGroupNotificationShowResponse) SetAutoScalingGroupId(v string)`

SetAutoScalingGroupId sets AutoScalingGroupId field to given value.


### GetCreatedAt

`func (o *AutoScalingGroupNotificationShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutoScalingGroupNotificationShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutoScalingGroupNotificationShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AutoScalingGroupNotificationShowResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AutoScalingGroupNotificationShowResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AutoScalingGroupNotificationShowResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *AutoScalingGroupNotificationShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoScalingGroupNotificationShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoScalingGroupNotificationShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *AutoScalingGroupNotificationShowResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AutoScalingGroupNotificationShowResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AutoScalingGroupNotificationShowResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *AutoScalingGroupNotificationShowResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *AutoScalingGroupNotificationShowResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *AutoScalingGroupNotificationShowResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetNotificationEvents

`func (o *AutoScalingGroupNotificationShowResponse) GetNotificationEvents() []string`

GetNotificationEvents returns the NotificationEvents field if non-nil, zero value otherwise.

### GetNotificationEventsOk

`func (o *AutoScalingGroupNotificationShowResponse) GetNotificationEventsOk() (*[]string, bool)`

GetNotificationEventsOk returns a tuple with the NotificationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEvents

`func (o *AutoScalingGroupNotificationShowResponse) SetNotificationEvents(v []string)`

SetNotificationEvents sets NotificationEvents field to given value.


### GetNotificationState

`func (o *AutoScalingGroupNotificationShowResponse) GetNotificationState() string`

GetNotificationState returns the NotificationState field if non-nil, zero value otherwise.

### GetNotificationStateOk

`func (o *AutoScalingGroupNotificationShowResponse) GetNotificationStateOk() (*string, bool)`

GetNotificationStateOk returns a tuple with the NotificationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationState

`func (o *AutoScalingGroupNotificationShowResponse) SetNotificationState(v string)`

SetNotificationState sets NotificationState field to given value.


### GetUserEmail

`func (o *AutoScalingGroupNotificationShowResponse) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *AutoScalingGroupNotificationShowResponse) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *AutoScalingGroupNotificationShowResponse) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetUserId

`func (o *AutoScalingGroupNotificationShowResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AutoScalingGroupNotificationShowResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AutoScalingGroupNotificationShowResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserName

`func (o *AutoScalingGroupNotificationShowResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AutoScalingGroupNotificationShowResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AutoScalingGroupNotificationShowResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


