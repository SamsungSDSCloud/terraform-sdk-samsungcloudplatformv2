# VolumeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Volumes** | [**[]VolumeShowResponse**](VolumeShowResponse.md) |  | 

## Methods

### NewVolumeListResponse

`func NewVolumeListResponse(volumes []VolumeShowResponse, ) *VolumeListResponse`

NewVolumeListResponse instantiates a new VolumeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeListResponseWithDefaults

`func NewVolumeListResponseWithDefaults() *VolumeListResponse`

NewVolumeListResponseWithDefaults instantiates a new VolumeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VolumeListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VolumeListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VolumeListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *VolumeListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *VolumeListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *VolumeListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetVolumes

`func (o *VolumeListResponse) GetVolumes() []VolumeShowResponse`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *VolumeListResponse) GetVolumesOk() (*[]VolumeShowResponse, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *VolumeListResponse) SetVolumes(v []VolumeShowResponse)`

SetVolumes sets Volumes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


