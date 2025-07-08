# ServerVolumesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnTermination** | Pointer to **NullableBool** |  | [optional] 
**VolumeId** | **string** | Volume ID | 

## Methods

### NewServerVolumesUpdateRequest

`func NewServerVolumesUpdateRequest(volumeId string, ) *ServerVolumesUpdateRequest`

NewServerVolumesUpdateRequest instantiates a new ServerVolumesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerVolumesUpdateRequestWithDefaults

`func NewServerVolumesUpdateRequestWithDefaults() *ServerVolumesUpdateRequest`

NewServerVolumesUpdateRequestWithDefaults instantiates a new ServerVolumesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnTermination

`func (o *ServerVolumesUpdateRequest) GetDeleteOnTermination() bool`

GetDeleteOnTermination returns the DeleteOnTermination field if non-nil, zero value otherwise.

### GetDeleteOnTerminationOk

`func (o *ServerVolumesUpdateRequest) GetDeleteOnTerminationOk() (*bool, bool)`

GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnTermination

`func (o *ServerVolumesUpdateRequest) SetDeleteOnTermination(v bool)`

SetDeleteOnTermination sets DeleteOnTermination field to given value.

### HasDeleteOnTermination

`func (o *ServerVolumesUpdateRequest) HasDeleteOnTermination() bool`

HasDeleteOnTermination returns a boolean if a field has been set.

### SetDeleteOnTerminationNil

`func (o *ServerVolumesUpdateRequest) SetDeleteOnTerminationNil(b bool)`

 SetDeleteOnTerminationNil sets the value for DeleteOnTermination to be an explicit nil

### UnsetDeleteOnTermination
`func (o *ServerVolumesUpdateRequest) UnsetDeleteOnTermination()`

UnsetDeleteOnTermination ensures that no value is present for DeleteOnTermination, not even an explicit nil
### GetVolumeId

`func (o *ServerVolumesUpdateRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *ServerVolumesUpdateRequest) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *ServerVolumesUpdateRequest) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


