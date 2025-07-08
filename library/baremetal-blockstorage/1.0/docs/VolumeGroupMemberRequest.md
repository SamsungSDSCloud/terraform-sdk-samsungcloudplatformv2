# VolumeGroupMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeIds** | **[]string** | List of volume(Block Storage) id | 

## Methods

### NewVolumeGroupMemberRequest

`func NewVolumeGroupMemberRequest(volumeIds []string, ) *VolumeGroupMemberRequest`

NewVolumeGroupMemberRequest instantiates a new VolumeGroupMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupMemberRequestWithDefaults

`func NewVolumeGroupMemberRequestWithDefaults() *VolumeGroupMemberRequest`

NewVolumeGroupMemberRequestWithDefaults instantiates a new VolumeGroupMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeIds

`func (o *VolumeGroupMemberRequest) GetVolumeIds() []string`

GetVolumeIds returns the VolumeIds field if non-nil, zero value otherwise.

### GetVolumeIdsOk

`func (o *VolumeGroupMemberRequest) GetVolumeIdsOk() (*[]string, bool)`

GetVolumeIdsOk returns a tuple with the VolumeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeIds

`func (o *VolumeGroupMemberRequest) SetVolumeIds(v []string)`

SetVolumeIds sets VolumeIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


