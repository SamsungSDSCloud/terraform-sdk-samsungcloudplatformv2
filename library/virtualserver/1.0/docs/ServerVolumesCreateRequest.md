# ServerVolumesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnTermination** | Pointer to **NullableBool** |  | [optional] 
**Device** | Pointer to **NullableString** |  | [optional] 
**VolumeId** | **string** | Volume ID | 

## Methods

### NewServerVolumesCreateRequest

`func NewServerVolumesCreateRequest(volumeId string, ) *ServerVolumesCreateRequest`

NewServerVolumesCreateRequest instantiates a new ServerVolumesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerVolumesCreateRequestWithDefaults

`func NewServerVolumesCreateRequestWithDefaults() *ServerVolumesCreateRequest`

NewServerVolumesCreateRequestWithDefaults instantiates a new ServerVolumesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnTermination

`func (o *ServerVolumesCreateRequest) GetDeleteOnTermination() bool`

GetDeleteOnTermination returns the DeleteOnTermination field if non-nil, zero value otherwise.

### GetDeleteOnTerminationOk

`func (o *ServerVolumesCreateRequest) GetDeleteOnTerminationOk() (*bool, bool)`

GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnTermination

`func (o *ServerVolumesCreateRequest) SetDeleteOnTermination(v bool)`

SetDeleteOnTermination sets DeleteOnTermination field to given value.

### HasDeleteOnTermination

`func (o *ServerVolumesCreateRequest) HasDeleteOnTermination() bool`

HasDeleteOnTermination returns a boolean if a field has been set.

### SetDeleteOnTerminationNil

`func (o *ServerVolumesCreateRequest) SetDeleteOnTerminationNil(b bool)`

 SetDeleteOnTerminationNil sets the value for DeleteOnTermination to be an explicit nil

### UnsetDeleteOnTermination
`func (o *ServerVolumesCreateRequest) UnsetDeleteOnTermination()`

UnsetDeleteOnTermination ensures that no value is present for DeleteOnTermination, not even an explicit nil
### GetDevice

`func (o *ServerVolumesCreateRequest) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ServerVolumesCreateRequest) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ServerVolumesCreateRequest) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ServerVolumesCreateRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *ServerVolumesCreateRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *ServerVolumesCreateRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetVolumeId

`func (o *ServerVolumesCreateRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *ServerVolumesCreateRequest) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *ServerVolumesCreateRequest) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


