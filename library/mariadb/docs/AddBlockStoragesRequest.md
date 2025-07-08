# AddBlockStoragesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleType** | [**BlockStorageGroupRoleType**](BlockStorageGroupRoleType.md) | Role type | 
**SizeGb** | **int32** | Size in GB | 
**VolumeType** | Pointer to [**VolumeType**](VolumeType.md) | Volume type | [optional] [default to VOLUMETYPE_SSD]

## Methods

### NewAddBlockStoragesRequest

`func NewAddBlockStoragesRequest(roleType BlockStorageGroupRoleType, sizeGb int32, ) *AddBlockStoragesRequest`

NewAddBlockStoragesRequest instantiates a new AddBlockStoragesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlockStoragesRequestWithDefaults

`func NewAddBlockStoragesRequestWithDefaults() *AddBlockStoragesRequest`

NewAddBlockStoragesRequestWithDefaults instantiates a new AddBlockStoragesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleType

`func (o *AddBlockStoragesRequest) GetRoleType() BlockStorageGroupRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *AddBlockStoragesRequest) GetRoleTypeOk() (*BlockStorageGroupRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *AddBlockStoragesRequest) SetRoleType(v BlockStorageGroupRoleType)`

SetRoleType sets RoleType field to given value.


### GetSizeGb

`func (o *AddBlockStoragesRequest) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *AddBlockStoragesRequest) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *AddBlockStoragesRequest) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.


### GetVolumeType

`func (o *AddBlockStoragesRequest) GetVolumeType() VolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *AddBlockStoragesRequest) GetVolumeTypeOk() (*VolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *AddBlockStoragesRequest) SetVolumeType(v VolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *AddBlockStoragesRequest) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


