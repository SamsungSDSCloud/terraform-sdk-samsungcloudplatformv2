# VolumeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** | Total count | [optional] [default to 0]
**Volumes** | Pointer to [**[]VolumeListModel**](VolumeListModel.md) | List of volumes | [optional] 

## Methods

### NewVolumeListResponse

`func NewVolumeListResponse() *VolumeListResponse`

NewVolumeListResponse instantiates a new VolumeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeListResponseWithDefaults

`func NewVolumeListResponseWithDefaults() *VolumeListResponse`

NewVolumeListResponseWithDefaults instantiates a new VolumeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *VolumeListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *VolumeListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *VolumeListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *VolumeListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetVolumes

`func (o *VolumeListResponse) GetVolumes() []VolumeListModel`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *VolumeListResponse) GetVolumesOk() (*[]VolumeListModel, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *VolumeListResponse) SetVolumes(v []VolumeListModel)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *VolumeListResponse) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


