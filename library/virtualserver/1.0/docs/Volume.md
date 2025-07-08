# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootIndex** | **int32** | Boot index | 
**DeleteOnTermination** | Pointer to **NullableBool** |  | [optional] 
**Size** | **int32** | Volume size | 
**SourceType** | Pointer to [**NullableVolumeSourceType**](VolumeSourceType.md) |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVolume

`func NewVolume(bootIndex int32, size int32, ) *Volume`

NewVolume instantiates a new Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeWithDefaults

`func NewVolumeWithDefaults() *Volume`

NewVolumeWithDefaults instantiates a new Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootIndex

`func (o *Volume) GetBootIndex() int32`

GetBootIndex returns the BootIndex field if non-nil, zero value otherwise.

### GetBootIndexOk

`func (o *Volume) GetBootIndexOk() (*int32, bool)`

GetBootIndexOk returns a tuple with the BootIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootIndex

`func (o *Volume) SetBootIndex(v int32)`

SetBootIndex sets BootIndex field to given value.


### GetDeleteOnTermination

`func (o *Volume) GetDeleteOnTermination() bool`

GetDeleteOnTermination returns the DeleteOnTermination field if non-nil, zero value otherwise.

### GetDeleteOnTerminationOk

`func (o *Volume) GetDeleteOnTerminationOk() (*bool, bool)`

GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnTermination

`func (o *Volume) SetDeleteOnTermination(v bool)`

SetDeleteOnTermination sets DeleteOnTermination field to given value.

### HasDeleteOnTermination

`func (o *Volume) HasDeleteOnTermination() bool`

HasDeleteOnTermination returns a boolean if a field has been set.

### SetDeleteOnTerminationNil

`func (o *Volume) SetDeleteOnTerminationNil(b bool)`

 SetDeleteOnTerminationNil sets the value for DeleteOnTermination to be an explicit nil

### UnsetDeleteOnTermination
`func (o *Volume) UnsetDeleteOnTermination()`

UnsetDeleteOnTermination ensures that no value is present for DeleteOnTermination, not even an explicit nil
### GetSize

`func (o *Volume) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Volume) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Volume) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSourceType

`func (o *Volume) GetSourceType() VolumeSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Volume) GetSourceTypeOk() (*VolumeSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Volume) SetSourceType(v VolumeSourceType)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *Volume) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *Volume) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *Volume) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetType

`func (o *Volume) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Volume) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Volume) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Volume) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Volume) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Volume) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


