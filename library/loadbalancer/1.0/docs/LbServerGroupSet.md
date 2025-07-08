# LbServerGroupSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**LbHealthCheckId** | Pointer to **NullableString** |  | [optional] 
**LbMethod** | Pointer to [**NullableLbServerGroupLbMethod**](LbServerGroupLbMethod.md) |  | [optional] 

## Methods

### NewLbServerGroupSet

`func NewLbServerGroupSet() *LbServerGroupSet`

NewLbServerGroupSet instantiates a new LbServerGroupSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbServerGroupSetWithDefaults

`func NewLbServerGroupSetWithDefaults() *LbServerGroupSet`

NewLbServerGroupSetWithDefaults instantiates a new LbServerGroupSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LbServerGroupSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LbServerGroupSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LbServerGroupSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LbServerGroupSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LbServerGroupSet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LbServerGroupSet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLbHealthCheckId

`func (o *LbServerGroupSet) GetLbHealthCheckId() string`

GetLbHealthCheckId returns the LbHealthCheckId field if non-nil, zero value otherwise.

### GetLbHealthCheckIdOk

`func (o *LbServerGroupSet) GetLbHealthCheckIdOk() (*string, bool)`

GetLbHealthCheckIdOk returns a tuple with the LbHealthCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbHealthCheckId

`func (o *LbServerGroupSet) SetLbHealthCheckId(v string)`

SetLbHealthCheckId sets LbHealthCheckId field to given value.

### HasLbHealthCheckId

`func (o *LbServerGroupSet) HasLbHealthCheckId() bool`

HasLbHealthCheckId returns a boolean if a field has been set.

### SetLbHealthCheckIdNil

`func (o *LbServerGroupSet) SetLbHealthCheckIdNil(b bool)`

 SetLbHealthCheckIdNil sets the value for LbHealthCheckId to be an explicit nil

### UnsetLbHealthCheckId
`func (o *LbServerGroupSet) UnsetLbHealthCheckId()`

UnsetLbHealthCheckId ensures that no value is present for LbHealthCheckId, not even an explicit nil
### GetLbMethod

`func (o *LbServerGroupSet) GetLbMethod() LbServerGroupLbMethod`

GetLbMethod returns the LbMethod field if non-nil, zero value otherwise.

### GetLbMethodOk

`func (o *LbServerGroupSet) GetLbMethodOk() (*LbServerGroupLbMethod, bool)`

GetLbMethodOk returns a tuple with the LbMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbMethod

`func (o *LbServerGroupSet) SetLbMethod(v LbServerGroupLbMethod)`

SetLbMethod sets LbMethod field to given value.

### HasLbMethod

`func (o *LbServerGroupSet) HasLbMethod() bool`

HasLbMethod returns a boolean if a field has been set.

### SetLbMethodNil

`func (o *LbServerGroupSet) SetLbMethodNil(b bool)`

 SetLbMethodNil sets the value for LbMethod to be an explicit nil

### UnsetLbMethod
`func (o *LbServerGroupSet) UnsetLbMethod()`

UnsetLbMethod ensures that no value is present for LbMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


