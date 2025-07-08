# BlockStorageGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID | 
**Name** | **string** | Name | 
**RoleType** | [**BlockStorageGroupRoleType**](BlockStorageGroupRoleType.md) | Role type | 
**SizeGb** | **int32** | Size in GB | 
**VolumeType** | [**VolumeType**](VolumeType.md) | Volume type | 

## Methods

### NewBlockStorageGroupResponse

`func NewBlockStorageGroupResponse(id string, name string, roleType BlockStorageGroupRoleType, sizeGb int32, volumeType VolumeType, ) *BlockStorageGroupResponse`

NewBlockStorageGroupResponse instantiates a new BlockStorageGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockStorageGroupResponseWithDefaults

`func NewBlockStorageGroupResponseWithDefaults() *BlockStorageGroupResponse`

NewBlockStorageGroupResponseWithDefaults instantiates a new BlockStorageGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlockStorageGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlockStorageGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlockStorageGroupResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BlockStorageGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlockStorageGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlockStorageGroupResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRoleType

`func (o *BlockStorageGroupResponse) GetRoleType() BlockStorageGroupRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *BlockStorageGroupResponse) GetRoleTypeOk() (*BlockStorageGroupRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *BlockStorageGroupResponse) SetRoleType(v BlockStorageGroupRoleType)`

SetRoleType sets RoleType field to given value.


### GetSizeGb

`func (o *BlockStorageGroupResponse) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *BlockStorageGroupResponse) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *BlockStorageGroupResponse) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.


### GetVolumeType

`func (o *BlockStorageGroupResponse) GetVolumeType() VolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *BlockStorageGroupResponse) GetVolumeTypeOk() (*VolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *BlockStorageGroupResponse) SetVolumeType(v VolumeType)`

SetVolumeType sets VolumeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


