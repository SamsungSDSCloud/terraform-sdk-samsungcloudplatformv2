# VolumeGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** | Total count | [optional] [default to 0]
**VolumeGroups** | Pointer to [**[]VolumeGroupListModel**](VolumeGroupListModel.md) | List of volume groups | [optional] 

## Methods

### NewVolumeGroupListResponse

`func NewVolumeGroupListResponse() *VolumeGroupListResponse`

NewVolumeGroupListResponse instantiates a new VolumeGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupListResponseWithDefaults

`func NewVolumeGroupListResponseWithDefaults() *VolumeGroupListResponse`

NewVolumeGroupListResponseWithDefaults instantiates a new VolumeGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *VolumeGroupListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *VolumeGroupListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *VolumeGroupListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *VolumeGroupListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetVolumeGroups

`func (o *VolumeGroupListResponse) GetVolumeGroups() []VolumeGroupListModel`

GetVolumeGroups returns the VolumeGroups field if non-nil, zero value otherwise.

### GetVolumeGroupsOk

`func (o *VolumeGroupListResponse) GetVolumeGroupsOk() (*[]VolumeGroupListModel, bool)`

GetVolumeGroupsOk returns a tuple with the VolumeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroups

`func (o *VolumeGroupListResponse) SetVolumeGroups(v []VolumeGroupListModel)`

SetVolumeGroups sets VolumeGroups field to given value.

### HasVolumeGroups

`func (o *VolumeGroupListResponse) HasVolumeGroups() bool`

HasVolumeGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


