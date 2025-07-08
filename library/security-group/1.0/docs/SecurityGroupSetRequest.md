# SecurityGroupSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Loggable** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSecurityGroupSetRequest

`func NewSecurityGroupSetRequest() *SecurityGroupSetRequest`

NewSecurityGroupSetRequest instantiates a new SecurityGroupSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupSetRequestWithDefaults

`func NewSecurityGroupSetRequestWithDefaults() *SecurityGroupSetRequest`

NewSecurityGroupSetRequestWithDefaults instantiates a new SecurityGroupSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SecurityGroupSetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroupSetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroupSetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityGroupSetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityGroupSetRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityGroupSetRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLoggable

`func (o *SecurityGroupSetRequest) GetLoggable() bool`

GetLoggable returns the Loggable field if non-nil, zero value otherwise.

### GetLoggableOk

`func (o *SecurityGroupSetRequest) GetLoggableOk() (*bool, bool)`

GetLoggableOk returns a tuple with the Loggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggable

`func (o *SecurityGroupSetRequest) SetLoggable(v bool)`

SetLoggable sets Loggable field to given value.

### HasLoggable

`func (o *SecurityGroupSetRequest) HasLoggable() bool`

HasLoggable returns a boolean if a field has been set.

### SetLoggableNil

`func (o *SecurityGroupSetRequest) SetLoggableNil(b bool)`

 SetLoggableNil sets the value for Loggable to be an explicit nil

### UnsetLoggable
`func (o *SecurityGroupSetRequest) UnsetLoggable()`

UnsetLoggable ensures that no value is present for Loggable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


