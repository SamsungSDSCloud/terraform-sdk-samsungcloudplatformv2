# ServerTypeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Core** | Pointer to **NullableString** |  | [optional] 
**GpuName** | Pointer to **NullableString** |  | [optional] 
**InstanceType** | Pointer to **NullableString** |  | [optional] 
**MemoryGb** | Pointer to **NullableString** |  | [optional] 
**ScaleUpYn** | Pointer to **bool** | Scale up or not | [optional] [default to true]
**ServerType** | Pointer to **NullableString** |  | [optional] 
**ServerTypeDescription** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServerTypeDTO

`func NewServerTypeDTO() *ServerTypeDTO`

NewServerTypeDTO instantiates a new ServerTypeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTypeDTOWithDefaults

`func NewServerTypeDTOWithDefaults() *ServerTypeDTO`

NewServerTypeDTOWithDefaults instantiates a new ServerTypeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCore

`func (o *ServerTypeDTO) GetCore() string`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *ServerTypeDTO) GetCoreOk() (*string, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *ServerTypeDTO) SetCore(v string)`

SetCore sets Core field to given value.

### HasCore

`func (o *ServerTypeDTO) HasCore() bool`

HasCore returns a boolean if a field has been set.

### SetCoreNil

`func (o *ServerTypeDTO) SetCoreNil(b bool)`

 SetCoreNil sets the value for Core to be an explicit nil

### UnsetCore
`func (o *ServerTypeDTO) UnsetCore()`

UnsetCore ensures that no value is present for Core, not even an explicit nil
### GetGpuName

`func (o *ServerTypeDTO) GetGpuName() string`

GetGpuName returns the GpuName field if non-nil, zero value otherwise.

### GetGpuNameOk

`func (o *ServerTypeDTO) GetGpuNameOk() (*string, bool)`

GetGpuNameOk returns a tuple with the GpuName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuName

`func (o *ServerTypeDTO) SetGpuName(v string)`

SetGpuName sets GpuName field to given value.

### HasGpuName

`func (o *ServerTypeDTO) HasGpuName() bool`

HasGpuName returns a boolean if a field has been set.

### SetGpuNameNil

`func (o *ServerTypeDTO) SetGpuNameNil(b bool)`

 SetGpuNameNil sets the value for GpuName to be an explicit nil

### UnsetGpuName
`func (o *ServerTypeDTO) UnsetGpuName()`

UnsetGpuName ensures that no value is present for GpuName, not even an explicit nil
### GetInstanceType

`func (o *ServerTypeDTO) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ServerTypeDTO) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ServerTypeDTO) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ServerTypeDTO) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *ServerTypeDTO) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *ServerTypeDTO) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetMemoryGb

`func (o *ServerTypeDTO) GetMemoryGb() string`

GetMemoryGb returns the MemoryGb field if non-nil, zero value otherwise.

### GetMemoryGbOk

`func (o *ServerTypeDTO) GetMemoryGbOk() (*string, bool)`

GetMemoryGbOk returns a tuple with the MemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryGb

`func (o *ServerTypeDTO) SetMemoryGb(v string)`

SetMemoryGb sets MemoryGb field to given value.

### HasMemoryGb

`func (o *ServerTypeDTO) HasMemoryGb() bool`

HasMemoryGb returns a boolean if a field has been set.

### SetMemoryGbNil

`func (o *ServerTypeDTO) SetMemoryGbNil(b bool)`

 SetMemoryGbNil sets the value for MemoryGb to be an explicit nil

### UnsetMemoryGb
`func (o *ServerTypeDTO) UnsetMemoryGb()`

UnsetMemoryGb ensures that no value is present for MemoryGb, not even an explicit nil
### GetScaleUpYn

`func (o *ServerTypeDTO) GetScaleUpYn() bool`

GetScaleUpYn returns the ScaleUpYn field if non-nil, zero value otherwise.

### GetScaleUpYnOk

`func (o *ServerTypeDTO) GetScaleUpYnOk() (*bool, bool)`

GetScaleUpYnOk returns a tuple with the ScaleUpYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleUpYn

`func (o *ServerTypeDTO) SetScaleUpYn(v bool)`

SetScaleUpYn sets ScaleUpYn field to given value.

### HasScaleUpYn

`func (o *ServerTypeDTO) HasScaleUpYn() bool`

HasScaleUpYn returns a boolean if a field has been set.

### GetServerType

`func (o *ServerTypeDTO) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ServerTypeDTO) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ServerTypeDTO) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *ServerTypeDTO) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### SetServerTypeNil

`func (o *ServerTypeDTO) SetServerTypeNil(b bool)`

 SetServerTypeNil sets the value for ServerType to be an explicit nil

### UnsetServerType
`func (o *ServerTypeDTO) UnsetServerType()`

UnsetServerType ensures that no value is present for ServerType, not even an explicit nil
### GetServerTypeDescription

`func (o *ServerTypeDTO) GetServerTypeDescription() string`

GetServerTypeDescription returns the ServerTypeDescription field if non-nil, zero value otherwise.

### GetServerTypeDescriptionOk

`func (o *ServerTypeDTO) GetServerTypeDescriptionOk() (*string, bool)`

GetServerTypeDescriptionOk returns a tuple with the ServerTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeDescription

`func (o *ServerTypeDTO) SetServerTypeDescription(v string)`

SetServerTypeDescription sets ServerTypeDescription field to given value.

### HasServerTypeDescription

`func (o *ServerTypeDTO) HasServerTypeDescription() bool`

HasServerTypeDescription returns a boolean if a field has been set.

### SetServerTypeDescriptionNil

`func (o *ServerTypeDTO) SetServerTypeDescriptionNil(b bool)`

 SetServerTypeDescriptionNil sets the value for ServerTypeDescription to be an explicit nil

### UnsetServerTypeDescription
`func (o *ServerTypeDTO) UnsetServerTypeDescription()`

UnsetServerTypeDescription ensures that no value is present for ServerTypeDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


