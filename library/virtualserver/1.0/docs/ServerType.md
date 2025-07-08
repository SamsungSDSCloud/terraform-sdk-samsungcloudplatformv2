# ServerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** |  | 
**Disk** | **int32** | Size of root disk | 
**Ephemeral** | **int32** | Size of ephemeral disk | 
**ExtraSpecs** | **map[string]interface{}** | Extra specs | 
**Id** | **string** | Server type ID | 
**Name** | **string** | Server type name | 
**Ram** | **int32** | Amount of RAM | 
**Swap** | **int32** | Size of dedicated swap disk | 
**Vcpus** | **int32** | Number of CPUs | 

## Methods

### NewServerType

`func NewServerType(description NullableString, disk int32, ephemeral int32, extraSpecs map[string]interface{}, id string, name string, ram int32, swap int32, vcpus int32, ) *ServerType`

NewServerType instantiates a new ServerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTypeWithDefaults

`func NewServerTypeWithDefaults() *ServerType`

NewServerTypeWithDefaults instantiates a new ServerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServerType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerType) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ServerType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServerType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisk

`func (o *ServerType) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ServerType) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ServerType) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetEphemeral

`func (o *ServerType) GetEphemeral() int32`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *ServerType) GetEphemeralOk() (*int32, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *ServerType) SetEphemeral(v int32)`

SetEphemeral sets Ephemeral field to given value.


### GetExtraSpecs

`func (o *ServerType) GetExtraSpecs() map[string]interface{}`

GetExtraSpecs returns the ExtraSpecs field if non-nil, zero value otherwise.

### GetExtraSpecsOk

`func (o *ServerType) GetExtraSpecsOk() (*map[string]interface{}, bool)`

GetExtraSpecsOk returns a tuple with the ExtraSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraSpecs

`func (o *ServerType) SetExtraSpecs(v map[string]interface{})`

SetExtraSpecs sets ExtraSpecs field to given value.


### GetId

`func (o *ServerType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerType) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ServerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerType) SetName(v string)`

SetName sets Name field to given value.


### GetRam

`func (o *ServerType) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ServerType) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ServerType) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetSwap

`func (o *ServerType) GetSwap() int32`

GetSwap returns the Swap field if non-nil, zero value otherwise.

### GetSwapOk

`func (o *ServerType) GetSwapOk() (*int32, bool)`

GetSwapOk returns a tuple with the Swap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwap

`func (o *ServerType) SetSwap(v int32)`

SetSwap sets Swap field to given value.


### GetVcpus

`func (o *ServerType) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *ServerType) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *ServerType) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


