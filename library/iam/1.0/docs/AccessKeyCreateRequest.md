# AccessKeyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyType** | [**AccessKeyTypeCreateRequestEnum**](AccessKeyTypeCreateRequestEnum.md) | Access key type | 
**AccountId** | Pointer to **string** | Account Id | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Duration** | Pointer to **NullableString** |  | [optional] 
**ParentAccessKeyId** | Pointer to **NullableString** |  | [optional] 
**Passcode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAccessKeyCreateRequest

`func NewAccessKeyCreateRequest(accessKeyType AccessKeyTypeCreateRequestEnum, ) *AccessKeyCreateRequest`

NewAccessKeyCreateRequest instantiates a new AccessKeyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyCreateRequestWithDefaults

`func NewAccessKeyCreateRequestWithDefaults() *AccessKeyCreateRequest`

NewAccessKeyCreateRequestWithDefaults instantiates a new AccessKeyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyType

`func (o *AccessKeyCreateRequest) GetAccessKeyType() AccessKeyTypeCreateRequestEnum`

GetAccessKeyType returns the AccessKeyType field if non-nil, zero value otherwise.

### GetAccessKeyTypeOk

`func (o *AccessKeyCreateRequest) GetAccessKeyTypeOk() (*AccessKeyTypeCreateRequestEnum, bool)`

GetAccessKeyTypeOk returns a tuple with the AccessKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyType

`func (o *AccessKeyCreateRequest) SetAccessKeyType(v AccessKeyTypeCreateRequestEnum)`

SetAccessKeyType sets AccessKeyType field to given value.


### GetAccountId

`func (o *AccessKeyCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccessKeyCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccessKeyCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccessKeyCreateRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDescription

`func (o *AccessKeyCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessKeyCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessKeyCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessKeyCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccessKeyCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccessKeyCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDuration

`func (o *AccessKeyCreateRequest) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AccessKeyCreateRequest) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AccessKeyCreateRequest) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AccessKeyCreateRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *AccessKeyCreateRequest) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *AccessKeyCreateRequest) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetParentAccessKeyId

`func (o *AccessKeyCreateRequest) GetParentAccessKeyId() string`

GetParentAccessKeyId returns the ParentAccessKeyId field if non-nil, zero value otherwise.

### GetParentAccessKeyIdOk

`func (o *AccessKeyCreateRequest) GetParentAccessKeyIdOk() (*string, bool)`

GetParentAccessKeyIdOk returns a tuple with the ParentAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccessKeyId

`func (o *AccessKeyCreateRequest) SetParentAccessKeyId(v string)`

SetParentAccessKeyId sets ParentAccessKeyId field to given value.

### HasParentAccessKeyId

`func (o *AccessKeyCreateRequest) HasParentAccessKeyId() bool`

HasParentAccessKeyId returns a boolean if a field has been set.

### SetParentAccessKeyIdNil

`func (o *AccessKeyCreateRequest) SetParentAccessKeyIdNil(b bool)`

 SetParentAccessKeyIdNil sets the value for ParentAccessKeyId to be an explicit nil

### UnsetParentAccessKeyId
`func (o *AccessKeyCreateRequest) UnsetParentAccessKeyId()`

UnsetParentAccessKeyId ensures that no value is present for ParentAccessKeyId, not even an explicit nil
### GetPasscode

`func (o *AccessKeyCreateRequest) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *AccessKeyCreateRequest) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *AccessKeyCreateRequest) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *AccessKeyCreateRequest) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.

### SetPasscodeNil

`func (o *AccessKeyCreateRequest) SetPasscodeNil(b bool)`

 SetPasscodeNil sets the value for Passcode to be an explicit nil

### UnsetPasscode
`func (o *AccessKeyCreateRequest) UnsetPasscode()`

UnsetPasscode ensures that no value is present for Passcode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


