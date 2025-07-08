# SecurityGroupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Loggable** | Pointer to **NullableBool** |  | [optional] 
**Name** | **string** | Security Group Name | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]

## Methods

### NewSecurityGroupCreateRequest

`func NewSecurityGroupCreateRequest(name string, ) *SecurityGroupCreateRequest`

NewSecurityGroupCreateRequest instantiates a new SecurityGroupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupCreateRequestWithDefaults

`func NewSecurityGroupCreateRequestWithDefaults() *SecurityGroupCreateRequest`

NewSecurityGroupCreateRequestWithDefaults instantiates a new SecurityGroupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SecurityGroupCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroupCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroupCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityGroupCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityGroupCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityGroupCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLoggable

`func (o *SecurityGroupCreateRequest) GetLoggable() bool`

GetLoggable returns the Loggable field if non-nil, zero value otherwise.

### GetLoggableOk

`func (o *SecurityGroupCreateRequest) GetLoggableOk() (*bool, bool)`

GetLoggableOk returns a tuple with the Loggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggable

`func (o *SecurityGroupCreateRequest) SetLoggable(v bool)`

SetLoggable sets Loggable field to given value.

### HasLoggable

`func (o *SecurityGroupCreateRequest) HasLoggable() bool`

HasLoggable returns a boolean if a field has been set.

### SetLoggableNil

`func (o *SecurityGroupCreateRequest) SetLoggableNil(b bool)`

 SetLoggableNil sets the value for Loggable to be an explicit nil

### UnsetLoggable
`func (o *SecurityGroupCreateRequest) UnsetLoggable()`

UnsetLoggable ensures that no value is present for Loggable, not even an explicit nil
### GetName

`func (o *SecurityGroupCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *SecurityGroupCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecurityGroupCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SecurityGroupCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SecurityGroupCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


