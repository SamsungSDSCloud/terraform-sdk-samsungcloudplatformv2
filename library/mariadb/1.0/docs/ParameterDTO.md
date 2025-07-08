# ParameterDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowsValue** | **string** | Parameter allows value | 
**AppliedValue** | **string** | Parameter applied value | 
**ApplyType** | **string** | Parameter apply type | 
**DataType** | **string** | Parameter data type | 
**DefaultValue** | **string** | Parameter default value | 
**Description** | **NullableString** |  | 
**Id** | **string** | Parameter id | 
**IsModifiable** | **bool** | Parameter is_modifiable | 
**Name** | **string** | Parameter name | 
**SoftwareType** | **NullableString** |  | 

## Methods

### NewParameterDTO

`func NewParameterDTO(allowsValue string, appliedValue string, applyType string, dataType string, defaultValue string, description NullableString, id string, isModifiable bool, name string, softwareType NullableString, ) *ParameterDTO`

NewParameterDTO instantiates a new ParameterDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterDTOWithDefaults

`func NewParameterDTOWithDefaults() *ParameterDTO`

NewParameterDTOWithDefaults instantiates a new ParameterDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowsValue

`func (o *ParameterDTO) GetAllowsValue() string`

GetAllowsValue returns the AllowsValue field if non-nil, zero value otherwise.

### GetAllowsValueOk

`func (o *ParameterDTO) GetAllowsValueOk() (*string, bool)`

GetAllowsValueOk returns a tuple with the AllowsValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsValue

`func (o *ParameterDTO) SetAllowsValue(v string)`

SetAllowsValue sets AllowsValue field to given value.


### GetAppliedValue

`func (o *ParameterDTO) GetAppliedValue() string`

GetAppliedValue returns the AppliedValue field if non-nil, zero value otherwise.

### GetAppliedValueOk

`func (o *ParameterDTO) GetAppliedValueOk() (*string, bool)`

GetAppliedValueOk returns a tuple with the AppliedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedValue

`func (o *ParameterDTO) SetAppliedValue(v string)`

SetAppliedValue sets AppliedValue field to given value.


### GetApplyType

`func (o *ParameterDTO) GetApplyType() string`

GetApplyType returns the ApplyType field if non-nil, zero value otherwise.

### GetApplyTypeOk

`func (o *ParameterDTO) GetApplyTypeOk() (*string, bool)`

GetApplyTypeOk returns a tuple with the ApplyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyType

`func (o *ParameterDTO) SetApplyType(v string)`

SetApplyType sets ApplyType field to given value.


### GetDataType

`func (o *ParameterDTO) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ParameterDTO) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ParameterDTO) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetDefaultValue

`func (o *ParameterDTO) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ParameterDTO) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ParameterDTO) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.


### GetDescription

`func (o *ParameterDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParameterDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParameterDTO) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ParameterDTO) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ParameterDTO) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *ParameterDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterDTO) SetId(v string)`

SetId sets Id field to given value.


### GetIsModifiable

`func (o *ParameterDTO) GetIsModifiable() bool`

GetIsModifiable returns the IsModifiable field if non-nil, zero value otherwise.

### GetIsModifiableOk

`func (o *ParameterDTO) GetIsModifiableOk() (*bool, bool)`

GetIsModifiableOk returns a tuple with the IsModifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModifiable

`func (o *ParameterDTO) SetIsModifiable(v bool)`

SetIsModifiable sets IsModifiable field to given value.


### GetName

`func (o *ParameterDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterDTO) SetName(v string)`

SetName sets Name field to given value.


### GetSoftwareType

`func (o *ParameterDTO) GetSoftwareType() string`

GetSoftwareType returns the SoftwareType field if non-nil, zero value otherwise.

### GetSoftwareTypeOk

`func (o *ParameterDTO) GetSoftwareTypeOk() (*string, bool)`

GetSoftwareTypeOk returns a tuple with the SoftwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareType

`func (o *ParameterDTO) SetSoftwareType(v string)`

SetSoftwareType sets SoftwareType field to given value.


### SetSoftwareTypeNil

`func (o *ParameterDTO) SetSoftwareTypeNil(b bool)`

 SetSoftwareTypeNil sets the value for SoftwareType to be an explicit nil

### UnsetSoftwareType
`func (o *ParameterDTO) UnsetSoftwareType()`

UnsetSoftwareType ensures that no value is present for SoftwareType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


