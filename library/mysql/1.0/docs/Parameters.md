# Parameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbaasParameterGroupId** | **string** | Parameter id | 
**DefaultValue** | **string** | Parameter default value | 
**Description** | **string** | Parameter description | 
**Id** | **string** | Parameter id | 
**Name** | **string** | Parameter name | 
**SoftwareType** | **NullableString** |  | 

## Methods

### NewParameters

`func NewParameters(dbaasParameterGroupId string, defaultValue string, description string, id string, name string, softwareType NullableString, ) *Parameters`

NewParameters instantiates a new Parameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParametersWithDefaults

`func NewParametersWithDefaults() *Parameters`

NewParametersWithDefaults instantiates a new Parameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbaasParameterGroupId

`func (o *Parameters) GetDbaasParameterGroupId() string`

GetDbaasParameterGroupId returns the DbaasParameterGroupId field if non-nil, zero value otherwise.

### GetDbaasParameterGroupIdOk

`func (o *Parameters) GetDbaasParameterGroupIdOk() (*string, bool)`

GetDbaasParameterGroupIdOk returns a tuple with the DbaasParameterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasParameterGroupId

`func (o *Parameters) SetDbaasParameterGroupId(v string)`

SetDbaasParameterGroupId sets DbaasParameterGroupId field to given value.


### GetDefaultValue

`func (o *Parameters) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *Parameters) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *Parameters) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.


### GetDescription

`func (o *Parameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Parameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Parameters) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *Parameters) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Parameters) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Parameters) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Parameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Parameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Parameters) SetName(v string)`

SetName sets Name field to given value.


### GetSoftwareType

`func (o *Parameters) GetSoftwareType() string`

GetSoftwareType returns the SoftwareType field if non-nil, zero value otherwise.

### GetSoftwareTypeOk

`func (o *Parameters) GetSoftwareTypeOk() (*string, bool)`

GetSoftwareTypeOk returns a tuple with the SoftwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareType

`func (o *Parameters) SetSoftwareType(v string)`

SetSoftwareType sets SoftwareType field to given value.


### SetSoftwareTypeNil

`func (o *Parameters) SetSoftwareTypeNil(b bool)`

 SetSoftwareTypeNil sets the value for SoftwareType to be an explicit nil

### UnsetSoftwareType
`func (o *Parameters) UnsetSoftwareType()`

UnsetSoftwareType ensures that no value is present for SoftwareType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


