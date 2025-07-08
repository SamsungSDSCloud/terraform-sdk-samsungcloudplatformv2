# BlockStorageGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleType** | [**BlockStorageGroupRoleType**](BlockStorageGroupRoleType.md) | Role type | 
**SizeGb** | **int32** | Size in GB | 
**VolumeType** | Pointer to [**VolumeType**](VolumeType.md) | Volume type | [optional] [default to VOLUMETYPE_SSD]

## Methods

### NewBlockStorageGroupRequest

`func NewBlockStorageGroupRequest(roleType BlockStorageGroupRoleType, sizeGb int32, ) *BlockStorageGroupRequest`

NewBlockStorageGroupRequest instantiates a new BlockStorageGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockStorageGroupRequestWithDefaults

`func NewBlockStorageGroupRequestWithDefaults() *BlockStorageGroupRequest`

NewBlockStorageGroupRequestWithDefaults instantiates a new BlockStorageGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleType

`func (o *BlockStorageGroupRequest) GetRoleType() BlockStorageGroupRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *BlockStorageGroupRequest) GetRoleTypeOk() (*BlockStorageGroupRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *BlockStorageGroupRequest) SetRoleType(v BlockStorageGroupRoleType)`

SetRoleType sets RoleType field to given value.


### GetSizeGb

`func (o *BlockStorageGroupRequest) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *BlockStorageGroupRequest) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *BlockStorageGroupRequest) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.


### GetVolumeType

`func (o *BlockStorageGroupRequest) GetVolumeType() VolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *BlockStorageGroupRequest) GetVolumeTypeOk() (*VolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *BlockStorageGroupRequest) SetVolumeType(v VolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *BlockStorageGroupRequest) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


