# VolumeShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootable** | **bool** | Bootable | 
**CreatedAt** | **time.Time** | Created at | 
**Encrypted** | **bool** | Encrypted | 
**Id** | **string** | Volume ID | 
**ModifiedAt** | **NullableTime** |  | 
**Multiattach** | **bool** | Multiattach | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Servers** | [**[]VolumeServer**](VolumeServer.md) | Servers attached | 
**Size** | **int32** | Volume size | 
**SnapshotId** | Pointer to **NullableString** |  | [optional] 
**State** | **string** | Volume state | 
**UserId** | **string** | User ID | 
**VolumeImageMetadata** | Pointer to [**NullableVolumeImageMetadataResponse**](VolumeImageMetadataResponse.md) |  | [optional] 
**VolumeType** | **string** | Volume type | 
**VolumeTypeId** | **string** | Volume type ID | 

## Methods

### NewVolumeShowResponse

`func NewVolumeShowResponse(bootable bool, createdAt time.Time, encrypted bool, id string, modifiedAt NullableTime, multiattach bool, servers []VolumeServer, size int32, state string, userId string, volumeType string, volumeTypeId string, ) *VolumeShowResponse`

NewVolumeShowResponse instantiates a new VolumeShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeShowResponseWithDefaults

`func NewVolumeShowResponseWithDefaults() *VolumeShowResponse`

NewVolumeShowResponseWithDefaults instantiates a new VolumeShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootable

`func (o *VolumeShowResponse) GetBootable() bool`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *VolumeShowResponse) GetBootableOk() (*bool, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *VolumeShowResponse) SetBootable(v bool)`

SetBootable sets Bootable field to given value.


### GetCreatedAt

`func (o *VolumeShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEncrypted

`func (o *VolumeShowResponse) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *VolumeShowResponse) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *VolumeShowResponse) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.


### GetId

`func (o *VolumeShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *VolumeShowResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VolumeShowResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VolumeShowResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *VolumeShowResponse) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *VolumeShowResponse) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
### GetMultiattach

`func (o *VolumeShowResponse) GetMultiattach() bool`

GetMultiattach returns the Multiattach field if non-nil, zero value otherwise.

### GetMultiattachOk

`func (o *VolumeShowResponse) GetMultiattachOk() (*bool, bool)`

GetMultiattachOk returns a tuple with the Multiattach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiattach

`func (o *VolumeShowResponse) SetMultiattach(v bool)`

SetMultiattach sets Multiattach field to given value.


### GetName

`func (o *VolumeShowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeShowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeShowResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeShowResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VolumeShowResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VolumeShowResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetServers

`func (o *VolumeShowResponse) GetServers() []VolumeServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *VolumeShowResponse) GetServersOk() (*[]VolumeServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *VolumeShowResponse) SetServers(v []VolumeServer)`

SetServers sets Servers field to given value.


### GetSize

`func (o *VolumeShowResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeShowResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeShowResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSnapshotId

`func (o *VolumeShowResponse) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *VolumeShowResponse) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *VolumeShowResponse) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *VolumeShowResponse) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *VolumeShowResponse) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *VolumeShowResponse) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetState

`func (o *VolumeShowResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VolumeShowResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VolumeShowResponse) SetState(v string)`

SetState sets State field to given value.


### GetUserId

`func (o *VolumeShowResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *VolumeShowResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *VolumeShowResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetVolumeImageMetadata

`func (o *VolumeShowResponse) GetVolumeImageMetadata() VolumeImageMetadataResponse`

GetVolumeImageMetadata returns the VolumeImageMetadata field if non-nil, zero value otherwise.

### GetVolumeImageMetadataOk

`func (o *VolumeShowResponse) GetVolumeImageMetadataOk() (*VolumeImageMetadataResponse, bool)`

GetVolumeImageMetadataOk returns a tuple with the VolumeImageMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeImageMetadata

`func (o *VolumeShowResponse) SetVolumeImageMetadata(v VolumeImageMetadataResponse)`

SetVolumeImageMetadata sets VolumeImageMetadata field to given value.

### HasVolumeImageMetadata

`func (o *VolumeShowResponse) HasVolumeImageMetadata() bool`

HasVolumeImageMetadata returns a boolean if a field has been set.

### SetVolumeImageMetadataNil

`func (o *VolumeShowResponse) SetVolumeImageMetadataNil(b bool)`

 SetVolumeImageMetadataNil sets the value for VolumeImageMetadata to be an explicit nil

### UnsetVolumeImageMetadata
`func (o *VolumeShowResponse) UnsetVolumeImageMetadata()`

UnsetVolumeImageMetadata ensures that no value is present for VolumeImageMetadata, not even an explicit nil
### GetVolumeType

`func (o *VolumeShowResponse) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VolumeShowResponse) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VolumeShowResponse) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.


### GetVolumeTypeId

`func (o *VolumeShowResponse) GetVolumeTypeId() string`

GetVolumeTypeId returns the VolumeTypeId field if non-nil, zero value otherwise.

### GetVolumeTypeIdOk

`func (o *VolumeShowResponse) GetVolumeTypeIdOk() (*string, bool)`

GetVolumeTypeIdOk returns a tuple with the VolumeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypeId

`func (o *VolumeShowResponse) SetVolumeTypeId(v string)`

SetVolumeTypeId sets VolumeTypeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


