# RoleSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**MaxSessionDuration** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewRoleSetRequest

`func NewRoleSetRequest() *RoleSetRequest`

NewRoleSetRequest instantiates a new RoleSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleSetRequestWithDefaults

`func NewRoleSetRequestWithDefaults() *RoleSetRequest`

NewRoleSetRequestWithDefaults instantiates a new RoleSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RoleSetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleSetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleSetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleSetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RoleSetRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RoleSetRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMaxSessionDuration

`func (o *RoleSetRequest) GetMaxSessionDuration() int32`

GetMaxSessionDuration returns the MaxSessionDuration field if non-nil, zero value otherwise.

### GetMaxSessionDurationOk

`func (o *RoleSetRequest) GetMaxSessionDurationOk() (*int32, bool)`

GetMaxSessionDurationOk returns a tuple with the MaxSessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionDuration

`func (o *RoleSetRequest) SetMaxSessionDuration(v int32)`

SetMaxSessionDuration sets MaxSessionDuration field to given value.

### HasMaxSessionDuration

`func (o *RoleSetRequest) HasMaxSessionDuration() bool`

HasMaxSessionDuration returns a boolean if a field has been set.

### SetMaxSessionDurationNil

`func (o *RoleSetRequest) SetMaxSessionDurationNil(b bool)`

 SetMaxSessionDurationNil sets the value for MaxSessionDuration to be an explicit nil

### UnsetMaxSessionDuration
`func (o *RoleSetRequest) UnsetMaxSessionDuration()`

UnsetMaxSessionDuration ensures that no value is present for MaxSessionDuration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


