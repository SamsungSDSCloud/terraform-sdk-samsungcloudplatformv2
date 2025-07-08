# LaunchConfigurationVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootIndex** | **int32** | Boot index | 
**Size** | **int32** | Volume size | 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLaunchConfigurationVolume

`func NewLaunchConfigurationVolume(bootIndex int32, size int32, ) *LaunchConfigurationVolume`

NewLaunchConfigurationVolume instantiates a new LaunchConfigurationVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchConfigurationVolumeWithDefaults

`func NewLaunchConfigurationVolumeWithDefaults() *LaunchConfigurationVolume`

NewLaunchConfigurationVolumeWithDefaults instantiates a new LaunchConfigurationVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootIndex

`func (o *LaunchConfigurationVolume) GetBootIndex() int32`

GetBootIndex returns the BootIndex field if non-nil, zero value otherwise.

### GetBootIndexOk

`func (o *LaunchConfigurationVolume) GetBootIndexOk() (*int32, bool)`

GetBootIndexOk returns a tuple with the BootIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootIndex

`func (o *LaunchConfigurationVolume) SetBootIndex(v int32)`

SetBootIndex sets BootIndex field to given value.


### GetSize

`func (o *LaunchConfigurationVolume) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LaunchConfigurationVolume) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LaunchConfigurationVolume) SetSize(v int32)`

SetSize sets Size field to given value.


### GetType

`func (o *LaunchConfigurationVolume) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LaunchConfigurationVolume) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LaunchConfigurationVolume) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LaunchConfigurationVolume) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *LaunchConfigurationVolume) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LaunchConfigurationVolume) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


