# CommandItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedValue** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**Modifiable** | **bool** |  | 
**Name** | **string** |  | 

## Methods

### NewCommandItem

`func NewCommandItem(id string, modifiable bool, name string, ) *CommandItem`

NewCommandItem instantiates a new CommandItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandItemWithDefaults

`func NewCommandItemWithDefaults() *CommandItem`

NewCommandItemWithDefaults instantiates a new CommandItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedValue

`func (o *CommandItem) GetAppliedValue() string`

GetAppliedValue returns the AppliedValue field if non-nil, zero value otherwise.

### GetAppliedValueOk

`func (o *CommandItem) GetAppliedValueOk() (*string, bool)`

GetAppliedValueOk returns a tuple with the AppliedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedValue

`func (o *CommandItem) SetAppliedValue(v string)`

SetAppliedValue sets AppliedValue field to given value.

### HasAppliedValue

`func (o *CommandItem) HasAppliedValue() bool`

HasAppliedValue returns a boolean if a field has been set.

### SetAppliedValueNil

`func (o *CommandItem) SetAppliedValueNil(b bool)`

 SetAppliedValueNil sets the value for AppliedValue to be an explicit nil

### UnsetAppliedValue
`func (o *CommandItem) UnsetAppliedValue()`

UnsetAppliedValue ensures that no value is present for AppliedValue, not even an explicit nil
### GetDescription

`func (o *CommandItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommandItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommandItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommandItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CommandItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CommandItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *CommandItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommandItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommandItem) SetId(v string)`

SetId sets Id field to given value.


### GetModifiable

`func (o *CommandItem) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *CommandItem) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *CommandItem) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.


### GetName

`func (o *CommandItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommandItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommandItem) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


