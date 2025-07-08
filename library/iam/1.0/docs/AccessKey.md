# AccessKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | Access Key | 
**AccessKeyType** | [**AccessKeyTypeEnum**](AccessKeyTypeEnum.md) | Access key type | 
**AccountId** | **string** | Account Id | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExpirationTimestamp** | **time.Time** | Access key expiration timestamp | 
**Id** | **string** | ID | 
**IsEnabled** | **bool** | 활성화 여부 | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**ParentAccessKeyId** | **NullableString** |  | 
**SecretKey** | **string** | Secret Key | 

## Methods

### NewAccessKey

`func NewAccessKey(accessKey string, accessKeyType AccessKeyTypeEnum, accountId string, createdAt time.Time, createdBy string, expirationTimestamp time.Time, id string, isEnabled bool, modifiedAt time.Time, modifiedBy string, parentAccessKeyId NullableString, secretKey string, ) *AccessKey`

NewAccessKey instantiates a new AccessKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyWithDefaults

`func NewAccessKeyWithDefaults() *AccessKey`

NewAccessKeyWithDefaults instantiates a new AccessKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AccessKey) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AccessKey) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AccessKey) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetAccessKeyType

`func (o *AccessKey) GetAccessKeyType() AccessKeyTypeEnum`

GetAccessKeyType returns the AccessKeyType field if non-nil, zero value otherwise.

### GetAccessKeyTypeOk

`func (o *AccessKey) GetAccessKeyTypeOk() (*AccessKeyTypeEnum, bool)`

GetAccessKeyTypeOk returns a tuple with the AccessKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyType

`func (o *AccessKey) SetAccessKeyType(v AccessKeyTypeEnum)`

SetAccessKeyType sets AccessKeyType field to given value.


### GetAccountId

`func (o *AccessKey) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccessKey) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccessKey) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *AccessKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AccessKey) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AccessKey) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AccessKey) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *AccessKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccessKey) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccessKey) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpirationTimestamp

`func (o *AccessKey) GetExpirationTimestamp() time.Time`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *AccessKey) GetExpirationTimestampOk() (*time.Time, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *AccessKey) SetExpirationTimestamp(v time.Time)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.


### GetId

`func (o *AccessKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessKey) SetId(v string)`

SetId sets Id field to given value.


### GetIsEnabled

`func (o *AccessKey) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AccessKey) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AccessKey) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetModifiedAt

`func (o *AccessKey) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AccessKey) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AccessKey) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *AccessKey) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *AccessKey) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *AccessKey) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetParentAccessKeyId

`func (o *AccessKey) GetParentAccessKeyId() string`

GetParentAccessKeyId returns the ParentAccessKeyId field if non-nil, zero value otherwise.

### GetParentAccessKeyIdOk

`func (o *AccessKey) GetParentAccessKeyIdOk() (*string, bool)`

GetParentAccessKeyIdOk returns a tuple with the ParentAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccessKeyId

`func (o *AccessKey) SetParentAccessKeyId(v string)`

SetParentAccessKeyId sets ParentAccessKeyId field to given value.


### SetParentAccessKeyIdNil

`func (o *AccessKey) SetParentAccessKeyIdNil(b bool)`

 SetParentAccessKeyIdNil sets the value for ParentAccessKeyId to be an explicit nil

### UnsetParentAccessKeyId
`func (o *AccessKey) UnsetParentAccessKeyId()`

UnsetParentAccessKeyId ensures that no value is present for ParentAccessKeyId, not even an explicit nil
### GetSecretKey

`func (o *AccessKey) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AccessKey) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AccessKey) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


