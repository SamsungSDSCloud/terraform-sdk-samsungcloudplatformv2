# ServerShowResponseServerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disk** | **int32** | Size of root disk | 
**Ephemeral** | **int32** | Size of ephemeral disk | 
**ExtraSpecs** | **map[string]interface{}** | Extra specs | 
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Server type name | 
**Ram** | **int32** | Amount of RAM | 
**Swap** | **int32** | Size of dedicated swap disk | 
**Vcpus** | **int32** | Number of CPUs | 

## Methods

### NewServerShowResponseServerType

`func NewServerShowResponseServerType(disk int32, ephemeral int32, extraSpecs map[string]interface{}, name string, ram int32, swap int32, vcpus int32, ) *ServerShowResponseServerType`

NewServerShowResponseServerType instantiates a new ServerShowResponseServerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerShowResponseServerTypeWithDefaults

`func NewServerShowResponseServerTypeWithDefaults() *ServerShowResponseServerType`

NewServerShowResponseServerTypeWithDefaults instantiates a new ServerShowResponseServerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisk

`func (o *ServerShowResponseServerType) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ServerShowResponseServerType) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ServerShowResponseServerType) SetDisk(v int32)`

SetDisk sets Disk field to given value.


### GetEphemeral

`func (o *ServerShowResponseServerType) GetEphemeral() int32`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *ServerShowResponseServerType) GetEphemeralOk() (*int32, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *ServerShowResponseServerType) SetEphemeral(v int32)`

SetEphemeral sets Ephemeral field to given value.


### GetExtraSpecs

`func (o *ServerShowResponseServerType) GetExtraSpecs() map[string]interface{}`

GetExtraSpecs returns the ExtraSpecs field if non-nil, zero value otherwise.

### GetExtraSpecsOk

`func (o *ServerShowResponseServerType) GetExtraSpecsOk() (*map[string]interface{}, bool)`

GetExtraSpecsOk returns a tuple with the ExtraSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraSpecs

`func (o *ServerShowResponseServerType) SetExtraSpecs(v map[string]interface{})`

SetExtraSpecs sets ExtraSpecs field to given value.


### GetId

`func (o *ServerShowResponseServerType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerShowResponseServerType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerShowResponseServerType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerShowResponseServerType) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ServerShowResponseServerType) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ServerShowResponseServerType) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ServerShowResponseServerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerShowResponseServerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerShowResponseServerType) SetName(v string)`

SetName sets Name field to given value.


### GetRam

`func (o *ServerShowResponseServerType) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ServerShowResponseServerType) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ServerShowResponseServerType) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetSwap

`func (o *ServerShowResponseServerType) GetSwap() int32`

GetSwap returns the Swap field if non-nil, zero value otherwise.

### GetSwapOk

`func (o *ServerShowResponseServerType) GetSwapOk() (*int32, bool)`

GetSwapOk returns a tuple with the Swap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwap

`func (o *ServerShowResponseServerType) SetSwap(v int32)`

SetSwap sets Swap field to given value.


### GetVcpus

`func (o *ServerShowResponseServerType) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *ServerShowResponseServerType) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *ServerShowResponseServerType) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


